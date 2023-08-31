package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// post new report request
func (env *Env) CreateReport(c *gin.Context) {
	// need to make a new record in database

	// and send message to rabbit mq

	c.String(http.StatusOK, "success")
}

// update report request status
func (env *Env) UpdateReport(c *gin.Context) {
	// need to update record in database

	// and make api call to application

	c.String(http.StatusOK, "success")
}

func (env *Env) Test(c *gin.Context) {
	res, err := env.requests.All()
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, err)
		//http.Error(w, http.StatusText(500), 500)
		return
	}

	c.JSON(http.StatusOK, res)
}
