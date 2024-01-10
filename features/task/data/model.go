package data

import (
	"lendra/todo/features/task"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name        string
	Description string
	Status      bool
	ProjectID   uint
	UserID      uint
}

func CoreToModel(input task.TaskCore) Task {
	return Task{
		Name:        input.Name,
		Description: input.Description,
		Status:      input.Status,
		ProjectID:   input.ProjectID,
		UserID:      input.UserID,
	}
}

func (t Task) ModelToCore() task.TaskCore {
	return task.TaskCore{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
		Status:      t.Status,
		ProjectID:   t.ProjectID,
		UserID:      t.UserID,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}