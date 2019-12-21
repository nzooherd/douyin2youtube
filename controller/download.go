package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"project/model"
	"project/service"
)

func DownloadFeed(ctx *gin.Context){
	body := ctx.DefaultPostForm("json", "null")
	var feedJson model.FeedJson
	err := json.Unmarshal([]byte(body), &feedJson)
	if err != nil{
		fmt.Println(err)
	}

	service.HandleFeed(feedJson)
	ctx.JSON(200, gin.H{"status": 0, "message": "success"})
}

func DownloadHot(ctx *gin.Context){

}

func DownloadSearch(ctx *gin.Context){

}
