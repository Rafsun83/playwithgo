package handlers

import (
	"net/http"

	"playwithgo/internal/repository"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	DB repository.Service
}


func (h *Handler) HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, h.DB.Health())
}
