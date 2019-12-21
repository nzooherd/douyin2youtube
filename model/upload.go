package model

import (
	_ "github.com/astaxie/beego/orm"
	"time"
)

type Upload struct{
	Id int `orm:"column(id)"`
	UploadId string `orm:"column(upload_id)"`
	CreateTime time.Time `orm:"column(create_time)"`
	UploadResult int `orm:"column(upload_result)"`
}

type UploadInfo struct{
	Id int `orm:"column(id)"`
	UploadId string `orm:"column(upload_id)"`
	Title string `orm:"column(title)"`
	Description string `orm:"column(description)"`
	Category string `orm:"column(category)"`
	Keywords string `orm:"column(keywords)"`
	Privacy string `orm:"column(privacy)"`
	VideoPath string `orm:"column(video_path)"`
}

type UploadVideoRef struct{
	Id int `orm:"column(id)"`
	UploadId string `orm:"column(upload_id)"`
	VideoId string `orm:"column(video_id)"`
	Order int `orm:"column(order)"`
}