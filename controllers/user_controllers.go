package controllers

import (
	"net/http"
	"restful-api/models"
	response "restful-api/responses"

	"github.com/labstack/echo/v4"
)

type UserModel interface {
	GetAllUsers() ([]models.User, error)
}

type UserController struct {
	model UserModel
}

func NewUserController(mod UserModel) UserController {
	return UserController{model: mod}
}

func (user_controller *UserController) GetAllUsersController(c echo.Context) error {
	users, err := user_controller.model.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success", users))
}
