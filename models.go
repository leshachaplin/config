package main

// SQL type with sql connection data
type SQL struct {
	Username string
	Password string
	DB       string
	Host     string
	Port     int
	Schema   string
}

// Kafka type with kafka connection data
type Kafka struct {
	Host  string
	Topic string
	Group string
	Port  int
}

// SMTP type with smtp connection data
type SMTP struct {
	Username string
	Password string
	Host     string
	Port     int
	SSL      string
	Email    string
}

// Redis type with redis connection data
type Redis struct {
	Username  string
	Password  string
	Host      string
	Port      int
	IsCluster bool
}

// Secret api key
type Secret struct {
	ApiKey   string
}

// MongoDB connection string
type MongoDB struct {
	ConnectionString string
}
