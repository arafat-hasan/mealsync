package api

import (
	"net/http"
	"strconv"

	"github.com/arafat-hasan/mealsync/internal/model"
	"github.com/arafat-hasan/mealsync/internal/service"
	"github.com/gin-gonic/gin"
)

// MealCommentHandler handles meal comment-related requests
type MealCommentHandler struct {
	mealCommentService service.MealCommentService
}

// NewMealCommentHandler creates a new MealCommentHandler
func NewMealCommentHandler(mealCommentService service.MealCommentService) *MealCommentHandler {
	return &MealCommentHandler{mealCommentService: mealCommentService}
}

// GetComments handles GET /api/meals/:meal_event_id/comments
// @Summary      List meal comments
// @Description  Get all comments for a specific meal event
// @Tags         meal-comments
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        meal_event_id  path      int  true  "Meal Event ID"
// @Success      200           {array}   model.MealComment
// @Failure      400           {object}  ErrorResponse
// @Failure      401           {object}  ErrorResponse
// @Failure      404           {object}  ErrorResponse
// @Failure      500           {object}  ErrorResponse
// @Router       /meals/{meal_event_id}/comments [get]
func (h *MealCommentHandler) GetComments(c *gin.Context) {
	mealEventID, err := strconv.ParseUint(c.Param("meal_event_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid meal event ID"})
		return
	}

	comments, err := h.mealCommentService.GetComments(c.Request.Context(), uint(mealEventID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

// GetCommentByID handles GET /api/meals/:meal_event_id/comments/:id
// @Summary      Get comment by ID
// @Description  Get a specific comment by its ID
// @Tags         meal-comments
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        meal_event_id  path      int  true  "Meal Event ID"
// @Param        id            path      int  true  "Comment ID"
// @Success      200          {object}  model.MealComment
// @Failure      400          {object}  ErrorResponse
// @Failure      401          {object}  ErrorResponse
// @Failure      404          {object}  ErrorResponse
// @Failure      500          {object}  ErrorResponse
// @Router       /meals/{meal_event_id}/comments/{id} [get]
func (h *MealCommentHandler) GetCommentByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid comment ID"})
		return
	}

	comment, err := h.mealCommentService.GetCommentByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, comment)
}

// CreateComment handles POST /api/meals/:meal_event_id/comments
// @Summary      Create comment
// @Description  Create a new comment for a meal event
// @Tags         meal-comments
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        meal_event_id  path           int             true  "Meal Event ID"
// @Param        comment        body           model.MealComment true  "Comment Data"
// @Success      201           {object}       model.MealComment
// @Failure      400           {object}       ErrorResponse
// @Failure      401           {object}       ErrorResponse
// @Failure      404           {object}       ErrorResponse
// @Failure      500           {object}       ErrorResponse
// @Router       /meals/{meal_event_id}/comments [post]
func (h *MealCommentHandler) CreateComment(c *gin.Context) {
	mealEventID, err := strconv.ParseUint(c.Param("meal_event_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid meal event ID"})
		return
	}

	var comment model.MealComment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid request format"})
		return
	}

	comment.MealEventID = uint(mealEventID)
	userID := uint(1) // TODO: Get from context after auth
	comment.UserID = userID

	if err := h.mealCommentService.CreateComment(c.Request.Context(), &comment, userID); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

// UpdateComment handles PUT /api/meals/:meal_event_id/comments/:id
// @Summary      Update comment
// @Description  Update an existing comment
// @Tags         meal-comments
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        meal_event_id  path           int             true  "Meal Event ID"
// @Param        id            path           int             true  "Comment ID"
// @Param        comment        body           model.MealComment true  "Comment Data"
// @Success      200          {object}       model.MealComment
// @Failure      400          {object}       ErrorResponse
// @Failure      401          {object}       ErrorResponse
// @Failure      403          {object}       ErrorResponse
// @Failure      404          {object}       ErrorResponse
// @Failure      500          {object}       ErrorResponse
// @Router       /meals/{meal_event_id}/comments/{id} [put]
func (h *MealCommentHandler) UpdateComment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid comment ID"})
		return
	}

	var comment model.MealComment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid request format"})
		return
	}

	userID := uint(1) // TODO: Get from context after auth

	if err := h.mealCommentService.UpdateComment(c.Request.Context(), uint(id), &comment, userID); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, comment)
}

// DeleteComment handles DELETE /api/meals/:meal_event_id/comments/:id
// @Summary      Delete comment
// @Description  Delete an existing comment
// @Tags         meal-comments
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        meal_event_id  path      int  true  "Meal Event ID"
// @Param        id            path      int  true  "Comment ID"
// @Success      200          {object}  SuccessResponse
// @Failure      400          {object}  ErrorResponse
// @Failure      401          {object}  ErrorResponse
// @Failure      403          {object}  ErrorResponse
// @Failure      404          {object}  ErrorResponse
// @Failure      500          {object}  ErrorResponse
// @Router       /meals/{meal_event_id}/comments/{id} [delete]
func (h *MealCommentHandler) DeleteComment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid comment ID"})
		return
	}

	userID := uint(1) // TODO: Get from context after auth

	if err := h.mealCommentService.DeleteComment(c.Request.Context(), uint(id), userID); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Comment deleted successfully"})
}

// GetReplies handles GET /api/meals/:meal_event_id/comments/:id/replies
// @Summary      List comment replies
// @Description  Get all replies to a specific comment
// @Tags         meal-comments
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        meal_event_id  path      int  true  "Meal Event ID"
// @Param        id            path      int  true  "Comment ID"
// @Success      200          {array}   model.MealComment
// @Failure      400          {object}  ErrorResponse
// @Failure      401          {object}  ErrorResponse
// @Failure      404          {object}  ErrorResponse
// @Failure      500          {object}  ErrorResponse
// @Router       /meals/{meal_event_id}/comments/{id}/replies [get]
func (h *MealCommentHandler) GetReplies(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid comment ID"})
		return
	}

	replies, err := h.mealCommentService.GetReplies(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, replies)
}

// GetUserComments handles GET /api/users/:user_id/comments
// @Summary      List user comments
// @Description  Get all comments made by a specific user
// @Tags         meal-comments
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        user_id  path      int  true  "User ID"
// @Success      200     {array}   model.MealComment
// @Failure      400     {object}  ErrorResponse
// @Failure      401     {object}  ErrorResponse
// @Failure      404     {object}  ErrorResponse
// @Failure      500     {object}  ErrorResponse
// @Router       /users/{user_id}/comments [get]
func (h *MealCommentHandler) GetUserComments(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid user ID"})
		return
	}

	comments, err := h.mealCommentService.GetUserComments(c.Request.Context(), uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}
