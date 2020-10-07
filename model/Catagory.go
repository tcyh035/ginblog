package model

import "ginblog/utils/errmsg"

//Catagory 目录model
type Catagory struct {
	ID   int    `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// CheckCatagoryExist 查询分类是否存在
func CheckCatagoryExist(name string) int {
	var catagory Catagory
	db.First(&catagory, "name = ?", name)
	if catagory.ID > 0 {
		return errmsg.ErrorCatagoryUsed
	}

	return errmsg.Success
}

// CreateCatagory 创建分类
func CreateCatagory(data *Catagory) int {
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.Error
	}

	return errmsg.Success
}

// GetCatagories 查询分类列表
func GetCatagories(pageSize int, pageNum int) []Catagory {
	var catagories []Catagory
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&catagories).Error
	if err != nil {
		return nil
	}

	return catagories
}

// EditCatagory 编辑分类
func EditCatagory(id int, data *Catagory) int {
	catagoryMap := make(map[string]interface{})
	catagoryMap["Name"] = data.Name

	var catagory Catagory
	err = db.Where("id = ?", id).Model(&catagory).Update(catagoryMap).Error
	if err != nil {
		return errmsg.Error
	}

	return errmsg.Success
}

// DeleteCatagory 删除分类
func DeleteCatagory(id int) int {
	var catagory Catagory
	err = db.Where("id = ?", id).Delete(&catagory).Error
	if err != nil {
		return errmsg.Error
	}

	return errmsg.Success
}
