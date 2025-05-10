package api

import (
	"net/http"
	"strconv"

	apperrors "github.com/arafat-hasan/mealsync/internal/errors"
	"github.com/arafat-hasan/mealsync/internal/middleware"
	"github.com/arafat-hasan/mealsync/internal/model"
	"github.com/arafat-hasan/mealsync/internal/service"
	"github.com/gin-gonic/gin"
)

// MenuItemHandler handles menu item-related requests
type MenuItemHandler struct {
	menuItemService service.MenuItemService
}

// NewMenuItemHandler creates a new MenuItemHandler
func NewMenuItemHandler(menuItemService service.MenuItemService) *MenuItemHandler {
	return &MenuItemHandler{menuItemService: menuItemService}
}

// GetMenuItems handles GET /api/menu-items
// @Summary      List menu items
// @Description  Get all menu items for the authenticated user
// @Tags         menu-items
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   model.MenuItem
// @Failure      401  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /menu-items [get]
func (h *MenuItemHandler) GetMenuItems(c *gin.Context) {
	items, err := h.menuItemService.GetMenuItems(c.Request.Context())
	if err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	c.JSON(http.StatusOK, items)
}

// GetMenuItemByID handles GET /api/menu-items/:id
// @Summary      Get menu item by ID
// @Description  Get a specific menu item by its ID
// @Tags         menu-items
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Menu Item ID"
// @Success      200  {object}  model.MenuItem
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      401  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /menu-items/{id} [get]
func (h *MenuItemHandler) GetMenuItemByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		middleware.HandleAppError(c, apperrors.NewValidationError("Invalid menu item ID", err))
		return
	}

	item, err := h.menuItemService.GetMenuItemByID(c.Request.Context(), uint(id))
	if err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	c.JSON(http.StatusOK, item)
}

// CreateMenuItem handles POST /api/menu-items
// @Summary      Create menu item
// @Description  Create a new menu item
// @Tags         menu-items
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        menuItem  body      model.MenuItem  true  "Menu Item object"
// @Success      201      {object}  model.MenuItem
// @Failure      400   {object}  dto.ErrorResponse
// @Failure      401   {object}  dto.ErrorResponse
// @Failure      500   {object}  dto.ErrorResponse
// @Router       /menu-items [post]
func (h *MenuItemHandler) CreateMenuItem(c *gin.Context) {
	var item model.MenuItem
	if err := c.ShouldBindJSON(&item); err != nil {
		middleware.HandleAppError(c, apperrors.NewValidationError("Invalid request format", err))
		return
	}

	userID := uint(1) // TODO: Get from context after auth

	if err := h.menuItemService.CreateMenuItem(c.Request.Context(), &item, userID); err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	c.JSON(http.StatusCreated, item)
}

// UpdateMenuItem handles PUT /api/menu-items/:id
// @Summary      Update menu item
// @Description  Update an existing menu item
// @Tags         menu-items
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id        path      int           true  "Menu Item ID"
// @Param        menuItem  body      model.MenuItem  true  "Menu Item object"
// @Success      200      {object}  model.MenuItem
// @Failure      400   {object}  dto.ErrorResponse
// @Failure      401   {object}  dto.ErrorResponse
// @Failure      403   {object}  dto.ErrorResponse
// @Failure      404   {object}  dto.ErrorResponse
// @Failure      500   {object}  dto.ErrorResponse
// @Router       /menu-items/{id} [put]
func (h *MenuItemHandler) UpdateMenuItem(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		middleware.HandleAppError(c, apperrors.NewValidationError("Invalid menu item ID", err))
		return
	}

	var item model.MenuItem
	if err := c.ShouldBindJSON(&item); err != nil {
		middleware.HandleAppError(c, apperrors.NewValidationError("Invalid request format", err))
		return
	}

	userID := uint(1) // TODO: Get from context after auth

	if err := h.menuItemService.UpdateMenuItem(c.Request.Context(), uint(id), &item, userID); err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	c.JSON(http.StatusOK, item)
}

// DeleteMenuItem handles DELETE /api/menu-items/:id
// @Summary      Delete menu item
// @Description  Delete an existing menu item
// @Tags         menu-items
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Menu Item ID"
// @Success      200  {object}  SuccessResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      401  {object}  dto.ErrorResponse
// @Failure      403  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /menu-items/{id} [delete]
func (h *MenuItemHandler) DeleteMenuItem(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		middleware.HandleAppError(c, apperrors.NewValidationError("Invalid menu item ID", err))
		return
	}

	userID := uint(1) // TODO: Get from context after auth

	if err := h.menuItemService.DeleteMenuItem(c.Request.Context(), uint(id), userID); err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Menu item deleted successfully"})
}

// GetMenuItemsByCategory godoc
// @Summary Get menu items by category
// @Description Retrieves all menu items belonging to a specific category
// @Tags menu-items
// @Accept json
// @Produce json
// @Param category path string true "Category name"
// @Success 200 {array} model.MenuItem
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Router /menu-items/category/{category} [get]
func (h *MenuItemHandler) GetMenuItemsByCategory(c *gin.Context) {
	category := c.Param("category")
	if category == "" {
		middleware.HandleAppError(c, apperrors.NewValidationError("Category is required", nil))
		return
	}

	items, err := h.menuItemService.GetMenuItemsByCategory(c.Request.Context(), category)
	if err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	c.JSON(http.StatusOK, items)
}

// GetMenuItemsByMenuSet godoc
// @Summary Get menu items by menu set
// @Description Retrieves all menu items belonging to a specific menu set
// @Tags menu-items
// @Accept json
// @Produce json
// @Param menu_set_id path string true "Menu Set ID"
// @Success 200 {array} model.MenuItem
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /menu-items/menu-set/{menu_set_id} [get]
func (h *MenuItemHandler) GetMenuItemsByMenuSet(c *gin.Context) {
	menuSetID, err := strconv.ParseUint(c.Param("menu_set_id"), 10, 32)
	if err != nil {
		middleware.HandleAppError(c, apperrors.NewValidationError("Invalid menu set ID", err))
		return
	}

	items, err := h.menuItemService.GetMenuItemsByMenuSet(c.Request.Context(), uint(menuSetID))
	if err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	c.JSON(http.StatusOK, items)
}
