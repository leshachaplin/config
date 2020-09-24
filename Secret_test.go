package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestSecret(t *testing.T) {
	err := Init()
	if err != nil {
		log.Error(err)
	}
	url, ok := os.LookupEnv("secret")
	if !ok {
		log.Error("key not found")
	}
	result, err := GetSecret(url)
	if err == nil {
		log.Println(result)
	} else {
		want, err := json.Marshal(Secret{})
		t.Errorf("parsing env failed, got:%s  , want:%s", err, want)
	}
}
