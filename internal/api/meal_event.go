package api

import (
	"net/http"
	"strconv"
	"time"

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

// GetMealEvents handles GET /api/meals
// @Summary      List meal events
// @Description  Get all meal events for the authenticated user
// @Tags         meals
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   model.MealEvent
// @Failure      401  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /meals [get]
func (h *MealEventHandler) GetMealEvents(c *gin.Context) {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	isAdmin := utils.IsAdminFromContext(c)

	meals, err := h.mealService.GetMeals(c.Request.Context(), userID, isAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, meals)
}

// GetMealEventByID handles GET /api/meals/:id
// @Summary      Get meal event by ID
// @Description  Get a specific meal event by its ID
// @Tags         meals
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Meal Event ID"
// @Success      200  {object}  model.MealEvent
// @Failure      400  {object}  ErrorResponse
// @Failure      401  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /meals/{id} [get]
func (h *MealEventHandler) GetMealEventByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid meal event ID"})
		return
	}

	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	isAdmin := utils.IsAdminFromContext(c)

	meal, err := h.mealService.GetMealByID(c.Request.Context(), uint(id), userID, isAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, meal)
}

// CreateMealEvent handles POST /api/meals
// @Summary      Create meal event
// @Description  Create a new meal event
// @Tags         meals
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        meal  body      model.MealEvent  true  "Meal Event Data"
// @Success      201   {object}  model.MealEvent
// @Failure      400   {object}  ErrorResponse
// @Failure      401   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /meals [post]
func (h *MealEventHandler) CreateMealEvent(c *gin.Context) {
	var meal model.MealEvent
	if err := c.ShouldBindJSON(&meal); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid request format"})
		return
	}

	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := h.mealService.CreateMeal(c.Request.Context(), &meal, userID); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, meal)
}

// UpdateMealEvent handles PUT /api/meals/:id
// @Summary      Update meal event
// @Description  Update an existing meal event
// @Tags         meals
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      int             true  "Meal Event ID"
// @Param        meal  body      model.MealEvent true  "Meal Event Data"
// @Success      200   {object}  model.MealEvent
// @Failure      400   {object}  ErrorResponse
// @Failure      401   {object}  ErrorResponse
// @Failure      403   {object}  ErrorResponse
// @Failure      404   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /meals/{id} [put]
func (h *MealEventHandler) UpdateMealEvent(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid meal event ID"})
		return
	}

	var meal model.MealEvent
	if err := c.ShouldBindJSON(&meal); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid request format"})
		return
	}

	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := h.mealService.UpdateMeal(c.Request.Context(), uint(id), &meal, userID); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, meal)
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
// @Failure      400  {object}  ErrorResponse
// @Failure      401  {object}  ErrorResponse
// @Failure      403  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /meals/{id} [delete]
func (h *MealEventHandler) DeleteMealEvent(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid meal event ID"})
		return
	}

	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := h.mealService.DeleteMeal(c.Request.Context(), uint(id), userID); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Meal event deleted successfully"})
}

// GetMealEventsByDateRange handles GET /api/meals/daterange
// @Summary      List meal events by date range
// @Description  Get meal events within a date range
// @Tags         meals
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        start_date  query     string  true  "Start Date (YYYY-MM-DD)"
// @Param        end_date    query     string  true  "End Date (YYYY-MM-DD)"
// @Success      200  {array}   model.MealEvent
// @Failure      400  {object}  ErrorResponse
// @Failure      401  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /meals/daterange [get]
func (h *MealEventHandler) GetMealEventsByDateRange(c *gin.Context) {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	if startDateStr == "" || endDateStr == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Both start_date and end_date are required"})
		return
	}

	// Parse dates
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid start date format. Use YYYY-MM-DD"})
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid end date format. Use YYYY-MM-DD"})
		return
	}

	// Set end date to the end of the day
	endDate = endDate.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	// Validate date range
	if startDate.After(endDate) {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Start date must be before end date"})
		return
	}

	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	isAdmin := utils.IsAdminFromContext(c)

	// Get meals in the date range
	meals, err := h.mealService.FindByDateRange(c.Request.Context(), startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	// Filter by user access if not admin
	if !isAdmin {
		userMeals, err := h.mealService.FindByUserID(c.Request.Context(), userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}

		// Create a map for fast lookup of user's meal IDs
		userMealIDs := make(map[uint]bool)
		for _, meal := range userMeals {
			userMealIDs[meal.ID] = true
		}

		// Filter meals to only include those the user has access to
		filteredMeals := []model.MealEvent{}
		for _, meal := range meals {
			if userMealIDs[meal.ID] {
				filteredMeals = append(filteredMeals, meal)
			}
		}

		meals = filteredMeals
	}

	c.JSON(http.StatusOK, meals)
}
