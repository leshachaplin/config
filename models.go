package main

type Sql struct {
	Username string
	Password string
	DB       string
	Host     string
	Port     string
	Schema   string
}

type Kafka struct {
	Host  string
	Topic string
	Group string
	Port  string
}

type SMTP struct {
	Username string
	Password string
	Host     string
	Port     string
	SSL      string
	Email    string
}

type Redis struct {
	Username string
	Password string
	Host     string
	Port     string
}

type Secret struct {
	Username string
	ApiKey   string
	ApiHost  string
	ApiPort  string
}

type MongoDB struct {
	ConnectionString string
}