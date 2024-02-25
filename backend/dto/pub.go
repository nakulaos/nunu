/**
 ******************************************************************************
 * @file           : pub.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/13
 ******************************************************************************
 */

package dto

type GetCaptchaResp struct {
	Id      string `json:"id"`
	Content string `json:"b64s"`
}

type SendCaptchaWithPhoneReq struct {
	Phone        string `json:"phone" form:"phone" binding:"required"`
	ImgCaptcha   string `json:"img_captcha" form:"img_captcha" binding:"required"`
	ImgCaptchaID string `json:"img_captcha_id" form:"img_captcha_id" binding:"required"`
}

type SendCaptchaWithMailReq struct {
	Mail         string `json:"mail" form:"mail" binding:"required"`
	ImgCaptcha   string `json:"img_captcha" form:"img_captcha" binding:"required"`
	ImgCaptchaID string `json:"img_captcha_id" form:"img_captcha_id" binding:"required"`
}

type LoginReqWithUserName struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginResp struct {
	Token string `json:"token"`
}

type RegisterReq struct {
	Username     string `json:"username" form:"username" binding:"required,gt=3,lt=12,alphanumunicode" msg:"WrongUserFormat"`
	Password     string `json:"password" form:"password" binding:"required,gt=8,lt=18,password" msg:"WrongPasswordFormat"`
	Mail         string `json:"mail" form:"mail" binding:"email,required" msg:"IncorrectMail"`
	Phone        string `json:"phone" form:"phone" binding:"required,omitempty" msg:"LackOfPhone"`
	MailCaptcha  string `json:"mail_captcha" form:"mail_captcha"`
	PhoneCaptcha string `json:"phone_captcha" form:"phone_captcha"`
}

type RegisterResp struct {
	UserID   uint64 `json:"id"`
	Username string `json:"username"`
	Mail     string `json:"mail"`
	Phone    string `json:"phone"`
}
