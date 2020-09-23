package main

import (
	"fmt"
	"github.com/leshachaplin/url-parser/models"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestSMTPResolve(t *testing.T) {
	const smtpUrlExample = "smtp://lesha.chaplin@gmail.com:su@smtp.gmail.com/lesha.chaplin@gmail.com"
	result, err := Resolve(smtpUrlExample)
	if err != nil {
		log.Error(err)
	}
	fmt.Println( result.(models.SMTP))
}
