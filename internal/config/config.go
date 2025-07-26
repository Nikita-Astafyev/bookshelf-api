package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Debug        bool
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type JWTConfig struct {
	SecretKey string
	ExpiresIn time.Duration
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	readTimeout, err := strconv.Atoi(getEnv("SERVER_READ_TIMEOUT", "10"))
	if err != nil {
		return nil, err
	}

	writeTimeout, err := strconv.Atoi(getEnv("SERVER_WRITE_TIMEOUT", "10"))
	if err != nil {
		return nil, err
	}

	jwtExpiresIn, err := strconv.Atoi(getEnv("JWT_EXPIRES_IN", "86400")) // 24 часа
	if err != nil {
		return nil, err
	}

	config := &Config{
		Server: ServerConfig{
			Port:         getEnv("SERVER_PORT", "8080"),
			ReadTimeout:  time.Duration(readTimeout) * time.Second,
			WriteTimeout: time.Duration(writeTimeout) * time.Second,
			Debug:        getEnv("SERVER_DEBUG", "false") == "true",
		},
		Postgres: PostgresConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "bookshelf"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			SecretKey: getEnv("JWT_SECRET_KEY", "very-secret-key"),
			ExpiresIn: time.Duration(jwtExpiresIn) * time.Second,
		},
	}

	return config, err
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func (pc *PostgresConfig) GetPostgresDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		pc.Host, pc.Port, pc.User, pc.Password, pc.DBName, pc.SSLMode)
}
