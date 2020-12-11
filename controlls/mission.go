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
	starStr := context.PostForm("star")
	duration, _ := strconv.Atoi(durationStr)
	now := time.Now().Format("2006/01/02 15:04:05")

	star, err := strconv.Atoi(starStr)
	if err != nil {
		star = 0
	}
	m := &dto.MissionResult{
		Time:        now,
		RoleName:    roleName,
		MissionId:   missionId,
		Platform:    platform,
		State:       state,
		Cards:       cards,
		DeviceModel: deviceModel,
		Star:        star,
		Duration:    duration,
	}

	resultStr, _ := json.Marshal(m)
	fmt.Println(string(resultStr))
	Send2MQ("mission", resultStr)
}
