package handlers

import (
	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary      Вход пользователя
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        data  body  dto.LoginReq  true  "Логин и пароль"
// @Success      200   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Router       /api/auth/login [post]
func (h *Handlers) Login(c *gin.Context) {
}
