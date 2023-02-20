package service

import (
	"fmt"

	"MyGO.com/m/dto"
	"MyGO.com/m/model"
	"MyGO.com/m/repository"
	"github.com/mashingan/smapping"
)

type UserService interface {
	CreateUser(user dto.RegisterDTO) model.User
	IsDuplicateEmail(email string) bool
	VerifyLogin(name string, password string) interface{}
	GetAllUsers() []model.User
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (service userService) CreateUser(user dto.RegisterDTO) model.User {
	userToCreate := model.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		fmt.Println("--------Here is error in repository ------", err)
	}
	res := service.userRepo.InsertUser(userToCreate)
	return res
}

func (service userService) IsDuplicateEmail(email string) bool {
	res := service.userRepo.IsDuplicateEmail(email)
	fmt.Println("____________res____________", res.Error)

	return (res.Error == nil)
}

func (service userService) VerifyLogin(name string, password string) interface{} {
	res := service.userRepo.VerifyLogin(name)
	if v, ok := res.(model.User); ok {
		isPassword := comparePassword(password, v.Password)
		if v.Name == name && isPassword {
			return res
		}
		return false
	}
	return false
}

func comparePassword(enterPass string, resPassword string) bool {
	if enterPass == resPassword {
		return true
	}
	return false
}

func (service userService) GetAllUsers() []model.User {
	res := service.userRepo.GetAllUser()
	return res
}
