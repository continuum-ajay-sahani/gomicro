package model

import "testing"

func TestAppHealth(t *testing.T) {
	health := AppHealth()
	if health.MemoryCache != true {
		t.Fatal("memory cache should be true")
	}
}
