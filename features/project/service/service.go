package service

import (
	"lendra/todo/features/project"
)

type projectService struct {
	projectData project.ProjectDataInterface
}
// dependency injection
func New(repo project.ProjectDataInterface) project.ProjectServiceInterface {
	return &projectService{
		projectData: repo,
	}
}

// Create implements project.ProjectServiceInterface.
func (ps *projectService) Create(input project.ProjectCore) error {
	err := ps.projectData.Insert(input.UserID, input)
	if err != nil {
		return err
	}
	return nil
}

// GetAll implements project.ProjectServiceInterface.
func (ps *projectService) GetAll(UserID uint) ([]project.ProjectCore, error) {
	projects, err := ps.projectData.SelectAll(UserID)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

// Update implements project.ProjectServiceInterface.
func (ps *projectService) Update(UserID uint, input project.ProjectCore) error {
	err := ps.projectData.Update(UserID, input)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements project.ProjectServiceInterface.
func (ps *projectService) Delete(UserID uint, ProjectID uint) error {
	err := ps.projectData.Delete(UserID, ProjectID)
	if err != nil {
		return err
	}

	return nil
}

// GetProjectById implements project.ProjectServiceInterface.
func (ps *projectService) GetProjectById(UserID uint, ProjectID uint) (project.ProjectCore, error) {
	projectData, err := ps.projectData.SelectProjectById(UserID, ProjectID)
	if err != nil {
		return project.ProjectCore{}, err
	}

	return projectData, nil
}

