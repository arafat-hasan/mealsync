package api

import (
	"net/http"
	"strconv"

	"github.com/arafat-hasan/mealsync/internal/model"
	"github.com/arafat-hasan/mealsync/internal/service"
	"github.com/gin-gonic/gin"
)

// MenuItemCommentHandler handles menu item comment-related requests
type MenuItemCommentHandler struct {
	menuItemCommentService service.MenuItemCommentService
}

// NewMenuItemCommentHandler creates a new MenuItemCommentHandler
func NewMenuItemCommentHandler(menuItemCommentService service.MenuItemCommentService) *MenuItemCommentHandler {
	return &MenuItemCommentHandler{menuItemCommentService: menuItemCommentService}
}

// GetComments handles GET /api/meals/:meal_event_id/comments
// @Summary      List menu item comments
// @Description  Get all comments for a specific meal event
// @Tags         menu-item-comments
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        meal_event_id  path      int  true  "Meal Event ID"
// @Success      200           {array}   model.MenuItemComment
// @Failure      400           {object}  ErrorResponse
// @Failure      401           {object}  ErrorResponse
// @Failure      404           {object}  ErrorResponse
// @Failure      500           {object}  ErrorResponse
// @Router       /meals/{meal_event_id}/comments [get]
func (h *MenuItemCommentHandler) GetComments(c *gin.Context) {
	mealEventID, err := strconv.ParseUint(c.Param("meal_event_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid meal event ID"})
		return
	}

	comments, err := h.menuItemCommentService.GetComments(c.Request.Context(), uint(mealEventID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

// GetCommentByID handles GET /api/meals/:meal_event_id/comments/:id
// @Summary      Get comment by ID
// @Description  Get a specific comment by its ID
// @Tags         menu-item-comments
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        meal_event_id  path      int  true  "Meal Event ID"
// @Param        id            path      int  true  "Comment ID"
// @Success      200          {object}  model.MenuItemComment
// @Failure      400          {object}  ErrorResponse
// @Failure      401          {object}  ErrorResponse
// @Failure      404          {object}  ErrorResponse
// @Failure      500          {object}  ErrorResponse
// @Router       /meals/{meal_event_id}/comments/{id} [get]
func (h *MenuItemCommentHandler) GetCommentByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid comment ID"})
		return
	}

	comment, err := h.menuItemCommentService.GetCommentByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, comment)
}

// CreateComment handles POST /api/meals/:meal_event_id/comments
// @Summary      Create comment
// @Description  Create a new comment for a meal event
// @Tags         menu-item-comments
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        meal_event_id  path           int             true  "Meal Event ID"
// @Param        comment        body           model.MenuItemComment true  "Comment Data"
// @Success      201           {object}       model.MenuItemComment
// @Failure      400           {object}       ErrorResponse
// @Failure      401           {object}       ErrorResponse
// @Failure      404           {object}       ErrorResponse
// @Failure      500           {object}       ErrorResponse
// @Router       /meals/{meal_event_id}/comments [post]
func (h *MenuItemCommentHandler) CreateComment(c *gin.Context) {
	mealEventID, err := strconv.ParseUint(c.Param("meal_event_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid meal event ID"})
		return
	}

	var comment model.MenuItemComment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid request format"})
		return
	}

	comment.MealEventID = uint(mealEventID)
	userID := uint(1) // TODO: Get from context after auth
	comment.UserID = userID

	if err := h.menuItemCommentService.CreateComment(c.Request.Context(), &comment, userID); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

// UpdateComment handles PUT /api/meals/:meal_event_id/comments/:id
// @Summary      Update comment
// @Description  Update an existing comment
// @Tags         menu-item-comments
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        meal_event_id  path           int             true  "Meal Event ID"
// @Param        id            path           int             true  "Comment ID"
// @Param        comment        body           model.MenuItemComment true  "Comment Data"
// @Success      200          {object}       model.MenuItemComment
// @Failure      400          {object}       ErrorResponse
// @Failure      401          {object}       ErrorResponse
// @Failure      403          {object}       ErrorResponse
// @Failure      404          {object}       ErrorResponse
// @Failure      500          {object}       ErrorResponse
// @Router       /meals/{meal_event_id}/comments/{id} [put]
func (h *MenuItemCommentHandler) UpdateComment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid comment ID"})
		return
	}

	var comment model.MenuItemComment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid request format"})
		return
	}

	userID := uint(1) // TODO: Get from context after auth

	if err := h.menuItemCommentService.UpdateComment(c.Request.Context(), uint(id), &comment, userID); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, comment)
}

// DeleteComment handles DELETE /api/meals/:meal_event_id/comments/:id
// @Summary      Delete comment
// @Description  Delete an existing comment
// @Tags         menu-item-comments
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
func (h *MenuItemCommentHandler) DeleteComment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid comment ID"})
		return
	}

	userID := uint(1) // TODO: Get from context after auth

	if err := h.menuItemCommentService.DeleteComment(c.Request.Context(), uint(id), userID); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Comment deleted successfully"})
}

// GetUserComments handles GET /api/users/:user_id/comments
// @Summary      List user comments
// @Description  Get all comments made by a specific user
// @Tags         menu-item-comments
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        user_id  path      int  true  "User ID"
// @Success      200     {array}   model.MenuItemComment
// @Failure      400     {object}  ErrorResponse
// @Failure      401     {object}  ErrorResponse
// @Failure      404     {object}  ErrorResponse
// @Failure      500     {object}  ErrorResponse
// @Router       /users/{user_id}/comments [get]
func (h *MenuItemCommentHandler) GetUserComments(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid user ID"})
		return
	}

	comments, err := h.menuItemCommentService.GetUserComments(c.Request.Context(), uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

// GetMenuItemComments handles GET /api/menu-items/:menu_item_id/comments
// @Summary      List menu item comments
// @Description  Get all comments for a specific menu item
// @Tags         menu-item-comments
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        menu_item_id  path      int  true  "Menu Item ID"
// @Success      200     {array}   model.MenuItemComment
// @Failure      400     {object}  ErrorResponse
// @Failure      401     {object}  ErrorResponse
// @Failure      404     {object}  ErrorResponse
// @Failure      500     {object}  ErrorResponse
// @Router       /menu-items/{menu_item_id}/comments [get]
func (h *MenuItemCommentHandler) GetMenuItemComments(c *gin.Context) {
	menuItemID, err := strconv.ParseUint(c.Param("menu_item_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid menu item ID"})
		return
	}

	comments, err := h.menuItemCommentService.GetMenuItemComments(c.Request.Context(), uint(menuItemID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

// GetReplies handles GET /api/comments/:comment_id/replies
// @Summary      Get comment replies
// @Description  Get all replies to a specific comment
// @Tags         menu-item-comments
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        comment_id  path      int  true  "Comment ID"
// @Success      200     {array}   model.MenuItemComment
// @Failure      400     {object}  ErrorResponse
// @Failure      401     {object}  ErrorResponse
// @Failure      404     {object}  ErrorResponse
// @Failure      500     {object}  ErrorResponse
// @Router       /comments/{comment_id}/replies [get]
func (h *MenuItemCommentHandler) GetReplies(c *gin.Context) {
	commentID, err := strconv.ParseUint(c.Param("comment_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid comment ID"})
		return
	}

	replies, err := h.menuItemCommentService.GetReplies(c.Request.Context(), uint(commentID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, replies)
}
