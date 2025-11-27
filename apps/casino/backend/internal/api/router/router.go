package router

import (
	"net/http"
	"time"

	"casino.sandboxforge.tech/casino-api/internal/api/handlers"
	"casino.sandboxforge.tech/casino-api/internal/api/token"
	"casino.sandboxforge.tech/casino-api/internal/app/config"
	"casino.sandboxforge.tech/casino-api/internal/service"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(cfg *config.Config, svc service.Auth) *gin.Engine {
	r := gin.Default()

	// Простой CORS для локалки
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Header("Access-Control-Allow-Credentials", "true")

		c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
		c.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})
	h := &handlers.Handlers{Cfg: cfg, Svc: svc}

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", h.Register)
			auth.POST("/login", h.Login)
		}

		protected := api.Group("")
		protected.Use(token.AuthMiddleware(cfg))
		{
			protected.GET("/me", h.Me)
		}
	}

	// healthcheck
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": true, "ts": time.Now().UTC()})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
