package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSecret(t *testing.T) {
	url := "secret://chaplin/"
	actual, err := GetSecret(url)
	if err != nil {
		t.Error(err)
	}
	expected := Secret{
		ApiKey:   "chaplin",
	}
	assert.Equal(t, expected, *actual)
}
