package database

import (
	"context"
	"fmt"

	"log/slog"

	"bimdb/internal/database/compute"
)

type ComputeLayerInterface interface {
	Compute(ctx context.Context, query string) (compute.Query, error)
}

type StorageLayerInterface interface {
	Set(ctx context.Context, key string, value string) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
}

type DB struct {
	computer ComputeLayerInterface
	storage  StorageLayerInterface
	logger   *slog.Logger
}

func NewDB(c ComputeLayerInterface, s StorageLayerInterface, l *slog.Logger) (*DB, error) {
	if c == nil {
		return nil, fmt.Errorf("there is invalid compute layer value")
	}

	if s == nil {
		return nil, fmt.Errorf("there is invalid storage layer value")
	}

	if l == nil {
		return nil, fmt.Errorf("there is invalid logger value")
	}

	return &DB{
		computer: c,
		storage:  s,
		logger:   l,
	}, nil
}

func (db *DB) Handle(ctx context.Context, rawQuery string) string {
	db.logger.Debug("handling query", slog.String("query", rawQuery))

	query, err := db.computer.Compute(ctx, rawQuery)
	if err != nil {
		db.logger.Error("compute layer error", slog.String("message", err.Error()))
		return fmt.Sprintf("[error] %s", err.Error())
	}

	switch query.Command() {
	case compute.SetCommand:
		return db.handleSet(ctx, query)
	case compute.GetCommand:
		return db.handleGet(ctx, query)
	case compute.DelCommand:
		return db.handleDel(ctx, query)
	default:
		db.logger.Error(
			"incorrect compute layer configuration",
			slog.String("message", fmt.Sprintf("invalid command %s", query.Command())),
		)

		return "[error] internal error"
	}
}

func (db *DB) handleSet(ctx context.Context, query compute.Query) string {
	arguments := query.Arguments()
	if err := db.storage.Set(ctx, arguments[0], arguments[1]); err != nil {
		return fmt.Sprintf("[error] %s", err.Error())
	}

	return "[ok]"
}

func (db *DB) handleGet(ctx context.Context, query compute.Query) string {
	arguments := query.Arguments()
	value, err := db.storage.Get(ctx, arguments[0])
	if err != nil {
		return fmt.Sprintf("[error] %s", err.Error())
	}

	return fmt.Sprintf("[ok] %s", value)
}

func (db *DB) handleDel(ctx context.Context, query compute.Query) string {
	arguments := query.Arguments()
	if err := db.storage.Del(ctx, arguments[0]); err != nil {
		return fmt.Sprintf("[error] %s", err.Error())
	}

	return "[ok]"
}
