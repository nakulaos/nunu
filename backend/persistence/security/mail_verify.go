/**
 ******************************************************************************
 * @file           : mail_verify.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/17
 ******************************************************************************
 */

package security

import (
	"gopkg.in/gomail.v2"
	"nunu/backend/repo"
	"time"
)

var _ repo.IMailVerifyServiceRepo = (*MailVerifyServiceRepo)(nil)

type MailVerifyServiceRepo struct {
	HttpHost         string
	HttpPort         int
	User             string
	Password         string
	DefaultFromEmail string
	UseSSL           bool
	UserTLS          bool
}
type Subject string

const (
	Code  Subject = "平台验证码"
	Note  Subject = "操作通知"
	Alarm Subject = "警告通知"
)

func NewMailVerifyServiceRepo(httpHost string, httpPort int, user string, password string, defaultFromEmail string, useSSL bool, userTLS bool) repo.IMailVerifyServiceRepo {
	return &MailVerifyServiceRepo{HttpHost: httpHost, HttpPort: httpPort, User: user, Password: password, DefaultFromEmail: defaultFromEmail, UseSSL: useSSL, UserTLS: userTLS}
}

func (s *MailVerifyServiceRepo) SendMailCaptcha(mail string, captcha string, expire time.Duration) error {
	m := gomail.NewMessage()
	// 发件人邮箱，发件人名字
	m.SetHeader("From", m.FormatAddress(s.User, s.DefaultFromEmail))
	m.SetHeader("To", mail)
	m.SetHeader("Subject", string(Code))
	m.SetBody("text/html", captcha)
	d := gomail.NewDialer(s.HttpHost, s.HttpPort, s.User, s.Password)
	err := d.DialAndSend(m)
	return err
}
