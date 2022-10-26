package services

import (
	"reaven-server/src/dao"
	"reaven-server/src/pogo"
)

func ChkToken(item *pogo.AccTokenStruct) bool {
	if dao.GetToken(item) == item.Token {
		return true
	}
	return false
}
