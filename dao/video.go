package dao

import (
	"fmt"
	"project/model"
	"github.com/astaxie/beego/orm"
)

func InsertVideos(videos []model.Video, o orm.Ormer) error{
	inserter, _ := o.QueryTable("video").PrepareInsert()
	var err error = nil
	for index, _ := range videos{
		_, err = inserter.Insert(&videos[index])
		if err != nil && !CheckDuplicateKey(err){
			fmt.Println(err)
		}
	}
	return err
}

func InsertVideoInfos(videoInfos []model.VideoInfo, o orm.Ormer) error{
	inserter, _ := o.QueryTable("video_info").PrepareInsert()
	var err error = nil
	for index, _ := range videoInfos{
		_, err = inserter.Insert(&videoInfos[index])
		if err != nil && !CheckDuplicateKey(err){
			fmt.Println(err)
		}
	}
	return err
}


//查询没在范围内的video id
func QueryNotUploadVideo(videoIds []string, o orm.Ormer) []model.VideoInfo{
	filterVideosQuery := "select * from video_info where has_download = 1"
	filterVideoIds := "("
	for index, videoId := range videoIds{
		if index != len(videoIds) - 1 {
			filterVideoIds += videoId + ", "
		}else{
			filterVideoIds += videoId + ")"
		}
	}
	if len(videoIds) > 0{
		filterVideosQuery += "and video_id not in" + filterVideoIds
	}
	var videoInfos []model.VideoInfo
	_, err := o.Raw( filterVideosQuery).QueryRows(&videoInfos)
	if err != nil{
		fmt.Print(err)
	}
	return videoInfos
}

//
func QueryVideoInfosById(videoIds []string, o orm.Ormer) []model.VideoInfo{
	filterVideosQuery := "("
	for index, videoId := range videoIds{
		if index != len(videoIds) - 1 {
			filterVideosQuery += videoId + ", "
		}else{
			filterVideosQuery += videoId + ")"
		}
	}
	var videoInfos []model.VideoInfo
	_, err := o.Raw(filterVideosQuery).QueryRows(&filterVideosQuery)
	if err != nil{
		fmt.Print(err)
	}
	return videoInfos
}