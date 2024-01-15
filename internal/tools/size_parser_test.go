package tools

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseSizeNegative(t *testing.T) {
	t.Run("invalid units", func(t *testing.T) {
		res, err := ParseSize("10BB")
		require.Error(t, err, "invalid units")
		require.Equal(t, 0, res)
	})

	t.Run("empty string", func(t *testing.T) {
		res, err := ParseSize("")
		require.Error(t, err, "invalid input size format")
		require.Equal(t, 0, res)
	})

	t.Run("wrong string", func(t *testing.T) {
		res, err := ParseSize("")
		require.Error(t, err, "invalid input size format")
		require.Equal(t, 0, res)
	})
}

func TestParseSizeSuccess(t *testing.T) {
	t.Run("parse bytes", func(t *testing.T) {
		res, err := ParseSize("10B")
		require.NoError(t, err)
		require.Equal(t, 10, res)

		res, err = ParseSize("10b")
		require.NoError(t, err)
		require.Equal(t, 10, res)

		res, err = ParseSize("10")
		require.NoError(t, err)
		require.Equal(t, 10, res)
	})

	t.Run("parse kilobytes", func(t *testing.T) {
		res, err := ParseSize("10KB")
		require.NoError(t, err)
		require.Equal(t, 10000, res)

		res, err = ParseSize("10Kb")
		require.NoError(t, err)
		require.Equal(t, 10000, res)

		res, err = ParseSize("10kb")
		require.NoError(t, err)
		require.Equal(t, 10000, res)
	})

	t.Run("parse megabytes", func(t *testing.T) {
		res, err := ParseSize("10MB")
		require.NoError(t, err)
		require.Equal(t, 10000000, res)

		res, err = ParseSize("10Mb")
		require.NoError(t, err)
		require.Equal(t, 10000000, res)

		res, err = ParseSize("10mb")
		require.NoError(t, err)
		require.Equal(t, 10000000, res)
	})

	t.Run("parse gigabytes", func(t *testing.T) {
		res, err := ParseSize("10GB")
		require.NoError(t, err)
		require.Equal(t, 10000000000, res)

		res, err = ParseSize("10Gb")
		require.NoError(t, err)
		require.Equal(t, 10000000000, res)

		res, err = ParseSize("10gb")
		require.NoError(t, err)
		require.Equal(t, 10000000000, res)
	})
}
