package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort   string
	DBURL     string
	jWTSecret string
	jWTTTL    time.Duration
	JWT       string
}

func (c *Config) JWTSecret() string {
    return c.jWTSecret // см. заметку ниже о совпадении имён
}

func (c *Config) JWTTTL() time.Duration {
    return c.jWTTTL
}

func Load() *Config {
	_ = godotenv.Load()

	port := getEnv("APP_PORT", "8080")
	db := must("DB_URL")
	secret := must("JWT_SECRET")
	ttlMin := getEnv("JWT_TTL_MINUTES", "60")
	ttlInt, err := strconv.Atoi(ttlMin)
	if err != nil {
		log.Fatalf("bad JWT_TTL_MINUTES: %v", err)
	}

	return &Config{
		AppPort:   port,
		DBURL:     db,
		jWTSecret: secret,
		jWTTTL:    time.Duration(ttlInt) * time.Minute,
	}
}

func must(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("missing required env %s", k)
	}
	return v
}

func getEnv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
