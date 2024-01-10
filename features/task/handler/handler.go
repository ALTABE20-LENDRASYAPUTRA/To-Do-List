package handler

import (
	"lendra/todo/features/task"
	"lendra/todo/utils/middlewares"
	"lendra/todo/utils/responses"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	taskService task.TaskServiceInterface
}

func New(ps task.TaskServiceInterface) *TaskHandler {
	return &TaskHandler{
		taskService: ps,
	}
}

func (handler *TaskHandler) CreateTask(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	userID, err := middlewares.ExtractTokenUserId(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Unauthorized User",
		})
	}

	newTask := TaskRequest{}
	errBind := c.Bind(&newTask)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	taskCore := RequestToCore(newTask)
	taskCore.UserID = userID

	errInsert := handler.taskService.Create(taskCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data", nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success insert data", nil))
}

func (handler *TaskHandler) DeleteTask(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	userID, err := middlewares.ExtractTokenUserId(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Unauthorized User",
		})
	}

	taskID, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error parsing task id", nil))
	}

	errDelete := handler.taskService.Delete(userID, uint(taskID))

	if errDelete != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error delete data", nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success delete data", nil))
}

func (handler *TaskHandler) UpdateStatusTask(c echo.Context) error {
token := c.Get("user").(*jwt.Token)
	userID, err := middlewares.ExtractTokenUserId(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Unauthorized User",
		})
	}

	taskID, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error parsing task id", nil))
	}

	updateTask := TaskPutRequest{}
	errBind := c.Bind(&updateTask)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	taskCore := RequestPutToCore(updateTask)
	taskCore.ID = uint(taskID)
	taskCore.UserID = userID

	errUpdate := handler.taskService.Update(userID, taskCore.ID, taskCore)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error update data", nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success update data", nil))
}