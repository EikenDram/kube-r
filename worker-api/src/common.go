package main

import (
	"log"
	"os"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v2"
)

// configuration
var config ConfigYaml

// server plumber api credentials
var serverUser string
var serverPass string

// rabbitmq credentials
var rabbitUser string
var rabbitPass string

// manager api credentials
var managerUser string
var managerPass string

// load configuration from env and config.json
func loadConfig() {
	// read env variables
	serverUser = os.Getenv("SERVER_USER")
	serverPass = os.Getenv("SERVER_PASS")
	rabbitUser = os.Getenv("RABBIT_USER")
	rabbitPass = os.Getenv("RABBIT_PASS")
	managerUser = os.Getenv("MANAGER_USER")
	managerPass = os.Getenv("MANAGER_PASS")

	// read config.json
	if err := defaults.Set(&config); err != nil {
		panic(err)
	}
	file, err := os.ReadFile("./config.yaml")
	check(err)
	err = yaml.Unmarshal(file, &config)
	check(err)
}

// fail on error
func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

// check for error
func check(err error) {
	if err != nil {
		panic(err)
	}
}
