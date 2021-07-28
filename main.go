package main

import (
	"github.com/gin-gonic/gin"
	"github.com/github-hewei/go-gin-demo/user"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		log.Println("use middleware")
	})

	r.LoadHTMLGlob("views/*")

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Gin")
	})

	r.GET("/ping", Ping)

	r.GET("/user", user.Lists)
	r.GET("/user/create", user.Create)
	r.GET("/user/edit/:id", user.Edit)
	r.POST("/user/save", user.Save)
	r.POST("/user/delete", user.Delete)

	r.Run(":8080")
}

func Ping(c *gin.Context) {
	c.Header("Content-type", "application/json;charset=utf-8")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
