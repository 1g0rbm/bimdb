package provider

import (
	"log/slog"
	"os"
)

const defaultFilePath = "output.log"

func (p *Provider) getLogger() (*slog.Logger, error) {
	if p.logger == nil {
		filePath := defaultFilePath

		cfg := p.config.LoggerConfig
		if cfg != nil {
			if cfg.FilePath != "" {
				filePath = cfg.FilePath
			}
		}

		f, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}

		p.logger = slog.New(slog.NewJSONHandler(f, nil))
	}

	return p.logger, nil
}
