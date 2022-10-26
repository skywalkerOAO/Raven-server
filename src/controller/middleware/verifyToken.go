package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reaven-server/src/pogo"
	"reaven-server/src/services"
	"reaven-server/utils"
)

func VerifyToken(c *gin.Context) {
	defer func() {
		token := c.Request.Header.Get("token")
		uuid := c.Request.Header.Get("uuid")
		_, err := utils.DecryptToken(uuid, token)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusForbidden, gin.H{"code": 403, "msg": "token不正确", "detail": err.Error()})
			return
		}
		item := &pogo.AccTokenStruct{
			Uuid:  uuid,
			Token: token,
		}
		if !services.ChkToken(item) {
			c.Abort()
			c.JSON(http.StatusForbidden, gin.H{"code": 403, "msg": "token已过期,请重新申请"})
			return
		}
		c.Next()
	}()

}
