package services

import (
	"reaven-server/src/dao"
	"reaven-server/src/pogo"
	"reaven-server/utils"
)

func Login(item *pogo.BindAccPsw) (string, error) {
	token, err := utils.CreateToken(item.Uuid)
	if err != nil {
		return "", err
	}
	atStruct := &pogo.AccTokenStruct{
		Uuid:  item.Uuid,
		Token: token,
	}
	sql, err := dao.LoginSql(item)
	if !sql {
		return "", err
	}
	_, err = dao.LoginRedis(atStruct)
	if err != nil {
		return "", err
	}
	return token, nil
}
