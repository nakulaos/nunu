/**
 ******************************************************************************
 * @file           : jwt.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/24
 ******************************************************************************
 */

package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/golang-jwt/jwt/v4"
	"nunu/backend/init/conf"
	"nunu/backend/model"
	"time"
)

type Claims struct {
	UID      uint64 `json:"uid"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GetJWTSecret() []byte {
	jwtConf := conf.GetConfig().JwtConf
	return []byte(jwtConf.Secret)
}

func GenerateToken(user *model.User) (string, error) {
	jwtConf := conf.GetConfig().JwtConf
	expireTime := time.Now().Add(jwtConf.Expire)
	claims := Claims{
		UID:      user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    IssuerFrom(user.Salt),
		},
	}
	// 使用指定的签名方法创建签名对象
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

func ParseToken(token string) (res *Claims, err error) {
	// 解释token
	var tokenClaims *jwt.Token
	tokenClaims, err = jwt.ParseWithClaims(token, &Claims{}, func(_ *jwt.Token) (any, error) {
		return GetJWTSecret(), nil
	})
	// 校验token
	if err == nil && tokenClaims != nil && tokenClaims.Valid {
		res, _ = tokenClaims.Claims.(*Claims)
	} else {
		err = jwt.ErrTokenNotValidYet
	}
	return
}

func IssuerFrom(data string) string {
	jwtConf := conf.GetConfig().JwtConf
	// 将issuer+data 转化为[]byte
	contents := make([]byte, 0, len(jwtConf.Issuer)+len(data))
	copy(contents, []byte(jwtConf.Issuer))
	contents = append(contents, []byte(data)...)
	// 获取MD5的校验值
	res := md5.Sum(contents)
	// 转化成16进制
	return hex.EncodeToString(res[:])
}
