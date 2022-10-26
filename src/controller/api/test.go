package api

import "github.com/gin-gonic/gin"

func Test(c *gin.Context) {
	c.JSON(200, gin.H{"code": 200, "msg": "测试成功"})
}
