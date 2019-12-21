package model

import (
	_ "github.com/astaxie/beego/orm"
)

type Video struct {
	Id int `orm:"column(id)"`
	VideoId string `orm:"column(video_id)"`
	AuthorId string `orm:"column(author_id)"`
}

type VideoInfo struct{
	Id int `orm:"column(id)"`
	VideoId string `orm:"column(video_id)"`
	UrlList string `orm:"column(url_list)"`
	Duration int `orm:"column(duration)"`
	CommentCount int `orm:"column(comment_count)"`
	LikeCount int `orm:"column(like_count)"`
	Description string `orm:"column(description)"`
	HasDownload int `orm:"column(has_download)"`
	VideoPath string `orm:"column(video_path)"`
}

type VideoTagRef struct{
	Id int `orm:"column(id)"`
	VideoId int `orm:"column(video_id)"`
	TagId int `orm:"column(tag_id)"`
}

