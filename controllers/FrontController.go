package controllers

import (
	"blog-go/models"
	"blog-go/util"
	"fmt"
	"time"
)

type FrontController struct {
	BaseController
}

//首页
func (c *FrontController) Index() {

	//查询所有的类目
	c.Data["classes"] = models.SelectAllClass()

	//分页查询文章详情
	var (
		page     int
		pageSize int = 6
		offset   int
		classId  int
		count    int64
	)
	if page, _ = c.GetInt("page"); page < 1 {
		page = 1
	}
	offset = (page - 1) * pageSize
	classId, _ = c.GetInt("classId")
	c.Data["detailList"], count = models.SelectDetailByPage(classId, "", "", pageSize, offset)
	c.Data["pagebar"] = util.NewPager(page, int(count), pageSize, "/"+"index", true).ToString()
	c.Data["actionName"] = "home"
	c.TplName = "front/home.html"
}

//查询文章列表
func (c *FrontController) SelectArticleList() {
	//查询所有的类目
	c.Data["classes"] = models.SelectAllClass()

	//分页查询文章详情
	var (
		page     int
		pageSize int = 6
		offset   int
		classId  int
		count    int64
	)
	if page, _ = c.GetInt("page"); page < 1 {
		page = 1
	}
	offset = (page - 1) * pageSize
	classId, _ = c.GetInt("classId")
	keyword := c.GetString("keyword")
	c.Data["detailList"], count = models.SelectDetailByPage(classId, "", keyword, pageSize, offset)
	c.Data["pagebar"] = util.NewPager(page, int(count), pageSize, "/"+"index", true).ToString()
	c.Data["actionName"] = "home"
	c.TplName = "front/article.html"
}

//查询文章详情
func (c *FrontController) SelectArticleDetail() {
	id, _ := c.GetInt("id")
	fmt.Println(id)
	detail := &models.ArticleDetail{Id: id}
	models.SelectDetailById(detail)
	c.Data["detail"] = detail
	fmt.Println(detail)
	c.Data["actionName"] = "detail"
	//	c.Data["comments"] = models.SelectCommentsByDetailId(id)
	c.TplName = "front/detail.html"
}

//插入评论
func (c *FrontController) InsertComment() {
	detailId, _ := c.GetInt("detailId")
	username := c.GetString("username")
	content := c.GetString("content")
	comment := &models.Comment{DetailId: detailId, UserName: username, CreateTime: time.Now(), Content: content}
	err := models.InsertComment(comment)
	if err == nil {
		c.Interactive("detail", "评论成功")
	}

}
