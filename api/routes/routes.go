package routes

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/cat/repositories"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, es *elasticsearch.Client, jwtSecret string) *mux.Router {
	router := mux.NewRouter()

	songRepo := repositories.
}
