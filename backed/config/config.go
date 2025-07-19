package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`

	MongoHost     string `yaml:"mongoHost"`
	MongoPort     int    `yaml:"mongoPort"`
	MongoUser     string `yaml:"mongoUser"`
	MongoPassword string `yaml:"mongoPassword"`
	MongoDBName   string `yaml:"mongoDBName"`
}

type AppConfig struct {
	Name      string `yaml:"name"`
	AppHost   string `yaml:"appHost"`
	FrontHost string `yaml:"frontHost"`
}

type SMTPConfig struct {
	SMTPHost     string `yaml:"smtpHost"`
	SMTPPort     int    `yaml:"smtpPort"`
	SMTPUser     string `yaml:"smtpUser"`
	SMTPPassword string `yaml:"smtpPassword"`
}

type Config struct {
	Database DBConfig   `yaml:"database"`
	App      AppConfig  `yaml:"app"`
	SMTP     SMTPConfig `yaml:"smtp"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
