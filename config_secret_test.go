package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSecret(t *testing.T) {
	url := "secret://lesha:chaplin@localhost:1234/"
	actual, err := GetSecret(url)
	if err != nil {
		t.Error(err)
	}
	expected := Secret{
		Username: "lesha",
		ApiKey:   "chaplin",
		ApiHost:  "localhost",
		ApiPort:  "1234",
	}
	assert.Equal(t, expected, *actual)
}
