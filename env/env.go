package env

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type envVariables struct {
	DbHost      string
	DbUser      string
	DbName      string
	DbPort      string
	DbPassword  string
	DatabaseUrl string
}

var (
	instance *envVariables
	once     sync.Once
)

// loadAndValidateEnv loads environment variables and validates their presence.
func loadAndValidateEnv() *envVariables {
	var missingEnv []string
	requiredEnv := []string{"DB_HOST", "DB_USER", "DB_PASSWORD", "DB_NAME", "DB_PORT", "DATABASE_URL"}

	// load .env if it's not in production
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error while loading file .env")
		}
	}

	// validate required envs
	for _, v := range requiredEnv {
		if os.Getenv(v) == "" {
			missingEnv = append(missingEnv, v)
		}
	}
	if missingEnv != nil {
		log.Fatal("❌ MISSING ENVIRONMENT VARIABLES: ", missingEnv)
	}

	// return the env struct
	return &envVariables{
		DbHost:      os.Getenv("DB_HOST"),
		DbUser:      os.Getenv("DB_USER"),
		DbPassword:  os.Getenv("DB_PASSWORD"),
		DbName:      os.Getenv("DB_NAME"),
		DbPort:      os.Getenv("DB_PORT"),
		DatabaseUrl: os.Getenv("DATABASE_URL"),
	}
}

// Get provides a thread-safe singleton instance of envVariables.
func Get() *envVariables {
	once.Do(func() {
		instance = loadAndValidateEnv()
	})
	return instance
}
