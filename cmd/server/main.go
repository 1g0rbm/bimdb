package main

import (
	"bimdb/internal/config"
	"bimdb/internal/provider"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var configFileName = os.Getenv("CONFIG_FILE_NAME")

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := &config.Config{}
	if configFileName != "" {
		var err error
		cfg, err = config.Load(configFileName)
		if err != nil {
			log.Fatal(err)
		}
	}

	dic, err := provider.NewProvider(cfg)
	if err != nil {
		log.Fatal(err)
	}

	db, err := dic.GetDB()
	if err != nil {
		log.Fatal(err)
	}

	server, err := dic.GetServer()
	if err != nil {
		log.Fatal(err)
	}

	err = server.Handle(ctx, func(ctx context.Context, bytes []byte) []byte {
		return []byte(db.Handle(ctx, string(bytes)))
	})
	if err != nil {
		log.Fatal(err)
	}
}
