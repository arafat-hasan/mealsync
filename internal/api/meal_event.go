package api

import (
	"net/http"
	"strconv"

	"github.com/arafat-hasan/mealsync/internal/model"
	"github.com/arafat-hasan/mealsync/internal/service"
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
	userID := uint(1) // TODO: Get from context after auth
	isAdmin := false  // TODO: Get from context after auth

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

	userID := uint(1) // TODO: Get from context after auth
	isAdmin := false  // TODO: Get from context after auth

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

	userID := uint(1) // TODO: Get from context after auth

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

	userID := uint(1) // TODO: Get from context after auth

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

	userID := uint(1) // TODO: Get from context after auth

	if err := h.mealService.DeleteMeal(c.Request.Context(), uint(id), userID); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Meal event deleted successfully"})
}
