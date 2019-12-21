package youtube

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"project/config"
	"project/dao"
	_ "project/dao"
	"project/model"
	"project/utils"
	"strconv"
	"time"
)

func Task(){
	o := orm.NewOrm()
	var source = new(Source)
	var videoInfos []model.VideoInfo
	var random RandomChoose
	var uploadInfo model.UploadInfo

	source.videoInfos = &videoInfos

	filterVideos(o, &videoInfos)

	chooseVideoInfos := random.choose(source)
	var videoPaths = make([]string, len(chooseVideoInfos))
	for index, videoInfo := range chooseVideoInfos{
		videoPaths[index] = videoInfo.VideoPath
	}
	uploadVideoId := generateUploadId(&chooseVideoInfos)

	var toPath = config.GetString("youtube.combine_video_path") +
		uploadVideoId + config.GetString("douyin.video_type")

	err := utils.CombineVideosUseFfmpeg(videoPaths, toPath)
	if err != nil{
		panic(err)
	}
	uploadInfo.UploadId = uploadVideoId
	uploadInfo.Description = "[TikTok]" + chooseVideoInfos[0].Description
	uploadInfo.Title = "欢迎关注，每日抖音热搜，要你好看，一起抖不停。"
	uploadInfo.Category = "10"
	uploadInfo.Keywords = "抖音,TikTok,Musicly"
	uploadInfo.Privacy = "public"
	uploadInfo.VideoPath = toPath

	writeDataBase(chooseVideoInfos, uploadInfo, o)

	//Upload(uploadInfo)
}


func generateUploadId(videoInfos *[]model.VideoInfo) string{
	nowTime := time.Now()
	date := strconv.Itoa(nowTime.Year()) + "_" +
		strconv.Itoa(int(nowTime.Month())) + "_" + strconv.Itoa(nowTime.Day())

	var sumVideoIds  = ""
	for _, videoInfo := range *videoInfos{
		sumVideoIds += videoInfo.VideoId
	}

	hashcode :=  utils.HashCode(sumVideoIds)
	return date + "_" + strconv.Itoa(hashcode)
}

func filterVideos(o orm.Ormer, videoInfos *[]model.VideoInfo){
	uploadedVideoIds := dao.QueryUploadedVideo(o)
	filteredVideoInfos := dao.QueryNotUploadVideo(uploadedVideoIds, o)
	*videoInfos = filteredVideoInfos
}

func writeDataBase(chooseVideoInfos []model.VideoInfo, uploadInfo model.UploadInfo, o orm.Ormer) {
	var uploadVideoRefs []model.UploadVideoRef
	uploadVideoRefs = make([]model.UploadVideoRef, len(chooseVideoInfos))
	for index, videoInfo := range chooseVideoInfos{
		var uploadVideoRef model.UploadVideoRef
		uploadVideoRef.UploadId = uploadInfo.UploadId
		uploadVideoRef.VideoId = videoInfo.VideoId
		uploadVideoRef.Order = index
		uploadVideoRefs[index] = uploadVideoRef
	}
	_, err := o.Insert(&uploadInfo)
	if err != nil{
		fmt.Println(err)
	}
	err = dao.InsertUploadVideoRefs(uploadVideoRefs, o)
	if err != nil{
		fmt.Println(err)
	}
}