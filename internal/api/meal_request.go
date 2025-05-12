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

// MealRequestHandler handles meal request-related requests
type MealRequestHandler struct {
	mealRequestService service.MealRequestService
}

// NewMealRequestHandler creates a new MealRequestHandler
func NewMealRequestHandler(mealRequestService service.MealRequestService) *MealRequestHandler {
	return &MealRequestHandler{mealRequestService: mealRequestService}
}

// GetMealRequests handles GET /api/meal-requests
//	@Summary		List meal requests
//	@Description	Get all meal requests for the authenticated user
//	@Tags			meal-requests
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{array}		model.MealRequest
//	@Failure		401	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/meal-requests [get]
func (h *MealRequestHandler) GetMealRequests(c *gin.Context) {
	userID := uint(1) // TODO: Get from context after auth
	isAdmin := false  // TODO: Get from context after auth

	requests, err := h.mealRequestService.GetMealRequests(c.Request.Context(), userID, isAdmin)
	if err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	c.JSON(http.StatusOK, requests)
}

// GetMealRequestByID handles GET /api/meal-requests/:id
//	@Summary		Get meal request by ID
//	@Description	Get a specific meal request by its ID
//	@Tags			meal-requests
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		int	true	"Meal Request ID"
//	@Success		200	{object}	model.MealRequest
//	@Failure		400	{object}	dto.ErrorResponse
//	@Failure		401	{object}	dto.ErrorResponse
//	@Failure		404	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/meal-requests/{id} [get]
func (h *MealRequestHandler) GetMealRequestByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		middleware.HandleAppError(c, apperrors.NewValidationError("Invalid meal request ID", err))
		return
	}

	userID := uint(1) // TODO: Get from context after auth
	isAdmin := false  // TODO: Get from context after auth

	request, err := h.mealRequestService.GetMealRequestByID(c.Request.Context(), uint(id), userID, isAdmin)
	if err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	c.JSON(http.StatusOK, request)
}

// CreateMealRequest handles POST /api/meal-requests
//	@Summary		Create meal request
//	@Description	Create a new meal request
//	@Tags			meal-requests
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			request	body		model.MealRequest	true	"Meal Request Data"
//	@Success		201		{object}	model.MealRequest
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		401		{object}	dto.ErrorResponse
//	@Failure		500		{object}	dto.ErrorResponse
//	@Router			/meal-requests [post]
func (h *MealRequestHandler) CreateMealRequest(c *gin.Context) {
	var request model.MealRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		middleware.HandleAppError(c, apperrors.NewValidationError("Invalid request format", err))
		return
	}

	userID := uint(1) // TODO: Get from context after auth

	if err := h.mealRequestService.CreateMealRequest(c.Request.Context(), &request, userID); err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	c.JSON(http.StatusCreated, request)
}

// UpdateMealRequest handles PUT /api/meal-requests/:id
//	@Summary		Update meal request
//	@Description	Update an existing meal request
//	@Tags			meal-requests
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id		path		int					true	"Meal Request ID"
//	@Param			request	body		model.MealRequest	true	"Meal Request Data"
//	@Success		200		{object}	model.MealRequest
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		401		{object}	dto.ErrorResponse
//	@Failure		403		{object}	dto.ErrorResponse
//	@Failure		404		{object}	dto.ErrorResponse
//	@Failure		500		{object}	dto.ErrorResponse
//	@Router			/meal-requests/{id} [put]
func (h *MealRequestHandler) UpdateMealRequest(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		middleware.HandleAppError(c, apperrors.NewValidationError("Invalid meal request ID", err))
		return
	}

	var request model.MealRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		middleware.HandleAppError(c, apperrors.NewValidationError("Invalid request format", err))
		return
	}

	userID := uint(1) // TODO: Get from context after auth
	isAdmin := false  // TODO: Get from context after auth

	if err := h.mealRequestService.UpdateMealRequest(c.Request.Context(), uint(id), &request, userID, isAdmin); err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	c.JSON(http.StatusOK, request)
}

// DeleteMealRequest handles DELETE /api/meal-requests/:id
//	@Summary		Delete meal request
//	@Description	Delete an existing meal request
//	@Tags			meal-requests
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		int	true	"Meal Request ID"
//	@Success		200	{object}	SuccessResponse
//	@Failure		400	{object}	dto.ErrorResponse
//	@Failure		401	{object}	dto.ErrorResponse
//	@Failure		403	{object}	dto.ErrorResponse
//	@Failure		404	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/meal-requests/{id} [delete]
func (h *MealRequestHandler) DeleteMealRequest(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		middleware.HandleAppError(c, apperrors.NewValidationError("Invalid meal request ID", err))
		return
	}

	userID := uint(1) // TODO: Get from context after auth
	isAdmin := false  // TODO: Get from context after auth

	if err := h.mealRequestService.DeleteMealRequest(c.Request.Context(), uint(id), userID, isAdmin); err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Meal request deleted successfully"})
}

// AddRequestItem handles POST /api/meal-requests/:id/items
//	@Summary		Add item to meal request
//	@Description	Add a new item to an existing meal request
//	@Tags			meal-requests
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id		path		int						true	"Meal Request ID"
//	@Param			item	body		model.MealRequestItem	true	"Meal Request Item Data"
//	@Success		200		{object}	model.MealRequestItem
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		401		{object}	dto.ErrorResponse
//	@Failure		403		{object}	dto.ErrorResponse
//	@Failure		404		{object}	dto.ErrorResponse
//	@Failure		500		{object}	dto.ErrorResponse
//	@Router			/meal-requests/{id}/items [post]
func (h *MealRequestHandler) AddRequestItem(c *gin.Context) {
	requestID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		middleware.HandleAppError(c, apperrors.NewValidationError("Invalid meal request ID", err))
		return
	}

	var item model.MealRequestItem
	if err := c.ShouldBindJSON(&item); err != nil {
		middleware.HandleAppError(c, apperrors.NewValidationError("Invalid request format", err))
		return
	}

	userID := uint(1) // TODO: Get from context after auth
	isAdmin := false  // TODO: Get from context after auth

	if err := h.mealRequestService.AddRequestItem(c.Request.Context(), uint(requestID), &item, userID, isAdmin); err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	c.JSON(http.StatusOK, item)
}

// RemoveRequestItem handles DELETE /api/meal-requests/:id/items/:item_id
//	@Summary		Remove item from meal request
//	@Description	Remove an item from an existing meal request
//	@Tags			meal-requests
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id		path		int	true	"Meal Request ID"
//	@Param			item_id	path		int	true	"Item ID"
//	@Success		200		{object}	SuccessResponse
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		401		{object}	dto.ErrorResponse
//	@Failure		403		{object}	dto.ErrorResponse
//	@Failure		404		{object}	dto.ErrorResponse
//	@Failure		500		{object}	dto.ErrorResponse
//	@Router			/meal-requests/{id}/items/{item_id} [delete]
func (h *MealRequestHandler) RemoveRequestItem(c *gin.Context) {
	requestID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		middleware.HandleAppError(c, apperrors.NewValidationError("Invalid meal request ID", err))
		return
	}

	itemID, err := strconv.ParseUint(c.Param("item_id"), 10, 32)
	if err != nil {
		middleware.HandleAppError(c, apperrors.NewValidationError("Invalid item ID", err))
		return
	}

	userID := uint(1) // TODO: Get from context after auth
	isAdmin := false  // TODO: Get from context after auth

	if err := h.mealRequestService.RemoveRequestItem(c.Request.Context(), uint(requestID), uint(itemID), userID, isAdmin); err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Item removed from meal request successfully"})
}

// GetRequestItems handles GET /api/meal-requests/:id/items
//	@Summary		List meal request items
//	@Description	Get all items in a specific meal request
//	@Tags			meal-requests
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		int	true	"Meal Request ID"
//	@Success		200	{array}		model.MealRequestItem
//	@Failure		400	{object}	dto.ErrorResponse
//	@Failure		401	{object}	dto.ErrorResponse
//	@Failure		404	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/meal-requests/{id}/items [get]
func (h *MealRequestHandler) GetRequestItems(c *gin.Context) {
	requestID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		middleware.HandleAppError(c, apperrors.NewValidationError("Invalid meal request ID", err))
		return
	}

	userID := uint(1) // TODO: Get from context after auth
	isAdmin := false  // TODO: Get from context after auth

	items, err := h.mealRequestService.GetRequestItems(c.Request.Context(), uint(requestID), userID, isAdmin)
	if err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	c.JSON(http.StatusOK, items)
}

// UpdateRequestStatus handles PUT /api/requests/:id/status
//	@Summary		Update meal request status
//	@Description	Update the status of a specific meal request
//	@Tags			meal-requests
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id		path		int		true	"Meal Request ID"
//	@Param			status	body		string	true	"New Status"	Enums(pending, accepted, rejected, completed)
//	@Success		200		{object}	model.MealRequest
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		401		{object}	dto.ErrorResponse
//	@Failure		404		{object}	dto.ErrorResponse
//	@Failure		500		{object}	dto.ErrorResponse
//	@Router			/meal-requests/{id}/status [put]
func (h *MealRequestHandler) UpdateRequestStatus(c *gin.Context) {
	requestID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		middleware.HandleAppError(c, apperrors.NewValidationError("Invalid meal request ID", err))
		return
	}

	var status model.RequestStatus
	if err := c.ShouldBindJSON(&status); err != nil {
		middleware.HandleAppError(c, apperrors.NewValidationError("Invalid request format", err))
		return
	}

	userID := uint(1) // TODO: Get from context after auth

	if err := h.mealRequestService.UpdateRequestStatus(c.Request.Context(), uint(requestID), status, userID); err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Meal request status updated successfully"})
}
