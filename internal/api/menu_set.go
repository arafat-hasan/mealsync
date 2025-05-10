package api

import (
	"net/http"
	"strconv"

	"github.com/arafat-hasan/mealsync/internal/dto"
	"github.com/arafat-hasan/mealsync/internal/model"
	"github.com/arafat-hasan/mealsync/internal/service"
	"github.com/gin-gonic/gin"
)

// MenuSetHandler handles menu set-related requests
type MenuSetHandler struct {
	menuSetService service.MenuSetService
}

// NewMenuSetHandler creates a new MenuSetHandler
func NewMenuSetHandler(menuSetService service.MenuSetService) *MenuSetHandler {
	return &MenuSetHandler{menuSetService: menuSetService}
}

// GetMenuSets handles GET /api/menu-sets
// @Summary      List menu sets
// @Description  Get all menu sets for the authenticated user
// @Tags         menu-sets
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   model.MenuSet
// @Failure      401  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /menu-sets [get]
func (h *MenuSetHandler) GetMenuSets(c *gin.Context) {
	menuSets, err := h.menuSetService.GetMenuSets(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, menuSets)
}

// GetMenuSetByID handles GET /api/menu-sets/:id
// @Summary      Get menu set by ID
// @Description  Get a specific menu set by its ID
// @Tags         menu-sets
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Menu Set ID"
// @Success      200  {object}  model.MenuSet
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      401  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /menu-sets/{id} [get]
func (h *MenuSetHandler) GetMenuSetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid menu set ID"})
		return
	}

	menuSet, err := h.menuSetService.GetMenuSetByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, menuSet)
}

// CreateMenuSet handles POST /api/menu-sets
// @Summary      Create menu set
// @Description  Create a new menu set
// @Tags         menu-sets
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        menuSet  body      model.MenuSet  true  "Menu Set Data"
// @Success      201     {object}  model.MenuSet
// @Failure      400     {object}  dto.ErrorResponse
// @Failure      401     {object}  dto.ErrorResponse
// @Failure      500     {object}  dto.ErrorResponse
// @Router       /menu-sets [post]
func (h *MenuSetHandler) CreateMenuSet(c *gin.Context) {
	var menuSet model.MenuSet
	if err := c.ShouldBindJSON(&menuSet); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request format"})
		return
	}

	userID := uint(1) // TODO: Get from context after auth

	if err := h.menuSetService.CreateMenuSet(c.Request.Context(), &menuSet, userID); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, menuSet)
}

// UpdateMenuSet handles PUT /api/menu-sets/:id
// @Summary      Update menu set
// @Description  Update an existing menu set
// @Tags         menu-sets
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id       path      int          true  "Menu Set ID"
// @Param        menuSet  body      model.MenuSet true  "Menu Set Data"
// @Success      200     {object}  model.MenuSet
// @Failure      400     {object}  dto.ErrorResponse
// @Failure      401     {object}  dto.ErrorResponse
// @Failure      403     {object}  dto.ErrorResponse
// @Failure      404     {object}  dto.ErrorResponse
// @Failure      500     {object}  dto.ErrorResponse
// @Router       /menu-sets/{id} [put]
func (h *MenuSetHandler) UpdateMenuSet(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid menu set ID"})
		return
	}

	var menuSet model.MenuSet
	if err := c.ShouldBindJSON(&menuSet); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request format"})
		return
	}

	userID := uint(1) // TODO: Get from context after auth

	if err := h.menuSetService.UpdateMenuSet(c.Request.Context(), uint(id), &menuSet, userID); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, menuSet)
}

// DeleteMenuSet handles DELETE /api/menu-sets/:id
// @Summary      Delete menu set
// @Description  Delete an existing menu set
// @Tags         menu-sets
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Menu Set ID"
// @Success      200  {object}  SuccessResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      401  {object}  dto.ErrorResponse
// @Failure      403  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /menu-sets/{id} [delete]
func (h *MenuSetHandler) DeleteMenuSet(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid menu set ID"})
		return
	}

	userID := uint(1) // TODO: Get from context after auth

	if err := h.menuSetService.DeleteMenuSet(c.Request.Context(), uint(id), userID); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{Message: "Menu set deleted successfully"})
}

// GetMenuSetItems handles GET /api/menus/:id/items
// @Summary      List menu set items
// @Description  Get all items in a specific menu set
// @Tags         menu-sets
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Menu Set ID"
// @Success      200  {array}   model.MenuItem
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      401  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /menu-sets/{id}/items [get]
func (h *MenuSetHandler) GetMenuSetItems(c *gin.Context) {
	menuSetID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid menu set ID"})
		return
	}

	items, err := h.menuSetService.GetMenuSetItems(c.Request.Context(), uint(menuSetID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

// AddMenuItemToMenuSet handles POST /api/menus/:id/items
// @Summary      Add menu item to set
// @Description  Add a menu item to an existing menu set
// @Tags         menu-sets
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id         path      int              true  "Menu Set ID"
// @Param        menuItem   body      model.MenuSetItem true  "Menu Item Data"
// @Success      201       {object}  model.MenuSetItem
// @Failure      400       {object}  dto.ErrorResponse
// @Failure      401       {object}  dto.ErrorResponse
// @Failure      403       {object}  dto.ErrorResponse
// @Failure      404       {object}  dto.ErrorResponse
// @Failure      500       {object}  dto.ErrorResponse
// @Router       /menu-sets/{id}/items [post]
func (h *MenuSetHandler) AddMenuItemToMenuSet(c *gin.Context) {
	menuSetID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid menu set ID"})
		return
	}

	var itemID uint
	if err := c.ShouldBindJSON(&itemID); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request format"})
		return
	}

	userID := uint(1) // TODO: Get from context after auth

	if err := h.menuSetService.AddItemToMenuSet(c.Request.Context(), uint(menuSetID), itemID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{Message: "Item added to menu set successfully"})
}

// RemoveMenuItemFromMenuSet handles DELETE /api/menus/:id/items/:item_id
// @Summary      Remove menu item from set
// @Description  Remove a menu item from an existing menu set
// @Tags         menu-sets
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id       path      int  true  "Menu Set ID"
// @Param        item_id  path      int  true  "Menu Item ID"
// @Success      200     {object}  SuccessResponse
// @Failure      400     {object}  dto.ErrorResponse
// @Failure      401     {object}  dto.ErrorResponse
// @Failure      403     {object}  dto.ErrorResponse
// @Failure      404     {object}  dto.ErrorResponse
// @Failure      500     {object}  dto.ErrorResponse
// @Router       /menu-sets/{id}/items/{item_id} [delete]
func (h *MenuSetHandler) RemoveMenuItemFromMenuSet(c *gin.Context) {
	menuSetID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid menu set ID"})
		return
	}

	itemID, err := strconv.ParseUint(c.Param("item_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid item ID"})
		return
	}

	userID := uint(1) // TODO: Get from context after auth

	if err := h.menuSetService.RemoveItemFromMenuSet(c.Request.Context(), uint(menuSetID), uint(itemID), userID); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{Message: "Item removed from menu set successfully"})
}
