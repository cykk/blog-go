package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//文章类目
type ArticleClass struct {
	Id         int
	ClassName  string
	CreateTime time.Time
	UpdateTime time.Time
}

//查询所有的类目
func SelectAllClass() []*ArticleClass {
	o := orm.NewOrm()
	var classes []*ArticleClass
	o.QueryTable("article_class").All(&classes)
	return classes
}

//删除类目
func DeleteClass(id int) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("article_class").Filter("id", id).Delete()
	return err
}

//修改类目名称
func UpdateClass(articleClass *ArticleClass) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("article_class").Filter("id", articleClass.Id).Update(orm.Params{"class_name": articleClass.ClassName, "update_time": articleClass.UpdateTime})
	return err
}

//新建类目
func InsertClass(articleClass *ArticleClass) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(articleClass)
	fmt.Println(err)
	return id, err
}

//根据类目ID查询类目
func SelectClassById(id int) ArticleClass {
	o := orm.NewOrm()
	articleClass := ArticleClass{Id: id}
	o.Read(&articleClass)
	return articleClass
}
