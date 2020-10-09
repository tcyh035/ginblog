package model

import (
	"ginblog/utils/errmsg"

	"github.com/jinzhu/gorm"
)

//Article 文章model
type Article struct {
	Category Category `gorm:"ForeignKey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200);not null" json:"desc"`
	Content string `gorm:"type:longtext;not null" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

// GetArticlesByCategory 查询分类下的所有文章
func GetArticlesByCategory(category int, pageSize int, pageNum int) ([]Article, int) {
	var articles []Article
	err = db.Preload("Category").Where("Cid = ?", category).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles).Error
	if err != nil {
		return nil, errmsg.Error
	}

	return articles, errmsg.Success
}

// GetArticle 查询单个文章
func GetArticle(id int) (Article, int) {
	var article Article
	err = db.Preload("Category").Where("id = ?", id).Find(&article).Error
	if err != nil {
		return article, errmsg.ErrorArticleNotExist
	}

	return article, errmsg.Success
}

// CreateArticle 新增文章
func CreateArticle(data *Article) int {
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.Error
	}

	return errmsg.Success
}

// GetArticles 查询文章列表
func GetArticles(pageSize int, pageNum int) ([]Article, int) {
	var articles []Article
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles).Error
	if err != nil {
		return nil, errmsg.Error
	}

	return articles, errmsg.Success
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
