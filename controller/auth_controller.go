package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"MyGO.com/m/dto"
	"MyGO.com/m/helper"
	"MyGO.com/m/model"
	"MyGO.com/m/service"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type authController struct {
	userService service.UserService
	jwtService  service.JwtService
}

func NewAuthContrller(userService service.UserService, jwtService service.JwtService) AuthController {
	return &authController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *authController) Register(ctx *gin.Context) {
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
	generateToken := c.jwtService.GenerateToken(strconv.FormatUint(createUser.ID, 10))
	createUser.Token = generateToken
	response := helper.ResponseData(0, "Success", createUser)

	ctx.JSON(http.StatusOK, response)
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.ResponseErrorData(503, errDTO.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}
	loginResult := c.userService.VerifyLogin(loginDTO.Name, loginDTO.Password)
	if v, ok := loginResult.(model.User); ok {
		generateToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generateToken
		response := helper.ResponseData(0, "Login successfull", v)
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helper.ResponseErrorData(504, "Invalid uesr name or password")
	ctx.JSON(http.StatusOK, response)

}
