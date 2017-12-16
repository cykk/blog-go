package routers

import (
	"blog-go/controllers"

	"github.com/astaxie/beego"
)

func init() {

	//前台
	beego.Router("/", &controllers.FrontController{}, "GET:Index")
	beego.Router("/home", &controllers.FrontController{}, "GET:Index")
	beego.Router("/articleList", &controllers.FrontController{}, "GET:SelectArticleList")
	beego.Router("/detail", &controllers.FrontController{}, "GET:SelectArticleDetail")
	beego.Router("/comment", &controllers.FrontController{}, "POST:InsertComment")

	//后台
	beego.AutoRouter(&controllers.BackController{})
}
