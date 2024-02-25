/**
 ******************************************************************************
 * @file           : xerror.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/3
 ******************************************************************************
 */

package xerror

import "nunu/backend/i18n"

type Error interface {
	error
	GetErrCode() int
}

type XError struct {
	Code int // 错误代码
	// 多国际化需要用下面的信息
	Map map[string]interface{} // 换算key
	Key string                 // lang对应的键
}

func (e XError) GetErrCode() int {
	return e.Code
}

func (e XError) Error() string {
	content := ""
	content = i18n.GetMsgWithMap(e.Key, e.Map)
	return content
}

func NewError(code int, key string) Error {
	return XError{
		Code: code,
		Key:  key,
		Map:  nil,
	}
}

func NewErrorWithMap(code int, key string, maps map[string]interface{}) Error {
	return XError{
		Code: code,
		Map:  maps,
		Key:  key,
	}
}

func NewErrorWithMapAndXError(e Error, maps map[string]interface{}) Error {
	err, ok := e.(XError)
	if ok {
		return XError{
			Code: err.Code,
			Map:  maps,
			Key:  err.Key,
		}
	} else {
		return XError{
			Code: 70000,
			Map:  maps,
			Key:  "ErrUnKnownError",
		}
	}

}
