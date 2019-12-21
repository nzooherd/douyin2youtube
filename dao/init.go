package dao

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"project/config"
	"project/model"
	"strings"
)

func init(){
	_ = orm.RegisterDriver(config.GetString("database.driver_name"), orm.DRMySQL)
	dataBase := config.GetString("database.user_name") + ":" +
		config.GetString("database.password") + "@tcp(" + config.GetString("database.ip") + ":" +
		config.GetString("database.port") + ")/" + config.GetString("database.db_name")
	fmt.Println(dataBase)
	err := orm.RegisterDataBase("default", config.GetString("database.driver_name"), dataBase)
	if err != nil{
		fmt.Println(err)
	}

	orm.RegisterModel(new(model.Video), new(model.VideoInfo), new(model.VideoTagRef), new(model.Tag),
		new(model.UploadVideoRef), new(model.Upload), new(model.UploadInfo))
}


func CheckDuplicateKey(err error) bool{
	return strings.Contains(err.Error(), "duplicate")
}