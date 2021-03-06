package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSMTPAws(t *testing.T) {
	url := "aws-ssm-smtp://SMTP/"
	s, err := NewForAws("us-west-2")
	if err != nil {
		t.Error(err)
	}
	actual, err := s.GetSMTP(url)
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
