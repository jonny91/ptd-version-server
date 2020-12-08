package controlls

import (
	"PTDVersionServer/dto"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/nsqio/go-nsq"
	"strconv"
	"time"
)

func HandleMission(context *gin.Context) {
	missionId := context.PostForm("missionId")
	roleName := context.PostForm("roleName")
	state := context.PostForm("state")
	platform := context.PostForm("platform")
	cards := context.PostForm("cards")
	deviceModel := context.PostForm("deviceModel")
	durationStr := context.PostForm("duration")
	duration, _ := strconv.Atoi(durationStr)
	now := time.Now().Format("2006/01/02 15:04:05")

	m := &dto.MissionResult{
		Time:        now,
		RoleName:    roleName,
		MissionId:   missionId,
		Platform:    platform,
		State:       state,
		Cards:       cards,
		DeviceModel: deviceModel,
		Duration:    duration,
	}

	resultStr, _ := json.Marshal(m)
	fmt.Println(string(resultStr))
	Send2MQ("mission", resultStr)
}
