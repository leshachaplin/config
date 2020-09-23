package main

import (
	"fmt"
	"github.com/leshachaplin/url-parser/models"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestKafkaResolve(t *testing.T) {
	const kafkaUrlExample = "kafka://topic:group@hostname.with.domain:9092/"
	result, err := Resolve(kafkaUrlExample)
	if err != nil {
		log.Error(err)
	}
	fmt.Println( result.(models.Kafka))
}
