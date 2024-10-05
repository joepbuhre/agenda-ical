package models

import (
	"database/sql"
	"joepbuhre/amphia-agenda-ical/v2/utils"
	"log"
)

type Agenda struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AgendaUrlResponse struct {
	Url string `json:"url"`
}

func CreateAgenda(db *sql.DB, agenda *Agenda) error {
	_, err := db.Exec("INSERT INTO agendas (name) VALUES (?)", agenda.Name)
	log.Printf("inserted record into agenda with [%s]", agenda.Name)
	return err
}

func DeleteAgenda(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM agendas WHERE id = ?", id)
	log.Printf("deleted record from agenda with [%v]", id)
	return err
}

func GetAgendas(db *sql.DB) ([]Agenda, error) {
	rows, err := db.Query("SELECT id, name FROM agendas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	agendas := []Agenda{}
	for rows.Next() {
		var agenda Agenda
		if err := rows.Scan(&agenda.ID, &agenda.Name); err != nil {
			return nil, err
		}
		agendas = append(agendas, agenda)
	}
	return agendas, nil
}

func GetAgendaUrl(db *sql.DB, id int) string {
	str, err := utils.GenerateSecureString(32)
	if err != nil {
		return ""
	}
	_, err = db.Exec("INSERT INTO agendaurls (agenda_id, code) VALUES (?, ?)", id, str)
	log.Printf("inserted record into agenda with id [%s] and code [%v]", id, str)

	return str
}
