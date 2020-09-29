package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedisResolve(t *testing.T) {
	url := "redis://lesha:chaplin@localhost:6379/?isCluster=true"
	actual, err := GetRedis(url)
	if err != nil {
		t.Error(err)
	}
	expected := Redis{
		Username:  "lesha",
		Password:  "chaplin",
		Host:      "localhost",
		Port:      6379,
		IsCluster: true,
	}
	assert.Equal(t, expected, *actual)
}
