package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"

	"bimdb/internal/database"
	"bimdb/internal/database/compute"
	"bimdb/internal/database/storage"
	"bimdb/internal/database/storage/engine/in_memory"
)

func main() {
	logger := slog.Default()

	parser, err := compute.NewParser(logger)
	if err != nil {
		log.Fatalln(err)
	}

	storage, err := storage.NewStorage(in_memory.NewEngine(), logger)
	if err != nil {
		log.Fatalln(err)
	}

	computer, err := compute.NewComputer(parser, logger)
	if err != nil {
		log.Fatalln(err)
	}

	db, err := database.NewDB(computer, storage, logger)
	if err != nil {
		log.Fatalln(err)
	}

	logger.Info("database started")
	for {
		fmt.Print("> ")

		reader := bufio.NewReader(os.Stdin)

		text, err := reader.ReadString('\n')
		if err != nil {
			logger.Error(err.Error())
			continue
		}

		if text == "exit\n" {
			os.Exit(0)
		}

		res := db.Handle(context.Background(), strings.Trim(text, "\n "))
		fmt.Printf("-> %s\n", res)
	}
}
