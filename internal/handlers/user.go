package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) UserHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "User Handler"
	c.JSON(200, resp)

}