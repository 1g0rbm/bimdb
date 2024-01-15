package provider

import (
	"bimdb/internal/tools"
	"time"

	"bimdb/internal/network"
)

const (
	defaultAddress        = "127.0.0.1:3223"
	defaultIdleTimeout    = time.Minute * 2
	defaultMaxMessageSize = 2048
	defaultMaxConnections = 10
)

func (p *Provider) getServer() (*network.TCPServer, error) {
	if p.server == nil {
		address := defaultAddress
		idleTimeout := defaultIdleTimeout
		maxMessageSize := defaultMaxMessageSize
		maxConnections := defaultMaxConnections

		cfg := p.config.NetworkConfig
		if cfg != nil {
			if cfg.Address != "" {
				address = cfg.Address
			}

			if cfg.IdleTimeout != 0 {
				idleTimeout = cfg.IdleTimeout
			}

			if cfg.MaxMessageSize != "" {
				var err error
				maxMessageSize, err = tools.ParseSize(cfg.MaxMessageSize)
				if err != nil {
					return nil, err
				}
			}

			if cfg.MaxConnections != 0 {
				maxConnections = cfg.MaxConnections
			}
		}

		logger, err := p.getLogger()
		if err != nil {
			return nil, err
		}

		server, err := network.NewTCPServer(address, idleTimeout, maxMessageSize, maxConnections, logger)
		if err != nil {
			return nil, err
		}

		p.server = server
	}

	return p.server, nil
}
