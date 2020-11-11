package store

import (
	"database/sql"

	"firlus.dev/firl.us/internal/model"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

// Store can manipulate shortcuts in the physical database
type Store struct {
	db *sql.DB
}

// CreateShortcut puts a new shortcut into the database
func (store *Store) CreateShortcut(shortcut *model.Shortcut) error {
	stmt, err := store.db.Prepare("INSERT INTO Shortcuts (path, url) VALUES (?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(shortcut.Path, shortcut.URL)
	if err != nil {
		return err
	}
	return nil
}

// ReadShortcut reads a shortcut from the database
func (store *Store) ReadShortcut(path string) (*model.Shortcut, error) {
	shortcut := model.Shortcut{}
	err := store.db.QueryRow("SELECT url FROM Shortcuts WHERE path = ?", path).Scan(&shortcut.URL)
	shortcut.Path = path
	return &shortcut, err
}

// ShortcutExists checks, whether a given shortcut exists
func (store *Store) ShortcutExists(path string) bool {
	shortcut, _ := store.ReadShortcut(path)
	return shortcut.IsValid()
}

// UpdateShortcut updates an existing shortcut in the database
func (store *Store) UpdateShortcut(shortcut *model.Shortcut) error {
	stmt, err := store.db.Prepare("UPDATE Shortcuts SET path=?, url=? WHERE path=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(shortcut.Path, shortcut.URL, shortcut.Path)
	if err != nil {
		return err
	}
	return nil
}

/*
func (store *Store) DeleteShortcut(path string) error {

}*/
