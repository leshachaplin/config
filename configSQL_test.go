package main

import (
	"fmt"
	"github.com/leshachaplin/url-parser/models"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestSQLResolve(t *testing.T) {
	const sqlUrlExample = "sql://lesha:su@localhost:5432/postgres"
	result, err := Resolve(sqlUrlExample)
	if err != nil {
		log.Error(err)
	}
	fmt.Println( result.(models.Sql))
}
