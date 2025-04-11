package services

import (
	"catify/internal/db"
	"catify/internal/elastic"
	"catify/internal/models"
	"catify/internal/redis"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"gorm.io/gorm"
)

type MusicServices struct {
	db     *db.DB
	client *redis.Client
	es     *elastic.Client
}

func NewMusicServices(db *db.DB, client *redis.Client, esClient *elastic.Client) *MusicServices {
	return &MusicServices{
		db:     db,
		client: client,
		es:     esClient,
	}
}

func (s *MusicServices) SelectMusic() *gorm.DB {
	return s.db.Select("id, title, artist, genre, file_name, file_size, content_type, duration, description, create_at, update_at")
}

func (s *MusicServices) GetAllMusic(ctx context.Context) ([]models.Music, error) {
	var musics []models.Music
	if err := s.SelectMusic().Find(&musics).Error; err != nil {
		return nil, err
	}

	return musics, nil
}

func (s *MusicServices) GetMusicFileData(ctx context.Context, id uint) ([]byte, string, string, error) {
	var music models.Music

	if err := s.db.Select("file_data", "file_name", "content_type").First(&music, id).Error; err != nil {
		return nil, "", "", err
	}

	return music.FileData, music.FileName, music.ContentType, nil
}

func (s *MusicServices) GetMusicById(id uint) (*models.Music, error) {
	var music models.Music
	if err := s.SelectMusic().First(&music, id).Error; err != nil {
		return nil, err
	}

	return &music, nil
}

func (s *MusicServices) SearchMusic(ctx context.Context, query string) ([]models.Music, error) {
	var musics []models.Music

	cacheKey := fmt.Sprintf("search:" + query)
	cacheValue, err := s.client.Get(ctx, cacheKey).Result()
	if err == nil {
		var cacheMusics []models.Music
		if err := json.Unmarshal([]byte(cacheValue), &cacheMusics); err == nil {
			return cacheMusics, nil
		}
	}

	// Elasticsearch query
	searchBody := fmt.Sprintf(`{
        "query": {
            "multi_match": {
                "query": "%s",
                "fields": "title"
            }
        }
    }`, query)

	res, err := s.es.Search(
		s.es.Search.WithContext(ctx),
		s.es.Search.WithIndex("musics"),
		s.es.Search.WithBody(strings.NewReader(searchBody)),
		s.es.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error in Elasticsearch response: %s", res.String())
	}

	var esResponse struct {
		Hits struct {
			Hits []struct {
				Source models.Music `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&esResponse); err != nil {
		return nil, err
	}

	for _, hit := range esResponse.Hits.Hits {
		musics = append(musics, hit.Source)
	}

	// Cache the result in Redis
	if len(musics) > 0 {
		if data, err := json.Marshal(musics); err == nil {
			s.client.Set(ctx, cacheKey, data, 10*time.Minute)
		}
	}

	return musics, nil
}

func (s *MusicServices) CreateNewMusic(ctx context.Context, music *models.Music) error {
	// Save to the database
	if err := s.db.Save(music).Error; err != nil {
		return err
	}

	// Index the music into Elasticsearch
	musicJSON, err := json.Marshal(music)
	if err != nil {
		return fmt.Errorf("failed to marshal music for Elasticsearch: %w", err)
	}

	_, err = s.es.Index(
		"musics",
		strings.NewReader(string(musicJSON)),
	)
	if err != nil {
		return fmt.Errorf("failed to index music into Elasticsearch: %w", err)
	}

	return nil
}

func (s *MusicServices) Delete(ctx context.Context, id uint) error {
	var music models.Music

	if err := s.db.First(&music, id).Error; err != nil {
		return err
	}

	return s.db.Delete(&music).Error
}

func (s *MusicServices) UpdateMusicInfo(ctx context.Context, music *models.Music) error {
	// Update in the database
	result := s.db.Model(&models.Music{}).Where("id = ?", music.ID).Updates(map[string]interface{}{
		"title":       music.Title,
		"artist":      music.Artist,
		"album":       music.Album,
		"genre":       music.Genre,
		"description": music.Description,
	})
	if result.Error != nil {
		return result.Error
	}

	// Update in Elasticsearch
	musicJSON, err := json.Marshal(music)
	if err != nil {
		return fmt.Errorf("failed to marshal music for Elasticsearch: %w", err)
	}

	_, err = s.es.Index(
		"musics",
		strings.NewReader(string(musicJSON)),
		s.es.Index.WithContext(ctx),
		s.es.Index.WithDocumentID(fmt.Sprintf("%d", music.ID)),
		s.es.Index.WithRefresh("true"),
	)
	if err != nil {
		return fmt.Errorf("failed to update music in Elasticsearch: %w", err)
	}

	return nil
}

func (s *MusicServices) UploadMusic(ctx context.Context, file *multipart.FileHeader, music *models.Music) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	fileData, err := io.ReadAll(src)
	if err != nil {
		return err
	}

	ext := filepath.Ext(file.Filename)
	contentType := getContentTypeFromExt(ext)

	music.ContentType = contentType
	music.FileName = file.Filename
	music.FileSize = file.Size
	music.FileData = fileData

	// Save to the database and Elasticsearch
	return s.CreateNewMusic(ctx, music)
}

func (s *MusicServices) UpdateMusicSound(ctx context.Context, id uint, file *multipart.FileHeader) error {
	var music models.Music
	if err := s.db.First(&music, id).Error; err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	fileData, err := io.ReadAll(src)
	if err != nil {
		return err
	}

	ext := filepath.Ext(file.Filename)
	contentType := getContentTypeFromExt(ext)

	// Update in the database
	result := s.db.Model(&music).Updates(map[string]interface{}{
		"file_name":    file.Filename,
		"file_data":    fileData,
		"file_size":    file.Size,
		"content_type": contentType,
	})
	if result.Error != nil {
		return result.Error
	}

	// Update in Elasticsearch
	music.FileName = file.Filename
	music.FileData = fileData
	music.FileSize = file.Size
	music.ContentType = contentType

	musicJSON, err := json.Marshal(music)
	if err != nil {
		return fmt.Errorf("failed to marshal music for Elasticsearch: %w", err)
	}

	_, err = s.es.Index(
		"musics",
		strings.NewReader(string(musicJSON)),
		s.es.Index.WithContext(ctx),
		s.es.Index.WithDocumentID(fmt.Sprintf("%d", music.ID)),
		s.es.Index.WithRefresh("true"),
	)
	if err != nil {
		return fmt.Errorf("failed to update music in Elasticsearch: %w", err)
	}

	return nil
}

func getContentTypeFromExt(ext string) string {
	switch strings.ToLower(ext) {
	case ".mp3":
		return "audio/mpeg"
	case ".wav":
		return "audio/wav"
	case ".flac":
		return "audio/flac"
	case ".aac":
		return "audio/aac"
	case ".ogg":
		return "audio/ogg"
	default:
		return "application/octet-stream"
	}
}
