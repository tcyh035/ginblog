package model

import "ginblog/utils/errmsg"

//Category 目录model
type Category struct {
	ID   int    `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// CheckCategoryExist 查询分类是否存在
func CheckCategoryExist(name string) int {
	var category Category
	db.First(&category, "name = ?", name)
	if category.ID > 0 {
		return errmsg.ErrorCategoryUsed
	}

	return errmsg.Success
}

// todo 查询分类下的所有文章

// CreateCategory 创建分类
func CreateCategory(data *Category) int {
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.Error
	}

	return errmsg.Success
}

// GetCatagories 查询分类列表
func GetCatagories(pageSize int, pageNum int) []Category {
	var catagories []Category
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&catagories).Error
	if err != nil {
		return nil
	}

	return catagories
}

// EditCategory 编辑分类
func EditCategory(id int, data *Category) int {
	categoryMap := make(map[string]interface{})
	categoryMap["Name"] = data.Name

	var category Category
	err = db.Where("id = ?", id).Model(&category).Update(categoryMap).Error
	if err != nil {
		return errmsg.Error
	}

	return errmsg.Success
}

// DeleteCategory 删除分类
func DeleteCategory(id int) int {
	var category Category
	err = db.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return errmsg.Error
	}

	return errmsg.Success
}
