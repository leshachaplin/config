package main

import (
	"errors"
	"net/url"
	"strings"
)

// convert result from resolve to sql structure type
func GetSQL(sqlUrl string) (*Sql, error) {
	res, err := resolve(sqlUrl)
	sql, ok := res.(Sql)
	if !ok {
		return nil, errors.New("your should use other method to parse this url or use other url")
	}
	return &sql, err
}

// convert result from resolve to secretApiKey structure type
func GetSecret(secretUrl string) (*Secret, error) {
	res, err := resolve(secretUrl)
	secret, ok := res.(Secret)
	if !ok {
		return nil, errors.New("your should use other method to parse this url or use other url")
	}
	return &secret, err
}

// convert result from resolve to mongo structure type
func GetMongo(mongoUrl string) (*MongoDB, error) {
	res, err := resolve(mongoUrl)
	mongo, ok := res.(MongoDB)
	if !ok {
		return nil, errors.New("your should use other method to parse this url or use other url")
	}
	return &mongo, err
}

// convert result from resolve to smtp structure type
func GetSMTP(smtpUrl string) (*SMTP, error) {
	res, err := resolve(smtpUrl)
	smtp, ok := res.(SMTP)
	if !ok {
		return nil, errors.New("your should use other method to parse this url or use other url")
	}
	return &smtp, err
}

// convert result from resolve to kafka structure type
func GetKafka(kafkaUrl string) (*Kafka, error) {
	res, err := resolve(kafkaUrl)
	kafka, ok := res.(Kafka)
	if !ok {
		return nil, errors.New("your should use other method to parse this url or use other url")
	}
	return &kafka, err
}

// convert result from resolve to redis structure type
func GetRedis(redisUrl string) (*Redis, error) {
	res, err := resolve(redisUrl)
	redis, ok := res.(Redis)
	if !ok {
		return nil, errors.New("your should use other method to parse this url or use other url")
	}
	return &redis, err
}

//parse connection url in to a specific structure that is indicated in the url scheme
func resolve(sqlUrl string) (interface{}, error) {
	data, err := url.Parse(sqlUrl)
	if err != nil {
		return nil, err
	}

	switch data.Scheme {
	case "sql":
		{
			res := Sql{}
			path := strings.Split(data.Path, "/")[1:]
			res.Username = data.User.Username()
			res.Password, _ = data.User.Password()
			res.Schema = path[0]
			res.Host = strings.Split(data.Host, ":")[0]
			res.Port = strings.Split(data.Host, ":")[1]
			res.DB = path[1]
			return res, nil
		}
	case "kafka":
		{
			group, _ := data.User.Password()
			res := Kafka{}
			res.Topic = data.User.Username()
			res.Host = strings.Split(data.Host, ":")[0]
			res.Port = strings.Split(data.Host, ":")[1]
			res.Group = group
			return res, nil
		}
	case "redis":
		{
			res := Redis{}
			res.Username = data.User.Username()
			res.Password, _ = data.User.Password()
			res.Host = strings.Split(data.Host, ":")[0]
			res.Port = strings.Split(data.Host, ":")[1]
			return res, nil
		}
	case "smtp":
		{
			res := SMTP{}
			res.Username = data.User.Username()
			res.Password, _ = data.User.Password()
			path := strings.Split(data.Path, "/")[1:]
			res.Host = strings.Split(data.Host, ":")[0]
			res.Port = strings.Split(data.Host, ":")[1]
			res.Email = path[0]
			res.SSL = path[1]

			return res, nil
		}
	case "secret":
		{
			key, _ := data.User.Password()
			res := Secret{}
			res.Username = data.User.Username()
			res.ApiKey = key
			res.ApiHost = strings.Split(data.Host, ":")[0]
			res.ApiPort = strings.Split(data.Host, ":")[1]

			return res, nil
		}
	case "mongo":
		{
			res := MongoDB{}
			res.ConnectionString = "mongodb://" + data.Host

			return res, nil
		}
	}
	return nil, err
}
