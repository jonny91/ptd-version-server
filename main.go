package main

import (
	"PTDVersionServer/config"
	"PTDVersionServer/controlls"
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	err error
	wg  sync.WaitGroup
)

func main() {
	if !config.ReadAll() {
		return
	}

	err = controlls.InitMQ()
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 3; i++ {
		go controlls.ReceiveFromMQ("mission")
	}

	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	g.LoadHTMLGlob("./html/*")

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
