package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"MyGO.com/m/dto"
	"MyGO.com/m/helper"
	"MyGO.com/m/model"
	"MyGO.com/m/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetWelcome(ctx *gin.Context)
	GetAllUsers(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JwtService
}

func NewUserController(userService service.UserService, jwtService service.JwtService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
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
	generateToken := c.jwtService.GenerateToken(strconv.FormatUint(createUser.ID, 10))
	createUser.Token = generateToken
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

type ResponseWelcomeStruct struct {
	ID uint64
}

func (c *userController) GetWelcome(ctx *gin.Context) {

	authHeader := ctx.GetHeader("Authorization")
	splitToken := strings.Split(authHeader, "Bearer ")
	authHeader = splitToken[1]
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		response := helper.ResponseErrorData(401, "Token error !")
		ctx.JSON(http.StatusOK, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		res := helper.ResponseErrorData(400, err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var responseUser ResponseWelcomeStruct
	responseUser.ID = id

	response := helper.ResponseData(0, "Success", responseUser)

	ctx.JSON(http.StatusOK, response)
}

type UserListData struct {
	List  []model.User `json:"list"`
	Total int64        `json:"total"`
}

func (c *userController) GetAllUsers(ctx *gin.Context) {

	authHeader := ctx.GetHeader("Authorization")
	splitToken := strings.Split(authHeader, "Bearer ")
	authHeader = splitToken[1]
	_, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		response := helper.ResponseErrorData(401, "Token error !")
		ctx.JSON(http.StatusOK, response)
		return
	}

	req := &dto.UserGetRequest{}

	if err := ctx.ShouldBind(&req); err != nil {
		response := helper.ResponseErrorData(500, "Internal server error !")
		ctx.JSON(http.StatusOK, response)
		return
	}

	result, count, err := c.userService.GetAllUsers(req)

	if err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	var responseList UserListData

	responseList.List = result
	responseList.Total = count

	response := helper.ResponseData(0, "success", responseList)
	ctx.JSON(http.StatusOK, response)

}
