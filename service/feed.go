package service

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"project/dao"
	"project/model"
	"project/utils"
)

func HandleFeed(feed model.FeedJson){
	var awemeList = feed.AwemeList

	var videos []model.Video
	var videoInfos []model.VideoInfo
	for _, aweme := range awemeList {
		video := new(model.Video)
		videoInfo := new(model.VideoInfo)
		aweme.LoadVideo(video)
		aweme.LoadVideoInfo(videoInfo)
		videoInfo.VideoPath = utils.GenerateVideoPath(videoInfo.VideoId)
		videos = append(videos, *video)
		videoInfos = append(videoInfos, *videoInfo)
	}
	o := orm.NewOrm()
	_ = o.Begin()
	insertVideoErr := dao.InsertVideos(videos, o)
	if insertVideoErr != nil{
		fmt.Println(insertVideoErr)
	}
	insertVideoInfoErr := dao.InsertVideoInfos(videoInfos, o)
	if insertVideoInfoErr != nil{
		fmt.Println(insertVideoInfoErr)
	}
	if insertVideoErr != nil || insertVideoInfoErr != nil {
		_ = o.Rollback()
		return
	}
	_ = o.Commit()

	go DownloadVideo(videoInfos)
}