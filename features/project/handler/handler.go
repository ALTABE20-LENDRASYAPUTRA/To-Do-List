package handler

import (
	"lendra/todo/features/project"
	"lendra/todo/utils/middlewares"
	"lendra/todo/utils/responses"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ProjectHandler struct {
	projectService project.ProjectServiceInterface
}

func New(ps project.ProjectServiceInterface) *ProjectHandler {
	return &ProjectHandler{
		projectService: ps,
	}
}

func (handler *ProjectHandler) CreateProject(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	userID, err := middlewares.ExtractTokenUserId(token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message": "Unauthorized User",
		})
	}

	newProject := ProjectRequest{}
	errBind := c.Bind(&newProject)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	projectCore := RequestToCore(newProject)
	projectCore.UserID = userID

	errInsert := handler.projectService.Create(projectCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data", nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success insert data", nil))
}

func (handler *ProjectHandler) GetAllProject(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	userID, err := middlewares.ExtractTokenUserId(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Unauthorized User",
		})
	}

	projects, err := handler.projectService.GetAll(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("Error getting projects", nil))
	}

	var projectResponses []ProjectResponse
	for _, project := range projects {
		projectResponses = append(projectResponses, CoreToResponse(project))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Success", projectResponses))
}

func (handler *ProjectHandler) UpdateProject(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	userID, err := middlewares.ExtractTokenUserId(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Unauthorized User",
		})
	}

	projectID, err := strconv.Atoi(c.Param("project_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error parsing project id", nil))
	}

	updateProject := ProjectRequest{}
	errBind := c.Bind(&updateProject)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	projectCore := RequestToCore(updateProject)
	projectCore.ID = uint(projectID)
	projectCore.UserID = userID

	errUpdate := handler.projectService.Update(userID, projectCore)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error update data", nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success update data", nil))
}

func (handler *ProjectHandler) DeleteProject(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	userID, err := middlewares.ExtractTokenUserId(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Unauthorized User",
		})
	}

	projectID, err := strconv.Atoi(c.Param("project_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error parsing project id", nil))
	}

	errDelete := handler.projectService.Delete(userID, uint(projectID))

	if errDelete != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error delete data", nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success delete data", nil))
}

func (handler *ProjectHandler) GetProjectById(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	userID, err := middlewares.ExtractTokenUserId(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Unauthorized User",
		})
	}

	projectID, err := strconv.Atoi(c.Param("project_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error parsing project id", nil))
	}

	// Panggil method GetProjectById dari service
	projectData, err := handler.projectService.GetProjectById(userID, uint(projectID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("Error getting project details", nil))
	}

	projectResponse := CoreToResponseTask(projectData)

	return c.JSON(http.StatusOK, responses.WebResponse("Success", projectResponse))
}