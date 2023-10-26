package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/subosito/gotenv"
)

const (
	defaultServerPort               = "8000"
	defaultServerRWTimeout          = 10 * time.Second
	defaultServerMaxHeaderMegabytes = 1
	TTLCache                        = 10 * time.Minute
)

type Config struct {
	Server   serverConfig
	Database databaseConfig
	Redis    RedisConfig
	Kafka    KafkaConfig
	Logger   LoggerConfig
	Handler  string
}

type LoggerConfig struct {
	Path  string
	Level string
}

type serverConfig struct {
	Port               string
	ReadTImeout        time.Duration
	WriteTimeout       time.Duration
	MaxHeaderMegabytes int
}

type databaseConfig struct {
	Driver string
	DSN    string
}

type RedisConfig struct {
	DB       int
	Host     string
	Port     string
	Password string
	Ttl      time.Duration
}

type KafkaConfig struct {
	Brokers []string
}

func Init() (*Config, error) {
	err := gotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load environment variables from file: %v", err)
	}
	driver := os.Getenv("DRIVER")
	dsn := os.Getenv("DSN")

	host := os.Getenv("HOST_REDIS")
	port := os.Getenv("PORT_REDIS")
	password_redis := os.Getenv("PASSWORD_REDIS")
	db_redis, err := strconv.Atoi(os.Getenv("DB_REDIS"))
	if err != nil {
		return nil, err
	}

	brokerStr := os.Getenv("KAFKA_BROKERS")
	brokers := strings.Split(brokerStr, ",")

	logPath := os.Getenv("LOG_PATH")
	level := os.Getenv("LOG_LEVEL")

	handler := os.Getenv("HANDLER")

	return &Config{
		Server: serverConfig{
			Port:               defaultServerPort,
			ReadTImeout:        defaultServerRWTimeout,
			WriteTimeout:       defaultServerRWTimeout,
			MaxHeaderMegabytes: defaultServerMaxHeaderMegabytes,
		},
		Database: databaseConfig{
			DSN:    dsn,
			Driver: driver,
		},
		Redis: RedisConfig{
			Host:     host,
			DB:       db_redis,
			Password: password_redis,
			Port:     port,
			Ttl:      TTLCache,
		},
		Kafka: KafkaConfig{
			Brokers: brokers,
		},
		Logger: LoggerConfig{
			Path:  logPath,
			Level: level,
		},
		Handler: handler,
	}, nil
}
