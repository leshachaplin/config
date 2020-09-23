package main

import (
	"github.com/leshachaplin/url-parser/models"
	"net/url"
	"strings"
)

func Resolve(sqlUrl string) (interface{}, error) {
	data, err := url.Parse(sqlUrl)
	if err != nil {
		return nil, err
	}

	switch data.Scheme {
	case "sql":
		{
			res := models.Sql{}
			res.Username = data.User.Username()
			res.Password, _ = data.User.Password()
			res.Schema = strings.Trim(data.Path, "/")
			res.DBname = data.Host
			return res, nil
		}
	case "kafka":
		{
			group, _ := data.User.Password()
			res := models.Kafka{
				Host:  data.Host,
				Topic: data.User.Username(),
				Group: group,
			}
			return res, nil
		}
	case "redis":
		{
			res := models.Redis{Host: data.Host}
			return res, nil
		}
	case "smtp":
		{
			password, _ := data.User.Password()
			res := models.SMTP{
				Username: data.User.Username(),
				Password: password,
				Host:     data.Host,
				Email:    strings.Trim(data.Path, "/"),
			}
			return res, nil
		}
	}
	return nil, err
}
