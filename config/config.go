package config

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

func GetEnv(envPath string, logger *logrus.Logger) map[string]string {
	_, err := os.Stat(envPath)
	if errors.Is(err, os.ErrNotExist) {
		logger.Fatalf("EnvFile doesnt exist, error:%v", err)
	}
	EnvMap, err := godotenv.Read(envPath)
	godotenv.Load()

	if err != nil {
		logger.Panicf("\nUnable to parse env.conf file, error:%v\n", err)
	}

	return EnvMap

}
