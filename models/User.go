package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int
	Name     string
	Password string
	Token    string
}

func SelectUserByToken(token string) (bool, *User) {
	o := orm.NewOrm()
	var user User
	sqlStr := "SELECT `id`,`name`,`password`,`token` FROM `user` where token = ?"
	err := o.Raw(sqlStr, token).QueryRow(&user)
	return err != orm.ErrNoRows, &user
}

func Login(name string, password string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable("user").Filter("name", name).Filter("password", password).One(&user)
	return err != orm.ErrNoRows, user
}
