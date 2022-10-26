package utils

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

/*CreateToken
 * @Author Hex575A
 * @生成token
 * @Date 13:44 2022/10/25
 * @Param
 * @return
 */
func CreateToken(Account string) (string, error) {
	e := errors.New("wrong account length")
	if len(Account) < 10 {
		return "", e
	}
	signUnixTime := time.Now().Unix()
	expireUnixTime := signUnixTime + 86400*1
	SUT := strconv.FormatInt(signUnixTime, 10)
	EUT := strconv.FormatInt(expireUnixTime, 10)
	cptres, err := encrypt(Account+"000000", Account+SUT+EUT)
	if err != nil {
		return "", err
	}
	return cptres, nil
}

/*DecryptToken
 * @Author Hex575A
 * @Description 解析Token并返回uid
 * @Date 15:07 2022/10/25
 * @Param
 * @return
 */
func DecryptToken(Account string, token string) (string, error) {
	var e error
	if len(Account) < 10 {
		e = errors.New("wrong account length")
		return "", e
	}
	if len(token) < 10 {
		e = errors.New("wrong token length")
		return "", e
	}
	res, err := decrypt(Account+"000000", token)
	if err != nil {
		return "", err
	}
	return res, nil
}

func encrypt(key string, enVal string) (string, error) {
	cryptText, err := EncryptByAes([]byte(enVal), []byte(key))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return cryptText, nil
}

func decrypt(key string, deVal string) (string, error) {
	decryptText, err := DecryptByAes(deVal, []byte(key))
	if err != nil {
		return "", err
	}
	d := string(decryptText)
	if d == "" {
		var e error
		e = errors.New("wrong token with uuid ")
		return "", e
	}
	return d, nil
}
