package handler

import (
	"errors"
	"lendra/todo/features/user"
	"lendra/todo/utils/middlewares"
	"lendra/todo/utils/responses"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func New(service user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) CreateUser(c echo.Context) error {
	newUser := UserRequest{}
	errBind := c.Bind(&newUser) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(newUser); err != nil {
		c.Echo().Logger.Error("Input error :", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
			"data":    nil,
		})
	}

	//mapping dari request ke core
	userCore := RequestToCore(newUser)

	errInsert := handler.userService.Create(userCore)
	if errInsert != nil {
		c.Logger().Error("ERROR Register, explain:", errInsert.Error())
		if strings.Contains(errInsert.Error(), "Duplicate entry") {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "data yang diinputkan sudah terdaftar pada sistem",
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "terjadi permasalahan ketika memproses data",
		})
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success insert data", nil))
}

func (handler *UserHandler) Login(c echo.Context) error {
	var reqData = LoginRequest{}
	errBind := c.Bind(&reqData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	result, token, err := handler.userService.Login(reqData.Email, reqData.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error login "+err.Error(), nil))
	}
	responseData := map[string]any{
		"token": token,
		"nama":  result.Name,
	}
	return c.JSON(http.StatusOK, responses.WebResponse("success login", responseData))
}

func (handler *UserHandler) UpdateUser(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	userID, err := middlewares.ExtractTokenUserId(token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message": "Unauthorized user",
		})
	}

	var userData = UserRequest{}
	errBind := c.Bind(&userData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	userCore := RequestToCore(userData)
	_, err = handler.userService.Update(userID, userCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error update data "+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success update data", nil))
}

func (handler *UserHandler) DeleteUser(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	userID, err := middlewares.ExtractTokenUserId(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Unauthorized User",
		})
	}

	err = handler.userService.Delete(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Error deleting user",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "User deleted successfully",
	})
}

func (handler *UserHandler) GetUserProfil(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	userID, err := middlewares.ExtractTokenUserId(token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Unauthorized User",
		})
	}

	userData, err := handler.userService.GetUserById(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, responses.WebResponse("user not found", nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error getting user data", nil))
	}

	userResult := CoreToResponse(userData)

	return c.JSON(http.StatusOK, responses.WebResponse("success get user data", userResult))
}
