package config

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

// GetSQL converts result from resolve to sql structure type
func GetSQL(sqlURL string) (*SQL, error) {
	res, err := resolve(sqlURL)
	if err != nil {
		return nil, err
	}
	sql, ok := res.(SQL)
	if !ok {
		return nil, errors.New("your should use other method to parse this url or use other url")
	}
	return &sql, nil
}

// GetSecret converts result from resolve to secretApiKey structure type
func GetSecret(secretURL string) (*Secret, error) {
	res, err := resolve(secretURL)
	secret, ok := res.(Secret)
	if !ok {
		return nil, errors.New("your should use other method to parse this url or use other url")
	}
	return &secret, err
}

// GetMongo converts result from resolve to mongo structure type
func GetMongo(mongoURL string) (*MongoDB, error) {
	res, err := resolve(mongoURL)
	mongo, ok := res.(MongoDB)
	if !ok {
		return nil, errors.New("your should use other method to parse this url or use other url")
	}
	return &mongo, err
}

// GetSMTP converts result from resolve to smtp structure type
func GetSMTP(smtpURL string) (*SMTP, error) {
	res, err := resolve(smtpURL)
	if err != nil {
		return nil, err
	}
	smtp, ok := res.(SMTP)
	if !ok {
		return nil, errors.New("your should use other method to parse this url or use other url")
	}
	return &smtp, nil
}

// GetKafka converts result from resolve to kafka structure type
func GetKafka(kafkaURL string) (*Kafka, error) {
	res, err := resolve(kafkaURL)
	if err != nil {
		return nil, err
	}
	kafka, ok := res.(Kafka)
	if !ok {
		return nil, errors.New("your should use other method to parse this url or use other url")
	}
	return &kafka, nil
}

// GetRedis convert result from resolve to redis structure type
func GetRedis(redisURL string) (*Redis, error) {
	res, err := resolve(redisURL)
	if err != nil {
		return nil, err
	}
	redis, ok := res.(Redis)
	if !ok {
		return nil, errors.New("your should use other method to parse this url or use other url")
	}
	return &redis, nil
}

// resolve parse connection url in to a specific structure that is indicated in the url scheme
func resolve(URL string) (interface{}, error) {
	data, err := url.Parse(URL)
	if err != nil {
		return nil, err
	}

	switch data.Scheme {
	case "sql":
		{
			res := SQL{}
			path := strings.Split(data.Path, "/")[1:]
			if strings.ContainsAny(data.Host, ":") {
				res.Username = data.User.Username()
				res.Password, _ = data.User.Password()
				res.Schema = path[0]
				res.DB = path[1]
				res.Host = strings.Split(data.Host, ":")[0]
				res.Port, err = strconv.Atoi(strings.Split(data.Host, ":")[1])
				if err != nil {
					return nil, err
				}
			} else {
				err = errors.New("url doesn't contain port")
				return nil, err
			}
			return res, nil
		}
	case "kafka":
		{
			group, _ := data.User.Password()
			res := Kafka{}
			if strings.ContainsAny(data.Host, ":") {
				res.Topic = data.User.Username()
				res.Host = strings.Split(data.Host, ":")[0]
				res.Port, err = strconv.Atoi(strings.Split(data.Host, ":")[1])
				if err != nil {
					return nil, err
				}
				res.Group = group
			} else {
				err = errors.New("url doesn't contain port")
				return nil, err
			}
			return res, nil
		}
	case "redis":
		{
			res := Redis{}
			if strings.ContainsAny(data.Host, ":") {
				res.Username = data.User.Username()
				res.Password, _ = data.User.Password()
				res.Host = strings.Split(data.Host, ":")[0]
				res.Port, err = strconv.Atoi(strings.Split(data.Host, ":")[1])
				if err != nil {
					return nil, err
				}
				res.IsCluster, err = strconv.ParseBool(strings.Split(data.RawQuery, "=")[1])
				if err != nil {
					return nil, err
				}
			} else {
				err = errors.New("url doesn't contain port")
				return nil, err
			}
			return res, nil
		}
	case "smtp":
		{
			res := SMTP{}
			path := strings.Split(data.Path, "/")[1:]
			if strings.ContainsAny(data.Host, ":") {
				res.Username = data.User.Username()
				res.Password, _ = data.User.Password()
				res.Host = strings.Split(data.Host, ":")[0]
				res.Port, err = strconv.Atoi(strings.Split(data.Host, ":")[1])
				if err != nil {
					return nil, err
				}
				res.Email = path[0]
				res.SSL = path[1]
			} else {
				err = errors.New("url doesn't contain port")
				return nil, err
			}

			return res, nil
		}
	case "secret":
		{
			res := Secret{}
			res.ApiKey = data.Host

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