package project

import (
	"lendra/todo/features/task"
	"time"
)

type ProjectCore struct {
	ID          uint
	Name        string
	Description string
	UserID      uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Tasks       []task.TaskCore
}

// interface untuk Data Layer
type ProjectDataInterface interface {
	Insert(UserID uint, input ProjectCore) error
	SelectAll(UserID uint) ([]ProjectCore, error)
	Update(UserID uint, input ProjectCore) error
	Delete(UserID uint, ProjectID uint) error
	SelectProjectById(UserID uint, ProjectID uint) (ProjectCore, error)
}

// interface untuk Service Layer
type ProjectServiceInterface interface {
	Create(input ProjectCore) error
	GetAll(UserID uint) ([]ProjectCore, error)
	Update(UserID uint, input ProjectCore) error
	Delete(UserID uint, ProjectID uint) error
	GetProjectById(UserID uint, ProjectID uint) (ProjectCore, error)
}
