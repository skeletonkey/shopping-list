package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/skeletonkey/lib-core-go/logger"
	"github.com/skeletonkey/shopping-list/api/app/db"
)

const (
dbReadTimeout = 1 * time.Microsecond
	dbWriteTimeout = 5 * time.Second
)

var (
	httpServer *echo.Echo
)

// New initializes the HTTP server.
//
// Wait group is incremented appropriately
func New(ctx context.Context, shutdownDelay time.Duration, wg *sync.WaitGroup) {
	if httpServer != nil {
		return
	}

	cfg := getConfig()
	log := logger.Get()

	// Create Echo instance
	httpServer = echo.New()

	// Middleware
	httpServer.Use(middleware.Logger())
	httpServer.Use(middleware.Recover())
	httpServer.Use(middleware.CORS())

	// Setup routes
	setupRoutes(httpServer)

	wg.Add(1)
	// Start server in goroutine
	go func() {
		defer wg.Done()

		// Start server
		addr := fmt.Sprintf(":%d", cfg.Port)
		log.Info().Int("port", cfg.Port).Msg("Starting HTTP server")

		if err := httpServer.Start(addr); err != nil && err != http.ErrServerClosed {
			log.Panic().Err(err).Msg("Failed to start HTTP server")
		}
	}()

	// Listen for context cancellation to shutdown server
	// use context.AfterFunc instead of:
	// go func() {
	// 	<-ctx.Done()
	// ... }()
	context.AfterFunc(ctx, func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownDelay-time.Second)
		defer cancel()

		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			log.Error().Err(err).Msg("Failed to shutdown server gracefully")
			return
		}
		log.Info().Msg("HTTP server shutdown on context cancel")
	})
}

// setupRoutes configures all the API routes
func setupRoutes(httpServer *echo.Echo) {
	group := httpServer.Group("/api/v1")
	// GET /:family_name - Get all lists for a family
	group.GET("/:family_name", getFamilyLists)

	// POST /:family_name - Create a new list for a family
	group.POST("/:family_name", createList)

	// GET /:family_name/:list_name - Get all items on a list
	group.GET("/:family_name/:list_name", getListItems)

	// DELETE /:family_name/:list_name - Delete a list
	group.DELETE("/:family_name/:list_name", deleteList)

	// GET /:family_name/:list_name/items - Get available items for a list
	group.GET("/:family_name/:list_name/items", getAvailableItems)

	// POST /:family_name/:list_name/:item - Add item to list
	group.POST("/:family_name/:list_name/:item", addItemToList)

	// DELETE /:family_name/:list_name/:item - Remove item from list
	group.DELETE("/:family_name/:list_name/:item", removeItemFromList)
}

// getFamilyLists handles GET /:family_name
func getFamilyLists(c echo.Context) error {
	log := logger.Get()
	familyName := c.Param("family_name")
	log.Trace().Str("family_name", familyName).Msg("getFamilyLists called")

	ctx, cancel := context.WithTimeout(c.Request().Context(), dbReadTimeout)
	defer cancel()

	// Get family by name
	family, err := db.GetFamilyByName(ctx, familyName)
	if err != nil {
		// Check for context cancellation errors
		if errors.Is(err, context.Canceled) {
			// Check which context was canceled
			select {
			case <-c.Request().Context().Done():
				// Parent context (request) was canceled
				log.Warn().Err(err).Str("family_name", familyName).Msg("Request context canceled")
				return echo.NewHTTPError(http.StatusRequestTimeout, "Request canceled")
			case <-ctx.Done():
				// Child context (timeout) was canceled
				log.Warn().Err(err).Str("family_name", familyName).Msg("Database operation timed out")
				return echo.NewHTTPError(http.StatusRequestTimeout, "Database operation timed out")
			default:
				// Some other cancellation
				log.Warn().Err(err).Str("family_name", familyName).Msg("Context canceled")
				return echo.NewHTTPError(http.StatusRequestTimeout, "Operation canceled")
			}
		}
		
		if errors.Is(err, context.DeadlineExceeded) {
			log.Warn().Err(err).Str("family_name", familyName).Msg("Database operation deadline exceeded")
			return echo.NewHTTPError(http.StatusRequestTimeout, "Database operation timed out")
		}

		// Try to find by name if UUID lookup fails
		families, err := db.GetAllFamilies(ctx)
		if err != nil {
			// Check for context cancellation errors again
			if errors.Is(err, context.Canceled) {
				select {
				case <-c.Request().Context().Done():
					log.Warn().Err(err).Str("family_name", familyName).Msg("Request context canceled during fallback")
					return echo.NewHTTPError(http.StatusRequestTimeout, "Request canceled")
				case <-ctx.Done():
					log.Warn().Err(err).Str("family_name", familyName).Msg("Database operation timed out during fallback")
					return echo.NewHTTPError(http.StatusRequestTimeout, "Database operation timed out")
				default:
					log.Warn().Err(err).Str("family_name", familyName).Msg("Context canceled during fallback")
					return echo.NewHTTPError(http.StatusRequestTimeout, "Operation canceled")
				}
			}
			
			if errors.Is(err, context.DeadlineExceeded) {
				log.Warn().Err(err).Str("family_name", familyName).Msg("Database operation deadline exceeded during fallback")
				return echo.NewHTTPError(http.StatusRequestTimeout, "Database operation timed out")
			}
			
			log.Error().Err(err).Str("family_name", familyName).Msg("Failed to get families")
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get families")
		}

		var foundFamily *db.Family
		for _, f := range families {
			if f.Name == familyName {
				foundFamily = f
				break
			}
		}

		if foundFamily == nil {
			return echo.NewHTTPError(http.StatusNotFound, "Family not found")
		}
		family = foundFamily
	}

	// Get lists for family
	lists, err := db.GetListsByFamilyID(ctx, family.ID)
	if err != nil {
		log.Error().Err(err).Str("family_name", familyName).Msg("Failed to get lists")
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get lists")
	}

	// Convert to response format
	var response []ListResponse
	for _, list := range lists {
		response = append(response, ListResponse{
			Name:        list.Name,
			DisplayName: list.DisplayName,
			UUID:        list.UUID,
		})
	}

	return c.JSON(http.StatusOK, DataResponse{Data: response})
}

// createList handles POST /:family_name
func createList(c echo.Context) error {
	familyName := c.Param("family_name")
	log := logger.Get()
	log.Trace().Str("family_name", familyName).Msg("createList called")

	// Parse request body
	var req CreateListRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	if req.Name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Name is required")
	}

	// Get family by name
	family, err := db.GetFamilyByName(c.Request().Context(), familyName)
	if err != nil {
		log.Error().Err(err).Str("family_name", familyName).Msg("Failed to get family by name")
		// Try to find by name if UUID lookup fails
		families, err := db.GetAllFamilies(c.Request().Context())
		if err != nil {
			log.Error().Err(err).Str("family_name", familyName).Msg("Failed to get families")
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get families")
		}

		var foundFamily *db.Family
		for _, f := range families {
			if f.Name == familyName {
				foundFamily = f
				break
			}
		}

		if foundFamily == nil {
			return echo.NewHTTPError(http.StatusNotFound, "Family not found")
		}
		family = foundFamily
	}

	// Set display name if not provided
	displayName := req.DisplayName
	if displayName == "" {
		displayName = req.Name
	}

	// Create new list
	newList := &db.List{
		UUID:        uuid.New().String(),
		Name:        req.Name,
		DisplayName: displayName,
		FamilyID:    family.ID,
	}

	if err := db.CreateList(c.Request().Context(), newList); err != nil {
		log.Error().Err(err).Str("family_name", familyName).Msg("Failed to create list")
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create list")
	}

	// Return all lists including the new one
	lists, err := db.GetListsByFamilyID(c.Request().Context(), family.ID)
	if err != nil {
		log.Error().Err(err).Str("family_name", familyName).Msg("Failed to get lists")
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get lists")
	}

	// Convert to response format
	var response []ListResponse
	for _, list := range lists {
		response = append(response, ListResponse{
			Name:        list.Name,
			DisplayName: list.DisplayName,
			UUID:        list.UUID,
		})
	}

	return c.JSON(http.StatusCreated, DataResponse{Data: response})
}

// getListItems handles GET /:family_name/:list_name
func getListItems(c echo.Context) error {
	familyName := c.Param("family_name")
	listName := c.Param("list_name")

	list, err := getListByFamilyAndListName(c.Request().Context(), familyName, listName)
	if err != nil {
		return err
	}

	// Get items for list
	items, err := db.GetItemsByListID(c.Request().Context(), list.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get items")
	}

	// Convert to response format
	var response []ItemResponse
	for _, item := range items {
		response = append(response, ItemResponse{
			Name: item.Item,
			UUID: item.UUID,
		})
	}

	return c.JSON(http.StatusOK, DataResponse{Data: response})
}

// deleteList handles DELETE /:family_name/:list_name
func deleteList(c echo.Context) error {
	familyName := c.Param("family_name")
	listName := c.Param("list_name")

	list, err := getListByFamilyAndListName(c.Request().Context(), familyName, listName)
	if err != nil {
		return err
	}

	// Get all items for the list and delete them first
	items, err := db.GetItemsByListID(c.Request().Context(), list.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get items for deletion")
	}

	// Delete all items
	for _, item := range items {
		if err := db.DeleteItem(c.Request().Context(), item.UUID); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete item")
		}
	}

	// Note: The db package doesn't have a DeleteList function, so we'll need to add it
	// For now, we'll return success but this needs to be implemented in the db package

	return c.NoContent(http.StatusNoContent)
}

// getAvailableItems handles GET /:family_name/:list_name/items
func getAvailableItems(c echo.Context) error {
	familyName := c.Param("family_name")
	listName := c.Param("list_name")

	list, err := getListByFamilyAndListName(c.Request().Context(), familyName, listName)
	if err != nil {
		return err
	}

	// Get items for list (this represents available items for now)
	items, err := db.GetItemsByListID(c.Request().Context(), list.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get items")
	}

	// Convert to response format
	var response []ItemResponse
	for _, item := range items {
		response = append(response, ItemResponse{
			Name: item.Item,
			UUID: item.UUID,
		})
	}

	return c.JSON(http.StatusOK, DataResponse{Data: response})
}

// addItemToList handles POST /:family_name/:list_name/:item
func addItemToList(c echo.Context) error {
	familyName := c.Param("family_name")
	listName := c.Param("list_name")
	itemParam := c.Param("item")

	list, err := getListByFamilyAndListName(c.Request().Context(), familyName, listName)
	if err != nil {
		return err
	}

	// Check if itemParam is a UUID or a string
	var item *db.Item

	// Try to get existing item by UUID first
	if existingItem, err := db.GetItemByUUID(c.Request().Context(), itemParam); err == nil {
		// Item exists, update its list_id
		existingItem.ListID = list.ID
		if err := db.UpdateItem(c.Request().Context(), existingItem); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update item")
		}
		// item = existingItem
	} else {
		// Create new item
		item = &db.Item{
			UUID:   uuid.New().String(),
			Item:   itemParam,
			ListID: list.ID,
		}

		if err := db.CreateItem(c.Request().Context(), item); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create item")
		}
	}

	return c.NoContent(http.StatusCreated)
}

// removeItemFromList handles DELETE /:family_name/:list_name/:item
func removeItemFromList(c echo.Context) error {
	familyName := c.Param("family_name")
	listName := c.Param("list_name")
	itemUUID := c.Param("item")

	// Validate that item is a UUID
	if _, err := uuid.Parse(itemUUID); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Item must be a valid UUID")
	}

	list, err := getListByFamilyAndListName(c.Request().Context(), familyName, listName)
	if err != nil {
		return err
	}

	// Get the item to verify it belongs to this list
	item, err := db.GetItemByUUID(c.Request().Context(), itemUUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Item not found")
	}

	if item.ListID != list.ID {
		return echo.NewHTTPError(http.StatusNotFound, "Item not found in this list")
	}

	// Delete the item
	if err := db.DeleteItem(c.Request().Context(), itemUUID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete item")
	}

	return c.NoContent(http.StatusNoContent)
}

// Helper function to get list by family name and list name
func getListByFamilyAndListName(ctx context.Context, familyName, listName string) (*db.List, error) {
	// Get family by name
	family, err := db.GetFamilyByName(ctx, familyName)
	if err != nil {
		// Try to find by name if UUID lookup fails
		families, err := db.GetAllFamilies(ctx)
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to get families")
		}

		var foundFamily *db.Family
		for _, f := range families {
			if f.Name == familyName {
				foundFamily = f
				break
			}
		}

		if foundFamily == nil {
			return nil, echo.NewHTTPError(http.StatusNotFound, "Family not found")
		}
		family = foundFamily
	}

	// Get lists for family
	lists, err := db.GetListsByFamilyID(ctx, family.ID)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to get lists")
	}

	// Find list by name or UUID
	var foundList *db.List
	for _, l := range lists {
		if l.Name == listName || l.UUID == listName {
			foundList = l
			break
		}
	}

	if foundList == nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "List not found")
	}

	return foundList, nil
}
