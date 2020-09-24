package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedisResolve(t *testing.T) {
	url := "redis://lesha:chaplin@localhost:6379/"
	actual, err := GetRedis(url)
	if err != nil {
		t.Error(err)
	}
	expected := Redis{
		Username: "lesha",
		Password: "chaplin",
		Host:     "localhost",
		Port:     "6379",
	}
	assert.Equal(t, expected, *actual)
}
