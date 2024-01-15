package storage

import (
	"context"
	"log/slog"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"bimdb/internal/database/storage/engine"
	"bimdb/internal/database/storage/engine/in_memory"
)

func TestNewStorage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEngine := engine.NewMockIEngine(ctrl)

	s, err := NewStorage(nil, nil)
	require.Error(t, err, "invalid engine")
	require.Nil(t, s)

	s, err = NewStorage(mockEngine, nil)
	require.Error(t, err, "invalid logger")
	require.Nil(t, s)

	s, err = NewStorage(in_memory.NewEngine(), slog.Default())
	require.Nil(t, err)
	require.NotNil(t, s)
}

func TestSuccessSet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockEngine := engine.NewMockIEngine(ctrl)
	mockEngine.EXPECT().Set("key", "value").Return()

	s, _ := NewStorage(mockEngine, slog.Default())

	err := s.Set(ctx, "key", "value")

	require.Nil(t, err)
}

func TestContextCanceledSet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	mockEngine := engine.NewMockIEngine(ctrl)

	s, _ := NewStorage(mockEngine, slog.Default())

	err := s.Set(ctx, "key", "value")

	require.ErrorContains(t, err, "context canceled")
}

func TestSuccessGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockEngine := engine.NewMockIEngine(ctrl)
	mockEngine.EXPECT().Get("key").Return("value", true)

	s, _ := NewStorage(mockEngine, slog.Default())

	val, err := s.Get(ctx, "key")
	require.Nil(t, err)
	require.Equal(t, val, "value")
}

func TestContextCanceledGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	mockEngine := engine.NewMockIEngine(ctrl)

	s, _ := NewStorage(mockEngine, slog.Default())

	val, err := s.Get(ctx, "key")
	require.ErrorContains(t, err, "context canceled")
	require.Equal(t, "", val)
}

func TestNotFoundKeyGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockEngine := engine.NewMockIEngine(ctrl)
	mockEngine.EXPECT().Get("key").Return("", false)

	s, _ := NewStorage(mockEngine, slog.Default())

	val, err := s.Get(ctx, "key")
	require.ErrorContains(t, err, "key key not found")
	require.Equal(t, "", val)
}

func TestSuccessDel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockEngine := engine.NewMockIEngine(ctrl)
	mockEngine.EXPECT().Del("key").Return()

	s, _ := NewStorage(mockEngine, slog.Default())

	err := s.Del(ctx, "key")
	require.Nil(t, err)
}

func TestContextCanceledDel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	mockEngine := engine.NewMockIEngine(ctrl)

	s, _ := NewStorage(mockEngine, slog.Default())

	err := s.Del(ctx, "key")
	require.ErrorContains(t, err, "context canceled")
}
