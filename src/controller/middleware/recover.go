package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			//封装通用json返回
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "服务器内部错误",
				"data": errorToString(r),
			})
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		k := v.Error()
		return recoverType(k)
	default:
		k := r.(string)
		return recoverType(k)
	}
}

func recoverType(r string) string {
	if strings.Contains(r, "index out of range") {
		return "数组越界"
	}
	return "未知错误"
}
