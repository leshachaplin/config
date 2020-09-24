package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestGRPC(t *testing.T) {
	err := Init()
	if err != nil {
		log.Error(err)
	}
	url, ok := os.LookupEnv("grpc")
	if !ok {
		log.Error("key not found")
	}
	result, err := GetGRPC(url)
	if err == nil {
		log.Println(result)
	} else {
		want, err := json.Marshal(GRPC{})
		t.Errorf("parsing env failed, got:%s  , want:%s", err, want)
	}
}
