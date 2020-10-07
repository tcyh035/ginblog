package model

import (
	"ginblog/utils/errmsg"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//User 用户model
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(200);not null" json:"password"`
	Role     int    `gorm:"type:int " json:"role"`
}

// CheckUserExist 查询用户是否存在
func CheckUserExist(name string) int {
	var user User
	db.First(&user, "username = ?", name)
	if user.ID > 0 {
		return errmsg.ErrorUserNameUsed
	}

	return errmsg.Success
}

// CreateUser 创建用户
func CreateUser(data *User) int {
	if data.Password == "" {
		return errmsg.Error
	}

	err = db.Create(&data).Error
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
func EditUser(id int, data *User) int {
	userMap := make(map[string]interface{})
	userMap["username"] = data.Username
	userMap["role"] = data.Role

	var user User
	err = db.Where("id = ?", id).Model(&user).Update(userMap).Error
	if err != nil {
		return errmsg.Error
	}

	return errmsg.Success
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.Error
	}

	return errmsg.Success
}

// BeforeSave 加密
func (u *User) BeforeSave() {
	u.Password = GetSrcyptPassword(u.Password)
}

// GetSrcyptPassword 获取密码
func GetSrcyptPassword(password string) string {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return ""
	}

	return string(pw[:])
}
