package database

import (
	"database/sql"
	"fmt"
)

func ConnectToDB() (*sql.DB, error) {
	user := "root"
	password := "root"
	port := "3306"
	ip := "localhost"
	// Change the name of the schema
	schema := "my-schema"

	url := user + ":" + password + "@tcp(" + ip + ":" + port + ")/" + schema

	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, fmt.Errorf("error conectando a la base de datos: %w", err)
	}

	return db, nil
}
