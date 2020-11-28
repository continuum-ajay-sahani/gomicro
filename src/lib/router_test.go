package lib

import "testing"

func TestGetRouter(t *testing.T) {

	t.Run("router object check", func(t *testing.T) {
		router := GetReqRouter()
		if router == nil {
			t.Fatal("req router object can not be nil")
		}
	})
}
