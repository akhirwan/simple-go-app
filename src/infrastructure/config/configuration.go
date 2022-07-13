package config

import (
	"os"
	"simple-go-app/src/domain/exception"

	"github.com/joho/godotenv"
)

type Configuration interface {
	Get(key string) string
}

type configImpl struct {
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) Configuration {
	err := godotenv.Load(filenames...)
	exception.PanicIfNeeded(err)
	return &configImpl{}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
