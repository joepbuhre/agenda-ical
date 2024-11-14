package handlers

import (
	"log"
	"net/http"
	"time"

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

func HandleJson(c *gin.Context) {
	log.Printf("Incoming request with method [%s] on path %s", c.Request.Method, c.Request.URL.Path)

	agendaurl := c.Query("agenda")

	fromDate := c.Query("from_date")
	if fromDate == "" {
		// Default to high date
		fromDate = time.Now().AddDate(-99, 0, 0).Format("2006-01-02")
	}
	toDate := c.Query("to_date")
	if toDate == "" {
		// Default to high date
		toDate = time.Now().AddDate(99, 0, 0).Format("2006-01-02")
	}

	log.Println(fromDate)
	log.Println(toDate)

	// Parse the start and end datetime
	fromDT, err := time.Parse(time.DateOnly, fromDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start datetime format"})
		return
	}

	toDT, err := time.Parse(time.DateOnly, toDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end datetime format"})
		return
	}

	meetings, err := models.GetJson(agendaurl, fromDT, toDT)
	if err != nil {
		c.String(http.StatusOK, "") // Or handle the error as appropriate
	} else {
		c.JSON(http.StatusOK, meetings)
	}
}
