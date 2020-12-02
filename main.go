package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Message struct {
	Message string
}

func main() {
	if !readAll() {
		return
	}

	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	g.LoadHTMLGlob("./html/*")
	g.GET("/ptd", func(context *gin.Context) {
		platform, _ := context.GetQuery("platform")

		v, ok := Versions.Data[platform]
		if !ok {
			context.String(http.StatusBadRequest, "")
		} else {
			c, _ := json.Marshal(v)
			context.String(http.StatusOK, string(c))
		}
	})

	g.GET("/edit", func(context *gin.Context) {
		context.HTML(http.StatusOK, "edit.html", Message{
			Message: "111",
		})
	})

	g.GET("/reload", func(context *gin.Context) {
		readAll()
		context.String(http.StatusOK, "success!")
	})

	err := g.Run("0.0.0.0:19999")
	if err != nil {
		fmt.Println(err)
		return
	}
}
