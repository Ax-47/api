package main

import (
	p "Axapi/path"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.POST("/q", func(c *gin.Context) {
		p.M(c)
	})
	r.POST("/register", func(c *gin.Context) {
		p.Register(c)
	})
	r.POST("/login", func(c *gin.Context) {
		p.Login(c)
	})
	r.Run(":8080")
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	/*
		#Example
		import request
		requets.post("localhost:8080/q",data={"leve":"1","ams":"hackerman"})

	*/

}
