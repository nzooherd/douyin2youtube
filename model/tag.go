package model

import (
	_ "github.com/astaxie/beego/orm"
)

type Tag struct{
	Id int `orm:"column(id)"`
	Name string `orm:"column(name)"`
}
