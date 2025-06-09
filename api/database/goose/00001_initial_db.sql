-- +goose Up
-- +goose StatementBegin
-- Create family table
CREATE TABLE family (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    display_name TEXT
);

-- Create list table
CREATE TABLE list (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    uuid TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    display_name TEXT,
    family_id INTEGER NOT NULL,
    FOREIGN KEY (family_id) REFERENCES family(id)
);

-- Create item table
CREATE TABLE item (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    uuid TEXT NOT NULL UNIQUE,
    item TEXT NOT NULL,
    list_id INTEGER NOT NULL DEFAULT 0,
    UNIQUE(item, list_id)
);

-- Create indexes for better performance
CREATE INDEX idx_list_uuid ON list(uuid);
CREATE INDEX idx_list_family_id ON list(family_id);
CREATE INDEX idx_item_uuid ON item(uuid);
CREATE INDEX idx_item_list_id ON item(list_id);
CREATE INDEX idx_item_name_list ON item(item, list_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_item_name_list;
DROP INDEX IF EXISTS idx_item_list_id;
DROP INDEX IF EXISTS idx_item_uuid;
DROP INDEX IF EXISTS idx_list_family_id;
DROP INDEX IF EXISTS idx_list_uuid;
DROP INDEX IF EXISTS idx_family_uuid;
DROP TABLE IF EXISTS item;
DROP TABLE IF EXISTS list;
DROP TABLE IF EXISTS family;
-- +goose StatementEnd
