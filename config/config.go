package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	App         string
	AppVersion  string
	Environment string // development, staging, production
	HTTPPort    string

	CategoryServiceGrpcHost string
	CategoryServiceGrpcPort string

	ProductServiceGrpcHost string
	ProductServiceGrpcPort string
}

// Load ...
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.App = cast.ToString(getOrReturnDefaultValue("APP", "article"))
	config.AppVersion = cast.ToString(getOrReturnDefaultValue("APP_VERSION", "1.0.0"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "development"))
	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":7070"))

	config.CategoryServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("CATEGORY_SERVICE_HOST", "localhost"))
	config.CategoryServiceGrpcPort = cast.ToString(getOrReturnDefaultValue("CATEGORY_SERVICE_PORT", ":9000"))

	config.ProductServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("PRODUCT_SERVICE_HOST", "localhost"))
	config.ProductServiceGrpcPort = cast.ToString(getOrReturnDefaultValue("PRODUCT_SERVICE_PORT", ":9000"))
	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
