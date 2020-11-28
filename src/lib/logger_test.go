package lib

import (
	"testing"
)

func TestGetLogger(t *testing.T) {

	t.Run("error check", func(t *testing.T) {
		_, err := GetLogger()
		if err != nil {
			t.Fatalf("unexpected error during log object creation = %v", err)
		}
	})

	t.Run("logger object check", func(t *testing.T) {
		logger, _ := GetLogger()
		if logger == nil {
			t.Fatal("logger object can not be nil")
		}
	})
}
