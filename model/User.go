package model

import (
	"ginblog/utils/errmsg"

	"github.com/jinzhu/gorm"
)

//User 用户model
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username"`
	Password string `gorm:"type:varchar(20);not null " json:"password"`
	Role     int    `gorm:"type:int " json:"role"`
}

// CheckUserExist 查询用户是否存在
func CheckUserExist(name string) int {
	var user User
	db.First(&user, "Username = ?", name)
	if user.ID > 0 {
		return errmsg.ErrorUserNameUsed
	}

	return errmsg.Success
}

// CreateUser 创建用户
func CreateUser(data *User) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.Error
	}

	return errmsg.Success
}

// GetUsers 查询用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil {
		return nil
	}

	return users
}

// EditUser 编辑用户

// DeleteUser 删除用户
