package main

import (
	"fmt"
	"github.com/leshachaplin/url-parser/models"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestRedisResolve(t *testing.T) {
	const redisUrlExample = "redis://username:password@hostname.with.domain:6379/"
	result, err := Resolve(redisUrlExample)
	if err != nil {
		log.Error(err)
	}
	fmt.Println( result.(models.Redis))
}
