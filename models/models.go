package models

type Sql struct {
	Username string
	Password string
	DBname   string
	Host     string
	Schema   string
}

type Kafka struct {
	Host  string
	Topic string
	Group string
}

type SMTP struct {
	Username string
	Password string
	Host     string
	Email    string
}

type Redis struct {
	Host string
}
