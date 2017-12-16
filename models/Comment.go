package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Comment struct {
	Id         int
	DetailId   int
	UserName   string
	Content    string
	CreateTime time.Time
}

func SelectCommentsByDetailId(detailId int) []*Comment {
	var comments []*Comment
	o := orm.NewOrm()
	o.Raw("select id,detail_id,username,content,create_time from comment where id = ?", detailId).QueryRows(&comments)
	//	o.QueryTable("comment").Filter("detail_id", detailId).All(&comments)
	return comments
}

func InsertComment(comment *Comment) error {
	o := orm.NewOrm()
	_, err := o.Insert(&comment)
	return err
}
