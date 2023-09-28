package redis

import (
	"gotest.tools/assert"
	"testing"
)

func TestNewClient(t *testing.T) {
	addr := "127.0.0.1:30379"
	_, err := NewClient(addr, 1, "")
	assert.NilError(t, err)
}

func TestNewCluster(t *testing.T) {
	addrs := []string{"127.0.0.1:30379"}
	_, err := NewCluster(addrs)
	assert.NilError(t, err)
}
