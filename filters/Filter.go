package filters

import (
	"blog-go/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func IsLogin(ctx *context.Context) (bool, *models.User) {
	token, flag := ctx.GetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"))

	var user *models.User
	if flag {
		flag, user = models.SelectUserByToken(token)
	}
	return flag, user
}
