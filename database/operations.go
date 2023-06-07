package database

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mssql", "server=localhost;database=InfoCheck_Dev;trusted_connection=true")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func closeDb(db *sql.DB) {
	db.Close()
}

func fetchRespondent(db *sql.DB) (Respondent, error) {

	rows, err := db.Query("SELECT * FROM Respondent")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var result []Respondent

	for rows.Next() {
		var respondent Respondent
		err := rows.Scan(&respondent)
		if err != nil {
			panic(err)
		}

		result = append(result, respondent)
	}

	return result
}

func insertRespondent(db *sql.DB, respondent Respondent) error {
	query := "INSERT INTO info_check (ID, Name, Email, Age, CreatedAt) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(query, respondent.Id, respondent.Name, respondent.Email, respondent.Age, respondent.CreatedAt)
	return err
}
