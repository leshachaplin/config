package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKafkaResolve(t *testing.T) {
	url := "kafka://library:books@localhost:9092/"
	actual, err := GetKafka(url)
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
