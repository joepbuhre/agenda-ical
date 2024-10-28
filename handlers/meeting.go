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

func RegisterMeetingRoutes(r *gin.RouterGroup) {
	meeting := r.Group("/meeting")
	{
		meeting.GET("/:agenda-id", GetMeetings)
		meeting.PUT("/", PutMeeting)
		meeting.DELETE("/", DeleteMeeting)
	}
}

func GetMeetings(c *gin.Context) {
	db := utils.GetDB()

	agenda_id, found := c.Params.Get("agenda-id")
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid agenda id provided"})
		return
	}

	meetings, err := models.GetMeetings(db, agenda_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong with fetching the meetings"})
		return
	}

	c.JSON(http.StatusOK, meetings)
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

func DeleteMeeting(c *gin.Context) {
	db := utils.GetDB()

	fromDate := c.Query("from_date")
	toDate := c.Query("to_date")

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

	log.Println("Deleting: ", fromDT, toDT)

	models.DeleteMeetingWithDates(db, fromDT, toDT)
}
