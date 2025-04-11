package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/arafat-hasan/mealsync/internal/model"
	"github.com/arafat-hasan/mealsync/internal/service"
	"github.com/gin-gonic/gin"
)

// MenuHandler handles menu-related HTTP requests
type MenuHandler struct {
	menuService service.MenuService
}

// NewMenuHandler creates a new MenuHandler
func NewMenuHandler(menuService service.MenuService) *MenuHandler {
	return &MenuHandler{
		menuService: menuService,
	}
}

// GetMenus godoc
// @Summary Get all menus
// @Description Get all menus, optionally filtered by date
// @Tags menus
// @Accept json
// @Produce json
// @Param date query string false "Filter by date (YYYY-MM-DD)"
// @Success 200 {array} model.MealMenu
// @Router /menus [get]
// @Security BearerAuth
func (h *MenuHandler) GetMenus(c *gin.Context) {
	dateStr := c.Query("date")
	var date *time.Time

	if dateStr != "" {
		parsedDate, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
			return
		}
		date = &parsedDate
	}

	menus, err := h.menuService.GetMenus(c.Request.Context(), date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get menus"})
		return
	}

	c.JSON(http.StatusOK, menus)
}

// GetMenuByID godoc
// @Summary Get a menu by ID
// @Description Get a menu by its ID
// @Tags menus
// @Accept json
// @Produce json
// @Param id path int true "Menu ID"
// @Success 200 {object} model.MealMenu
// @Failure 404 {object} map[string]string
// @Router /menus/{id} [get]
// @Security BearerAuth
func (h *MenuHandler) GetMenuByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu ID"})
		return
	}

	menu, err := h.menuService.GetMenuByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
		return
	}

	c.JSON(http.StatusOK, menu)
}

// CreateMenu godoc
// @Summary Create a new menu
// @Description Create a new menu
// @Tags menus
// @Accept json
// @Produce json
// @Param menu body model.MealMenu true "Menu to create"
// @Success 201 {object} model.MealMenu
// @Failure 400 {object} map[string]string
// @Router /menus [post]
// @Security BearerAuth
func (h *MenuHandler) CreateMenu(c *gin.Context) {
	var menu model.MealMenu
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu data"})
		return
	}

	// Set the creator to the current user
	userID, _ := c.Get("user_id")
	menu.CreatedBy = userID.(uint)

	if err := h.menuService.CreateMenu(c.Request.Context(), &menu); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create menu"})
		return
	}

	c.JSON(http.StatusCreated, menu)
}

// UpdateMenu godoc
// @Summary Update a menu
// @Description Update an existing menu
// @Tags menus
// @Accept json
// @Produce json
// @Param id path int true "Menu ID"
// @Param menu body model.MealMenu true "Menu to update"
// @Success 200 {object} model.MealMenu
// @Failure 400,404 {object} map[string]string
// @Router /menus/{id} [put]
// @Security BearerAuth
func (h *MenuHandler) UpdateMenu(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu ID"})
		return
	}

	var menu model.MealMenu
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu data"})
		return
	}

	menu.ID = uint(id)

	if err := h.menuService.UpdateMenu(c.Request.Context(), &menu); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update menu"})
		return
	}

	c.JSON(http.StatusOK, menu)
}

// DeleteMenu godoc
// @Summary Delete a menu
// @Description Delete a menu by its ID
// @Tags menus
// @Accept json
// @Produce json
// @Param id path int true "Menu ID"
// @Success 204 "No Content"
// @Failure 400,404 {object} map[string]string
// @Router /menus/{id} [delete]
// @Security BearerAuth
func (h *MenuHandler) DeleteMenu(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu ID"})
		return
	}

	if err := h.menuService.DeleteMenu(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete menu"})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetMenuItems godoc
// @Summary Get menu items
// @Description Get all menu items for a specific menu
// @Tags menus
// @Accept json
// @Produce json
// @Param id path int true "Menu ID"
// @Param date query string false "Filter by date (YYYY-MM-DD)"
// @Success 200 {array} model.MenuItem
// @Failure 400,404 {object} map[string]string
// @Router /menus/{id}/items [get]
// @Security BearerAuth
func (h *MenuHandler) GetMenuItems(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu ID"})
		return
	}

	dateStr := c.Query("date")
	var date *time.Time

	if dateStr != "" {
		parsedDate, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
			return
		}
		date = &parsedDate
	}

	items, err := h.menuService.GetMenuItems(c.Request.Context(), uint(id), date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get menu items"})
		return
	}

	c.JSON(http.StatusOK, items)
}

// AddMenuItem godoc
// @Summary Add a menu item to a menu
// @Description Add a menu item to a menu
// @Tags menus
// @Accept json
// @Produce json
// @Param id path int true "Menu ID"
// @Param item body map[string]interface{} true "Menu item to add"
// @Success 201 {object} model.MealMenuItem
// @Failure 400,404 {object} map[string]string
// @Router /menus/{id}/items [post]
// @Security BearerAuth
func (h *MenuHandler) AddMenuItem(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu ID"})
		return
	}

	var request struct {
		MenuItemID uint   `json:"menu_item_id" binding:"required"`
		SetName    string `json:"set_name"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if err := h.menuService.AddMenuItem(c.Request.Context(), uint(id), request.MenuItemID, request.SetName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add menu item"})
		return
	}

	c.Status(http.StatusCreated)
}

// RemoveMenuItem godoc
// @Summary Remove a menu item from a menu
// @Description Remove a menu item from a menu
// @Tags menus
// @Accept json
// @Produce json
// @Param id path int true "Menu ID"
// @Param item_id path int true "Menu Item ID"
// @Success 204 "No Content"
// @Failure 400,404 {object} map[string]string
// @Router /menus/{id}/items/{item_id} [delete]
// @Security BearerAuth
func (h *MenuHandler) RemoveMenuItem(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu ID"})
		return
	}

	itemIDStr := c.Param("item_id")
	itemID, err := strconv.ParseUint(itemIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu item ID"})
		return
	}

	if err := h.menuService.RemoveMenuItem(c.Request.Context(), uint(id), uint(itemID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove menu item"})
		return
	}

	c.Status(http.StatusNoContent)
}
