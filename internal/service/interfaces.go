package service

import (
	"context"
	"time"

	"github.com/arafat-hasan/mealsync/internal/model"
)

// UserService defines user-related business operations
type UserService interface {
	Register(ctx context.Context, user *model.User) error
	Login(ctx context.Context, email, password string) (string, string, error)
	RefreshToken(ctx context.Context, refreshToken string) (string, string, error)
	GetUserByID(ctx context.Context, id uint) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id uint) error
}

// MealEventService defines meal event-specific business logic
type MealEventService interface {
	Create(ctx context.Context, meal *model.MealEvent) error
	FindByID(ctx context.Context, id uint) (*model.MealEvent, error)
	FindAll(ctx context.Context) ([]model.MealEvent, error)
	FindActive(ctx context.Context) ([]model.MealEvent, error)
	Update(ctx context.Context, meal *model.MealEvent) error
	Delete(ctx context.Context, meal *model.MealEvent) error
	HardDelete(ctx context.Context, meal *model.MealEvent) error

	// Methods used by the API handlers
	GetMeals(ctx context.Context, userID uint, isAdmin bool) ([]model.MealEvent, error)
	GetMealByID(ctx context.Context, id uint, userID uint, isAdmin bool) (*model.MealEvent, error)
	CreateMeal(ctx context.Context, meal *model.MealEvent, userID uint) error
	UpdateMeal(ctx context.Context, id uint, meal *model.MealEvent, userID uint) error
	DeleteMeal(ctx context.Context, id uint, userID uint) error

	// MealEventMenuSet management with label and note
	AddMenuSetToEvent(ctx context.Context, MealEventMenuSet *model.MealEventMenuSet) error
	UpdateMenuSetInEvent(ctx context.Context, MealEventMenuSet *model.MealEventMenuSet) error
	RemoveMenuSetFromEvent(ctx context.Context, mealEventID uint, menuSetID uint) error
	FindMenuSetsByEventID(ctx context.Context, mealEventID uint) ([]model.MealEventMenuSet, error)

	// MealRequest specific operations
	CreateMealRequest(ctx context.Context, request *model.MealRequest) error
	FindRequestByID(ctx context.Context, id uint) (*model.MealRequest, error)
	FindAllRequests(ctx context.Context) ([]model.MealRequest, error)
	FindRequestsByUserID(ctx context.Context, userID uint) ([]model.MealRequest, error)
	DeleteMealRequest(ctx context.Context, request *model.MealRequest) error

	// MealEventAddress operations
	AddAddressToEvent(ctx context.Context, eventAddressID uint, mealEventID uint) error
	RemoveAddressFromEvent(ctx context.Context, eventAddressID uint, mealEventID uint) error
	FindAddressesByEventID(ctx context.Context, mealEventID uint) ([]model.MealEventAddress, error)

	CreateComment(ctx context.Context, comment *model.MenuItemComment) error
	FindCommentsByMealEventID(ctx context.Context, mealEventID uint) ([]model.MenuItemComment, error)

	FindByUserID(ctx context.Context, userID uint) ([]model.MealEvent, error)
	FindUpcomingAndActive(ctx context.Context) ([]model.MealEvent, error)
	FindByDateRange(ctx context.Context, startDate, endDate time.Time) ([]model.MealEvent, error)
}

// MenuSetService defines menu set-related business operations
type MenuSetService interface {
	GetMenuSets(ctx context.Context) ([]model.MenuSet, error)
	GetMenuSetByID(ctx context.Context, id uint) (*model.MenuSet, error)
	CreateMenuSet(ctx context.Context, menuSet *model.MenuSet, userID uint) error
	UpdateMenuSet(ctx context.Context, id uint, menuSet *model.MenuSet, userID uint) error
	DeleteMenuSet(ctx context.Context, id uint, userID uint) error
	GetMenuItems(ctx context.Context) ([]model.MenuItem, error)
	GetMenuItemByID(ctx context.Context, id uint) (*model.MenuItem, error)
	CreateMenuItem(ctx context.Context, menuItem *model.MenuItem, userID uint) error
	UpdateMenuItem(ctx context.Context, id uint, menuItem *model.MenuItem, userID uint) error
	DeleteMenuItem(ctx context.Context, id uint, userID uint) error
	AddItemToMenuSet(ctx context.Context, menuSetID uint, menuItemID uint, userID uint) error
	RemoveItemFromMenuSet(ctx context.Context, menuSetID uint, menuItemID uint, userID uint) error
	GetMenuSetItems(ctx context.Context, menuSetID uint) ([]model.MenuItem, error)
}

// MenuItemService defines the interface for menu item operations
type MenuItemService interface {
	GetMenuItems(ctx context.Context) ([]model.MenuItem, error)
	GetMenuItemByID(ctx context.Context, id uint) (*model.MenuItem, error)
	CreateMenuItem(ctx context.Context, menuItem *model.MenuItem, userID uint) error
	UpdateMenuItem(ctx context.Context, id uint, menuItem *model.MenuItem, userID uint) error
	DeleteMenuItem(ctx context.Context, id uint, userID uint) error
	GetMenuItemsByCategory(ctx context.Context, category string) ([]model.MenuItem, error)
	GetMenuItemsByMenuSet(ctx context.Context, menuSetID uint) ([]model.MenuItem, error)
}

// NotificationService defines notification-related business operations
type NotificationService interface {
	GetNotifications(ctx context.Context, userID uint) ([]model.Notification, error)
	CreateNotification(ctx context.Context, notification *model.Notification, userID uint) error
	MarkNotificationAsRead(ctx context.Context, notificationID uint, userID uint) error
	MarkNotificationAsDelivered(ctx context.Context, notificationID uint, userID uint) error
	DeleteNotification(ctx context.Context, notificationID uint, userID uint) error
	GetUnreadNotificationCount(ctx context.Context, userID uint) (int64, error)
	GetUndeliveredNotificationCount(ctx context.Context, userID uint) (int64, error)
	GetNotificationsByType(ctx context.Context, userID uint, notificationType model.NotificationType) ([]model.Notification, error)
	GetUnreadNotifications(ctx context.Context, userID uint) ([]model.Notification, error)
	CreateMealConfirmationNotification(ctx context.Context, userID uint, mealEventID uint, message string) error
	CreateMealReminderNotification(ctx context.Context, userID uint, mealEventID uint, message string, deadline time.Time) error
	CreateMealCancellationNotification(ctx context.Context, userID uint, mealEventID uint, message string) error
	CreateAdminNotification(ctx context.Context, userID uint, message string, importance string) error
}

// MealRequestService defines the interface for meal request operations
type MealRequestService interface {
	GetMealRequests(ctx context.Context, userID uint, isAdmin bool) ([]model.MealRequest, error)
	GetMealRequestByID(ctx context.Context, id uint, userID uint, isAdmin bool) (*model.MealRequest, error)
	CreateMealRequest(ctx context.Context, request *model.MealRequest, userID uint) error
	UpdateMealRequest(ctx context.Context, id uint, request *model.MealRequest, userID uint, isAdmin bool) error
	DeleteMealRequest(ctx context.Context, id uint, userID uint, isAdmin bool) error
	AddRequestItem(ctx context.Context, requestID uint, item *model.MealRequestItem, userID uint, isAdmin bool) error
	RemoveRequestItem(ctx context.Context, requestID uint, itemID uint, userID uint, isAdmin bool) error
	GetRequestItems(ctx context.Context, requestID uint, userID uint, isAdmin bool) ([]model.MealRequestItem, error)
	UpdateRequestStatus(ctx context.Context, requestID uint, status model.RequestStatus, userID uint) error
}

// EventAddressService defines event address-related business operations
type EventAddressService interface {
	GetAddresses(ctx context.Context) ([]model.MealEventAddress, error)
	GetAddressByID(ctx context.Context, id uint) (*model.MealEventAddress, error)
	CreateAddress(ctx context.Context, address *model.MealEventAddress, userID uint) error
	UpdateAddress(ctx context.Context, id uint, address *model.MealEventAddress, userID uint) error
	DeleteAddress(ctx context.Context, id uint, userID uint) error
	GetAddressesByType(ctx context.Context, addressType string) ([]model.MealEventAddress, error)
	GetAddressesByCapacity(ctx context.Context, minCapacity, maxCapacity int) ([]model.MealEventAddress, error)
	GetAddressesByLocation(ctx context.Context, latitude, longitude, radius float64) ([]model.MealEventAddress, error)
	GetAvailableAddresses(ctx context.Context, date time.Time) ([]model.MealEventAddress, error)
}

// MenuItemCommentService defines menu item comment-related business operations
type MenuItemCommentService interface {
	GetComments(ctx context.Context, mealEventID uint) ([]model.MenuItemComment, error)
	GetCommentByID(ctx context.Context, id uint) (*model.MenuItemComment, error)
	CreateComment(ctx context.Context, comment *model.MenuItemComment, userID uint) error
	UpdateComment(ctx context.Context, id uint, comment *model.MenuItemComment, userID uint) error
	DeleteComment(ctx context.Context, id uint, userID uint) error
	GetUserComments(ctx context.Context, userID uint) ([]model.MenuItemComment, error)
	GetMenuItemComments(ctx context.Context, menuItemID uint) ([]model.MenuItemComment, error)
	GetReplies(ctx context.Context, commentID uint) ([]model.MenuItemComment, error) // Added for comment replies
}
