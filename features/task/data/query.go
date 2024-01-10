package data

import (
	"errors"
	"lendra/todo/features/task"

	"gorm.io/gorm"
)

type taskQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) task.TaskDataInterface {
	return &taskQuery{
		db: db,
	}
}

// Insert implements task.TaskDataInterface.
func (repo *taskQuery) Insert(UserID uint, input task.TaskCore) error {
	taskInputGorm := Task{
		UserID:      UserID,
		ProjectID:   input.ProjectID,
		Name:        input.Name,
		Description: input.Description,
	}

	// simpan ke DB
	tx := repo.db.Create(&taskInputGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}
	return nil
}

// Delete implements task.TaskDataInterface.
func (repo *taskQuery) Delete(UserID uint, TaskID uint) error {
	var tasks Task

	if err := repo.db.First(&tasks, TaskID).Error; err != nil {
		return errors.New("comment not found")
	}

	if tasks.UserID != UserID {
		return errors.New("you are not authorized to delete this task")
	}

	if err := repo.db.Delete(&tasks).Error; err != nil {
		return err
	}

	return nil
}

// Update implements task.TaskDataInterface.
func (repo *taskQuery) Update(UserID uint, TaskID uint, input task.TaskCore) error {
	var tasks Task

	if err := repo.db.First(&tasks, input.ID).Error; err != nil {
		return errors.New("project not found")
	}

	if tasks.UserID != UserID {
		return errors.New("you are not authorized to update this task")
	}

	tasks.Status = input.Status

	if err := repo.db.Save(&tasks).Error; err != nil {
		return err
	}

	return nil
}
