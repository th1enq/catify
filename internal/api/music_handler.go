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

func NewMusicHandler(s *services.MusicServices) (*MusicHandler, error) {
	return &MusicHandler{
		musicServices: s,
	}, nil
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

	FileData, FileName, ContentType, err := h.musicServices.GetMusicFileData(c, uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	music := &models.Music{
		
	}
}
