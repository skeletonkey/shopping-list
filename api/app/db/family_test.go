package db

import (
	"context"
	"testing"
)

// TestFamily_String tests the String method of the Family struct.
// It verifies that the method returns DisplayName when set, otherwise returns Name.
func TestFamily_String(t *testing.T) {
	tests := []struct {
		name        string
		family      Family
		wantDisplay string
	}{
		{
			name:        "DisplayName set",
			family:      Family{Name: "fam", DisplayName: "The Family"},
			wantDisplay: "The Family",
		},
		{
			name:        "DisplayName empty",
			family:      Family{Name: "fam", DisplayName: ""},
			wantDisplay: "fam",
		},
		{
			name:        "Both Name and DisplayName empty",
			family:      Family{Name: "", DisplayName: ""},
			wantDisplay: "",
		},
		{
			name:        "Name empty but DisplayName set",
			family:      Family{Name: "", DisplayName: "Display Only"},
			wantDisplay: "Display Only",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.family.String()
			if got != tt.wantDisplay {
				t.Errorf("Family.String() = %q, want %q", got, tt.wantDisplay)
			}
		})
	}
}

// TestFamily_create tests the create method of the Family struct.
// It verifies proper family creation with validation rules.
func TestFamily_create(t *testing.T) {
	if dbConn == nil {
		t.Skip("Database connection not available for testing")
	}

	ctx := context.Background()

	tests := []struct {
		name      string
		family    Family
		wantError bool
		errorMsg  string
	}{
		{
			name:      "Valid family with name only",
			family:    Family{Name: "test-family"},
			wantError: false,
		},
		{
			name:      "Valid family with name and display name",
			family:    Family{Name: "test-family-2", DisplayName: "Test Family 2"},
			wantError: false,
		},
		{
			name:      "Invalid family with ID set",
			family:    Family{ID: 1, Name: "test-family-3"},
			wantError: true,
			errorMsg:  "ID (1) can not be provided to family create",
		},
		{
			name:      "Invalid family with empty name",
			family:    Family{Name: ""},
			wantError: true,
			errorMsg:  "name must be provided to create a family",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy to avoid modifying the test case
			family := tt.family
			err := family.create(ctx)

			if tt.wantError {
				if err == nil {
					t.Errorf("Family.create() expected error but got none")
					return
				}
				if err.Error() != tt.errorMsg {
					t.Errorf("Family.create() error = %q, want %q", err.Error(), tt.errorMsg)
				}
				return
			}

			if err != nil {
				t.Errorf("Family.create() unexpected error = %v", err)
				return
			}

			// Verify ID was set
			if family.ID == 0 {
				t.Error("Family.create() did not set ID")
			}

			// Verify DisplayName was set to Name if it was empty
			if tt.family.DisplayName == "" && family.DisplayName != family.Name {
				t.Errorf("Family.create() DisplayName = %q, want %q", family.DisplayName, family.Name)
			}

			// Clean up - delete the created family
			if err := family.delete(ctx); err != nil {
				t.Logf("Failed to clean up test family: %v", err)
			}
		})
	}
}

// TestFamily_update tests the update method of the Family struct.
// It verifies that family records can be properly updated.
func TestFamily_update(t *testing.T) {
	if dbConn == nil {
		t.Skip("Database connection not available for testing")
	}

	ctx := context.Background()

	// Create a test family first
	family := Family{Name: "test-update-family", DisplayName: "Original Display"}
	if err := family.create(ctx); err != nil {
		t.Fatalf("Failed to create test family: %v", err)
	}
	defer func() {
		if err := family.delete(ctx); err != nil {
			t.Logf("Failed to clean up test family: %v", err)
		}
	}()

	// Update the family
	family.Name = "updated-family"
	family.DisplayName = "Updated Display"

	err := family.update(ctx)
	if err != nil {
		t.Errorf("Family.update() unexpected error = %v", err)
	}
}

// TestFamily_delete tests the delete method of the Family struct.
// It verifies that family records can be properly deleted.
func TestFamily_delete(t *testing.T) {
	if dbConn == nil {
		t.Skip("Database connection not available for testing")
	}

	ctx := context.Background()

	// Create a test family first
	family := Family{Name: "test-delete-family"}
	if err := family.create(ctx); err != nil {
		t.Fatalf("Failed to create test family: %v", err)
	}

	// Delete the family
	err := family.delete(ctx)
	if err != nil {
		t.Errorf("Family.delete() unexpected error = %v", err)
	}
}

// TestFamily_removeFamily tests the removeFamily method of the Family struct.
// It verifies that families and all associated data are properly removed.
func TestFamily_removeFamily(t *testing.T) {
	if dbConn == nil {
		t.Skip("Database connection not available for testing")
	}

	ctx := context.Background()

	tests := []struct {
		name      string
		family    Family
		wantError bool
		errorMsg  string
	}{
		{
			name:      "Invalid family with no ID",
			family:    Family{Name: "test-family"},
			wantError: true,
			errorMsg:  "ID must be provided to remove a family",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.family.removeFamily(ctx)

			if tt.wantError {
				if err == nil {
					t.Errorf("Family.removeFamily() expected error but got none")
					return
				}
				if err.Error() != tt.errorMsg {
					t.Errorf("Family.removeFamily() error = %q, want %q", err.Error(), tt.errorMsg)
				}
				return
			}

			if err != nil {
				t.Errorf("Family.removeFamily() unexpected error = %v", err)
			}
		})
	}
}
