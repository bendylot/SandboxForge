package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"casino.sandboxforge.tech/casino-api/internal/config"
	repo "casino.sandboxforge.tech/casino-api/internal/repo/postgres"
)

type Handlers struct {
	Cfg  *config.Config
	Repo *repo.UserRepo
}

// Register godoc
// @Summary      Регистрация нового пользователя
// @Description  Создает пользователя в системе и возвращает JWT
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        data  body  RegisterReq  true  "Логин и пароль"
// @Success      201   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]string
// @Failure      409   {object}  map[string]string
// @Router       /api/auth/register [post]
func (h *Handlers) Register(c *gin.Context) {
	var req RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad json"})
		return
	}
	login := strings.ToLower(strings.TrimSpace(req.Login))
	if login == "" || len(req.Password) < 2 || len(req.Password) > 24 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "login/password invalid"})
		return
	}
	existing, err := h.Repo.ByLogin(c, login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}
	if existing != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "email already registered"})
		return
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	uid, err := h.Repo.Create(context.Background(), login, string(hash))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db insert error"})
		return
	}

	token, err := makeToken(h.Cfg.JWTSecret, h.Cfg.JWTTTL, uid, login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "jwt error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":    uid,
		"email": login,
		"token": token,
	})
}

// Login godoc
// @Summary      Вход пользователя
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        data  body  LoginReq  true  "Логин и пароль"
// @Success      200   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Router       /api/auth/login [post]
func (h *Handlers) Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad json"})
		return
	}
	login := strings.ToLower(strings.TrimSpace(req.Login))
	u, err := h.Repo.ByLogin(c, login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}
	if u == nil || bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	token, err := makeToken(h.Cfg.JWTSecret, h.Cfg.JWTTTL, u.ID, u.Login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "jwt error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

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
