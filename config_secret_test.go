package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSecret(t *testing.T) {
	url := "aws-ssm-secret://Secret/"
	s, err := New("us-west-2")
	if err != nil {
		t.Error(err)
	}
	actual, err := s.GetSecret(url)
	if err != nil {
		t.Error(err)
	}
	expected := Secret{
		ApiKey:   "chaplin",
	}
	assert.Equal(t, expected, *actual)
}
