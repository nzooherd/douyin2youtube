package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
	"os"
	"os/exec"
	"project/adb"
	"project/config"
	"project/controller"
	"project/dao"
	_ "project/dao"
	"project/model"
	"project/youtube"
	"strconv"
	"time"
)

func WebStart(){
	gin.SetMode(config.GetString("gin.model"))
	r := gin.Default()
	r.POST("/feed/", controller.DownloadFeed)
	_ = r.Run(":" + strconv.Itoa(config.GetInt("gin.addr")))
}

func AnyProxyStart(){
	cmd := exec.Command("anyproxy", "--intercept", "--rule",
		config.GetString("anyproxy.rule_js_path"))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

func AdbStart(){
	adb.CloseApp(config.GetString("app.packageName"))
	adb.RunApp(config.GetString("app.packageName") + "/" + config.GetString("app.startPath"))
	for {
		adb.Swipe(config.GetString("swipe.startX"), config.GetString("swipe.startY"),
			config.GetString("swipe.endX"), config.GetString("swipe.endY"))
		time.Sleep(config.GetDuration("swipe.sleep") * time.Millisecond)
	}
}

func CombineCronJob(){
	ticker := time.NewTicker(30 * 60 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			youtube.Task()
		}
	}
}

func UploadCronJob(){
	o := orm.NewOrm()
	uploadInfo := dao.QueryOneNotUploadVideo(o)
	err := youtube.Upload(*uploadInfo)
	if err == nil{
		var upload = new(model.Upload)
		upload.UploadId = uploadInfo.UploadId
		upload.UploadResult = 0
		_, err = o.Insert(uploadInfo)
		if err != nil{
			fmt.Printf(err.Error())
		}
	}


}

func main(){
	go AnyProxyStart()
	go WebStart()
	//go AdbStart()
	//go CombineCronJob()
	//UploadCronJob()
	select {}
}
