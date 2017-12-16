package main

import (
	"blog-go/models"
	_ "blog-go/routers"
	_ "blog-go/templates"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("jdbc.username")+":"+beego.AppConfig.String("jdbc.password")+"@tcp(localhost:3306)/blog?charset=utf8", 30)
	orm.RegisterModel(new(models.ArticleClass), new(models.ArticleDetail), new(models.User), new(models.Comment))
}

func main() {
	beego.Run()
}
