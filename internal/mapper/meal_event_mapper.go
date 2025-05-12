package mapper

import (
	"github.com/arafat-hasan/mealsync/internal/dto"
	"github.com/arafat-hasan/mealsync/internal/model"
)

// ToMealEventResponse converts a MealEvent model to a MealEventResponse DTO
func ToMealEventResponse(event *model.MealEvent) *dto.MealEventResponse {
	if event == nil {
		return nil
	}

	response := &dto.MealEventResponse{
		BaseAuditResponse: dto.BaseAuditResponse{
			BaseResponse: dto.BaseResponse{
				ID:        event.ID,
				CreatedAt: event.CreatedAt,
				UpdatedAt: event.UpdatedAt,
				DeletedAt: event.DeletedAt,
			},
			CreatedByID: event.CreatedByID,
			CreatedBy:   ToUserResponse(event.CreatedBy),
			UpdatedByID: event.UpdatedByID,
			UpdatedBy:   ToUserResponse(event.UpdatedBy),
		},
		Name:          event.Name,
		Description:   event.Description,
		EventDate:     event.EventDate,
		EventDuration: event.EventDuration,
		CutoffTime:    event.CutoffTime,
		ConfirmedAt:   event.ConfirmedAt,
	}

	// Add menu sets if available
	if len(event.MenuSets) > 0 {
		response.MenuSets = make([]dto.MealEventSetResponse, len(event.MenuSets))
		for i, set := range event.MenuSets {
			response.MenuSets[i] = *ToMealEventSetResponse(&set)
		}
	}

	// Add addresses if available
	if len(event.Addresses) > 0 {
		response.Addresses = make([]dto.MealEventAddressResponse, len(event.Addresses))
		for i, addr := range event.Addresses {
			response.Addresses[i] = *ToMealEventAddressResponse(&addr)
		}
	}

	return response
}

// ToMealEventResponseList converts a slice of MealEvent models to a slice of MealEventResponse DTOs
func ToMealEventResponseList(events []model.MealEvent) []dto.MealEventResponse {
	if events == nil {
		return nil
	}

	responses := make([]dto.MealEventResponse, len(events))
	for i, event := range events {
		resp := ToMealEventResponse(&event)
		if resp != nil {
			responses[i] = *resp
		}
	}

	return responses
}

// ToMealEventSetResponse converts a MealEventSet model to a MealEventSetResponse DTO
func ToMealEventSetResponse(set *model.MealEventSet) *dto.MealEventSetResponse {
	if set == nil {
		return nil
	}

	response := &dto.MealEventSetResponse{
		BaseAuditResponse: dto.BaseAuditResponse{
			BaseResponse: dto.BaseResponse{
				ID:        set.ID,
				CreatedAt: set.CreatedAt,
				UpdatedAt: set.UpdatedAt,
				DeletedAt: set.DeletedAt,
			},
			CreatedByID: set.CreatedByID,
			CreatedBy:   ToUserResponse(set.CreatedBy),
			UpdatedByID: set.UpdatedByID,
			UpdatedBy:   ToUserResponse(set.UpdatedBy),
		},
		MealEventID: set.MealEventID,
		MenuSetID:   set.MenuSetID,
		Label:       set.Label,
		Note:        set.Note,
	}

	// Add menu set if available
	if set.MenuSet != nil {
		menuSetResponse := ToMenuSetResponse(set.MenuSet)
		response.MenuSet = menuSetResponse
	}

	return response
}

// ToMealEventAddressResponse converts a MealEventAddress model to a MealEventAddressResponse DTO
func ToMealEventAddressResponse(addr *model.MealEventAddress) *dto.MealEventAddressResponse {
	if addr == nil {
		return nil
	}

	response := &dto.MealEventAddressResponse{
		BaseAuditResponse: dto.BaseAuditResponse{
			BaseResponse: dto.BaseResponse{
				ID:        addr.ID,
				CreatedAt: addr.CreatedAt,
				UpdatedAt: addr.UpdatedAt,
				DeletedAt: addr.DeletedAt,
			},
			CreatedByID: addr.CreatedByID,
			CreatedBy:   ToUserResponse(addr.CreatedBy),
			UpdatedByID: addr.UpdatedByID,
			UpdatedBy:   ToUserResponse(addr.UpdatedBy),
		},
		MealEventID: addr.MealEventID,
		AddressID:   addr.AddressID,
	}

	// Add address if available
	if addr.Address.ID != 0 {
		addrResponse := ToEventAddressResponse(&addr.Address)
		response.Address = addrResponse
	}

	return response
}

// ToMenuSetResponse converts a MenuSet model to a MenuSetResponse DTO
func ToMenuSetResponse(set *model.MenuSet) *dto.MenuSetResponse {
	if set == nil {
		return nil
	}

	response := &dto.MenuSetResponse{
		BaseAuditResponse: dto.BaseAuditResponse{
			BaseResponse: dto.BaseResponse{
				ID:        set.ID,
				CreatedAt: set.CreatedAt,
				UpdatedAt: set.UpdatedAt,
				DeletedAt: set.DeletedAt,
			},
			CreatedByID: set.CreatedByID,
			CreatedBy:   ToUserResponse(set.CreatedBy),
			UpdatedByID: set.UpdatedByID,
			UpdatedBy:   ToUserResponse(set.UpdatedBy),
		},
		Name:        set.MenuSetName,
		Description: set.MenuSetDescription,
	}

	// Add menu items if available
	if len(set.MenuSetItems) > 0 {
		menuItems := make([]dto.MenuItemResponse, len(set.MenuSetItems))
		for i, item := range set.MenuSetItems {
			menuItems[i] = *ToMenuItemResponse(&item.MenuItem)
		}
		response.MenuItems = menuItems
	}

	return response
}

// ToEventAddressResponse converts an EventAddress model to an EventAddressResponse DTO
func ToEventAddressResponse(addr *model.EventAddress) *dto.EventAddressResponse {
	if addr == nil {
		return nil
	}

	return &dto.EventAddressResponse{
		BaseAuditResponse: dto.BaseAuditResponse{
			BaseResponse: dto.BaseResponse{
				ID:        addr.ID,
				CreatedAt: addr.CreatedAt,
				UpdatedAt: addr.UpdatedAt,
				DeletedAt: addr.DeletedAt,
			},
			CreatedByID: addr.CreatedByID,
			CreatedBy:   ToUserResponse(addr.CreatedBy),
			UpdatedByID: addr.UpdatedByID,
			UpdatedBy:   ToUserResponse(addr.UpdatedBy),
		},
		Name:        addr.Label,
		AddressLine: addr.Address,
	}
}

// ToMenuItemResponse converts a MenuItem model to a MenuItemResponse DTO
func ToMenuItemResponse(item *model.MenuItem) *dto.MenuItemResponse {
	if item == nil {
		return nil
	}

	return &dto.MenuItemResponse{
		BaseAuditResponse: dto.BaseAuditResponse{
			BaseResponse: dto.BaseResponse{
				ID:        item.ID,
				CreatedAt: item.CreatedAt,
				UpdatedAt: item.UpdatedAt,
				DeletedAt: item.DeletedAt,
			},
			CreatedByID: item.CreatedByID,
			CreatedBy:   ToUserResponse(item.CreatedBy),
			UpdatedByID: item.UpdatedByID,
			UpdatedBy:   ToUserResponse(item.UpdatedBy),
		},
		Name:          item.Name,
		Description:   item.Description,
		ImageURL:      item.ImageURL,
		AverageRating: item.AverageRating,
	}
}
