package handlers

import (
	"log"
	"net/http"

	"joepbuhre/amphia-agenda-ical/v2/models"

	"github.com/gin-gonic/gin"
)

func HandleIcal(c *gin.Context) {
	log.Printf("Incoming request with method [%s] on path %s", c.Request.Method, c.Request.URL.Path)

	agendaurl := c.Query("agenda")

	calendar, err := models.GetIcal(agendaurl)
	if err != nil {
		c.String(http.StatusOK, "") // Or handle the error as appropriate
	} else {
		c.String(http.StatusOK, calendar)
	}
}
