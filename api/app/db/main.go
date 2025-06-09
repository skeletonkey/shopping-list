package db

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/skeletonkey/lib-core-go/logger"
)

const (
	driverName  = "sqlite3"
	readTimeout = 5 * time.Second
)

var (
	dbConn *sql.DB
)

// New creates a singleton database connection
// The DB connection will be available internally, but should not be used outside
// of this package.
//
// Wait group is incremented appropriately
func New(ctx context.Context, wg *sync.WaitGroup) (err error) {
	if dbConn != nil {
		return err
	}

	dbCfg := getConfig()
	dbConn, err = sql.Open(driverName, dbCfg.DbFile)
	if err != nil {
		return fmt.Errorf("failed to open database (%s): %s", cfg.DbFile, err)
	}

	// Test the connection
	err = dbConn.Ping()
	if err != nil {
		dbConn.Close()
		dbConn = nil
		return fmt.Errorf("failed to ping database: %s", err)
	}

	// Enable foreign key constraints
	_, err = dbConn.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		dbConn.Close()
		dbConn = nil
		return fmt.Errorf("failed to enable foreign keys: %s", err)
	}

	// config.RegisterInitializer("db", dbCfg)

	wg.Add(1)
	context.AfterFunc(ctx, func() {
		defer wg.Done()
		log := logger.Get()
		// log.Trace().Msg("DB going to sleep for 10 seconds")
		// time.Sleep(10 * time.Second) // Allow some time for any pending operations to complete
		if dbConn != nil {
			if err := dbConn.Close(); err != nil {
				log.Error().Err(err).Msg("failed to close database connection on context cancel")
			} else {
				log.Info().Msg("Database connection closed on context cancel")
			}
		}
	})

	return err
}

// GetFamilyByName retrieves a family by UUID
func GetFamilyByName(ctx context.Context, name string) (*Family, error) {
	query := `SELECT id, display_name FROM family WHERE name = ?`
	readCtx, readCancel := context.WithTimeout(ctx, readTimeout)
	defer readCancel()
	row := dbConn.QueryRowContext(readCtx, query, name)

	family := &Family{Name: name}
	if err := row.Scan(&family.ID, &family.DisplayName); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("family (%s) not found", name)
		}
		return nil, fmt.Errorf("failed to get family by name (%s): %s", name, err)
	}

	return family, nil
}

// GetAllFamilies retrieves all families
func GetAllFamilies(ctx context.Context) ([]*Family, error) {
	query := `SELECT id, name, display_name FROM family ORDER BY name`
	readCtx, readCancel := context.WithTimeout(ctx, readTimeout)
	defer readCancel()
	rows, err := dbConn.QueryContext(readCtx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query families: %w", err)
	}
	defer rows.Close()

	var families []*Family
	for rows.Next() {
		family := &Family{}
		err := rows.Scan(&family.ID, &family.Name, &family.DisplayName)
		if err != nil {
			return nil, fmt.Errorf("failed to scan family: %w", err)
		}
		families = append(families, family)
	}

	return families, nil
}

// CreateList creates a new list record
func CreateList(ctx context.Context, list *List) error {
	query := `INSERT INTO list (uuid, name, display_name, family_id) VALUES (?, ?, ?, ?)`
	result, err := dbConn.ExecContext(ctx, query, list.UUID, list.Name, list.DisplayName, list.FamilyID)
	if err != nil {
		return fmt.Errorf("failed to create list: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get list ID: %w", err)
	}

	list.ID = int(id)
	return nil
}

// GetListsByFamilyID retrieves all lists for a family
func GetListsByFamilyID(ctx context.Context, familyID int) ([]*List, error) {
	query := `SELECT id, uuid, name, display_name, family_id FROM list WHERE family_id = ? ORDER BY name`
	readCtx, readCancel := context.WithTimeout(ctx, readTimeout)
	defer readCancel()
	rows, err := dbConn.QueryContext(readCtx, query, familyID)
	if err != nil {
		return nil, fmt.Errorf("failed to query lists: %w", err)
	}
	defer rows.Close()

	var lists []*List
	for rows.Next() {
		list := &List{}
		err := rows.Scan(&list.ID, &list.UUID, &list.Name, &list.DisplayName, &list.FamilyID)
		if err != nil {
			return nil, fmt.Errorf("failed to scan list: %w", err)
		}
		lists = append(lists, list)
	}

	return lists, nil
}

// CreateItem creates a new item record
func CreateItem(ctx context.Context, item *Item) error {
	query := `INSERT INTO item (uuid, item, list_id) VALUES (?, ?, ?)`
	result, err := dbConn.ExecContext(ctx, query, item.UUID, item.Item, item.ListID)
	if err != nil {
		return fmt.Errorf("failed to create item: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get item ID: %w", err)
	}

	item.ID = int(id)
	return nil
}

// GetItemByUUID retrieves an item by UUID
func GetItemByUUID(ctx context.Context, uuid string) (*Item, error) {
	query := `SELECT id, uuid, item, list_id FROM item WHERE uuid = ?`
	row := dbConn.QueryRowContext(ctx, query, uuid)

	item := &Item{}
	if err := row.Scan(&item.ID, &item.UUID, &item.Item, &item.ListID); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("item not found")
		}
		return nil, fmt.Errorf("failed to get item: %w", err)
	}

	return item, nil
}

// GetItemsByListID retrieves all items for a list
func GetItemsByListID(ctx context.Context, listID int) ([]*Item, error) {
	query := `SELECT id, uuid, item, list_id FROM item WHERE list_id = ? ORDER BY item`
	rows, err := dbConn.QueryContext(ctx, query, listID)
	if err != nil {
		return nil, fmt.Errorf("failed to query items: %w", err)
	}
	defer rows.Close()

	var items []*Item
	for rows.Next() {
		item := &Item{}
		err := rows.Scan(&item.ID, &item.UUID, &item.Item, &item.ListID)
		if err != nil {
			return nil, fmt.Errorf("failed to scan item: %w", err)
		}
		items = append(items, item)
	}

	return items, nil
}

// UpdateItem updates an existing item
func UpdateItem(ctx context.Context, item *Item) error {
	query := `UPDATE item SET item = ?, list_id = ? WHERE uuid = ?`
	result, err := dbConn.ExecContext(ctx, query, item.Item, item.ListID, item.UUID)
	if err != nil {
		return fmt.Errorf("failed to update item: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("item not found")
	}

	return nil
}

// DeleteItem deletes an item by UUID
func DeleteItem(ctx context.Context, uuid string) error {
	query := `DELETE FROM item WHERE uuid = ?`
	result, err := dbConn.ExecContext(ctx, query, uuid)
	if err != nil {
		return fmt.Errorf("failed to delete item: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("item not found")
	}

	return nil
}

// DeleteList deletes a list by UUID
func DeleteList(ctx context.Context, uuid string) error {
	query := `DELETE FROM list WHERE uuid = ?`
	result, err := dbConn.ExecContext(ctx, query, uuid)
	if err != nil {
		return fmt.Errorf("failed to delete list: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("list not found")
	}

	return nil
}
