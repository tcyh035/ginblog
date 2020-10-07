package model

import (
	"ginblog/utils/errmsg"

	"github.com/jinzhu/gorm"
)

//Article 文章model
type Article struct {
	Catagory Catagory
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200);not null" json:"desc"`
	Content string `gorm:"type:longtext;not null" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

// todo 查询分类下的所有文章

// todo 查询单个文章

// CreateArticle 新增文章
func CreateArticle(data *Article) int {
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.Error
	}

	return errmsg.Success
}

// GetArticles 查询文章列表
func GetArticles(pageSize int, pageNum int) []Article {
	var articles []Article
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles).Error
	if err != nil {
		return nil
	}

	return articles
}

// EditArticle 编辑文章
func EditArticle(id int, data *Article) int {
	articleMap := make(map[string]interface{})
	articleMap["Title"] = data.Title
	articleMap["Cid"] = data.Cid
	articleMap["Desc"] = data.Desc
	articleMap["Content"] = data.Content
	articleMap["Img"] = data.Img

	var article Article
	err = db.Where("id = ?", id).Model(&article).Update(articleMap).Error
	if err != nil {
		return errmsg.Error
	}

	return errmsg.Success
}

// DeleteArticle 删除文章
func DeleteArticle(id int) int {
	var article Article
	err = db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return errmsg.Error
	}

	return errmsg.Success
}
