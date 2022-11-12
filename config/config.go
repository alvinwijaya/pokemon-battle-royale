package config

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DbUser       string `envconfig:"DB_USER"`
	DbPassword   string `envconfig:"DB_PASSWORD"`
	DbHost       string `envconfig:"DB_HOST"`
	DbPort       string `envconfig:"DB_PORT"`
	DbName       string `envconfig:"DB_NAME"`
	DbConnection string `envconfig:"DB_CONNECTION"`

	AppPort string `envconfig:"APP_PORT"`

	PokeapiUrl             string `envconfig:"POKEAPI_URL"`
	PokeapiPaginationLimit string `envconfig:"POKEAPI_PAGINATION_LIMIT"`

	MatchParticipantCount int `envconfig:"MATCH_PARTICIPANT_COUNT"`
}

var once sync.Once
var instance Config

func GetConfig() Config {
	godotenv.Load()
	once.Do(func() {
		err := envconfig.Process("", &instance)
		if err != nil {
			log.Fatal(err.Error())
		}

	})

	return instance
}
