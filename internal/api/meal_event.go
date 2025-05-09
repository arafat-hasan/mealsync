package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/arafat-hasan/mealsync/internal/dto"
	apperrors "github.com/arafat-hasan/mealsync/internal/errors"
	"github.com/arafat-hasan/mealsync/internal/mapper"
	"github.com/arafat-hasan/mealsync/internal/middleware"
	"github.com/arafat-hasan/mealsync/internal/model"
	"github.com/arafat-hasan/mealsync/internal/service"
	"github.com/arafat-hasan/mealsync/internal/utils"
	"github.com/gin-gonic/gin"
)

// MealEventHandler handles meal event-related requests
type MealEventHandler struct {
	mealService service.MealEventService
}

// NewMealEventHandler creates a new MealEventHandler
func NewMealEventHandler(mealService service.MealEventService) *MealEventHandler {
	return &MealEventHandler{mealService: mealService}
}

// GetMealEventByID handles GET /api/meals/:id
// @Summary      Get meal event by ID
// @Description  Get a specific meal event by its ID
// @Tags         meals
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Meal Event ID"
// @Success      200  {object}  dto.MealEventResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      401  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /meals/{id} [get]
func (h *MealEventHandler) GetMealEventByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		middleware.HandleAppError(c, apperrors.NewValidationError("Invalid meal event ID", err))
		return
	}

	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		middleware.HandleAppError(c, apperrors.NewUnauthorizedError("Unauthorized", err))
		return
	}
	isAdmin := utils.IsAdminFromContext(c)

	meal, err := h.mealService.GetMealByID(c.Request.Context(), uint(id), userID, isAdmin)
	if err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	c.JSON(http.StatusOK, mapper.ToMealEventResponse(meal))
}

// CreateMealEvent handles POST /api/meals
// @Summary      Create meal event
// @Description  Create a new meal event
// @Tags         meals
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        meal  body      dto.MealEventCreateRequest  true  "Meal Event Data"
// @Success      201   {object}  dto.MealEventResponse
// @Failure      400   {object}  dto.ErrorResponse
// @Failure      401   {object}  dto.ErrorResponse
// @Failure      500   {object}  dto.ErrorResponse
// @Router       /meals [post]
func (h *MealEventHandler) CreateMealEvent(c *gin.Context) {
	var reqDTO dto.MealEventCreateRequest
	if err := c.ShouldBindJSON(&reqDTO); err != nil {
		middleware.HandleAppError(c, apperrors.NewValidationError("Invalid request format", err))
		return
	}

	// Convert DTO to model
	meal := &model.MealEvent{
		Name:          reqDTO.Name,
		Description:   reqDTO.Description,
		EventDate:     reqDTO.EventDate,
		EventDuration: reqDTO.EventDuration,
		CutoffTime:    reqDTO.CutoffTime,
	}

	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		middleware.HandleAppError(c, apperrors.NewUnauthorizedError("Unauthorized", err))
		return
	}

	if err := h.mealService.CreateMeal(c.Request.Context(), meal, userID); err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	// Return response as DTO
	c.JSON(http.StatusCreated, mapper.ToMealEventResponse(meal))
}

// UpdateMealEvent handles PUT /api/meals/:id
// @Summary      Update meal event
// @Description  Update an existing meal event
// @Tags         meals
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      int             true  "Meal Event ID"
// @Param        meal  body      dto.MealEventUpdateRequest true  "Meal Event Data"
// @Success      200   {object}  dto.MealEventResponse
// @Failure      400   {object}  dto.ErrorResponse
// @Failure      401   {object}  dto.ErrorResponse
// @Failure      403   {object}  dto.ErrorResponse
// @Failure      404   {object}  dto.ErrorResponse
// @Failure      500   {object}  dto.ErrorResponse
// @Router       /meals/{id} [put]
func (h *MealEventHandler) UpdateMealEvent(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		middleware.HandleAppError(c, apperrors.NewValidationError("Invalid meal event ID", err))
		return
	}

	var reqDTO dto.MealEventUpdateRequest
	if err := c.ShouldBindJSON(&reqDTO); err != nil {
		middleware.HandleAppError(c, apperrors.NewValidationError("Invalid request format", err))
		return
	}

	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		middleware.HandleAppError(c, apperrors.NewUnauthorizedError("Unauthorized", err))
		return
	}

	// First fetch existing meal
	existingMeal, err := h.mealService.GetMealByID(c.Request.Context(), uint(id), userID, utils.IsAdminFromContext(c))
	if err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	// Update the meal with values from DTO
	if reqDTO.Name != nil {
		existingMeal.Name = *reqDTO.Name
	}
	if reqDTO.Description != nil {
		existingMeal.Description = *reqDTO.Description
	}
	if reqDTO.EventDate != nil {
		existingMeal.EventDate = *reqDTO.EventDate
	}
	if reqDTO.EventDuration != nil {
		existingMeal.EventDuration = *reqDTO.EventDuration
	}
	if reqDTO.CutoffTime != nil {
		existingMeal.CutoffTime = *reqDTO.CutoffTime
	}

	if err := h.mealService.UpdateMeal(c.Request.Context(), uint(id), existingMeal, userID); err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	c.JSON(http.StatusOK, mapper.ToMealEventResponse(existingMeal))
}

// DeleteMealEvent handles DELETE /api/meals/:id
// @Summary      Delete meal event
// @Description  Delete an existing meal event
// @Tags         meals
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Meal Event ID"
// @Success      200  {object}  SuccessResponse
// @Failure      400   {object}  dto.ErrorResponse
// @Failure      401   {object}  dto.ErrorResponse
// @Failure      403   {object}  dto.ErrorResponse
// @Failure      404   {object}  dto.ErrorResponse
// @Failure      500   {object}  dto.ErrorResponse
// @Router       /meals/{id} [delete]
func (h *MealEventHandler) DeleteMealEvent(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		middleware.HandleAppError(c, apperrors.NewValidationError("Invalid meal event ID", err))
		return
	}

	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		middleware.HandleAppError(c, apperrors.NewUnauthorizedError("Unauthorized", err))
		return
	}

	if err := h.mealService.DeleteMeal(c.Request.Context(), uint(id), userID); err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Meal event deleted successfully"})
}

// GetMeals handles GET /api/meals
// @Summary      List meal events by date range
// @Description  Get meal events within a date range
// @Tags         meals
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        start_date  query     string  true  "Start Date (YYYY-MM-DD)"
// @Param        end_date    query     string  true  "End Date (YYYY-MM-DD)"
// @Success      200  {array}   dto.MealEventResponse
// @Failure      400   {object}  dto.ErrorResponse
// @Failure      401   {object}  dto.ErrorResponse
// @Failure      500   {object}  dto.ErrorResponse
// @Router       /meals [get]
func (h *MealEventHandler) GetMeals(c *gin.Context) {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	if startDateStr == "" || endDateStr == "" {
		middleware.HandleAppError(c, apperrors.NewValidationError("Both start_date and end_date are required", nil))
		return
	}

	// Parse dates
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		middleware.HandleAppError(c, apperrors.NewValidationError("Invalid start date format. Use YYYY-MM-DD", err))
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		middleware.HandleAppError(c, apperrors.NewValidationError("Invalid end date format. Use YYYY-MM-DD", err))
		return
	}

	// Set end date to the end of the day
	endDate = endDate.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	// Validate date range
	if startDate.After(endDate) {
		middleware.HandleAppError(c, apperrors.NewValidationError("Start date must be before end date", nil))
		return
	}

	// Check authentication
	_, err = utils.GetUserIDFromContext(c)
	if err != nil {
		middleware.HandleAppError(c, apperrors.NewUnauthorizedError("Unauthorized", err))
		return
	}

	meals, err := h.mealService.FindByDateRange(c.Request.Context(), startDate, endDate)
	if err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	// Convert to DTOs
	c.JSON(http.StatusOK, mapper.ToMealEventResponseList(meals))
}
