package routers

import (
	ud "lendra/todo/features/user/data"
	uh "lendra/todo/features/user/handler"
	us "lendra/todo/features/user/service"
	pd "lendra/todo/features/project/data"
	ph "lendra/todo/features/project/handler"
	ps "lendra/todo/features/project/service"
	td "lendra/todo/features/task/data"
	th "lendra/todo/features/task/handler"
	ts "lendra/todo/features/task/service"
	"lendra/todo/utils/encrypts"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userData := ud.New(db)
	hash := encrypts.New()
	userService := us.New(userData, hash)
	userHandlerAPI := uh.New(userService)

	projectData := pd.New(db)
	projectService := ps.New(projectData)
	projectHandlerAPI := ph.New(projectService)

	taskData := td.New(db)
	taskService := ts.New(taskData)
	taskHandlerAPI := th.New(taskService)

	// define routes/ endpoint user
	e.POST("/login", userHandlerAPI.Login)
	e.POST("/users", userHandlerAPI.CreateUser)
	e.PUT("/users", userHandlerAPI.UpdateUser, echojwt.JWT([]byte("K3YYKijS879!!")))
	e.DELETE("/users", userHandlerAPI.DeleteUser, echojwt.JWT([]byte("K3YYKijS879!!")))
	e.GET("/users", userHandlerAPI.GetUserProfil, echojwt.JWT([]byte("K3YYKijS879!!")))

	// define routes/ endpoint project
	e.POST("/projects", projectHandlerAPI.CreateProject, echojwt.JWT([]byte("K3YYKijS879!!")))
	e.GET("/projects", projectHandlerAPI.GetAllProject, echojwt.JWT([]byte("K3YYKijS879!!")))
	e.PUT("/projects/:project_id", projectHandlerAPI.UpdateProject, echojwt.JWT([]byte("K3YYKijS879!!")))
	e.DELETE("/projects/:project_id", projectHandlerAPI.DeleteProject, echojwt.JWT([]byte("K3YYKijS879!!")))
	e.GET("/projects/:project_id", projectHandlerAPI.GetProjectById, echojwt.JWT([]byte("K3YYKijS879!!")))

	// define routes/ endpoint task
	e.POST("tasks", taskHandlerAPI.CreateTask, echojwt.JWT([]byte("K3YYKijS879!!")))
	e.DELETE("tasks/:task_id", taskHandlerAPI.DeleteTask, echojwt.JWT([]byte("K3YYKijS879!!")))
	e.PUT("/tasks/:task_id", taskHandlerAPI.UpdateStatusTask, echojwt.JWT([]byte("K3YYKijS879!!")))
}
