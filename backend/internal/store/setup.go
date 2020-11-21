package store

import (
	"database/sql"
	"fmt"
)

// Setup connects with database and returns database object
func Setup(username string, password string, server string, databaseName string) (store Store) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, server, databaseName)
	db, errSetup := sql.Open("mysql", dataSourceName)
	if errSetup != nil {
		panic(errSetup)
	}
	// Check if Shortcut table exist
	_, errTest := db.Query("SELECT * FROM Shortcuts")
	if errTest != nil {
		// If we get an error, that can either mean that the table "Shortcuts" is yet to be created or that we do not have a connection.
		_, errCreate := db.Exec("CREATE TABLE Shortcuts (path VARCHAR(50) NOT NULL PRIMARY KEY, url TEXT NOT NULL)")
		if errCreate != nil {
			// Could not create table, aborting.
			panic(errCreate)
		}
	}
	return Store{
		db: db}
}
