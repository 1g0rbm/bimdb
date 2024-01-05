package network

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"log/slog"
)

type TCPHandler = func(context.Context, []byte) []byte

type TCPServer struct {
	address          string
	idleTimeout      time.Duration
	messageSize      int
	maxConnectionsCh chan struct{}
	logger           *slog.Logger
}

func NewTCPServer(
	address string,
	idleTimeout time.Duration,
	messageSize int,
	maxConnections int,
	logger *slog.Logger,
) (*TCPServer, error) {
	if maxConnections <= 0 {
		return nil, fmt.Errorf("max connections number should be more than 0")
	}

	if logger == nil {
		return nil, fmt.Errorf("invalid logger value")
	}

	return &TCPServer{
		address:          address,
		idleTimeout:      idleTimeout,
		messageSize:      messageSize,
		maxConnectionsCh: make(chan struct{}, maxConnections),
		logger:           logger,
	}, nil
}

func (s *TCPServer) Handle(ctx context.Context, handler TCPHandler) error {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			connection, err := listener.Accept()
			if err != nil {
				if errors.Is(err, net.ErrClosed) {
					return
				}

				s.logger.Error("failed to accept connection", slog.String("err", err.Error()))
				continue
			}

			s.maxConnectionsCh <- struct{}{}
			go func(connection net.Conn) {
				wg.Add(1)
				defer func() {
					wg.Done()
					<-s.maxConnectionsCh
				}()

				s.handleConnection(ctx, connection, handler)
			}(connection)
		}
	}()

	wg.Wait()
	close(s.maxConnectionsCh)

	return nil
}

func (s *TCPServer) handleConnection(ctx context.Context, connection net.Conn, handler TCPHandler) {
	request := make([]byte, s.messageSize)

	for {
		if err := connection.SetDeadline(time.Now().Add(s.idleTimeout)); err != nil {
			s.logger.Warn("failed to set read deadline", slog.String("err", err.Error()))
		}

		count, err := connection.Read(request)
		if err != nil {
			if err != io.EOF {
				s.logger.Warn("failed to read", slog.String("err", err.Error()))
			}

			break
		}

		resp := handler(ctx, request[:count])
		if _, err := connection.Write(resp); err != nil {
			s.logger.Warn("failed to write response", slog.String("err", err.Error()))
			break
		}
	}

	if err := connection.Close(); err != nil {
		s.logger.Warn("failed to close connection", slog.String("err", err.Error()))
	}
}
