package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	err := os.WriteFile(".env", []byte("TEST_KEY=TEST_VALUE"), 0644)
	if err != nil {
		return
	}
	defer os.Remove(".env")
	value := GetEnv("TEST_KEY")
	assert.Equal(t, "TEST_VALUE", value)
}
