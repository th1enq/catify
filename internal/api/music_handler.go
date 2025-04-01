package api

import (
	"catify/internal/models"
	"catify/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MusicHandler struct {
	musicServices *services.MusicServices
}

func NewMusicHandler(s *services.MusicServices) *MusicHandler {
	return &MusicHandler{
		musicServices: s,
	}
}

func (h *MusicHandler) GetAllMusic(c *gin.Context) {
	music, err := h.musicServices.GetAllMusic(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}
	c.JSON(http.StatusOK, music)
}

func (h *MusicHandler) GetMusicFileData(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	fileData, fileName, contentType, err := h.musicServices.GetMusicFileData(c, uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Data(http.StatusOK, contentType, fileData)
}

func (h *MusicHandler) GetMusicById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	music, err := h.musicServices.GetMusicById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, music)
}

func (h *MusicHandler) SearchMusic(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "search query is required"})
		return
	}

	musics, err := h.musicServices.SearchMusic(c, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, musics)
}

func (h *MusicHandler) CreateNewMusic(c *gin.Context) {
	var music models.Music
	if err := c.ShouldBindJSON(&music); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.musicServices.CreateNewMusic(c, &music); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, music)
}

func (h *MusicHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}
	if err := h.musicServices.Delete(c, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "music deleted successfully"})
}

func (h *MusicHandler) UpdateMusicInfo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}
	var music models.Music
	if err := c.ShouldBindJSON(&music); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	music.ID = uint(id)
	if err := h.musicServices.UpdateMusicInfo(c, &music); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "music information updated successfully"})
}

func (h *MusicHandler) UploadMusic(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file upload failed"})
		return
	}

	var music models.Music
	if err := c.ShouldBindJSON(&music); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.musicServices.UploadMusic(c, file, &music); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "music uploaded successfully",
		"music":   music,
	})
}

func (h *MusicHandler) UpdateMusicSound(c *gin.Context) {

}
