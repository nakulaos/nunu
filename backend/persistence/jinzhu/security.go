/**
 ******************************************************************************
 * @file           : security.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/17
 ******************************************************************************
 */

package jinzhu

import (
	"gorm.io/gorm"
	"math/rand"
	"nunu/backend/model"
	"nunu/backend/repo"
	"strconv"
	"time"
)

var _ repo.ISecurityServiceRepo = (*SecurityServiceRepo)(nil)
var (
	_captchaEmail model.CaptchaType = 2
	_captchaPhone model.CaptchaType = 1
)

type SecurityServiceRepo struct {
	db         *gorm.DB
	rand       *rand.Rand
	mailVerify repo.IMailVerifyServiceRepo
}

func NewSecurityServiceRepo(db *gorm.DB, rand *rand.Rand, mailVerify repo.IMailVerifyServiceRepo) repo.ISecurityServiceRepo {
	return &SecurityServiceRepo{db: db, mailVerify: mailVerify, rand: rand}
}

func (s *SecurityServiceRepo) GetLatestPhoneCaptcha(phone string) (*model.Captcha, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SecurityServiceRepo) UsePhoneCaptcha(captcha *model.Captcha) error {
	//TODO implement me
	panic("implement me")
}

func (s *SecurityServiceRepo) SendPhoneCaptcha(phone string) error {
	//TODO implement me
	panic("implement me")
}

func (s *SecurityServiceRepo) GetLatestMailCaptcha(mail string) (*model.Captcha, error) {
	capt := &model.Captcha{}
	if mail != "" {

		err := s.db.Where("destination = ? AND type = ? AND is_del =?", mail, _captchaEmail, 0).Last(&capt).Error
		if err != nil {
			return nil, err
		}
		return capt, err

	}
	return capt, nil
}

func (s *SecurityServiceRepo) UseMailCaptcha(captcha *model.Captcha) error {
	captcha.UseTimes++
	return s.db.Model(&model.Captcha{}).Where("id = ? AND is_del = ?", captcha.ID, 0).Save(captcha).Error

}

func (s *SecurityServiceRepo) SendMailCaptcha(mail string) error {
	expire := time.Duration(5)

	// 发送验证码
	captcha := strconv.Itoa(s.rand.Intn(900000) + 100000)
	if err := s.mailVerify.SendMailCaptcha(mail, captcha, expire); err != nil {
		return err
	}

	// 写入表
	captchaModel := &model.Captcha{
		Type:        _captchaEmail,
		Destination: mail,
		Captcha:     captcha,
		UseTimes:    0,
		ExpiredOn:   time.Now().Add(expire * time.Minute).Unix(),
	}

	err := s.db.Create(&captchaModel).Error
	return err

}
