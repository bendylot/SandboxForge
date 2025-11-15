//go:generate swag init -g main.go -o docs -d .,../internal/api

package main

// @title           Casino API
// @version         1.0
// @description     Бэкенд авторизации для казино-проекта.
// @host            localhost:8080
// @BasePath        /
// @schemes         http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
import (
	"fmt"
	"log"

	"casino.sandboxforge.tech/casino-api/internal/api"
	"casino.sandboxforge.tech/casino-api/internal/config"
	"casino.sandboxforge.tech/casino-api/internal/domain"
	"casino.sandboxforge.tech/casino-api/internal/repo"
)

func main() {
	cfg := config.Load()
	pool := domain.NewPool(cfg.DBURL)
	domain.AutoMigrate(pool)

	userRepo := repo.NewUserRepo(pool)
	router := api.NewRouter(cfg, userRepo)

	addr := fmt.Sprintf(":%s", cfg.AppPort)
	log.Printf("API listening on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
