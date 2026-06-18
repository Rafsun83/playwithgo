package handlers

import (
	"net/http"

	"playwithgo/internal/repository"

	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/cors"
	// "github.com/gin-gonic/gin"
)

func RegisterRoutes(db repository.Service) http.Handler {
	r := gin.Default()

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:5173"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
	// 	AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
	// 	AllowCredentials: true,
	// }))

	h := &Handler{DB: db}

	r.GET("/", h.HelloWorldHandler)
	r.GET("/health", h.HealthHandler)
	r.GET("/user", h.UserHandler)

	return r
}
