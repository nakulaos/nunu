/**
 ******************************************************************************
 * @file           : validation.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/16
 ******************************************************************************
 */

package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"nunu/backend/i18n"
	"reflect"
)

// GetValidMsg 用于自定义输出validate校验的国际化信息
func GetValidMsg(err error, obj any) string {
	// obj 的类型
	// TypeOf 返回i的动态类型信息，即reflect.Type类型
	getObj := reflect.TypeOf(obj)

	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		// 断言成功
		for _, e := range errs {
			// 循环每一个错误信息
			// 根据报错名，获取结构体的具体字段
			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
				msg := i18n.GetMsgWithMap("ErrValidation."+fmt.Sprintf(f.Tag.Get("msg")), nil)
				return msg
			}
		}
	}
	return err.Error()
}
