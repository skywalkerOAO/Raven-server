package dao

import (
	"errors"
	"github.com/skywalkerOAO/gosql"
	"reaven-server/src/pogo"
)

func LoginRedis(item *pogo.AccTokenStruct) (bool, error) {
	c := gosql.GetRedisCon("rds")
	_, err := c.Do("set", item.Uuid, item.Token)
	defer c.Close()
	if err != nil {
		return false, err
	}
	return true, nil
}
func LoginSql(item *pogo.BindAccPsw) (bool, error) {
	con, _ := gosql.GetDBCon("ms")
	res := gosql.Query(con, "SELECT * FROM [user] WHERE usr = ? AND psw = ?", item.Uuid, item.Password)
	defer con.Close()
	if len(res) > 0 {
		return true, nil
	}
	e := errors.New("sql wrong")
	return false, e
}
