package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestSQL(t *testing.T) {
	err := Init()
	if err != nil {
		log.Error(err)
	}
	url, ok := os.LookupEnv("sql")
	if !ok {
		log.Error("key not found")
	}
	result, err := GetSQL(url)
	if err == nil {
		log.Println(result)
	} else {
		want, err := json.Marshal(Sql{})
		t.Errorf("parsing env failed, got:%s  , want:%s", err, want)
	}
}
