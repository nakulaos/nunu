/**
 ******************************************************************************
 * @file           : admin.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/26
 ******************************************************************************
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"nunu/backend/api/v1/base"
	"nunu/backend/constant"
	"nunu/backend/model"
)

const (
	userStatusNormal = iota
	userStatusClosed
)

func Admin() gin.HandlerFunc {
	// 必须在jwtAuth 后使用
	return func(c *gin.Context) {
		if user, exist := c.Get("user"); exist {
			if userModel, ok := user.(*model.User); ok {
				if userModel.Status == userStatusNormal && userModel.IsAdmin {
					c.Next()
					return
				}
			}
		}

		base.Render(c, nil, constant.ErrNoAdminPermission)
		c.Abort()

	}
}
