package service

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"io"
	"net/http"
	"os"
	"project/model"
	"project/utils"
)

func filter(videoInfo model.VideoInfo) bool{
	return videoInfo.LikeCount > 50000
}

func DownloadVideo(videoInfos []model.VideoInfo){
	o := orm.NewOrm()

	for _, videoInfo := range videoInfos{
		if !filter(videoInfo){
			continue
		}
		var videoPath = videoInfo.VideoPath
		var urlList []string
		err := json.Unmarshal([]byte(videoInfo.UrlList), &urlList)
		if err != nil{
			fmt.Println(err)
		}
		for _, url := range urlList{
			tsVideoPath := utils.Mp4ToTsFileName(videoPath)
			for ; ; {
				downloadVideoErr := DownloadDo(url, videoPath)
				if downloadVideoErr != nil{
					fmt.Println(downloadVideoErr)
				}else{
					if utils.CheckVideo(videoPath) == nil && utils.Mp4ToTs(videoPath, tsVideoPath) == nil{
						utils.DeleteFile(videoPath)
						break
					}else{
						utils.DeleteFile(videoPath)
						utils.DeleteFile(tsVideoPath)
					}
				}
			}
			videoInfo.HasDownload = 1
			videoInfo.VideoPath = tsVideoPath
			_, err := o.Update(&videoInfo)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func DownloadHttpFile(avatarUrl, videoUrl string, coverUrl string) (string, string, string, error) {
	var localAvatar, localCover, localVideo string
	localAvatar = "download/avatar/" + utils.Md5(avatarUrl) + ".jpeg"
	localVideo = "download/video/" + utils.Md5(videoUrl) + ".mp4"
	localCover = "download/cover/" + utils.Md5(coverUrl) + ".jpeg"
	err := DownloadDo(avatarUrl, localAvatar)
	if err != nil {
		return "", "", "", err
	}
	err = DownloadDo(videoUrl, localVideo)
	if err != nil {
		return "", "", "", err
	}
	err = DownloadDo(coverUrl, localCover)
	if err != nil {
		return "", "", "", err
	}
	return localAvatar, localCover, localVideo, nil
}

func DownloadDo(url, saveFile string) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil{
		fmt.Printf("%s", err)
		return nil
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	f, err := os.Create(saveFile)
	if err != nil {
		_ = os.Remove(saveFile)
		return err
	}
	_, err = io.Copy(f, res.Body)
	if err != nil {
		_ = os.Remove(saveFile)
		return err
	}
	return nil
}
