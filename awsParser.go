package config

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"net/url"
	"strconv"
	"strings"
)

// ConfigAws struct which contain connection struct to aws service manager
type ConfigAws struct {
	conn *ssm.SSM
}

// NewForAws create new ConfigAws and initialize their field
func NewForAws(awsRegion string) (*ConfigAws, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion)},
	)
	if err != nil {
		return nil, err
	}

	return &ConfigAws{conn: ssm.New(sess)}, nil
}

// GetSQL converts result from resolveAws to sql structure type
func (c *ConfigAws) GetSQL(sqlURL string) (*SQL, error) {
	res, err := c.resolveAws(sqlURL)
	if err != nil {
		return nil, err
	}
	sql, ok := res.(SQL)
	if !ok {
		return nil, errors.New("your should use other method to parse this url or use other url")
	}
	return &sql, nil
}

// GetSecret converts result from resolveAws to secretApiKey structure type
func (c *ConfigAws) GetSecret(secretURL string) (*Secret, error) {
	res, err := c.resolveAws(secretURL)
	secret, ok := res.(Secret)
	if !ok {
		return nil, errors.New("your should use other method to parse this url or use other url")
	}
	return &secret, err
}

// GetMongo converts result from resolveAws to mongo structure type
func (c *ConfigAws) GetMongo(mongoURL string) (*MongoDB, error) {
	res, err := c.resolveAws(mongoURL)
	mongo, ok := res.(MongoDB)
	if !ok {
		return nil, errors.New("your should use other method to parse this url or use other url")
	}
	return &mongo, err
}

// GetSMTP converts result from resolveAws to smtp structure type
func (c *ConfigAws) GetSMTP(smtpURL string) (*SMTP, error) {
	res, err := c.resolveAws(smtpURL)
	if err != nil {
		return nil, err
	}
	smtp, ok := res.(SMTP)
	if !ok {
		return nil, errors.New("your should use other method to parse this url or use other url")
	}
	return &smtp, nil
}

// GetKafka converts result from resolveAws to kafka structure type
func (c *ConfigAws) GetKafka(kafkaURL string) (*Kafka, error) {
	res, err := c.resolveAws(kafkaURL)
	if err != nil {
		return nil, err
	}
	kafka, ok := res.(Kafka)
	if !ok {
		return nil, errors.New("your should use other method to parse this url or use other url")
	}
	return &kafka, nil
}

// GetRedis convert result from resolveAws to redis structure type
func (c *ConfigAws) GetRedis(redisURL string) (*Redis, error) {
	res, err := c.resolveAws(redisURL)
	if err != nil {
		return nil, err
	}
	redis, ok := res.(Redis)
	if !ok {
		return nil, errors.New("your should use other method to parse this url or use other url")
	}
	return &redis, nil
}

// Resolve parse connection url in to a specific structure that is indicated in the url scheme
func (c *ConfigAws) resolveAws(URL string) (interface{}, error) {

	awsURL, err := url.Parse(URL)
	if err != nil {
		return nil, err
	}
	param := ssm.GetParameterInput{}
	param.SetName(awsURL.Host)
	param.SetWithDecryption(true)
	paramOut, err := c.conn.GetParameter(&param)
	if err != nil {
		return nil, err
	}

	switch awsURL.Scheme {
	case "aws-ssm-sql":
		{
			data, err := url.Parse(*paramOut.Parameter.Value)
			if err != nil {
				return nil, err
			}
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
	case "aws-ssm-kafka":
		{
			data, err := url.Parse(*paramOut.Parameter.Value)
			if err != nil {
				return nil, err
			}
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
	case "aws-ssm-redis":
		{
			data, err := url.Parse(*paramOut.Parameter.Value)
			if err != nil {
				return nil, err
			}
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
	case "aws-ssm-smtp":
		{
			data, err := url.Parse(*paramOut.Parameter.Value)
			if err != nil {
				return nil, err
			}
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
	case "aws-ssm-secret":
		{
			data, err := url.Parse(*paramOut.Parameter.Value)
			if err != nil {
				return nil, err
			}
			res := Secret{}
			res.ApiKey = data.Host

			return res, nil
		}
	case "aws-ssm-mongo":
		{
			data, err := url.Parse(*paramOut.Parameter.Value)
			if err != nil {
				return nil, err
			}
			res := MongoDB{}
			res.ConnectionString = "mongodb://" + data.Host

			return res, nil
		}
	}
	return nil, err
}
