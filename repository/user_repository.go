package repository

import (
	"fmt"

	"MyGO.com/m/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user model.User) model.User
	IsDuplicateEmail(email string) (tx *gorm.DB)
	VerifyLogin(name string) interface{}
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user model.User) model.User {
	err := db.connection.Save(&user)
	if err != nil {
		fmt.Println("------------Here is error in user repository--------------", err)
	}
	return user
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user model.User
	return db.connection.Where("email = ?", email).Take(&user)

}

func (db *userConnection) VerifyLogin(name string) interface{} {
	var user model.User

	res := db.connection.Where("name = ?", name).Take(&user)

	if res.Error == nil {
		return user
	}
	return nil
}
