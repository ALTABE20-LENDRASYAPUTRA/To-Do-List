package handler

import "lendra/todo/features/project"

type ProjectRequest struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
}

func RequestToCore(input ProjectRequest) project.ProjectCore {
	return project.ProjectCore{
		Name:        input.Name,
		Description: input.Description,
	}
}
