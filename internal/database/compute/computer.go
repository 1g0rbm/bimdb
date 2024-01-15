package compute

import (
	"context"
	"fmt"
	"log/slog"
)

type Computer struct {
	parser *Parser
	logger *slog.Logger
}

func NewComputer(p *Parser, l *slog.Logger) (*Computer, error) {
	if p == nil {
		return nil, fmt.Errorf("there is invalid parser")
	}

	if l == nil {
		return nil, fmt.Errorf("there is invalid logger")
	}

	return &Computer{
		parser: p,
		logger: l,
	}, nil
}

func (c *Computer) Compute(_ context.Context, query string) (Query, error) {
	tokens, err := c.parser.Parse(query)
	if err != nil {
		return Query{}, err
	}

	q, err := AnalyzeQuery(tokens)
	if err != nil {
		return Query{}, err
	}

	return q, nil
}
