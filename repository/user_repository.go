package repository

import (
	"fmt"

	"MyGO.com/m/dto"
	"MyGO.com/m/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user model.User) model.User
	IsDuplicateEmail(email string) (tx *gorm.DB)
	VerifyLogin(name string) interface{}
	GetAllUser(req *dto.UserGetRequest) ([]model.User, int64, error)
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

func (db *userConnection) GetAllUser(req *dto.UserGetRequest) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	offset := (req.Page - 1) * req.PageSize
	pageSize := req.PageSize

	var filter string

	if req.ID != 0 {
		filter = fmt.Sprintf("where id = %v", req.ID)
	}

	sql := fmt.Sprintf("select * from users %s limit %v offset %v", filter, pageSize, offset)
	res := db.connection.Raw(sql).Scan(&users)

	countQuery := fmt.Sprintf("select count(1) from users %s", filter)
	if err := db.connection.Raw(countQuery).Scan(&total).Error; err != nil {
		return nil, 0, err
	}

	if res.Error == nil {
		return users, total, nil
	}

	return nil, 0, nil
}
