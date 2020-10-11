package model

import (
	"ginblog/utils/errmsg"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//User 用户model
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(200);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
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
func GetUsers(pageSize int, pageNum int) ([]User, int) {
	var users []User
	var total int
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total).Error
	if err != nil {
		return nil, 0
	}

	return users, total
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

// CheckLogin 登陆验证
func CheckLogin(username string, password string) int {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ErrorUserNotExist
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return errmsg.ErrorPasswordWrong
	}

	if user.Role != 1 {
		return errmsg.ErrorUserNoRight
	}

	return errmsg.Success
}
