package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"reaven-server/src/pogo"
	"reaven-server/src/services"
)

func Login(c *gin.Context) {
	item := &pogo.BindAccPsw{}
	if err := c.BindJSON(item); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "给定的关键字类型不正确或不存在，请检查参数"})
		return
	}
	res, err := services.Login(item)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"code": 500, "msg": "用户名或密码不正确！"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "登陆成功", "token": res})
}
