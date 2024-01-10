package handler

import "lendra/todo/features/task"

type TaskRequest struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	ProjectID   uint `json:"project_id" form:"project_id"`
}

type TaskPutRequest struct {
	Status        bool `json:"status" form:"status"`
}

func RequestToCore(input TaskRequest) task.TaskCore {
	return task.TaskCore{
		Name:        input.Name,
		Description: input.Description,
		ProjectID:   input.ProjectID,
	}
}


func RequestPutToCore(input TaskPutRequest) task.TaskCore {
	return task.TaskCore{
		Status: input.Status,
	}
}
