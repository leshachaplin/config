# config

in order to parse data from .env file in it, they must be presented in the following form:
* SQL: "sql://username:password@hostname.with.domain:5432/schema/dbname"
* Kafka: "kafka://topic:group@hostname.with.domain:9092/"
* Redis: "redis://username:password@hostname.with.domain:6379/?isCluster=false"
* SMTP: "smtp://username:password@hostname.with.domain:port/email/ssl"
* MongoDB: "mongo://hostname.with.domain:port/"
* SecretApiKey: "secret://apikey/"