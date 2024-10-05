package handlers

import (
	"log"
	"net/http"
	"time"

	"joepbuhre/amphia-agenda-ical/v2/models"
	"joepbuhre/amphia-agenda-ical/v2/utils"

	"github.com/gin-gonic/gin"
)

type MeetingRequest struct {
	ID            int    `json:"id"` // Optional for update
	AgendaID      int    `json:"agenda_id"`
	Summary       string `json:"summary"`
	Description   string `json:"description"`
	StartDateTime string `json:"start_datetime"`
	EndDateTime   string `json:"end_datetime"`
	Location      string `json:"location"`
	Color         string `json:"color"`
}

func PutMeeting(c *gin.Context) {
	db := utils.GetDB()

	var meetingReq MeetingRequest
	if err := c.ShouldBindJSON(&meetingReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Parse the start and end datetime
	startDT, err := time.Parse(time.RFC3339, meetingReq.StartDateTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start datetime format"})
		return
	}

	endDT, err := time.Parse(time.RFC3339, meetingReq.EndDateTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end datetime format"})
		return
	}

	// Add meeting
	err = models.AddOrUpdateMeeting(db, &models.DbMeeting{
		Id:             meetingReq.ID,
		CreateDateTime: time.Now(),
		StartDateTime:  startDT,
		EndDateTime:    endDT,
		Summary:        meetingReq.Summary,
		Description:    meetingReq.Description,
		Location:       meetingReq.Location,
		Color:          meetingReq.Color,
		AgendaId:       meetingReq.AgendaID,
	})

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	c.Status(http.StatusCreated)
}
