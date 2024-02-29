/**
 ******************************************************************************
 * @file           : bind.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/2
 ******************************************************************************
 */

package base

import (
	"github.com/gin-gonic/gin"
	"nunu/backend/constant"
	"nunu/backend/model"
	"nunu/backend/pkg/validation"
	"nunu/backend/pkg/xerror"
)

type SetterUser interface {
	SetUser(user *model.User)
}

type SetterUserUID interface {
	SetUserUID(uid uint64)
}

func Bind(c *gin.Context, req interface{}) xerror.Error {
	err := c.ShouldBind(req)

	if err != nil {
		// 只拿去检验的第一个错误
		return xerror.NewErrorWithMapAndXError(constant.ErrInvalidParams, map[string]interface{}{"detail": validation.GetValidMsg(err, req)})
	}

	// 设置用户
	if setter, ok := req.(SetterUser); ok {
		if user, ok := c.Get("user"); ok {
			if userModel, ok := user.(*model.User); ok {
				setter.SetUser(userModel)
			}
		}
	}

	// 设置uid
	if setter, ok := req.(SetterUserUID); ok {
		if uid, ok := c.Get("uid"); ok {
			if uidModel, ok := uid.(uint64); ok {
				setter.SetUserUID(uidModel)
			}
		}
	}

	return nil
}
