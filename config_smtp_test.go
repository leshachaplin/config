package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSMTPResolve(t *testing.T) {
	url := "smtp://lesha.chaplin@gmail.com:su@smtp.gmail.com:443/lesha.chaplin@gmail.com/disable"
	actual, err := GetSMTP(url)
	if err != nil {
		t.Error(err)
	}
	expected := SMTP{
		Username: "lesha.chaplin@gmail.com",
		Password: "su",
		Host:     "smtp.gmail.com",
		Port:     443,
		SSL:      "disable",
		Email:    "lesha.chaplin@gmail.com",
	}
	assert.Equal(t, expected, *actual)
}
