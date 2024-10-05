package utils

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

var db *sql.DB
var once sync.Once

func InitDB(filepath string) (*sql.DB, error) {
	var err error
	once.Do(func() {
		db, err = sql.Open("sqlite3", filepath)
		if err != nil {
			return
		}
		createTables()
	})
	return db, err
}

func GetDB() *sql.DB {
	return db
}

func createTables() {
	agendaTable := `
    CREATE TABLE IF NOT EXISTS agendas (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT
    );`

	meetingTable := `
	CREATE TABLE IF NOT EXISTS meetings (
		id TEXT, -- Unique identifier for the event
		create_datetime DATETIME NOT NULL,      -- Creation timestamp
		start_datetime DATETIME NOT NULL,       -- Start date and time of the event
		end_datetime DATETIME NOT NULL,         -- End date and time of the event
		summary TEXT NOT NULL,                   -- Brief title or summary of the event
		description TEXT,                        -- Detailed description of the event
		location TEXT,                           -- Location of the event
		color TEXT,                               -- Color associated with the event
		agenda_id integer,
		PRIMARY KEY (id, agenda_id) -- Unique constraint on the combination of id and agenda_id

	);`

	agendaUrlTable := `
    CREATE TABLE IF NOT EXISTS agendaurls (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        agenda_id INTEGER,
        code TEXT
    );`

	if _, err := db.Exec(agendaTable); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(meetingTable); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(agendaUrlTable); err != nil {
		log.Fatal(err)
	}
}
