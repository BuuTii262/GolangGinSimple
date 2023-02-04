package controller

import (
	"fmt"
	"net/http"

	"MyGO.com/m/dto"
	"MyGO.com/m/helper"
	"MyGO.com/m/model"
	"MyGO.com/m/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (c *userController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.ResponseErrorData(501, errDTO.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}
	err := c.userService.IsDuplicateEmail(registerDTO.Email)
	fmt.Println("Here log the return err is true or false-------", err)
	if err {
		response := helper.ResponseErrorData(502, "Email is duplicate")
		ctx.JSON(http.StatusOK, response)
		return
	}

	createUser := c.userService.CreateUser(registerDTO)
	response := helper.ResponseData(0, "Success", createUser)
	ctx.JSON(http.StatusOK, response)
}

func (c *userController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.ResponseErrorData(503, errDTO.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}
	loginResult := c.userService.VerifyLogin(loginDTO.Name, loginDTO.Password)
	if _, ok := loginResult.(model.User); ok {
		response := helper.ResponseData(0, "Login success full", helper.EmptyObj{})
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helper.ResponseErrorData(504, "Invalid uesr name or password")
	ctx.JSON(http.StatusOK, response)
}
