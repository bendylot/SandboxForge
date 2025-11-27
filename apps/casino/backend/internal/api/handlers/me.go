package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Me godoc
// @Summary      Информация о текущем пользователе
// @Tags         users
// @Security     BearerAuth
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      401  {object}  map[string]string
// @Router       /api/me [get]
func (h *Handlers) Me(c *gin.Context) {
	claimsVal, _ := c.Get("claims")
	claims := claimsVal.(map[string]interface{})
	c.JSON(http.StatusOK, gin.H{
		"id":    claims["sub"],
		"email": claims["email"],
	})
}
