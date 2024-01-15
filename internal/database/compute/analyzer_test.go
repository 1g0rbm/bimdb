package compute

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAnalyzeSuccess(t *testing.T) {
	t.Run("GET", func(t *testing.T) {
		tokens := []string{"GET", "arg1"}

		q, err := AnalyzeQuery(tokens)

		require.NoError(t, err)
		require.Equal(t, tokens[0], q.Command())
		require.Equal(t, tokens[1:], q.Arguments())
	})

	t.Run("SET", func(t *testing.T) {
		tokens := []string{"SET", "arg1", "arg2"}

		q, err := AnalyzeQuery(tokens)

		require.NoError(t, err)
		require.Equal(t, tokens[0], q.Command())
		require.Equal(t, tokens[1:], q.Arguments())
	})

	t.Run("DEL", func(t *testing.T) {
		tokens := []string{"DEL", "arg1"}

		q, err := AnalyzeQuery(tokens)

		require.NoError(t, err)
		require.Equal(t, tokens[0], q.Command())
		require.Equal(t, tokens[1:], q.Arguments())
	})
}

func TestAnalyzeNegative(t *testing.T) {
	t.Run("Invalid command", func(t *testing.T) {
		var expectedArguments []string
		var expectedCommand string

		tokens := []string{"WRONG", "arg1"}

		q, err := AnalyzeQuery(tokens)

		require.Error(t, err)
		require.Equal(t, expectedCommand, q.Command())
		require.Equal(t, expectedArguments, q.Arguments())
	})
}
