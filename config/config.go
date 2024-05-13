package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Environment string

const (
	Development Environment = "dev"
	Staging     Environment = "staging"
)

type Config struct {
	PG_HOST      string
	PG_PORT      string
	PG_NAME      string
	PG_USER      string
	PG_PASS      string
	DATABASE_URL string
	PORT         int
	ENVIRONMENT  Environment
}

// Get retrieves the env variable from your .env file. if it does not exist, a callback string is returned.
func Get(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// Get retrieves the env variable(int) from your .env file. if it does not exist, a callback int is returned.
func GetInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Printf("%s: %s", key, err)
			return fallback
		}
		return i
	}
	return fallback
}

// Get retrieves the environment from your .env file.
func GetEnvironment() Environment {
	if env := Get("ENV", ""); env == "" {
		return Development
	} else {
		return Environment(env)
	}
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("error loading env file: %s", err)
	}
	return &Config{
		PG_HOST:      Get("PG_HOST", "localhost"),
		PG_PORT:      Get("PG_PORT", "5432"),
		PG_NAME:      Get("PG_NAME", "achieve"),
		PG_USER:      Get("PG_USER", "postgres"),
		PG_PASS:      Get("PG_PASS", "password"),
		DATABASE_URL: Get("DATABASE_URL", ""),
		PORT:         GetInt("PORT", 8080),
		ENVIRONMENT:  GetEnvironment(),
	}
}
