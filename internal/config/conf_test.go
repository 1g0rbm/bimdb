package config

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestLoadNotExistFile(t *testing.T) {
	t.Parallel()

	cfg, err := Load("./test_data/non_exist_conf.yaml")
	require.Error(t, err)
	require.Nil(t, cfg)
}

func TestLoadEmptyConfig(t *testing.T) {
	cfg, err := Load("./test_data/non_exist_conf.yaml")
	require.Error(t, err)
	require.Nil(t, cfg)
}

func TestLoadEmptyFilenameConfig(t *testing.T) {
	cfg, err := Load("")
	require.NoError(t, err)
	require.NotNil(t, cfg)
}

func TestLoadConfig(t *testing.T) {
	cfg, err := Load("./test_data/conf.yaml")
	require.NoError(t, err)

	require.Equal(t, "127.0.0.1:3223", cfg.NetworkConfig.Address)
	require.Equal(t, 100, cfg.NetworkConfig.MaxConnections)
	require.Equal(t, "4KB", cfg.NetworkConfig.MaxMessageSize)
	require.Equal(t, time.Minute*5, cfg.NetworkConfig.IdleTimeout)
}
