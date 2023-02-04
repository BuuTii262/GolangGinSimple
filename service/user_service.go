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
	fmt.Println("____________res____________", res)
	if res.Error != nil {
		return false
	}
	return true
}
