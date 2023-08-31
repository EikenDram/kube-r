package main

import (
	"github.com/EikenDram/kube-r/manager/config"
	"github.com/EikenDram/kube-r/manager/database"
)

type Env struct {
	config *config.ConfigStruct

	requests     database.RequestModel
	applications database.ApplicationModel
	requestLogs  database.RequestLogModel
	users        database.UserModel
}

func Init() (*Env, error) {
	//load configuration
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	// Initialise the connection pool as normal.
	db, err := database.Init(cfg)
	if err != nil {
		return nil, err
	}

	// Initalise Env with a models.BookModel instance (which in turn wraps
	// the connection pool).
	env := &Env{
		config:       cfg,
		requests:     database.RequestModel{DB: db},
		applications: database.ApplicationModel{DB: db},
		requestLogs:  database.RequestLogModel{DB: db},
		users:        database.UserModel{DB: db},
	}

	return env, nil
}
