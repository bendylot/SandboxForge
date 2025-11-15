package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"casino.sandboxforge.tech/casino-api/internal/config"
	"casino.sandboxforge.tech/casino-api/internal/repo"
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
// @Param        data  body  RegisterReq  true  "Email и пароль"
// @Success      201   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]string
// @Failure      409   {object}  map[string]string
// @Router       /api/register [post]
func (h *Handlers) Register(c *gin.Context) {
	var req RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad json"})
		return
	}
	email := strings.ToLower(strings.TrimSpace(req.Email))
	if email == "" || len(req.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email/password invalid"})
		return
	}
	// exists?
	existing, err := h.Repo.ByEmail(c, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}
	if existing != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "email already registered"})
		return
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	uid, err := h.Repo.Create(context.Background(), email, string(hash))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db insert error"})
		return
	}

	token, err := makeToken(h.Cfg.JWTSecret, h.Cfg.JWTTTL, uid, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "jwt error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":    uid,
		"email": email,
		"token": token,
	})
}

// Login godoc
// @Summary      Вход пользователя
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        data  body  LoginReq  true  "Email и пароль"
// @Success      200   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Router       /api/login [post]
func (h *Handlers) Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad json"})
		return
	}
	email := strings.ToLower(strings.TrimSpace(req.Email))
	u, err := h.Repo.ByEmail(c, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}
	if u == nil || bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	token, err := makeToken(h.Cfg.JWTSecret, h.Cfg.JWTTTL, u.ID, u.Email)
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
