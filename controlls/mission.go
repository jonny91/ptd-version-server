package controlls

import (
	"PTDVersionServer/dto"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/nsqio/go-nsq"
	"net/http"
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

func SelectAllMissions(context *gin.Context) {
	var result []dto.MissionResult
	DB.Find(&result)
	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    result,
	})
}

func SelectMissionsByRoleName(context *gin.Context) {
	roleName, roleNameOk := context.GetQuery("rolename")
	t, timeOk := context.GetQuery("time")
	var result []dto.MissionResult
	var r *gorm.DB
	if roleNameOk {
		r = DB.Where("role_name = ?", roleName)
	}

	if timeOk {
		t1, err := time.Parse("2006-01-02", t)
		if err != nil {
			fmt.Println(err)
			context.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
			})
		}
		t2 := t1.AddDate(0, 0, 1)

		if r != nil {
			r = r.Where("time >= ? and time < ?", t1, t2)
		} else {
			r = DB.Where("time >= ? and time < ?", t1, t2)
		}
	}

	if r != nil {
		r.Find(&result)
		context.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"msg":    result,
		})
	}

}
