package db

import (
	"context"
	"fmt"
)

type Family struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

// String returns the display name of the Family if it is set; otherwise, it returns the Family's name.
func (f *Family) String() string {
	if f.DisplayName != "" {
		return f.DisplayName
	}
	return f.Name
}

// create creates a new family record
func (f *Family) create(ctx context.Context) error {
	if f.ID != 0 {
		return fmt.Errorf("ID (%d) can not be provided to family create", f.ID)
	}
	if f.Name == "" {
		return fmt.Errorf("name must be provided to create a family")
	}

	if f.DisplayName == "" {
		f.DisplayName = f.Name
	}

	query := `INSERT INTO family (name, display_name) VALUES (?, ?)`
	result, err := dbConn.ExecContext(ctx, query, f.Name, f.DisplayName)
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

func (f *Family) update(cfg context.Context) error {
	query := `UPDATE family SET name = ?, display_name = ? WHERE id = ?`
	_, err := dbConn.ExecContext(cfg, query, f.Name, f.DisplayName, f.ID)
	if err != nil {
		return fmt.Errorf("failed to update family: %s", err)
	}

	return nil
}

// delete removes the family from the database.
func (f *Family) delete(ctx context.Context) error {
	query := `DELETE FROM family WHERE id = ?`
	_, err := dbConn.ExecContext(ctx, query, f.ID)
	if err != nil {
		return fmt.Errorf("failed to delete family: %s", err)
	}

	return nil
}

// removeFamily removes the family from the database; this includes deleting all lists and items associated with the family.
func (f *Family) removeFamily(ctx context.Context) error {
	if f.ID == 0 {
		return fmt.Errorf("ID must be provided to remove a family")
	}

	queryCtx, cancel := context.WithTimeout(ctx, readTimeout)
	defer cancel()

	lists, err := GetListsByFamilyID(queryCtx, f.ID)
	if err != nil {
		return fmt.Errorf("failed to get lists for family %d: %w", f.ID, err)
	}
	for _, list := range lists {
		items, err := GetItemsByListID(queryCtx, list.ID)
		if err != nil {
			return fmt.Errorf("failed to get items for list %d: %w", list.ID, err)
		}
		for _, item := range items {
			if err := DeleteItem(queryCtx, item.UUID); err != nil {
				return fmt.Errorf("failed to delete item %d: %w", item.ID, err)
			}
		}
		if err := DeleteList(queryCtx, list.UUID); err != nil {
			return fmt.Errorf("failed to delete list %d: %w", list.ID, err)
		}
	}
	if err := f.delete(queryCtx); err != nil {
		return fmt.Errorf("failed to delete family %d: %w", f.ID, err)
	}
	return nil
}
