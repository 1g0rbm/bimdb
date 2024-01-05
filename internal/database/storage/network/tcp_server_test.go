package network

import (
	"context"
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"log/slog"
)

func TestTCPServer_Handle(t *testing.T) {
	t.Parallel()

	request := "Hello, server!"
	response := "Hello, client!"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	maxMessageSize := 2048
	maxConnections := 10
	idleTimeout := time.Minute

	server, err := NewTCPServer(":20001", idleTimeout, maxMessageSize, maxConnections, slog.Default())
	require.NoError(t, err)

	go func() {
		require.NoError(t, server.Handle(ctx, func(ctx context.Context, buffer []byte) []byte {
			require.True(t, reflect.DeepEqual([]byte(request), buffer))
			return []byte(response)
		}))
	}()

	connection, err := net.Dial("tcp", "localhost:20001")
	require.NoError(t, err)

	_, err = connection.Write([]byte(request))
	require.NoError(t, err)

	buffer := make([]byte, maxMessageSize)
	count, err := connection.Read(buffer)
	require.NoError(t, err)
	require.True(t, reflect.DeepEqual([]byte(response), buffer[:count]))
}
