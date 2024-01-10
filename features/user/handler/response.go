package handler

import (
	"lendra/todo/features/user"
	"time"
)

type UserResponse struct {
	ID          uint      `json:"id" form:"id"`
	Name        string    `json:"name" form:"name"`
	Email       string    `json:"email" form:"email"`
	Address     string    `json:"address" form:"address"`
	PhoneNumber string    `json:"phone_number" form:"phone_number"`
	Role        string    `json:"role" form:"role"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
}

type UserProResponse struct {
	ID   uint   `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Role string `json:"role" form:"role"`
}

func CoreToResponse(data user.Core) UserResponse {
	return UserResponse{
		ID:          data.ID,
		Name:        data.Name,
		Email:       data.Email,
		PhoneNumber: data.PhoneNumber,
		Role:        data.Role,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func CoreToResponseList(data []user.Core) []UserResponse {
	var results []UserResponse
	for _, v := range data {
		results = append(results, CoreToResponse(v))
	}
	return results
}
