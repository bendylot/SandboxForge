package main

// @title           Casino API Auth Service
// @version         1.0
// @description     Бэкенд авторизации для SandboxForge (Go + Gin + JWT).
// @host            localhost:8080
// @BasePath        /
// @schemes         http
import (
	"fmt"
	"log"

	_ "casino.sandboxforge.tech/casino-api/cmd/api/docs"
	"casino.sandboxforge.tech/casino-api/internal/config"
	"casino.sandboxforge.tech/casino-api/internal/db"
	"casino.sandboxforge.tech/casino-api/internal/http"
	"casino.sandboxforge.tech/casino-api/internal/repo"
)

func main() {
	cfg := config.Load()
	pool := db.NewPool(cfg.DBURL)
	db.AutoMigrate(pool)

	userRepo := repo.NewUserRepo(pool)
	router := http.NewRouter(cfg, userRepo)

	addr := fmt.Sprintf(":%s", cfg.AppPort)
	log.Printf("API listening on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
