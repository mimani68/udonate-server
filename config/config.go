package config

import (
	"os"
	"udonate/exception"

	"github.com/joho/godotenv"
)

type IConfig interface {
	Get(key string) string
	Set(key string, value string) bool
}

type Config struct {
	ADDRESS                    string
	PORT                       string
	MONGO_URI                  string
	MONGO_DATABASE             string
	MONGO_POOL_MIN             string
	MONGO_POOL_MAX             string
	MONGO_MAX_IDLE_TIME_SECOND string
	TZ                         string
}

func (c *Config) Get(key string) string {
	result := ""
	switch key {
	case "ADDRESS":
		result = c.ADDRESS
	case "PORT":
		result = c.PORT
	case "MONGO_URI":
		result = c.MONGO_URI
	case "MONGO_DATABASE":
		result = c.MONGO_DATABASE
	case "MONGO_POOL_MIN":
		result = c.MONGO_POOL_MIN
	case "MONGO_POOL_MAX":
		result = c.MONGO_POOL_MAX
	case "MONGO_MAX_IDLE_TIME_SECOND":
		result = c.MONGO_MAX_IDLE_TIME_SECOND
	case "TZ":
		result = c.TZ
	}
	return result
}

func (c *Config) Set(key string, value string) bool {
	return true
}

func New(filenames ...string) IConfig {
	err := godotenv.Load(filenames...)
	exception.PanicIfNeeded(err)
	return &Config{
		ADDRESS:                    os.Getenv("ADDRESS"),
		PORT:                       os.Getenv("PORT"),
		MONGO_URI:                  os.Getenv("MONGO_URI"),
		MONGO_DATABASE:             os.Getenv("MONGO_DATABASE"),
		MONGO_POOL_MIN:             os.Getenv("MONGO_POOL_MIN"),
		MONGO_POOL_MAX:             os.Getenv("MONGO_POOL_MAX"),
		MONGO_MAX_IDLE_TIME_SECOND: os.Getenv("MONGO_MAX_IDLE_TIME_SECOND"),
		TZ:                         "America/New_York",
	}
}
