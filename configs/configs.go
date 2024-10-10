package configs

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func DefaultReadTimeout() time.Duration {
	return 10 * time.Second
}
func DefaultWriteTimeout() time.Duration {
	return 10 * time.Second
}

func LoadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
}

func EnvPort() string {
	port, found := os.LookupEnv("PORT")
	if found {
		return port
	}
	return "9099"
}

func EnvHostname() string {
	hostname, found := os.LookupEnv("HOST")
	if found {
		return hostname
	}
	return "app.meetdicer.local"
}

func EnvMongoURI() string {
	LoadDotEnv()

	return os.Getenv("MONGODB_URI")
}
