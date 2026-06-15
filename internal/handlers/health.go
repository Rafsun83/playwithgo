package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"playwithgo/internal/repository"
)

type Handler struct {
	DB repository.Service
}

func (h *Handler) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, h.DB.Health())
}
