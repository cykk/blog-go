package util

import (
	"time"
)

//时间格式的转换 转换为 2006-01-02 15:04:05 这种格式
func DateFormat(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}
