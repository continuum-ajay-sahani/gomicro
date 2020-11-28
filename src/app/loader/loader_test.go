package loader

import (
	"errors"
	"practice/gomicro/src/config"
	"practice/gomicro/src/logger"

	"testing"
)

func TestLoadApplicationServices(t *testing.T) {

	t.Run("config load error", func(t *testing.T) {
		old := config.Load
		config.Load = func() error {
			return errors.New("config load error")
		}
		defer func() {
			config.Load = old
		}()

		err := LoadApplicationServices()
		if err == nil {
			t.Fatal("expected error while loading config, found nil")
		}
	})

	t.Run("log load error", func(t *testing.T) {
		oldconfig := config.Load
		config.Load = func() error {
			return nil
		}
		old := config.Load
		logger.Load = func() error {
			return errors.New("log load error")
		}
		defer func() {
			config.Load = old
			config.Load = oldconfig
		}()

		err := LoadApplicationServices()
		if err == nil {
			t.Fatal("expected error loading log, found nil")
		}
	})

	t.Run("app load success", func(t *testing.T) {
		oldconfig := config.Load
		config.Load = func() error {
			return nil
		}
		old := config.Load
		logger.Load = func() error {
			return nil
		}
		defer func() {
			config.Load = old
			config.Load = oldconfig
		}()

		err := LoadApplicationServices()
		if err != nil {
			t.Fatalf("expected success in loading app, found %v", err)
		}
	})
}
