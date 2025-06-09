package db

import (
	"database/sql"
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
		return fmt.Errorf("Name must be provided to create a family")
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

// getFamilyByID retrieves a family by ID
func getFamilyByID(id int) (Family, error) {
	query := `SELECT name, display_name FROM family WHERE id = ?`
	row := dbConn.QueryRow(query, id)

	family := &Family{ID: id}
	if err := row.Scan(&family.Name, &family.DisplayName); err != nil {
		if err == sql.ErrNoRows {
			return *family, fmt.Errorf("family not found by id")
		}
		return *family, fmt.Errorf("failed to get family by id: %s", err)
	}

	return *family, nil
}

// getFamilyByName retrieves a family by ID
func getFamilyByName(name string) (Family, error) {
	query := `SELECT id, display_name FROM family WHERE name = ?`
	row := dbConn.QueryRow(query, name)

	family := &Family{Name: name}
	if err := row.Scan(&family.ID, &family.DisplayName); err != nil {
		if err == sql.ErrNoRows {
			return *family, fmt.Errorf("family not found by name")
		}

		return *family, fmt.Errorf("failed to get family by name: %s", err)
	}

	return *family, nil
}

// validate ensure the validity of a family entry
// if an ID is provide then it must match what is in the database
// if ID is not provided then Name must be provided and exist in the DB
func (f *Family) validate() (err error) {
	var existingFamily Family
	if f.ID != 0 {
		existingFamily, err = getFamilyByID(f.ID)
		if err != nil {
			return fmt.Errorf("failed to validate existing family by ID (%d): %s", f.ID, err)
		}
	} else if f.Name != "" {
		existingFamily, _ = getFamilyByName(f.Name)
		if existingFamily.ID == 0 {
			return fmt.Errorf("Name (%s) already exist", f.Name)
		}
	}
	if existingFamily.ID != 0 {
		if existingFamily.ID != f.ID {
			return fmt.Errorf("provided ID (%d) does not match db ID (%d)", f.ID, existingFamily.ID)
		}

		return nil
	}

	if f.Name == "" {
		return fmt.Errorf("Name must be provided")
	}

	return nil
}
