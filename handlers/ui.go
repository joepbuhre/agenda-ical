package handlers

import (
	"html/template"
	"log"
	"net/http"

	"joepbuhre/amphia-agenda-ical/v2/models"
	"joepbuhre/amphia-agenda-ical/v2/utils"

	"github.com/gin-gonic/gin"
)

type PageData struct {
	Agendas  []models.Agenda
	Meetings []models.DbMeeting
}

func HandleUI(c *gin.Context) {
	db := utils.GetDB()

	// Fetch all agendas and meetings
	agendas, err := models.GetAgendas(db)
	log.Println(agendas)
	if err != nil {
		log.Println("Error fetching agendas:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching agendas"})
		return
	}

	// Load the template
	tmpl, err := template.ParseFiles("assets/index.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error loading page"})
		return
	}

	// Render the page with dynamic data
	pageData := PageData{
		Agendas: agendas,
	}

	err = tmpl.Execute(c.Writer, pageData)
	if err != nil {
		log.Println("Error executing template:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error rendering page"})
		return
	}
}
