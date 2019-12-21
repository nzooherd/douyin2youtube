package utils

import "project/config"

func GenerateVideoPath(videoId string) string{
	return config.GetString("douyin.store_path") + videoId +
		config.GetString("douyin.video_type")
}
