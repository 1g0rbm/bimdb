package provider

import (
	"fmt"

	"bimdb/internal/database/storage/engine"
	"bimdb/internal/database/storage/engine/in_memory"
)

const (
	inMemoryEngineType = "in_memory"
)

const defaultEngineType = inMemoryEngineType

func (p *Provider) getEngine() (engine.IEngine, error) {
	if p.engine == nil {
		engineType := defaultEngineType

		cfg := p.config.EngineConfig
		if cfg != nil {
			if cfg.Type != "" {
				engineType = cfg.Type
			}
		}

		switch engineType {
		case inMemoryEngineType:
			p.engine = in_memory.NewEngine()
		default:
			return nil, fmt.Errorf("invalid engine type %s", engineType)
		}
	}

	return p.engine, nil
}
