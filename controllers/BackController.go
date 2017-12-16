package controllers

import (
	"blog-go/models"
	"blog-go/util"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type BackController struct {
	BaseController
}

func (c *BackController) index() {
	c.Data["classes"] = models.SelectAllClass()
	c.Data["details"] = models.SelectAllDetail()
	c.TplName = "front/index.html"
}

//登录页
func (c *BackController) LoginPage() {
	c.TplName = "back/login.html"
}

//登录验证
func (c *BackController) Login() {
	flash := beego.NewFlash()
	username, password := c.Input().Get("username"), c.Input().Get("password")

	password = util.Md5(strings.Trim(password, " "))
	fmt.Println(password)
	if flag, user := models.Login(username, password); flag {
		c.SetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"), user.Token, 30*24*60*60, "/", beego.AppConfig.String("cookie.domain"), false, true)
		c.Redirect("backindex", 302)
	} else {
		flash.Error("用户名或者密码错误")
		flash.Store(&c.Controller)
		c.Redirect("loginpage", 302)
	}
}

//退出登录
func (c *BackController) Logout() {
	c.SetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"), "", -1, "/", beego.AppConfig.String("cookie.domain"), false, true)
	c.Redirect("/back/loginpage", 302)
}

//后台首页
func (c *BackController) BackIndex() {
	c.TplName = "back/index.html"
}

//查询所有的类目
func (c *BackController) ArticleClass() {
	classes := models.SelectAllClass()
	c.Data["classes"] = classes
	c.TplName = "back/articleclass.html"
}

//获取添加或修改类目的页面
func (c *BackController) ClassAdd() {
	id, _ := strconv.Atoi(c.Input().Get("id"))
	if id > 0 {
		articleClass := models.SelectClassById(id)
		c.Data["articleClass"] = articleClass
	}
	c.TplName = "back/class_add.html"
}

//保存类目
func (c *BackController) ClassSave() {
	className := c.Input().Get("name")
	id, _ := strconv.Atoi(c.Input().Get("id"))
	if id == 0 {
		//新增
		articleClass := models.ArticleClass{ClassName: className, CreateTime: time.Now()}
		_, err := models.InsertClass(&articleClass)
		if err != nil {
			c.Interactive("", "添加类目失败")
		} else {
			c.Interactive("articleclass", "添加成功")
		}
	} else {
		//更新
		articleClass := models.ArticleClass{Id: id, ClassName: className, UpdateTime: time.Now()}
		err := models.UpdateClass(&articleClass)
		if err != nil {
			c.Interactive("", "更新类目失败")
		} else {
			c.Interactive("articleclass", "更新成功")
		}
	}
}

//删除类目
func (c *BackController) ClassDel() {
	id, _ := strconv.Atoi(c.Input().Get("id"))
	err := models.DeleteClass(id)
	if err != nil {
		c.Interactive("", "删除失败")
	} else {
		c.Interactive("articleclass.html", "")
	}
}

//分页查询文章列表
func (c *BackController) DetailList() {
	classId, _ := c.GetInt("class_id")
	title := c.GetString("title")
	page, _ := c.GetInt("page")
	var pageSize = 8
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	details, count := models.SelectDetailByPage(classId, title, "", pageSize, offset)
	c.Data["title"] = title
	c.Data["count"] = count
	c.Data["list"] = details
	c.Data["classId"] = classId
	c.Data["pagebar"] = util.NewPager(page, int(count), pageSize,
		fmt.Sprintf("/back/index.html?keyword=%s", title), true).ToString()

	//查询所有的类目
	classes := models.SelectAllClass()
	c.Data["classes"] = classes
	c.TplName = "back/detailList.html"
}

//跳转到编辑文章页面
func (c *BackController) ArticleDetail() {
	id, _ := c.GetInt("id")
	if id != 0 {
		detail := models.ArticleDetail{Id: id}
		detailReturn := models.SelectDetailById(&detail)
		c.Data["detail"] = detailReturn
	}
	classes := models.SelectAllClass()
	c.Data["classes"] = classes
	c.TplName = "back/detailForm.html"
}

//保存编辑的文章
func (c *BackController) SaveDetail() {
	id, _ := c.GetInt("id")
	title := c.GetString("title")
	classId, _ := c.GetInt("classId")
	isTop, _ := c.GetInt8("isTop")
	keyword := c.GetString("keyword")
	content := c.GetString("content")
	detail := models.ArticleDetail{ClassId: classId, Title: title, IsTop: isTop, Keyword: keyword, Content: content, CreateTime: time.Now()}
	if id == 0 {
		//新增
		err := models.InsertDetail(&detail)
		if err != nil {
			c.Interactive("", "新增文章失败")
		} else {
			c.Interactive("detaillist", "")
		}
	} else {
		//更新
		detail.Id = id
		err := models.UpdateDetail(&detail)
		if err != nil {
			c.Interactive("", "更新文章失败")
		} else {
			c.Interactive("detaillist", "")
		}

	}

}

//删除文章详情
func (c *BackController) DeleteDetail() {
	id, _ := c.GetInt("id")
	models.DeleteDetail(id)
	c.Interactive("detaillist", "")
}

func (c *BackController) Hello() {
	c.TplName = "back/hello.html"
}
