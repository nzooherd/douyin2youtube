package model

import (
	"encoding/json"
	"fmt"
)

type FeedJson struct {
	StatusCode int    `json:"status_code"`
	MinCursor  int    `json:"min_cursor"`
	MaxCursor  int    `json:"max_cursor"`
	HasMore    int    `json:"has_more"`
	AwemeList  []Aweme `json:"aweme_list"`
}

type Aweme struct {
	AwemeId      string      `json:"aweme_id"`
	Desc         string      `json:"desc"`
	CreateTime   int         `json:"create_time"`
	Author       author      `json:"author"`
	Video        video       `json:"video"`
	IsAds        bool        `json:"is_ads"`
	Duration     int         `json:"duration"`
	GroupId      string      `json:"group_id"`
	AuthorUserId int64       `json:"author_user_id"`
	LongVideo    []longVideo `json:"long_video"`
	Statistics   statistics  `json:"statistics"`
	ShareInfo    shareInfo   `json:"share_info"`
}

type author struct {
	Uid         string `json:"uid"`
	Nickname    string `json:"nickname"`
	Gender      int    `json:"gender"`
	UniqueId    string `json:"unique_id"`
	AvatarThumb uriStr `json:"avatar_thumb"`
}

type video struct {
	PlayAddr uriStr `json:"play_addr"`
	Cover    uriStr `json:"cover"`
	Duration int `json:"duration"`
	DownloadAddr uriStr `json:"download_addr"`
}

type longVideo struct {
	Video            video `json:"video"`
	TrailerStartTime int   `json:"trailer_start_time"`
}

type uriStr struct {
	Uri     string   `json:"uri"`
	UrlList []string `json:"url_list"`
	Width   int      `json:"width"`
	Height  int      `json:"height"`
}

type statistics struct {
	AwemeId           string `json:"aweme_id"`
	CommentCount      int    `json:"comment_count"`
	LikeCount         int    `json:"digg_count"`
	DownloadCount     int    `json:"download_count"`
	PlayCount         int    `json:"play_count"`
	ShareCount        int    `json:"share_count"`
	ForwardCount      int    `json:"forward_count"`
	LoseCount         int    `json:"lose_count"`
	LoseCommentCount int    `json:"lose_comment_count"`
}

type shareInfo struct {
	ShareUrl   string `json:"share_url"`
	ShareTitle string `json:"share_title"`
}

func (aweme *Aweme) LoadVideo(video *Video){
	video.VideoId = aweme.Video.PlayAddr.Uri
	video.AuthorId = aweme.Author.Uid
}

func (aweme *Aweme) LoadVideoInfo(videoInfo *VideoInfo){
	videoInfo.VideoId = aweme.Video.PlayAddr.Uri
	videoInfo.Description = aweme.Desc
	videoInfo.Duration = aweme.Video.Duration
	videoInfo.CommentCount = aweme.Statistics.CommentCount
	videoInfo.LikeCount = aweme.Statistics.LikeCount
	urlListData, err := json.Marshal(aweme.Video.DownloadAddr.UrlList)
	if err == nil{
		videoInfo.UrlList = string(urlListData)
	} else{
		fmt.Println(err)
	}
}
