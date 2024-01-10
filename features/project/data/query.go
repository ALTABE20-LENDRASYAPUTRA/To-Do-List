package data

import (
	"errors"
	"lendra/todo/features/project"
	"lendra/todo/features/task"

	"gorm.io/gorm"
)

type projectQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) project.ProjectDataInterface {
	return &projectQuery{
		db: db,
	}
}

// Insert implements project.ProjectDataInterface.
func (repo *projectQuery) Insert(UserID uint, input project.ProjectCore) error {
	projectInputGorm := Project{
		UserID:      UserID,
		Name:        input.Name,
		Description: input.Description,
	}

	// simpan ke DB
	tx := repo.db.Create(&projectInputGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}
	return nil
}

// SelectAll implements project.ProjectDataInterface.
func (repo *projectQuery) SelectAll(UserID uint) ([]project.ProjectCore, error) {
	var projects []Project
	err := repo.db.Where("user_id = ?", UserID).Find(&projects).Error
	if err != nil {
		return nil, err
	}

	var projectCores []project.ProjectCore
	for _, p := range projects {
		projectCores = append(projectCores, p.ModelToCore())
	}

	return projectCores, nil
}

// Update implements project.ProjectDataInterface.
func (repo *projectQuery) Update(UserID uint, input project.ProjectCore) error {
	var projects Project

	if err := repo.db.First(&projects, input.ID).Error; err != nil {
		return errors.New("project not found")
	}

	if projects.UserID != UserID {
		return errors.New("you are not authorized to update this project")
	}

	if err := repo.db.Model(&projects).Updates(Project{
		Name:        input.Name,
		Description: input.Description,
	}).Error; err != nil {
		return err
	}

	return nil
}

// Delete implements project.ProjectDataInterface.
func (repo *projectQuery) Delete(UserID uint, ProjectID uint) error {
	var projects Project

	if err := repo.db.First(&projects, ProjectID).Error; err != nil {
		return errors.New("project not found")
	}

	if projects.UserID != UserID {
		return errors.New("you are not authorized to delete this project")
	}

	if err := repo.db.Delete(&projects).Error; err != nil {
		return err
	}

	return nil
}

// SelectProjectById implements project.ProjectDataInterface.
func (repo *projectQuery) SelectProjectById(UserID uint, ProjectID uint) (project.ProjectCore, error) {
	var projectModel Project

	if err := repo.db.Preload("Tasks").First(&projectModel, ProjectID).Error; err != nil {
		return project.ProjectCore{}, errors.New("project not found")
	}

	if projectModel.UserID != UserID {
		return project.ProjectCore{}, errors.New("you are not authorized to access this project")
	}

	var tasksDataCore []task.TaskCore
	for _, taskModel := range projectModel.Tasks {
		taskCore := taskModel.ModelToCore()
		tasksDataCore = append(tasksDataCore, taskCore)
	}

	// Proses mapping dari struct gorm model ke struct core
	projectCore := project.ProjectCore{
		ID:          projectModel.ID,
		Name:        projectModel.Name,
		Description: projectModel.Description,
		UserID:      projectModel.UserID,
		Tasks:       tasksDataCore,
	}

	return projectCore, nil
}
