package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestKafkaResolve(t *testing.T) {
	err := Init()
	if err != nil {
		log.Error(err)
	}
	url, ok := os.LookupEnv("kafka")
	if !ok {
		log.Error("key not found")
	}
	result, err := GetKafka(url)
	if err == nil {
		log.Println(result)
	} else {
		want, err := json.Marshal(Kafka{})
		t.Errorf("parsing env failed, got:%s  , want:%s", err, want)
	}
}
