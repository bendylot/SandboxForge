package http

import (
	"net/http"
	"time"

	"casino.sandboxforge.tech/casino-api/internal/config"
	"casino.sandboxforge.tech/casino-api/internal/repo"
	"github.com/gin-gonic/gin"
)

func NewRouter(cfg *config.Config, ur *repo.UserRepo) *gin.Engine {
	r := gin.Default()

	// Простой CORS для локалки
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
		c.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})

	h := &Handlers{Cfg: cfg, Repo: ur}

	api := r.Group("/api")
	{
		api.POST("/register", h.Register)
		api.POST("/login", h.Login)

		protected := api.Group("")
		protected.Use(AuthMiddleware(cfg))
		{
			protected.GET("/me", h.Me)
		}
	}

	// healthcheck
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": true, "ts": time.Now().UTC()})
	})

	return r
}
