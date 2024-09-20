package repository

import (
	"alger/model"
	login "alger/model/dto/login"
	"gorm.io/gorm"
)

func GetUser(tx *gorm.DB, req login.LoginReq) (resp model.User, err error) {

	if req.UserName != "" {
		tx = tx.Where("username = ?", req.UserName)
	}
	if req.Password != "" {
		tx = tx.Where("password = ?", req.Password)
	}

	err = tx.Find(&resp).Error
	return resp, err
}
