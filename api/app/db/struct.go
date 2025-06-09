package db

// List represents a list record
type List struct {
	ID          int    `json:"id"`
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	FamilyID    int    `json:"family_id"`
}

// Item represents an item record
type Item struct {
	ID     int    `json:"id"`
	UUID   string `json:"uuid"`
	Item   string `json:"item"`
	ListID int    `json:"list_id"`
}
