package mapper

import (
	"github.com/arafat-hasan/mealsync/internal/dto"
	"github.com/arafat-hasan/mealsync/internal/model"
)

// ToUserResponse converts a User model to a UserResponse DTO
func ToUserResponse(user *model.User) *dto.UserResponse {
	if user == nil {
		return nil
	}

	return &dto.UserResponse{
		BaseResponse: dto.BaseResponse{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		},
		EmployeeID:          user.EmployeeID,
		Username:            user.Username,
		Name:                user.Name,
		Email:               user.Email,
		Department:          user.Department,
		Role:                user.Role,
		NotificationEnabled: user.NotificationEnabled,
		IsActive:            user.IsActive,
	}
}

// ToUserResponseList converts a slice of User models to a slice of UserResponse DTOs
func ToUserResponseList(users []model.User) []dto.UserResponse {
	if users == nil {
		return nil
	}

	userResponses := make([]dto.UserResponse, len(users))
	for i, user := range users {
		userResp := ToUserResponse(&user)
		if userResp != nil {
			userResponses[i] = *userResp
		}
	}

	return userResponses
}

// ToTokenResponse converts a service.TokenPair and User model to a TokenResponse DTO
func ToTokenResponse(accessToken, refreshToken string, user *model.User) *dto.TokenResponse {
	if user == nil {
		return nil
	}

	return &dto.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         *ToUserResponse(user),
	}
}

// ToUser converts a UserCreateRequest DTO to a User model
func ToUser(request *dto.UserCreateRequest) *model.User {
	if request == nil {
		return nil
	}

	return &model.User{
		EmployeeID:          request.EmployeeID,
		Username:            request.Username,
		Password:            request.Password, // Note: password will be hashed in service layer
		Name:                request.Name,
		Email:               request.Email,
		Department:          request.Department,
		Role:                request.Role,
		NotificationEnabled: request.NotificationEnabled,
		IsActive:            true, // Default value
	}
}

// UpdateUserFromDTO updates a User model using a UserUpdateRequest DTO
func UpdateUserFromDTO(user *model.User, request *dto.UserUpdateRequest) {
	if user == nil || request == nil {
		return
	}

	if request.Username != nil {
		user.Username = *request.Username
	}

	if request.Password != nil {
		user.Password = *request.Password // Will be hashed in service layer
	}

	if request.Name != nil {
		user.Name = *request.Name
	}

	if request.Email != nil {
		user.Email = *request.Email
	}

	if request.Department != nil {
		user.Department = *request.Department
	}

	if request.Role != nil {
		user.Role = *request.Role
	}

	if request.NotificationEnabled != nil {
		user.NotificationEnabled = *request.NotificationEnabled
	}

	if request.IsActive != nil {
		user.IsActive = *request.IsActive
	}
}
