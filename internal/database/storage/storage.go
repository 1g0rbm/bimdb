package storage

import (
	"context"
	"errors"
	"fmt"

	"log/slog"

	"bimdb/internal/database/storage/engine"
)

var (
	errInvalidEngine = errors.New("invalid engine")
	errInvalidLogger = errors.New("invalid logger")
)

type Storage struct {
	engine engine.IEngine
	logger *slog.Logger
}

func NewStorage(e engine.IEngine, l *slog.Logger) (*Storage, error) {
	if e == nil {
		return nil, errInvalidEngine
	}

	if l == nil {
		return nil, errInvalidLogger
	}

	return &Storage{
		engine: e,
		logger: l,
	}, nil
}

func (s *Storage) Set(ctx context.Context, key string, value string) error {
	if ctx.Err() != nil {
		s.logger.Debug("Canceled action SET.")
		return ctx.Err()
	}

	s.engine.Set(key, value)
	s.logger.Debug(fmt.Sprintf("Action: SET | Key: %s | Value: %s", key, value))

	return nil
}

func (s *Storage) Get(ctx context.Context, key string) (string, error) {
	if ctx.Err() != nil {
		s.logger.Debug("Canceled action GET.")
		return "", ctx.Err()
	}

	value, ok := s.engine.Get(key)
	if !ok {
		return "", fmt.Errorf("key %s not found", key)
	}

	s.logger.Debug(fmt.Sprintf("Action: GET | Key: %s", key))

	return value, nil
}

func (s *Storage) Del(ctx context.Context, key string) error {
	if ctx.Err() != nil {
		s.logger.Debug("Canceled action DEL.")
		return ctx.Err()
	}

	s.logger.Debug(fmt.Sprintf("Action: DEL | Key: %s", key))
	s.engine.Del(key)

	return nil
}
