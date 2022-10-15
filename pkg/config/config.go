package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	ServiceName   string
	ServiceStatus string

	GRPC string

	TCP  string
	Port string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	return &Config{
		ServiceName:   getEnv("SERVICE_NAME"),
		ServiceStatus: getEnv("SERVICE_STATUS"),

		GRPC: getEnv("GRPC_HOST_PORT"),

		TCP:  getEnv("TCP"),
		Port: getEnv("PORT"),
	}
}

func getEnv(key string) string {
	value := os.Getenv(key)
	return value
}
