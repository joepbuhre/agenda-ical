package handlers

import (
	"log"
	"net/http"
	"strconv"

	"joepbuhre/amphia-agenda-ical/v2/models"
	"joepbuhre/amphia-agenda-ical/v2/utils"

	"github.com/gin-gonic/gin"
)

func CreateAgenda(c *gin.Context) {
	db := utils.GetDB()

	var agenda models.Agenda
	if err := c.ShouldBindJSON(&agenda); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := models.CreateAgenda(db, &agenda); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func DeleteAgenda(c *gin.Context) {
	db := utils.GetDB()

	id := c.Query("id")
	if err := models.DeleteAgenda(db, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func GetAgendaUrl(c *gin.Context) {
	idstr := c.Query("id")

	var err error
	var id int
	id, err = strconv.Atoi(idstr)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url := models.GetAgendaUrl(utils.GetDB(), id)
	c.String(http.StatusOK, url)
}
