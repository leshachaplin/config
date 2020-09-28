package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMongo(t *testing.T) {
	url := "aws-ssm-mongo://Mongo/"
	s, err := New("us-west-2")
	if err != nil {
		t.Error(err)
	}
	actual, err := s.GetMongo(url)
	if err != nil {
		t.Error(err)
	}
	expected := MongoDB{ConnectionString: "mongodb://localhost:27017"}
	assert.Equal(t, expected, *actual)
}
