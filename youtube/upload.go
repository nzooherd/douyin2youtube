package youtube


import (
	"fmt"
	"google.golang.org/api/youtube/v3"
	"log"
	"os"
	"project/model"
	"strings"
)


func Upload(uploadInfo model.UploadInfo) error{
	if uploadInfo.VideoPath == "" {
		log.Fatalf("You must provide a filename of a video file to upload")
	}

	client := getClient(youtube.YoutubeUploadScope)

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	upload := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       uploadInfo.Title,
			Description: uploadInfo.Description,
			CategoryId:  uploadInfo.Category,
		},
		Status: &youtube.VideoStatus{PrivacyStatus: uploadInfo.Privacy},
	}

	// The API returns a 400 Bad Request response if tags is an empty string.
	if strings.Trim(uploadInfo.Keywords, "") != "" {
		upload.Snippet.Tags = strings.Split(uploadInfo.Keywords, ",")
	}

	call := service.Videos.Insert("snippet,status", upload)

	file, err := os.Open(uploadInfo.VideoPath)
	defer file.Close()
	if err != nil {
		log.Fatalf("Error opening %v: %v", uploadInfo.VideoPath, err)
	}

	response, err := call.Media(file).Do()
	if err == nil{
		fmt.Printf("Upload successful! Video ID: %v\n", response.Id)
	}
	return err
}

func handleError(err error){
	fmt.Printf("%s", err)
}
