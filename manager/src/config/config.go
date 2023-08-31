package config

import (
	"os"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v2"
)

type ConfigStruct struct {
	Manager      ManagerStruct
	Database     DatabaseStruct
	RabbitMQ     RabbitMQStruct
	Applications []ApplicationStruct
}

type ManagerStruct struct {
	User string
	Pass string
}

type ConfigYaml struct {
	Database     DatabaseStruct      `yaml:"database"`
	RabbitMQ     RabbitMQStruct      `yaml:"rabbitmq"`
	Applications []ApplicationStruct `yaml:"applications"`
}

type DatabaseStruct struct {
	Host string `yaml:"host"`
	Port string `yaml:"port" default:"80"`
	Name string `yaml:"name"`
	User string `yaml:"user"`
	Pass string `yaml:"password"`
}

type RabbitMQStruct struct {
	Host string `yaml:"host"`
	Port string `yaml:"port" default:"80"`
	User string `yaml:"user"`
	Pass string `yaml:"password"`
}

type ApplicationStruct struct {
	Name   string `yaml:"name"`
	URL    string `yaml:"url"`
	Notify string `yaml:"notify"`
	Result string `yaml:"result"`
	Prefix string `yaml:"prefix"`
}

// load configuration from env and config.yaml
func Load() (*ConfigStruct, error) {
	// read env variables
	rabbitUser := os.Getenv("RABBIT_USER")
	rabbitPass := os.Getenv("RABBIT_PASS")
	managerUser := os.Getenv("MANAGER_USER")
	managerPass := os.Getenv("MANAGER_PASS")
	databaseUser := os.Getenv("DATABASE_USER")
	databasePass := os.Getenv("DATABASE_PASS")

	// read config.json
	var configYaml = new(ConfigYaml)
	err := defaults.Set(&configYaml)
	if err != nil {
		return nil, err
	}
	file, err := os.ReadFile("./config.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, &configYaml)
	if err != nil {
		return nil, err
	}

	//fill config
	var config = new(ConfigStruct)
	config.Applications = configYaml.Applications
	config.Database = configYaml.Database
	config.Database.User = databaseUser
	config.Database.Pass = databasePass
	config.Manager.User = managerUser
	config.Manager.Pass = managerPass
	config.RabbitMQ = configYaml.RabbitMQ
	config.RabbitMQ.User = rabbitUser
	config.RabbitMQ.Pass = rabbitPass
	return config, nil
}
