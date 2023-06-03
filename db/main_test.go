package db

import "testing"

func TestConnect(t *testing.T) {
	_, err := Connect()
	if err != nil {
		t.Error(err)
	}
}
