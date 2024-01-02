package compute

import (
	"context"
	"fmt"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewComputer(t *testing.T) {
	logger := slog.Default()
	parser, _ := NewParser(logger)

	t.Run("error when parser and logger are nil", func(t *testing.T) {
		c, err := NewComputer(nil, nil)

		require.Nil(t, c)
		require.Error(t, err, "there is parser is nil")
	})

	t.Run("error when logger are nil", func(t *testing.T) {
		c, err := NewComputer(parser, nil)

		require.Nil(t, c)
		require.Error(t, err, "there is logger is nil")
	})

	t.Run("success computer creation", func(t *testing.T) {
		c, err := NewComputer(parser, logger)

		require.NotNil(t, c)
		require.NoError(t, err)
	})
}

func TestComputeQuery(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name        string
		rawQuery    string
		parsedQuery Query
		err         error
	}{
		{
			name:        "GET with key return success",
			rawQuery:    "GET key",
			parsedQuery: NewQuery("GET", []string{"key"}),
			err:         nil,
		},
		{
			name:        "GET without key return error",
			rawQuery:    "GET",
			parsedQuery: Query{},
			err:         fmt.Errorf("invalid number command argument"),
		},
		{
			name:        "SET with key and value return success",
			rawQuery:    "SET key value",
			parsedQuery: NewQuery("SET", []string{"key", "value"}),
			err:         nil,
		},
		{
			name:        "SET with one arguments return error",
			rawQuery:    "SET",
			parsedQuery: Query{},
			err:         fmt.Errorf("invalid number command argument"),
		},
		{
			name:        "SET without arguments return error",
			rawQuery:    "SET",
			parsedQuery: Query{},
			err:         fmt.Errorf("invalid number command argument"),
		},
		{
			name:        "DEL with key return success",
			rawQuery:    "DEL key",
			parsedQuery: NewQuery("DEL", []string{"key"}),
			err:         nil,
		},
		{
			name:        "DEL without key return error",
			rawQuery:    "DEL",
			parsedQuery: Query{},
			err:         fmt.Errorf("invalid number command argument"),
		},
		{
			name:        "Invalid command return error",
			rawQuery:    "INVALID key value",
			parsedQuery: Query{},
			err:         fmt.Errorf("invalid command name"),
		},
	}

	for _, tCase := range cases {
		p, _ := NewParser(slog.Default())
		c, _ := NewComputer(p, slog.Default())
		tt := tCase

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			q, err := c.Compute(context.Background(), tt.rawQuery)

			require.Equal(t, tt.parsedQuery, q)
			require.Equal(t, tt.err, err)
		})
	}
}
