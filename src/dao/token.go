package dao

import (
	"fmt"
	"github.com/skywalkerOAO/gosql"
	"reaven-server/src/pogo"
)

func GetToken(item *pogo.AccTokenStruct) string {
	c := gosql.GetRedisCon("rds")
	fmt.Println(item.Uuid)
	r, err := c.Do("get", item.Uuid)
	defer c.Close()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if r != nil {
		return string(r.([]byte))
	}
	return ""
}
