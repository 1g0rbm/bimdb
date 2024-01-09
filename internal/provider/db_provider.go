package provider

import (
	"fmt"

	"bimdb/internal/database"
	"bimdb/internal/database/compute"
	"bimdb/internal/database/storage"
)

func (p *Provider) getDB() (*database.DB, error) {
	if p.db == nil {
		logger, err := p.getLogger()
		if err != nil {
			return nil, err
		}

		computer, err := p.getComputerLayer()
		if err != nil {
			return nil, err
		}

		store, err := p.getStorageLayer()
		if err != nil {
			return nil, err
		}

		db, err := database.NewDB(computer, store, logger)
		if err != nil {
			return nil, err
		}

		p.db = db
	}

	return p.db, nil
}

func (p *Provider) getComputerLayer() (*compute.Computer, error) {
	logger, err := p.getLogger()
	if err != nil {
		return nil, err
	}

	parser, err := compute.NewParser(logger)
	if err != nil {
		return nil, fmt.Errorf("failed to create parser: %w", err)
	}

	computer, err := compute.NewComputer(parser, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to create computer layer: %w", err)
	}

	return computer, nil
}

func (p *Provider) getStorageLayer() (*storage.Storage, error) {
	logger, err := p.getLogger()
	if err != nil {
		return nil, err
	}

	engine, err := p.getEngine()
	if err != nil {
		return nil, err
	}

	store, err := storage.NewStorage(engine, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to create storage layer: %w", err)
	}

	return store, nil
}
