package models

import (
	"joepbuhre/amphia-agenda-ical/v2/utils"
	"log"
	"strconv"
	"time"

	ics "github.com/arran4/golang-ical"
)

type IcalEvent struct {
	CreateDateTime time.Time
	StartDateTime  time.Time
	EndDateTime    time.Time
	Summary        string
	Description    string
	Location       string
	Color          string
	AgendaId       int
}

func GetIcal(agendaCode string) (string, error) {

	meetings, err := GetMeetingsByCode(utils.GetDB(), agendaCode)

	log.Println(meetings)

	if err != nil {
		log.Println(err)
		return "", err
	}
	cal := ics.NewCalendar()

	for _, meeting := range meetings {
		event := cal.AddEvent(strconv.Itoa(meeting.Id))

		event.SetCreatedTime(meeting.CreateDateTime)
		event.SetStartAt(meeting.StartDateTime)
		event.SetEndAt(meeting.EndDateTime)
		event.SetSummary(meeting.Summary)
		event.SetDescription(meeting.Description)
		event.SetLocation(meeting.Location)
		event.SetColor(meeting.Color)
		event.SetDtStampTime(time.Now())

	}

	return cal.Serialize(), nil
}

func GetJson(agendaCode string, fromDate time.Time, toDate time.Time) ([]DbMeeting, error) {
	// Because todate is less then (not equal)
	toDate = toDate.AddDate(0, 0, 1)
	meetings, err := GetMeetingsByCodeAndDate(utils.GetDB(), fromDate, toDate, agendaCode)

	if err != nil {
		log.Println(err)
		return []DbMeeting{}, err
	}

	return meetings, nil
}
