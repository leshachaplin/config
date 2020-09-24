package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"net/url"
	"strings"
)

func Init() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	return nil
}

func GetSQL(sqlUrl string) (Sql, error) {
	res, err := resolve(sqlUrl)
	return res.(Sql), err
}

func GetGRPC(grpcUrl string) (GRPC, error) {
	res, err := resolve(grpcUrl)
	return res.(GRPC), err
}

func GetSecret(secretUrl string) (Secret, error) {
	res, err := resolve(secretUrl)
	return res.(Secret), err
}

func GetMongo(mongoUrl string) (MongoDB, error) {
	res, err := resolve(mongoUrl)
	return res.(MongoDB), err
}

func GetSMTP(smtpUrl string) (SMTP, error) {
	res, err := resolve(smtpUrl)
	return res.(SMTP), err
}

func GetKafka(kafkaUrl string) (Kafka, error) {
	res, err := resolve(kafkaUrl)
	return res.(Kafka), err
}

func GetRedis(redisUrl string) (Redis, error) {
	res, err := resolve(redisUrl)
	return res.(Redis), err
}

func resolve(sqlUrl string) (interface{}, error) {
	data, err := url.Parse(sqlUrl)
	if err != nil {
		return nil, err
	}

	a, _ := json.Marshal(data)
	fmt.Println(string(a))

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
	case "grpc":
		{
			res := GRPC{}
			res.Host = strings.Split(data.Host, ":")[0]
			res.Port = strings.Split(data.Host, ":")[1]

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
