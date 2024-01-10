package handler

import (
	"lendra/todo/features/project"
	"lendra/todo/features/task/handler"
	"time"
)

type ProjectTaskResponse struct {
	ID          uint      `json:"id" form:"id"`
	Name        string    `json:"name" form:"name"`
	Description string    `json:"description" form:"description"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
	UserID      uint      `json:"user_id" form:"user_id"`
	Tasks       []handler.TaskResponse
}

type ProjectResponse struct {
	ID          uint      `json:"id" form:"id"`
	Name        string    `json:"name" form:"name"`
	Description string    `json:"description" form:"description"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
	UserID      uint      `json:"user_id" form:"user_id"`
}

func CoreToResponseTask(data project.ProjectCore) ProjectTaskResponse {
	var tasks []handler.TaskResponse
	for _, task := range data.Tasks {
		tasks = append(tasks, handler.CoreToResponse(task))
	}

	return ProjectTaskResponse{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		UserID:      data.UserID,
		Tasks:       tasks,
	}
}

func CoreToResponseTaskList(data []project.ProjectCore) []ProjectTaskResponse {
	var results []ProjectTaskResponse
	for _, v := range data {
		results = append(results, CoreToResponseTask(v))
	}
	return results
}

func CoreToResponse(data project.ProjectCore) ProjectResponse {
	return ProjectResponse{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		UserID:      data.UserID,
	}
}

func CoreToResponseList(data []project.ProjectCore) []ProjectResponse {
	var results []ProjectResponse
	for _, v := range data {
		results = append(results, CoreToResponse(v))
	}
	return results
}
