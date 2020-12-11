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
	g.GET("/edit", controlls.HandleEdit)
	g.GET("/reload", controlls.HandleReload)

	g.POST("/mission", controlls.HandleMission)

	err = g.Run("0.0.0.0:19999")
	if err != nil {
		fmt.Println(err)
		return
	}
}
