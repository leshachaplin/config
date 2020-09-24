package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestRedisResolve(t *testing.T) {
	err := Init()
	if err != nil {
		log.Error(err)
	}
	url, ok := os.LookupEnv("redis")
	if !ok {
		log.Error("key not found")
	}
	result, err := GetRedis(url)
	if err == nil {
		log.Println(result)
	} else {
		want, err := json.Marshal(Redis{})
		t.Errorf("parsing env failed, got:%s  , want:%s", err, want)
	}
}
