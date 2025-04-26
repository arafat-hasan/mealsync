package repository

import (
	"context"
	"time"

	"github.com/arafat-hasan/mealsync/internal/model"
)

// BaseRepository defines common CRUD operations
type BaseRepository[T any] interface {
	Create(ctx context.Context, entity *T) error
	FindByID(ctx context.Context, id uint) (*T, error)
	FindAll(ctx context.Context) ([]T, error)
	FindActive(ctx context.Context, conditions map[string]interface{}) ([]T, error)
	Update(ctx context.Context, entity *T) error
	Delete(ctx context.Context, entity *T) error
	HardDelete(ctx context.Context, entity *T) error
}

// UserRepository defines user-specific operations
type UserRepository interface {
	BaseRepository[model.User]
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	FindByUsername(ctx context.Context, username string) (*model.User, error)
	FindByEmployeeID(ctx context.Context, employeeID int) (*model.User, error)
}

// MealEventRepository defines meal event-specific operations
type MealEventRepository interface {
	BaseRepository[model.MealEvent]
	CreateRequest(ctx context.Context, request *model.MealRequest) error
	FindRequestByID(ctx context.Context, id uint) (*model.MealRequest, error)
	FindAllRequests(ctx context.Context) ([]model.MealRequest, error)
	FindRequestsByUserID(ctx context.Context, userID uint) ([]model.MealRequest, error)
	DeleteRequest(ctx context.Context, request *model.MealRequest) error
	CreateComment(ctx context.Context, comment *model.MenuItemComment) error
	FindCommentsByMealEventID(ctx context.Context, mealEventID uint) ([]model.MenuItemComment, error)
	FindUpcomingAndActive(ctx context.Context) ([]model.MealEvent, error)
	FindByDateRange(ctx context.Context, startDate, endDate time.Time) ([]model.MealEvent, error)
	AddMenuSetToEvent(ctx context.Context, MealEventSet *model.MealEventSet) error
	UpdateMenuSetInEvent(ctx context.Context, MealEventSet *model.MealEventSet) error
	RemoveMenuSetFromEvent(ctx context.Context, mealEventID uint, menuSetID uint) error
	FindMenuSetsByEventID(ctx context.Context, mealEventID uint) ([]model.MealEventSet, error)
}

// EventAddressRepository defines the interface for event address repository
type EventAddressRepository interface {
	Create(ctx context.Context, address *model.MealEventAddress) error
	FindByID(ctx context.Context, id uint) (*model.MealEventAddress, error)
	FindAll(ctx context.Context) ([]model.MealEventAddress, error)
	FindActive(ctx context.Context, conditions map[string]interface{}) ([]model.MealEventAddress, error)
	Update(ctx context.Context, address *model.MealEventAddress) error
	Delete(ctx context.Context, address *model.MealEventAddress) error
	HardDelete(ctx context.Context, address *model.MealEventAddress) error
	FindByMealEventID(ctx context.Context, mealEventID uint) ([]model.MealEventAddress, error)
	CountByMealEventID(ctx context.Context, mealEventID uint) (int64, error)
	FindAvailableAddresses(ctx context.Context, date time.Time) ([]model.MealEventAddress, error)
	FindByAddressType(ctx context.Context, addressType string) ([]model.MealEventAddress, error)
	FindByBuildingName(ctx context.Context, buildingName string) ([]model.MealEventAddress, error)
	FindByCapacity(ctx context.Context, minCapacity, maxCapacity int) ([]model.MealEventAddress, error)
	FindByLocation(ctx context.Context, lat, lng, radius float64) ([]model.MealEventAddress, error)
	FindWithEventDetails(ctx context.Context, mealEventID uint) (*model.MealEventAddress, error)
}

// MenuSetRepository handles menu set related database operations
type MenuSetRepository interface {
	Create(ctx context.Context, menuSet *model.MenuSet) error
	FindByID(ctx context.Context, id uint) (*model.MenuSet, error)
	FindAll(ctx context.Context) ([]model.MenuSet, error)
	FindActive(ctx context.Context, conditions map[string]interface{}) ([]model.MenuSet, error)
	Update(ctx context.Context, menuSet *model.MenuSet) error
	Delete(ctx context.Context, menuSet *model.MenuSet) error
	HardDelete(ctx context.Context, menuSet *model.MenuSet) error
	AddMenuItem(ctx context.Context, setItem *model.MenuSetItem) error
	RemoveMenuItem(ctx context.Context, setItem *model.MenuSetItem) error
	FindMenuItems(ctx context.Context, menuSetID uint) ([]model.MenuItem, error)
}

// MenuItemRepository handles menu item related database operations
type MenuItemRepository interface {
	Create(ctx context.Context, item *model.MenuItem) error
	FindByID(ctx context.Context, id uint) (*model.MenuItem, error)
	FindAll(ctx context.Context) ([]model.MenuItem, error)
	FindActive(ctx context.Context, conditions map[string]interface{}) ([]model.MenuItem, error)
	Update(ctx context.Context, item *model.MenuItem) error
	Delete(ctx context.Context, item *model.MenuItem) error
	HardDelete(ctx context.Context, item *model.MenuItem) error
}

// NotificationRepository defines notification-specific operations
type NotificationRepository interface {
	BaseRepository[model.Notification]
	FindByUserID(ctx context.Context, userID uint) ([]model.Notification, error)
	CountUnreadByUserID(ctx context.Context, userID uint) (int64, error)
	CountUndeliveredByUserID(ctx context.Context, userID uint) (int64, error)
	MarkAsRead(ctx context.Context, id uint) error
	MarkAsDelivered(ctx context.Context, id uint) error
	FindUnreadByUserID(ctx context.Context, userID uint) ([]model.Notification, error)
	FindByType(ctx context.Context, userID uint, notificationType model.NotificationType) ([]model.Notification, error)
}

// MealRequestRepository handles meal request related database operations
type MealRequestRepository interface {
	Create(ctx context.Context, request *model.MealRequest) error
	FindByID(ctx context.Context, id uint) (*model.MealRequest, error)
	FindAll(ctx context.Context) ([]model.MealRequest, error)
	FindActive(ctx context.Context, conditions map[string]interface{}) ([]model.MealRequest, error)
	Update(ctx context.Context, request *model.MealRequest) error
	Delete(ctx context.Context, request *model.MealRequest) error
	HardDelete(ctx context.Context, request *model.MealRequest) error
	FindByUserID(ctx context.Context, userID uint) ([]model.MealRequest, error)
	FindByMealEventID(ctx context.Context, mealEventID uint) ([]model.MealRequest, error)
	AddRequestItem(ctx context.Context, item *model.MealRequestItem) error
	RemoveRequestItem(ctx context.Context, item *model.MealRequestItem) error
	FindRequestItems(ctx context.Context, requestID uint) ([]model.MealRequestItem, error)
	// Additional methods
	FindByDateRange(ctx context.Context, startDate, endDate time.Time) ([]model.MealRequest, error)
	FindPendingRequests(ctx context.Context) ([]model.MealRequest, error)
	FindApprovedRequests(ctx context.Context) ([]model.MealRequest, error)
	FindRejectedRequests(ctx context.Context) ([]model.MealRequest, error)
	UpdateRequestStatus(ctx context.Context, requestID uint, status model.RequestStatus) error
	CountByMealEventID(ctx context.Context, mealEventID uint) (int64, error)
	FindByMenuSetID(ctx context.Context, menuSetID uint) ([]model.MealRequest, error)
	FindWithDetails(ctx context.Context, requestID uint) (*model.MealRequest, error)
}

// MenuItemCommentRepository handles menu item comment related database operations
type MenuItemCommentRepository interface {
	Create(ctx context.Context, comment *model.MenuItemComment) error
	FindByID(ctx context.Context, id uint) (*model.MenuItemComment, error)
	FindAll(ctx context.Context) ([]model.MenuItemComment, error)
	FindActive(ctx context.Context, conditions map[string]interface{}) ([]model.MenuItemComment, error)
	Update(ctx context.Context, comment *model.MenuItemComment) error
	Delete(ctx context.Context, comment *model.MenuItemComment) error
	HardDelete(ctx context.Context, comment *model.MenuItemComment) error
	FindByMealEventID(ctx context.Context, mealEventID uint) ([]model.MenuItemComment, error)
	FindByUserID(ctx context.Context, userID uint) ([]model.MenuItemComment, error)
	FindByMenuItemID(ctx context.Context, menuItemID uint) ([]model.MenuItemComment, error)
	// Additional methods
	FindByDateRange(ctx context.Context, startDate, endDate time.Time) ([]model.MenuItemComment, error)
	FindRecentComments(ctx context.Context, limit int) ([]model.MenuItemComment, error)
	CountByMealEventID(ctx context.Context, mealEventID uint) (int64, error)
	FindWithUserDetails(ctx context.Context, commentID uint) (*model.MenuItemComment, error)
	FindReplies(ctx context.Context, parentID uint) ([]model.MenuItemComment, error) // Added for replies
}
