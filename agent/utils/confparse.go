package utils

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	configPath = "/bin/"
)

type Config struct {
	Services struct {
		SSH    bool `yaml:"ssh"`
		Auditd bool `yaml:"auditd"`
	} `yaml:"services"`
	Stream struct {
		URL string `yaml:"url"`
	} `yaml:"stream"`
	Files   []string `yaml:"files"`
	Capture struct {
		Uploads bool `yaml:"upload"`
	} `yaml:"capture"`
	Webfiles string `yaml:"webfiles"`
}

func ParseConfig() (*Config, error) {
	yamlFile, err := os.ReadFile(configPath + "config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return &config, nil
}
