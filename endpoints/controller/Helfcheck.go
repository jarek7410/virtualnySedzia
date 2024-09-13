package controller

import (
	"afmib_server/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HalfCheck(c *gin.Context) {
	va := model.HealthCheck()
	var dbStatus string
	if va != nil {
		dbStatus = "down"
	} else {
		dbStatus = "up"
	}
	data := map[string]interface{}{
		"version":  "alpha",
		"database": dbStatus,
	}
	c.JSONP(http.StatusOK, data)
}
func Coffee(c *gin.Context) {
	c.String(http.StatusTeapot, "I am a teapot but programer need coffee!!!\nblik 735917147")
}
