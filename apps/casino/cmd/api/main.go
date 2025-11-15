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
	"casino.sandboxforge.tech/casino-api/internal/db"
	"casino.sandboxforge.tech/casino-api/internal/repo"

	_ "casino.sandboxforge.tech/casino-api/cmd/api/docs"
)

func main() {
	cfg := config.Load()
	pool := db.NewPool(cfg.DBURL)
	db.AutoMigrate(pool)

	userRepo := repo.NewUserRepo(pool)
	router := api.NewRouter(cfg, userRepo)

	addr := fmt.Sprintf(":%s", cfg.AppPort)
	log.Printf("API listening on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
