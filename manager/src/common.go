package main

import (
	"os"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v2"
)

// configuration
var config ConfigYaml

// database credentials
var databaseUser string
var databasePass string

// rabbitmq credentials
var rabbitUser string
var rabbitPass string

// manager api credentials
var managerUser string
var managerPass string

// load configuration from env and config.yaml
func loadConfig() {
	// read env variables
	rabbitUser = os.Getenv("RABBIT_USER")
	rabbitPass = os.Getenv("RABBIT_PASS")
	managerUser = os.Getenv("MANAGER_USER")
	managerPass = os.Getenv("MANAGER_PASS")
	databaseUser = os.Getenv("DATABASE_USER")
	databasePass = os.Getenv("DATABASE_PASS")

	// read config.json
	err := defaults.Set(&config)
	check(err)
	file, err := os.ReadFile("./config.yaml")
	check(err)
	err = yaml.Unmarshal(file, &config)
	check(err)
}

// check for error
func check(err error) {
	if err != nil {
		panic(err)
	}
}
