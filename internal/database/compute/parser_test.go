package compute

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewParser(t *testing.T) {
	p, err := NewParser(nil)
	require.Error(t, err, "invalid logger")
	require.Nil(t, p)

	p, err = NewParser(slog.Default())
	require.NoError(t, err, "invalid logger")
	require.NotNil(t, p)
}

func TestSuccessParse(t *testing.T) {
	t.Run("2 words", func(t *testing.T) {
		p, _ := NewParser(slog.Default())
		res, err := p.Parse("test test")

		require.NoError(t, err)
		require.Equal(t, []string{"test", "test"}, res)
	})

	t.Run("3 words with punctuation", func(t *testing.T) {
		p, _ := NewParser(slog.Default())
		res, err := p.Parse("test t_e_s_* /test/")

		require.NoError(t, err)
		require.Equal(t, []string{"test", "t_e_s_*", "/test/"}, res)
	})

	t.Run("call parse 2 times", func(t *testing.T) {
		p, _ := NewParser(slog.Default())
		res, err := p.Parse("test t_e_s_* /test/")

		require.NoError(t, err)
		require.Equal(t, []string{"test", "t_e_s_*", "/test/"}, res)

		res, err = p.Parse("test test")

		require.NoError(t, err)
		require.Equal(t, []string{"test", "test"}, res)
	})
}

func TestFailParse(t *testing.T) {
	p, _ := NewParser(slog.Default())

	res, err := p.Parse("test &uu")
	require.Error(t, err, "invalid char")
	require.Nil(t, res)
}
