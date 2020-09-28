package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKafkaResolve(t *testing.T) {
	url := "aws-ssm-kafka://Kafka/"
	s, err := New("us-west-2")
	if err != nil {
		t.Error(err)
	}
	actual, err := s.GetKafka(url)
	if err != nil {
		t.Error(err)
	}
	expected := Kafka{
		Host:  "localhost",
		Topic: "library",
		Group: "books",
		Port:  9092,
	}
	assert.Equal(t, expected, *actual)
}
