package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// post new report request
func CreateReport(c *gin.Context) {
	// need to make a new record in database

	// and send message to rabbit mq

	c.String(http.StatusOK, "success")
}

// update report request status
func UpdateReport(c *gin.Context) {
	// need to update record in database

	// and make api call to application

	c.String(http.StatusOK, "success")
}
