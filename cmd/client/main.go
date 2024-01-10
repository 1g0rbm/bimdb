package main

import (
	"bimdb/internal/network"
	"bimdb/internal/tools"
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"syscall"
	"time"
)

func main() {
	address := flag.String("address", "localhost:3223", "Address of the bimdb")
	idleTimeout := flag.Duration("idle_timeout", time.Minute, "Idle timeout for connection")
	maxMessageSizeStr := flag.String("max_message_size", "4KB", "Max message size for connection")
	flag.Parse()

	logger := slog.Default()

	maxMessageSize, err := tools.ParseSize(*maxMessageSizeStr)
	if err != nil {
		logger.Error("Failed to parse max message size", slog.String("err", err.Error()))
	}

	reader := bufio.NewReader(os.Stdin)
	client, err := network.NewTCPClient(*address, maxMessageSize, *idleTimeout)
	if err != nil {
		logger.Error("Failed to connect with server", slog.String("err", err.Error()))
	}

	for {
		fmt.Print("[bimdb]> ")
		request, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, syscall.EPIPE) {
				logger.Error("connection was closed", slog.String("err", err.Error()))
			}

			logger.Error("failed to read query", slog.String("err", err.Error()))
		}

		response, err := client.Send([]byte(request))
		if err != nil {
			if errors.Is(err, syscall.EPIPE) {
				logger.Error("connection was closed", slog.String("err", err.Error()))
			}

			logger.Error("failed to send query", slog.String("err", err.Error()))
		}

		fmt.Println(string(response))
	}
}
