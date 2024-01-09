package provider

import (
	"bimdb/internal/database"
	"fmt"
	"log/slog"

	"bimdb/internal/config"
	"bimdb/internal/database/storage/engine"
	"bimdb/internal/network"
)

type Provider struct {
	config *config.Config
	server *network.TCPServer
	engine engine.IEngine
	db     *database.DB
	logger *slog.Logger
}

func NewProvider(cfg *config.Config) (*Provider, error) {
	if cfg == nil {
		return nil, fmt.Errorf("invalid config value")
	}

	return &Provider{
		config: cfg,
	}, nil
}

func (p *Provider) GetServer() (*network.TCPServer, error) {
	return p.getServer()
}

func (p *Provider) GetDB() (*database.DB, error) {
	return p.getDB()
}
