package service

import "lendra/todo/features/task"

type taskService struct {
	taskData task.TaskDataInterface
}

// dependency injection
func New(repo task.TaskDataInterface) task.TaskServiceInterface {
	return &taskService{
		taskData: repo,
	}
}

// Create implements task.TaskServiceInterface.
func (ps *taskService) Create(input task.TaskCore) error {
	err := ps.taskData.Insert(input.UserID, input)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements task.TaskServiceInterface.
func (ps *taskService) Delete(UserID uint, TaskID uint) error {
	err := ps.taskData.Delete(UserID, TaskID)
	if err != nil {
		return err
	}

	return nil
}

// Update implements task.TaskServiceInterface.
func (ps *taskService) Update(UserID uint, TaskID uint, input task.TaskCore) error {
	err := ps.taskData.Update(UserID, TaskID, input)
	if err != nil {
		return err
	}
	return nil
}
