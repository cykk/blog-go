package templates

import (
	"blog-go/util"
	"time"

	"github.com/astaxie/beego"
	"github.com/xeonx/timeago"
)

//时间格式转换
func FormatTime(time time.Time) string {
	return timeago.Chinese.Format(time)
}

func init() {
	beego.AddFuncMap("timeago", FormatTime)
	beego.AddFuncMap("dateFormat", util.DateFormat)
}
