/**
 ******************************************************************************
 * @file           : render.go
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
	"nunu/backend/pkg/xerror"
	"reflect"
)

func Render(c *gin.Context, data interface{}, err xerror.Error) {
	if err == nil {
		if data == nil || reflect.ValueOf(data).IsNil() {
			c.JSON(constant.CodeSuccess, &Response{
				Code: 0,
				Msg:  "Success",
			})
		} else {
			c.JSON(constant.CodeSuccess, &Response{
				Code: 0,
				Msg:  "Success",
				Data: data,
			})
		}

	} else {
		res := &Response{
			Code: err.GetErrCode(),
			Msg:  err.Error(),
		}
		c.JSON(GetHttpStatusCode(err), res)
	}
}

func GetHttpStatusCode(e xerror.Error) int {
	switch e.GetErrCode() {
	case constant.Success.GetErrCode():
		return constant.CodeSuccess
	case constant.ErrInvalidParams.GetErrCode():
		return constant.CodeBadRequest
	case constant.ErrInternalServer.GetErrCode():
		return constant.CodeErrInternalServer
	case constant.ErrUnauthorizedAuthNotExist.GetErrCode():
		fallthrough
	case constant.ErrUnauthorizedAuthFailed.GetErrCode():
		return constant.CodeErrUnAuthorized
	}
	return constant.CodeErrInternalServer
}
