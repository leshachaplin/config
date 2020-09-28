package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSQLAws(t *testing.T) {
	url := "aws-ssm-sql://SQL/"
	s, err := NewForAws("us-west-2")
	if err != nil {
		t.Error(err)
	}
	actual, err := s.GetSQL(url)
	if err != nil {
		t.Error(err)
	}
	expected := SQL{
		Username: "lesha",
		Password: "su",
		DB:       "books",
		Host:     "localhost",
		Port:     5432,
		Schema:   "postgres",
	}
	assert.Equal(t, expected, *actual)
}
