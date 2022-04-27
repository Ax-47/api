package main

import (
	q "Axapi/question"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/q", func(c *gin.Context) {
		q.M(c)
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	/*
		#Example
		import request
		requets.post("localhost:8080/q",data={"leve":"1","ams":"hackerman"})

	*/

}
