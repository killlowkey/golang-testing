package store

import (
	"gorm.io/gorm"
)

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{db: db}
}

func (u *UserModel) List() ([]*User, error) {
	var res []*User
	err := u.db.Find(&res).Error
	return res, err
}

func (u *UserModel) GetUserById(id int) (*User, error) {
	var res User
	err := u.db.Where("id = ?", id).Find(&res).Error
	return &res, err
}

func (u *UserModel) CreateUser(user *User) error {
	return u.db.Create(user).Error
}

func (u *UserModel) UpdateUser(user *User) error {
	return u.db.Save(user).Error
}

func (u *UserModel) DeleteUser(id int) error {
	return u.db.Where("id = ?", id).Delete(&User{}).Error
}
