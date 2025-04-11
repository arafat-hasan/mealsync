package api

import (
	"net/http"
	"strconv"

	"github.com/arafat-hasan/mealsync/internal/model"
	"github.com/arafat-hasan/mealsync/internal/service"
	"github.com/gin-gonic/gin"
)

type MealHandler struct {
	mealService service.MealService
}

func NewMealHandler(mealService service.MealService) *MealHandler {
	return &MealHandler{
		mealService: mealService,
	}
}

// CreateMealRequest represents the request body for creating a meal
type CreateMealRequest struct {
	Date        string   `json:"date" binding:"required" example:"2024-03-20"`
	MenuItems   []string `json:"menu_items" binding:"required" example:"['Chicken Curry', 'Rice', 'Salad']"`
	Description string   `json:"description" example:"Special lunch menu for the week"`
}

// @Summary Create a new meal
// @Description Create a new meal with menu items
// @Tags meals
// @Accept json
// @Produce json
// @Param meal body CreateMealRequest true "Meal details"
// @Success 201 {object} model.Meal
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /meals [post]
func (h *MealHandler) CreateMeal(c *gin.Context) {
	var req CreateMealRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	meal := &model.Meal{
		Date:        req.Date,
		MenuItems:   req.MenuItems,
		Description: req.Description,
	}

	if err := h.mealService.CreateMeal(c.Request.Context(), meal); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, meal)
}

// @Summary Get all meals
// @Description Get a list of all meals
// @Tags meals
// @Produce json
// @Success 200 {array} model.Meal
// @Failure 401 {object} ErrorResponse
// @Router /meals [get]
func (h *MealHandler) GetMeals(c *gin.Context) {
	meals, err := h.mealService.GetMeals(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, meals)
}

// @Summary Get meal by ID
// @Description Get a specific meal by its ID
// @Tags meals
// @Produce json
// @Param id path int true "Meal ID"
// @Success 200 {object} model.Meal
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /meals/{id} [get]
func (h *MealHandler) GetMealByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid meal ID"})
		return
	}

	meal, err := h.mealService.GetMealByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Meal not found"})
		return
	}

	c.JSON(http.StatusOK, meal)
}

// @Summary Update a meal
// @Description Update an existing meal's details
// @Tags meals
// @Accept json
// @Produce json
// @Param id path int true "Meal ID"
// @Param meal body CreateMealRequest true "Updated meal details"
// @Success 200 {object} model.Meal
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /meals/{id} [put]
func (h *MealHandler) UpdateMeal(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid meal ID"})
		return
	}

	var req CreateMealRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	meal := &model.Meal{
		ID:          uint(id),
		Date:        req.Date,
		MenuItems:   req.MenuItems,
		Description: req.Description,
	}

	if err := h.mealService.UpdateMeal(c.Request.Context(), meal); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, meal)
}

// @Summary Delete a meal
// @Description Delete a meal by its ID
// @Tags meals
// @Produce json
// @Param id path int true "Meal ID"
// @Success 200 {object} SuccessResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /meals/{id} [delete]
func (h *MealHandler) DeleteMeal(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid meal ID"})
		return
	}

	if err := h.mealService.DeleteMeal(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Meal deleted successfully"})
}
