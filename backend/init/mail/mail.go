/**
 ******************************************************************************
 * @file           : mail.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/17
 ******************************************************************************
 */

package mail

import (
	"nunu/backend/init/conf"
	"nunu/backend/persistence/security"
	"nunu/backend/repo"
	"sync"
)

var (
	_mailVerify repo.IMailVerifyServiceRepo
	_once       sync.Once
)

func initMailVerify() {
	mailConf := conf.GetConfig().MailConf
	_mailVerify = security.NewMailVerifyServiceRepo(
		mailConf.HttpHost,
		mailConf.HttpPort,
		mailConf.User,
		mailConf.Password,
		mailConf.DefaultFromEmail,
		mailConf.UseSSL,
		mailConf.UserTLS)
}

func GetMailVerify() repo.IMailVerifyServiceRepo {
	_once.Do(func() {
		initMailVerify()
	})
	return _mailVerify
}
