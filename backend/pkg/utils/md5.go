/**
 ******************************************************************************
 * @file           : md5.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/17
 ******************************************************************************
 */

package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gofrs/uuid/v5"
)

// EncryptPasswordAndSalt 密码加密&生成salt
// 这个是加密算法
func EncryptPasswordAndSalt(password string) (string, string) {
	salt := uuid.Must(uuid.NewV4()).String()[:8]
	password = EncodeMD5(EncodeMD5(password) + salt)
	return password, salt
}

// EncodeMD5 对值进行MD5 加密
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
