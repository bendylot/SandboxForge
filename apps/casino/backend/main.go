//go:generate swag init -g main.go -o docs -d .,../internal/api

package main

// @title           Casino API
// @version         1.0
// @description     Бэкенд авторизации для казино-проекта.
// @host            localhost:3001
// @BasePath        /
// @schemes         http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
import (
	"fmt"
	"log"

	"casino.sandboxforge.tech/casino-api/docs"
	"casino.sandboxforge.tech/casino-api/internal/api/router"
	"casino.sandboxforge.tech/casino-api/internal/app/config"
	"casino.sandboxforge.tech/casino-api/internal/repo/postgres"
	"casino.sandboxforge.tech/casino-api/internal/service"
)

func main() {
	cfg := config.Load()
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.AppPort)
	docs.SwaggerInfo.BasePath = "/"

	pool := postgres.NewPool(cfg.DBURL)
	postgres.AutoMigrate(pool)

	userRepo := postgres.NewUserRepo(pool)
	authSrc := service.NewAuthService(userRepo, cfg)
	router := router.NewRouter(cfg, authSrc)

	addr := ":" + cfg.AppPort
	log.Printf("API listening on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
