package config

import (
	"log"
	"os"

	"github.com/alexpts/go-next/next"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Layers  []LayerConfig
	Version string
	Http    HttpConfig
}

type HttpConfig struct {
	Host string
	Port int
}

type LayerConfig struct {
	Path         string
	Methods      string
	Action       string
	Name         string
	Priority     int
	Handler      next.Handler
	Restrictions next.Restrictions
}

func Load(path string) Config {
	rawData, err := os.ReadFile(path)

	var config Config
	err = yaml.Unmarshal(rawData, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return config
}
