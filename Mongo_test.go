package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestMongo(t *testing.T) {
	err := Init()
	if err != nil {
		log.Error(err)
	}
	url, ok := os.LookupEnv("mongo")
	if !ok {
		log.Error("key not found")
	}
	result, err := GetMongo(url)
	if err == nil {
		log.Println(result)
	} else {
		want, err := json.Marshal(MongoDB{})
		t.Errorf("parsing env failed, got:%s  , want:%s", err, want)
	}
}
