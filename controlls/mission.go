package controlls

import (
	"PTDVersionServer/dto"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/nsqio/go-nsq"
	"time"
)

func HandleMission(context *gin.Context) {
	missionId := context.PostForm("missionId")
	state := context.PostForm("state")
	platform := context.PostForm("platform")
	cards := context.PostForm("cards")
	now := time.Now().Format("2006/01/02 15:04:05")

	m := &dto.MissionResult{
		Time:      now,
		MissionId: missionId,
		Platform:  platform,
		State:     state,
		Cards:     cards,
	}

	resultStr, _ := json.Marshal(m)
	fmt.Println(string(resultStr))
	Send2MQ("mission", resultStr)
}
