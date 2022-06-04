package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	ServerPort        string
	KafkaBroker       string
	ReceptCreateTopic string
}

var envConfig *EnvConfig

func GetEnv() EnvConfig {
	if envConfig == nil {
		getEnv()
	}
	return *envConfig
}

func getEnv() {
	err := godotenv.Load()

	if err != nil {
		fmt.Printf("error loading .env file")
	}

	envConfig = &EnvConfig{
		ServerPort:        os.Getenv("SERVER_PORT"),
		KafkaBroker:       os.Getenv("KAFKA_BROKER"),
		ReceptCreateTopic: os.Getenv("KAFKA_TOPIC_RECEPT_CREATE"),
	}
}
