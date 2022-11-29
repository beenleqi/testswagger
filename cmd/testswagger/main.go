package main

import "github.com/gin-gonic/gin"

func main()  {
	r := gin.New()

	r.GET("/", func(context *gin.Context) {
		context.Writer.Write([]byte("hello world"))
		return
	})
}
