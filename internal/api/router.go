package api

import (
	"catify/internal/db"
	"catify/internal/elastic"
	"catify/internal/redis"
	"catify/internal/services"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *db.DB, redisClient *redis.Client, esClient *elastic.Client) *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default())

	musicServices := services.NewMusicServices(db, redisClient, esClient)

	musicHandler := NewMusicHandler(musicServices)

	api := router.Group("/api")
	{
		music := api.Group("/music")
		{
			music.GET("", musicHandler.GetAllMusic)
			music.GET("/:id", musicHandler.GetMusicById)
			music.GET("/search", musicHandler.SearchMusic)
			music.GET("/download/:id", musicHandler.GetMusicFileData)

			music.POST("", musicHandler.CreateNewMusic)
			music.POST("/:id/sound", musicHandler.UploadMusic)

			music.DELETE("/:id", musicHandler.Delete)
		}
	}

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"user":   "th1enq",
		})
	})

	return router
}
