/**
 ******************************************************************************
 * @file           : captcha.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/17
 ******************************************************************************
 */

package model

// CaptchaType 验证码类型 1手机号 2 邮箱
type CaptchaType int

type Captcha struct {
	//*Model 表示 Captcha 结构体嵌套了一个指向 Model 类型的指针。
	//这意味着 Captcha 类型包含了 Model 类型的所有字段和方法，
	//并且在实例化时需要手动分配和初始化 Model 对象。
	//，Model 表示 Captcha 结构体嵌套了一个 Model 类型的实例。
	//这意味着 Captcha 类型包含了 Model 类型的所有字段和方法，
	//而在实例化时会自动创建和初始化 Model 对象。
	//选择使用指针或实例取决于你的设计需求。
	//使用指针可以实现更轻量的对象，因为不会复制 Model 对象的实例。
	//但同时，使用指针需要手动管理内存，确保 Model 对象被正确初始化。
	//使用实例则更方便，因为对象的初始化和清理都由 Go 语言的垃圾回收器负责。
	*Model
	Type        CaptchaType
	Destination string `json:"destination"`
	Captcha     string `json:"captcha"`
	UseTimes    int    `json:"use_times"`
	ExpiredOn   int64  `json:"expired_on"`
}
