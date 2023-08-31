package main

import (
	"log"
)

func main() {
	env, err := Init()
	if err != nil {
		log.Fatal(err)
	}

	r := setupRouter(env)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
