package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestSMTPResolve(t *testing.T) {
	err := Init()
	if err != nil {
		log.Error(err)
	}
	url, ok := os.LookupEnv("smtp")
	if !ok {
		log.Error("key not found")
	}
	result, err := GetSMTP(url)
	if err == nil {
		log.Println(result)
	} else {
		want, err := json.Marshal(SMTP{})
		t.Errorf("parsing env failed, got:%s  , want:%s", err, want)
	}
}
