package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSQL(t *testing.T) {
	url := "sql://lesha:su@localhost:5432/postgres/books"
	actual, err := GetSQL(url)
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
