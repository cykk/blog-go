package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type ArticleDetail struct {
	Id         int
	ClassId    int
	Title      string
	Keyword    string
	Content    string
	BrowseNum  int
	IsTop      int8
	CreateTime time.Time
}

//新增
func InsertDetail(detail *ArticleDetail) error {
	o := orm.NewOrm()
	_, err := o.Insert(detail)
	return err
}

//根据ID查询单条
func SelectDetailById(detail *ArticleDetail) *ArticleDetail {
	o := orm.NewOrm()
	o.Read(detail)
	return detail
}

//删除
func DeleteDetail(id int) {
	o := orm.NewOrm()
	o.QueryTable("article_detail").Filter("id", id).Delete()

}

//更新
func UpdateDetail(detail *ArticleDetail) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("article_detail").Filter("id", detail.Id).Update(orm.Params{"class_id": detail.ClassId, "keyword": detail.Keyword, "title": detail.Title, "content": detail.Content})
	return err
}

//查询所有
func SelectAllDetail() []*ArticleDetail {
	o := orm.NewOrm()
	var details []*ArticleDetail
	o.QueryTable("article_detail").All(&details)
	return details
}

//分页查询文章
func SelectDetailByPage(classId int, title string, keyword string, pageSize int, offset int) ([]*ArticleDetail, int64) {
	o := orm.NewOrm()
	var details []*ArticleDetail
	qs := o.QueryTable("article_detail")
	if classId != 0 {
		qs = qs.Filter("class_id", classId)
	}
	if title != "" {
		qs = qs.Filter("title", title)
	}
	if keyword != "" {
		qs = qs.Filter("keyword", keyword)
	}
	count, _ := qs.Count()
	qs.OrderBy("-is_top", "-create_time").Limit(pageSize, offset).All(&details)
	return details, count
}
