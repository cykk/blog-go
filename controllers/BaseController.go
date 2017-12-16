package controllers

import (
	"blog-go/filters"
	"strings"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
}

func (b *BaseController) Prepare() {
	controllerName, actionName := b.GetControllerAndAction()
	b.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	b.actionName = strings.ToLower(actionName)
	if b.controllerName == "back" && (b.actionName != "login" && b.actionName != "loginpage") {
		_, user := filters.IsLogin(b.Ctx)
		if user == nil {
			b.Interactive("loginpage", "未登录")
		}
	}

}

//与前台交互或重定向
func (b *BaseController) Interactive(path string, errMsg string) {
	if path == "" {
		b.Ctx.WriteString("<script>alert('" + errMsg + "');window.history.go(-1);</script>")
		b.StopRun()
	} else {
		b.Redirect(path, 302)
	}
}
