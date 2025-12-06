package controllers

import (
	"net/http"
	"task4/models"

	"github.com/gin-gonic/gin"
)

// 文章管理功能
// 实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
// 实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息。
// 实现文章的更新功能，只有文章的作者才能更新自己的文章。
// 实现文章的删除功能，只有文章的作者才能删除自己的文章。
type PostController struct {
	BaseController
}

func (postcontroller PostController) list(c *gin.Context) {
	var posts []models.Post
	result := models.DB.Order("Created_at desc").Find(&posts)
	if result.Error != nil {
		c.HTML(http.StatusOK, "", gin.H{
			"result": false,
			"msg":    "查询文章列表失败",
		})
	}
	c.HTML(http.StatusOK, "", gin.H{
		"result": true,
		"list":   posts,
	})
}

func (postcontroller PostController) create(c *gin.Context) {
	c.Get("user")
	var posts []models.Post
	result := models.DB.Order("Created_at desc").Find(&posts)
	if result.Error != nil {
		c.HTML(http.StatusOK, "", gin.H{
			"result": false,
			"msg":    "查询文章列表失败",
		})
	}
	c.HTML(http.StatusOK, "", gin.H{
		"result": true,
		"list":   posts,
	})
}

func (postcontroller PostController) read(c *gin.Context) {
	var posts []models.Post
	result := models.DB.Order("Created_at desc").Find(&posts)
	if result.Error != nil {
		c.HTML(http.StatusOK, "", gin.H{
			"result": false,
			"msg":    "查询文章列表失败",
		})
	}
	c.HTML(http.StatusOK, "", gin.H{
		"result": true,
		"list":   posts,
	})
}

func (postcontroller PostController) redit(c *gin.Context) {
	var posts []models.Post
	result := models.DB.Order("Created_at desc").Find(&posts)
	if result.Error != nil {
		c.HTML(http.StatusOK, "", gin.H{
			"result": false,
			"msg":    "查询文章列表失败",
		})
	}
	c.HTML(http.StatusOK, "", gin.H{
		"result": true,
		"list":   posts,
	})
}

func (postcontroller PostController) delete(c *gin.Context) {
	var posts []models.Post
	result := models.DB.Order("Created_at desc").Find(&posts)
	if result.Error != nil {
		c.HTML(http.StatusOK, "", gin.H{
			"result": false,
			"msg":    "查询文章列表失败",
		})
	}
	c.HTML(http.StatusOK, "", gin.H{
		"result": true,
		"list":   posts,
	})
}
