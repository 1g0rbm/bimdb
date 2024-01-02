package compute

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuery(t *testing.T) {
	command := "COMMAND"
	arguments := []string{"arg1", "arg2"}

	q := NewQuery(command, arguments)

	require.Equal(t, command, q.Command())
	require.Equal(t, arguments, q.Arguments())
}
