package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	DatabaseFile string `yaml:"database_file"`
}

var Config config

func init() {
	dat, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(dat, &Config)
	if err != nil {
		log.Fatal(err)
	}

}
