package test

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

/**
 * @description：TODO
 * @author     ：yangsen
 * @date       ：2021/6/29 下午2:23
 * @company    ：eeo.cn
 */

func TestJwt(t *testing.T) {
	//加密
	token, err := CreateToken(`10010`, `yang_sen@sohu.com`)
	t.Log(token, err)
	//解密
	res, err := ParseToken(token, `yang_sen@sohu.com`)
	t.Log(res, err)
	//过期时间
	exp := int64(res["exp"].(float64))
	guoqi := time.Unix(exp, 0).Format("2006-01-02 15:04:05")
	t.Log(guoqi)
}

func CreateToken(uid, secret string) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(token string, secret string) (map[string]interface{}, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return claim.Claims.(jwt.MapClaims), nil
}
