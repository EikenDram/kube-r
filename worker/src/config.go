package main

import "log"

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

// load configuration from env and config.json
func loadConfig() {
	//
}
