package api

import (
	"catify/internal/db"
	"catify/internal/redis"
	"catify/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *db.DB, redisClient *redis.Client) *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default())

	musicServices := services.NewMusicServices(db, redisClient)

	musicHandler := 
}
