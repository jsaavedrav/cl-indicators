package main

import (
	"cl-indicators/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/indicators", func(ctx *gin.Context) {
		ufVal := util.GetUF()
		utmVal := util.GetUTM()
		ctx.JSON(http.StatusOK, gin.H{
			"uf":  ufVal,
			"utm": utmVal,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
