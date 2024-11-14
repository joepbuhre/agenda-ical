package models

import (
	"database/sql"
	"log"
	"time"
)

type DbMeeting struct {
	Id             int
	CreateDateTime time.Time
	StartDateTime  time.Time
	EndDateTime    time.Time
	Summary        string
	Description    string
	Location       string
	Color          string
	AgendaId       int
}

func AddOrUpdateMeeting(db *sql.DB, meeting *DbMeeting) error {
	_, err := db.Exec(`
		INSERT OR REPLACE INTO  meetings (id, create_datetime,start_datetime,end_datetime,summary,description,location,color, agenda_id) 
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		meeting.Id, meeting.CreateDateTime, meeting.StartDateTime, meeting.EndDateTime, meeting.Summary, meeting.Description, meeting.Location, meeting.Color, meeting.AgendaId)

	if err != nil {
		return err
	}

	log.Printf("Inserted meeting on AgendaId [%v]", meeting)
	return err
}

func DeleteMeeting(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM meetings WHERE id = ?", id)
	log.Printf("Deleted meeting with id [%v]", id)
	return err
}

func DeleteMeetingWithDates(db *sql.DB, fromDate time.Time, toDate time.Time) error {
	_, err := db.Exec("DELETE FROM meetings WHERE start_datetime >= ? and end_datetime <= ?", fromDate, toDate)
	log.Printf("Deleted meeting with from [%v] to [%v]", fromDate, toDate)
	return err
}

func GetMeetings(db *sql.DB, id string) ([]DbMeeting, error) {
	rows, err := db.Query("SELECT COALESCE(id,0),create_datetime,start_datetime,end_datetime,summary,description,location,color FROM meetings where agenda_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	meetings := []DbMeeting{}
	for rows.Next() {
		var meeting DbMeeting
		if err := rows.Scan(&meeting.Id, &meeting.CreateDateTime, &meeting.StartDateTime, &meeting.EndDateTime, &meeting.Summary, &meeting.Description, &meeting.Location, &meeting.Color); err != nil {
			return nil, err
		}
		meetings = append(meetings, meeting)
	}
	return meetings, nil
}

func GetMeetingsByCode(db *sql.DB, code string) ([]DbMeeting, error) {
	log.Println(code)
	rows, err := db.Query("SELECT COALESCE(id, 0),create_datetime,start_datetime,end_datetime,summary,description,location,color FROM meetings where agenda_id = (select agenda_id from agendaurls where code = ? LIMIT 1)", code)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	meetings := []DbMeeting{}
	for rows.Next() {
		var meeting DbMeeting
		if err := rows.Scan(&meeting.Id, &meeting.CreateDateTime, &meeting.StartDateTime, &meeting.EndDateTime, &meeting.Summary, &meeting.Description, &meeting.Location, &meeting.Color); err != nil {
			return nil, err
		}
		meetings = append(meetings, meeting)
	}
	return meetings, nil
}

func GetMeetingsByCodeAndDate(db *sql.DB, fromDate time.Time, toDate time.Time, code string) ([]DbMeeting, error) {
	log.Println(code)
	log.Println(fromDate)
	log.Println(toDate)
	rows, err := db.Query("SELECT COALESCE(id, 0),create_datetime,start_datetime,end_datetime,summary,description,location,color FROM meetings where start_datetime >= ? and end_datetime < ? and agenda_id = (select agenda_id from agendaurls where code = ? LIMIT 1)", fromDate, toDate, code)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	meetings := []DbMeeting{}
	for rows.Next() {
		var meeting DbMeeting
		if err := rows.Scan(&meeting.Id, &meeting.CreateDateTime, &meeting.StartDateTime, &meeting.EndDateTime, &meeting.Summary, &meeting.Description, &meeting.Location, &meeting.Color); err != nil {
			return nil, err
		}
		meetings = append(meetings, meeting)
	}
	return meetings, nil
}
