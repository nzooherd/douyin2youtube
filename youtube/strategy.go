package youtube

import (
	"math/rand"
	"project/model"
	"time"
)

type Source struct{
	videoInfos *[]model.VideoInfo
}
type StrategyType int32
const (
	randomStrategy StrategyType = 1
	videoDuration int = 10 * 60 * 1000 // unit ms
)

// 策略模式
type chooseVideo interface {
	choose(source *Source) []model.VideoInfo
}
var strategyMap map[StrategyType]chooseVideo

// 随机算法
type RandomChoose struct{

}
func (random RandomChoose) choose(source *Source) ([]model.VideoInfo){
	var durationSum int
	var length = 0
	videoInfos := source.videoInfos
	rand.Seed(time.Now().Unix())
	for durationSum = 0; durationSum < videoDuration; length += 1{
		if length >= len(*videoInfos){
			break
		}
		index := rand.Intn(len(*videoInfos) - length)
		toIndex := len(*videoInfos) - length - 1

		durationSum += (*videoInfos)[index].Duration

		temp := (*videoInfos)[toIndex]
		(*videoInfos)[toIndex] = (*videoInfos)[index]
		(*videoInfos)[index] = temp
	}

	return (*videoInfos)[len(*videoInfos) - length:]
}


func init(){
	strategyMap = make(map[StrategyType]chooseVideo)
	strategyMap[randomStrategy] = RandomChoose{}
}
