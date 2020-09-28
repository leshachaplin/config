package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedisResolve(t *testing.T) {
	url := "aws-ssm-redis://Redis/"
	s, err := New("us-west-2")
	if err != nil {
		t.Error(err)
	}
	actual, err := s.GetRedis(url)
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
