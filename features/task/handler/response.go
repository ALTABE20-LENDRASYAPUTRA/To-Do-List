package handler

import (
	"lendra/todo/features/task"
	"time"
)

type TaskResponse struct {
	ID          uint      `json:"id" form:"id"`
	Name        string    `json:"name" form:"name"`
	Description string    `json:"description" form:"description"`
	Status      string    `json:"status" form:"status"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
	UserID      uint      `json:"user_id" form:"user_id"`
}

func statusToString(status bool) string {
	if status {
		return "completed"
	}
	return "not completed"
}

func CoreToResponse(data task.TaskCore) TaskResponse {
	return TaskResponse{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		Status:      statusToString(data.Status),
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		UserID:      data.UserID,
	}
}

func CoreToResponseList(data []task.TaskCore) []TaskResponse {
	var results []TaskResponse
	for _, v := range data {
		results = append(results, CoreToResponse(v))
	}
	return results
}
