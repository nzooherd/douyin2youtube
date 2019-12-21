package dao

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"project/model"
	"strconv"
)

func QueryUploadedVideo(o orm.Ormer) []string{
	query := "SELECT distinct(upload_video_ref.video_id) FROM upload inner join  upload_video_ref on " +
		"upload.upload_id = upload_video_ref.upload_id where upload.upload_result = 1"

	var res orm.ParamsList
	videoIds := make([]string, 0)
	_, err := o.Raw(query).ValuesFlat(&res)
	if err != nil{
		fmt.Println(err)
	}
	for _, singleRes := range res{
		videoIds = append(videoIds, singleRes.(string))
	}
	return videoIds

}

func InsertUploadVideoRefs(uploadVideoRefs []model.UploadVideoRef , o orm.Ormer) error{
	inserter, _ := o.QueryTable("upload_video_ref").PrepareInsert()
	var err error = nil
	for index, _ := range uploadVideoRefs{
		_, err = inserter.Insert(&uploadVideoRefs[index])
		if err != nil && !CheckDuplicateKey(err){
			fmt.Println(err)
		}
	}
	return err
}

func QueryOneNotUploadVideo(o orm.Ormer) *model.UploadInfo{
	query := "select * from upload_info left join upload " +
		"on upload.upload_id = upload_info.upload_id " +
		"where upload_result IS NULL or upload_result = 1 limit 1;"

	var uploadInfo = new(model.UploadInfo)

	var lists []orm.ParamsList
	num, err := o.Raw(query).ValuesList(&lists)
	if err != nil{
		panic(err)
	}
	if num <= 0 {
		return uploadInfo
	}
	uploadInfo.Id, _ = strconv.Atoi(lists[0][0].(string))
	err = o.Read(uploadInfo)
	return uploadInfo
}