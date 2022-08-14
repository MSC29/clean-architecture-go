package infrastructure

import (
	"clean-architecture/clean-architecture-go/src/domain"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

type ConfigurationMapper struct{}

func NewConfigurationMapper() domain.ConfigurationEntity {
	projectName := regexp.MustCompile(`^(.*clean-architecture-go)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	env := os.Getenv("ENV")
	err := godotenv.Load(fmt.Sprintf(`%s/.env.%s`, string(rootPath), env))

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return domain.ConfigurationEntity{
		DatabaseName:     os.Getenv("DATABASE_NAME"),
		DatabaseUsername: os.Getenv("DATABASE_USERNAME"),
		DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
		CatsSource:       os.Getenv("CATS_SOURCE"),
		Env:              env,
	}
}
