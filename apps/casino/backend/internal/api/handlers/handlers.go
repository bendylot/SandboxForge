package handlers

import (
	"casino.sandboxforge.tech/casino-api/internal/app/config"
	"casino.sandboxforge.tech/casino-api/internal/service"
)

type Handlers struct {
	Cfg  *config.Config
	Svc  service.Auth
}
