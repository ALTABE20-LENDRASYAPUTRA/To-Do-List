package data

import (
	"lendra/todo/features/project"
	td "lendra/todo/features/task/data"
	ud "lendra/todo/features/user/data"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name        string
	Description string
	User        ud.User `gorm:"foreignKey:UserID"`
	UserID      uint
	Tasks       []td.Task
}

func CoreToModel(input project.ProjectCore) Project {
	return Project{
		Name:        input.Name,
		Description: input.Description,
		UserID:      input.UserID,
	}
}

func (p Project) ModelToCore() project.ProjectCore {
	return project.ProjectCore{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		UserID:      p.UserID,
	}
}
