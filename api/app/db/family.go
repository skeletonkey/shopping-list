package db

import (
	"fmt"
)

type Family struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

// create creates a new family record
func (f *Family) create() error {
	if f.ID != 0 {
		return fmt.Errorf("ID (%d) can not be provided to family create", f.ID)
	}
	if f.Name == "" {
		return fmt.Errorf("name must be provided to create a family")
	}

	if f.DisplayName == "" {
		f.DisplayName = f.Name
	}

	query := `INSERT INTO family (name, display_name) VALUES (?, ?, ?)`
	result, err := dbConn.Exec(query, f.Name, f.DisplayName)
	if err != nil {
		return fmt.Errorf("failed to create family: %s", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get family ID: %s", err)
	}

	f.ID = int(id)

	return nil
}

func (f *Family) update() error {
	query := `UPDATE family SET name = ?, display_name = ? WHERE id = ?`
	_, err := dbConn.Exec(query, f.Name, f.DisplayName)
	if err != nil {
		return fmt.Errorf("failed to update family: %s", err)
	}

	return nil
}
