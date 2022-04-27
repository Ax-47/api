package question

import "github.com/gin-gonic/gin"

func M(c *gin.Context) {
	Leve := c.PostForm("leve")
	Asm := c.PostForm("asm")
	if Leve == "1" {
		if Asm == "hackerman" {
			c.JSON(200, gin.H{
				"message": Asm + "leve:" + Leve + " True",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "leve:" + Leve + " False",
			})
		}
	}

}
