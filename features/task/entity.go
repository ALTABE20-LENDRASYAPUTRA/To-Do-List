package task

import "time"

type TaskCore struct {
	ID          uint
	Name        string
	Description string
	Status      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ProjectID   uint
	UserID      uint
}

// interface untuk Data Layer
type TaskDataInterface interface {
	Insert(UserID uint, input TaskCore) error
	Delete(UserID uint, TaskID uint) error
	Update(UserID uint, TaskID uint, input TaskCore) error
}

// interface untuk Service Layer
type TaskServiceInterface interface {
	Create(input TaskCore) error
	Delete(UserID uint, TaskID uint) error
	Update(UserID uint, TaskID uint, input TaskCore) error
}
