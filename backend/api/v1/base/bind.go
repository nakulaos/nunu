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
	"nunu/backend/pkg/validation"
	"nunu/backend/pkg/xerror"
)

func Bind(c *gin.Context, req interface{}) xerror.Error {
	err := c.ShouldBind(req)

	if err != nil {
		// 只拿去检验的第一个错误
		return xerror.NewErrorWithMapAndXError(constant.ErrInvalidParams, map[string]interface{}{"detail": validation.GetValidMsg(err, req)})
	}
	return nil
}
