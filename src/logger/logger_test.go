package logger

import "testing"

func TestLoad(t *testing.T) {
	err := Load()
	if err != nil {
		t.Fatalf("%v", err)
	}
}
