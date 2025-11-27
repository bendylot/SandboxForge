package handlers

import (
	"net/http"

	"casino.sandboxforge.tech/casino-api/internal/api/validation"
	"casino.sandboxforge.tech/casino-api/internal/service"
	"github.com/gin-gonic/gin"

	"casino.sandboxforge.tech/casino-api/internal/api/dto"
	"casino.sandboxforge.tech/casino-api/internal/api/token"
)

// Register godoc
// @Summary      Регистрация нового пользователя
// @Description  Создает пользователя в системе и возвращает JWT
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        data  body  dto.RegisterReq  true  "Логин, email, пароль"
// @Success      201   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]string
// @Failure      409   {object}  map[string]string
// @Router       /api/auth/register [post]
// POST /api/auth/register
func (h *Handlers) Register(c *gin.Context) {
	var req dto.RegisterReq
	// конвертация входящего джсона в нужный обьект
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}
	// нормализация обьекта и валидация
	validation.NormalizeRegisterReq(&req)
	if err := validation.ValidateRegisterReq(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation failed"})
		return
	}
	uid, err := h.Svc.Register(c.Request.Context(), req.Login, req.Email, req.Password)
	if err != nil {
		switch err {
		case service.ErrLoginTaken:
			c.JSON(http.StatusConflict, gin.H{"error": "login already registered"})
		case service.ErrEmailTaken:
			c.JSON(http.StatusConflict, gin.H{"error": "email already registered"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		}
		return
	}
	// создание ответа - токена по джсону
	tok, err := token.MakeToken(h.Cfg.JWTSecret(), h.Cfg.JWTTTL(), uid, req.Login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "jwt error"})
		return
	}
	// ответ операции
	c.JSON(http.StatusCreated, gin.H{
		"id":    uid,
		"login": req.Login,
		"email": req.Email,
		"token": tok,
	})
}