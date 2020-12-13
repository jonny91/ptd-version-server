package main

import (
	"PTDVersionServer/config"
	"PTDVersionServer/controlls"
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	err error
)

func main() {
	if !config.ReadAll() {
		return
	}

	err = controlls.InitMQ()
	if err != nil {
		fmt.Println(err)
	}

	err = controlls.InitDB()
	if err != nil {
		fmt.Println(err)
	}

	defer controlls.DB.Close()
	for i := 0; i < 3; i++ {
		go controlls.ReceiveFromMQ("mission")
	}

	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	g.Static("/files", "./static/files")

	g.GET("/ptd", controlls.HandleVersion)
	g.GET("/reload", controlls.HandleReload)

	r := g.Group("/mission")
	{
		r.GET("/all", controlls.SelectAllMissions)
		r.GET("/select", controlls.SelectMissionsByRoleName)
	}
	g.POST("/mission", controlls.HandleMission)

	err = g.Run("0.0.0.0:19999")
	if err != nil {
		fmt.Println(err)
		return
	}
}
