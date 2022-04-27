package path

import (
	db "Axapi/db"
	h "Axapi/hashpaww"

	"github.com/gin-gonic/gin"
)

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
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	repassword := c.PostForm("repassword")
	if password == repassword {
		var masp = make(map[string]string)
		masp["username"] = username
		masp["password"] = h.Mhash(password)
		//{"":username,"password":password}
		var Database db.Db_mongo
		Database.Db_start()
		Database.Db_InsertOne(masp)
		c.JSON(200, gin.H{
			"message": "Register suss",
		})

	} else {
		c.JSON(404, gin.H{
			"message": "Register fail",
		})
	}
}
