package database

import "context"

type computeInterface interface {
	Handle(ctx context.Context, query string)
}

type DB struct {
}
