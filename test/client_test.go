package test

import (
	"blockchain/core"
	"testing"
)

func TestClient(t *testing.T) {
	client := core.NewClient("tester")
	client.Run()
}
