package compute

import (
	"fmt"
	"log/slog"
)

type Parser struct {
	logger  *slog.Logger
	machine *stateMachine
}

func NewParser(l *slog.Logger) (*Parser, error) {
	if l == nil {
		return nil, fmt.Errorf("invalid logger")
	}

	return &Parser{
		logger:  l,
		machine: newStateMachine(),
	}, nil
}

func (p *Parser) Parse(query string) ([]string, error) {
	tokens, err := p.machine.parse(query)
	if err != nil {
		return nil, err
	}

	p.logger.Debug("parsed query", slog.Any("query", query), slog.Any("tokens", tokens))

	return tokens, nil
}
