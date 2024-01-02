package database

import (
	"context"
	"fmt"
	"log/slog"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"bimdb/internal/database/compute"
)

func TestNewDB(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)

	computer := NewMockComputeLayerInterface(ctrl)
	storage := NewMockStorageLayerInterface(ctrl)

	t.Run("without arguments return error", func(t *testing.T) {
		db, err := NewDB(nil, nil, nil)

		require.Nil(t, db)
		require.Error(t, err, "there is invalid compute layer value")
	})

	t.Run("only with computer layer argument return error", func(t *testing.T) {
		db, err := NewDB(computer, nil, nil)

		require.Nil(t, db)
		require.Error(t, err, "there is invalid storage layer value")
	})

	t.Run("witout logger return error", func(t *testing.T) {
		db, err := NewDB(computer, storage, nil)

		require.Nil(t, db)
		require.Error(t, err, "there is invalid logger value")
	})

	t.Run("with all arguments return success", func(t *testing.T) {
		db, err := NewDB(computer, storage, slog.Default())

		require.NotNil(t, db)
		require.NoError(t, err)
	})
}

func TestHandleQuery(t *testing.T) {
	t.Parallel()

	t.Run("set success", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		command := "SET key value"

		computer := NewMockComputeLayerInterface(ctrl)
		computer.
			EXPECT().
			Compute(gomock.Any(), command).
			Return(compute.NewQuery("SET", []string{"key", "value"}), nil)

		storage := NewMockStorageLayerInterface(ctrl)
		storage.
			EXPECT().
			Set(context.TODO(), "key", "value").
			Return(nil)

		db, _ := NewDB(computer, storage, slog.Default())
		res := db.Handle(context.TODO(), command)

		require.Equal(t, "[ok]", res)
	})

	t.Run("get success", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		command := "GET key"

		computer := NewMockComputeLayerInterface(ctrl)
		computer.
			EXPECT().
			Compute(gomock.Any(), command).
			Return(compute.NewQuery("GET", []string{"key"}), nil)

		storage := NewMockStorageLayerInterface(ctrl)
		storage.
			EXPECT().
			Get(context.TODO(), "key").
			Return("value", nil)

		db, _ := NewDB(computer, storage, slog.Default())
		res := db.Handle(context.TODO(), command)

		require.Equal(t, "[ok] value", res)
	})

	t.Run("del success", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		command := "DEL key"

		computer := NewMockComputeLayerInterface(ctrl)
		computer.
			EXPECT().
			Compute(gomock.Any(), command).
			Return(compute.NewQuery("DEL", []string{"key"}), nil)

		storage := NewMockStorageLayerInterface(ctrl)
		storage.
			EXPECT().
			Del(context.TODO(), "key").
			Return(nil)

		db, _ := NewDB(computer, storage, slog.Default())
		res := db.Handle(context.TODO(), command)

		require.Equal(t, "[ok]", res)
	})

	t.Run("computer layer error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		command := "SET key"

		computer := NewMockComputeLayerInterface(ctrl)
		computer.
			EXPECT().
			Compute(gomock.Any(), command).
			Return(compute.Query{}, fmt.Errorf("some error"))

		storage := NewMockStorageLayerInterface(ctrl)

		db, _ := NewDB(computer, storage, slog.Default())
		res := db.Handle(context.TODO(), command)

		require.Equal(t, "[error] some error", res)
	})

	t.Run("storage layer error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		command := "DEL key"

		computer := NewMockComputeLayerInterface(ctrl)
		computer.
			EXPECT().
			Compute(gomock.Any(), command).
			Return(compute.NewQuery("DEL", []string{"key"}), nil)

		storage := NewMockStorageLayerInterface(ctrl)
		storage.
			EXPECT().
			Del(context.TODO(), "key").
			Return(fmt.Errorf("some error"))

		db, _ := NewDB(computer, storage, slog.Default())
		res := db.Handle(context.TODO(), command)

		require.Equal(t, "[error] some error", res)
	})
}
