/**
 ******************************************************************************
 * @file           : pub.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/13
 ******************************************************************************
 */

package service

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/gofrs/uuid/v5"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"image/color"
	"image/png"
	"nunu/backend/IService"
	"nunu/backend/assets"
	"nunu/backend/constant"
	"nunu/backend/dto"
	"nunu/backend/model"
	"nunu/backend/pkg/utils"
	"nunu/backend/pkg/xerror"
	"nunu/backend/repo"
	"time"
)

var _ IService.IPubService = (*PubService)(nil)
var _maxCountCaptcha int64 = 3 // 验证码发送次数
var _maxCaptchaTimes int = 5   // 单个验证码（手机，邮箱）最大使用次数

type PubService struct {
	UserManageRepo      repo.IUserManageRepo
	SecurityServiceRepo repo.ISecurityServiceRepo
	RedisCacheRepo      repo.IRedisCacheRepo
}

func NewPubService(userManageRepo repo.IUserManageRepo, securityServiceRepo repo.ISecurityServiceRepo, redisCacheRepo repo.IRedisCacheRepo) IService.IPubService {
	return &PubService{UserManageRepo: userManageRepo, SecurityServiceRepo: securityServiceRepo, RedisCacheRepo: redisCacheRepo}
}

func (s *PubService) SendCaptchaWithMailReq(req *dto.SendCaptchaWithMailReq) xerror.Error {
	// 验证图片验证码
	if imgCaptcha, err := s.RedisCacheRepo.GetImgCaptcha(req.ImgCaptchaID); err != nil || string(imgCaptcha) != req.ImgCaptcha {
		logrus.Debugf("get img captcha err:%s expect:%s got:%s", err, imgCaptcha, req.ImgCaptcha)
		return xerror.NewErrorWithMapAndXError(constant.ErrErrorCaptchaPassword, nil)
	}

	// 删除图片验证码
	s.RedisCacheRepo.DelImgCaptcha(req.ImgCaptchaID)

	// 判断是否验证码次数是否已达上限
	if count, _ := s.RedisCacheRepo.GetCountMailCaptcha(req.Mail); count > _maxCountCaptcha {
		logrus.Errorf("The number of times the email (%s)verification code has been obtained has reached the upper limit", req.Mail)
		return xerror.NewErrorWithMapAndXError(constant.ErrTooManyMailCaptchaSend, nil)
	}

	if err := s.SecurityServiceRepo.SendMailCaptcha(req.Mail); err != nil {
		logrus.Errorf("Failed to  sending email (%s) verification code:%s", req.Mail, err.Error())
		return xerror.NewErrorWithMapAndXError(constant.ErrInternalServer, nil)
	}

	if err := s.RedisCacheRepo.IncrCountMailCaptcha(req.Mail); err != nil {
		logrus.Errorf("Error calculating the number of times an email (%s) verification code was obtained", req.Mail)
	}

	return nil

}

func (s *PubService) SendCaptchaWithPhoneReq(req *dto.SendCaptchaWithPhoneReq) xerror.Error {
	//TODO implement me
	panic("implement me")
}

func (s *PubService) GetCaptcha() (*dto.GetCaptchaResp, xerror.Error) {
	// 获取图片验证码
	capt := captcha.New()
	if err := capt.AddFontFromBytes(assets.ComicBytes); err != nil {
		logrus.Errorf("cap.AddFontFromBytes err:%s", err)
		return nil, xerror.NewErrorWithMapAndXError(constant.ErrInternalServer, nil)
	}
	capt.SetSize(160, 64)
	capt.SetDisturbance(captcha.MEDIUM)
	capt.SetFrontColor(color.RGBA{0, 0, 0, 255})
	capt.SetBkgColor(color.RGBA{218, 240, 228, 255})

	img, password := capt.Create(6, captcha.NUM)
	fmt.Println(password)
	emptyBuff := bytes.NewBuffer(nil)
	if err := png.Encode(emptyBuff, img); err != nil {
		logrus.Errorf("png.Encode err:%s", err)
		return nil, xerror.NewErrorWithMapAndXError(constant.ErrInternalServer, nil)
	}

	key := utils.EncodeMD5(uuid.Must(uuid.NewV4()).String())

	s.RedisCacheRepo.SetImgCaptcha(key, password)
	return &dto.GetCaptchaResp{
		Id:      key,
		Content: "data:image/png;base64," + base64.StdEncoding.EncodeToString(emptyBuff.Bytes()),
	}, nil

}

func (s *PubService) Register(req *dto.RegisterReq) (*dto.RegisterResp, xerror.Error) {
	if !_allowNewUserRegister_ {
		return nil, constant.ErrDisallowUserRegister
	}

	// 检查账号是否重复
	user, err := s.UserManageRepo.GetUserByUsername(req.Username)
	if err == nil && user.Model != nil && user.ID > 0 {
		return nil, xerror.NewErrorWithMapAndXError(constant.ErrUsernameHasExisted,
			map[string]interface{}{"username": user.Username})
	}

	// 密码加密
	password, salt := utils.EncryptPasswordAndSalt(req.Password)

	// 验证邮箱
	// 邮箱重复性检查
	if req.Mail != "" {
		// 如果邮箱不为空
		// 判断邮箱是否被绑定
		user, err = s.UserManageRepo.GetUserByMail(req.Mail)
		if err == nil && user.Model != nil && user.ID > 0 && user.IsDel == 0 {
			return nil, xerror.NewErrorWithMapAndXError(constant.ErrExistedUserMail, map[string]interface{}{"mail": req.Mail})
		}

		c, err := s.SecurityServiceRepo.GetLatestMailCaptcha(req.Mail)
		if err != nil {
			logrus.Errorf("Internal server error was happened when getting email verification code:%s", errors.Wrapf(err, "req.mail:%s", req.Mail).Error())
			return nil, xerror.NewErrorWithMapAndXError(constant.ErrInternalServer, nil)
		}
		// 增加验证码使用次数
		s.SecurityServiceRepo.UseMailCaptcha(c)
		// 验证码错误
		if c.Captcha != req.MailCaptcha {
			logrus.Errorf("The email (%s) verification code is wrong:req.captcha:%s,internal captcha:%s", req.Mail, req.MailCaptcha, c.Captcha)
			return nil, xerror.NewErrorWithMapAndXError(constant.ErrErrorMailCaptcha, nil)
		}
		// 过期
		if c.ExpiredOn < time.Now().Unix() {
			logrus.Errorf("The email (%s) verification code is expired:req.captcha:%s,internal captcha:%s", req.Mail, req.MailCaptcha, c.Captcha)
			return nil, xerror.NewErrorWithMapAndXError(constant.ErrEmailVerificationCodeExpired, nil)
		}
		// 使用次数过多
		if c.UseTimes >= _maxCaptchaTimes {
			logrus.Errorf("\"The email verification code has reached the maximum number of times it can be used:req.mail:%s,internal.captcha:%s", req.Mail, c.Captcha)
			return nil, xerror.NewErrorWithMapAndXError(constant.ErrMaxMailCaptchaUseTimes, nil)
		}

	}

	// 验证手机号
	if req.Phone != "" {
		logrus.Debugf("This version does not support the mobile phone binding function for the time being")
	}

	user = &model.User{
		Nickname: req.Username,
		Username: req.Username,
		Phone:    req.Phone,
		Password: password,
		Salt:     salt,
		Status:   userStatusNormal,
		Avatar:   "",
		Balance:  0,
		IsAdmin:  false,
		Mail:     req.Mail,
	}

	user, err = s.UserManageRepo.CreateUser(user)
	if err != nil {
		logrus.Errorf("Ds.CreateUser err: %s", err)
		return nil, xerror.NewErrorWithMapAndXError(constant.ErrUserRegisterFailed, nil)
	}

	return &dto.RegisterResp{
		UserID:   user.ID,
		Username: user.Username,
		Mail:     user.Mail,
		Phone:    user.Phone,
	}, nil

}

func (s *PubService) Login(req *dto.LoginReq) (*dto.LoginResp, xerror.Error) {
	//TODO implement me
	panic("implement me")
}
