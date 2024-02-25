/**
 ******************************************************************************
 * @file           : cache.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/18
 ******************************************************************************
 */

package cache

import (
	"github.com/go-redis/redis"
	"nunu/backend/repo"
	"strconv"
	"time"
)

var _ repo.IRedisCacheRepo = (*RedisCacheRepo)(nil)

type RedisCacheRepo struct {
	c *redis.Client
}

func (r *RedisCacheRepo) GetCountMailCaptcha(mail string) (int64, error) {
	cmd := r.c.Get(_mailCaptchaKey + mail)
	if cmd.Err() == nil {
		return cmd.Int64()
	} else {
		currentTime := time.Now()
		endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(),
			23, 59, 59, 0, currentTime.Location())
		err := r.c.Set(_mailCaptchaKey+mail, 0, endTime.Sub(currentTime)).Err()
		return 0, err
	}
}

func (r *RedisCacheRepo) IncrCountMailCaptcha(mail string) (err error) {
	if err = r.c.Incr(_mailCaptchaKey + mail).Err(); err == nil {
		currentTime := time.Now()
		endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(),
			23, 59, 59, 0, currentTime.Location())
		err = r.c.Expire(_mailCaptchaKey+mail, endTime.Sub(currentTime)).Err()
	}
	return
}

func (r *RedisCacheRepo) SetPushToSearchJob() error {
	//TODO implement me
	panic("implement me")
}

func (r *RedisCacheRepo) DelPushToSearchJob() error {
	//TODO implement me
	panic("implement me")
}

func (r *RedisCacheRepo) SetImgCaptcha(id string, value string) error {
	return r.c.Set(_imgCaptchaKey+id, value, time.Second*300).Err()
}

func (r *RedisCacheRepo) GetImgCaptcha(id string) (string, error) {
	return r.c.Get(_imgCaptchaKey + id).Result()
}

func (r *RedisCacheRepo) DelImgCaptcha(id string) error {
	return r.c.Del(_imgCaptchaKey + id).Err()
}

func (r *RedisCacheRepo) GetCountPhoneCaptcha(phone string) (int64, error) {
	// TODO: 有bug
	cmd := r.c.Get(_phoneCaptchaKey + phone)
	err := cmd.Err()
	if err != nil {
		return 0, err
	}
	res, err := strconv.ParseInt(cmd.String(), 10, 64)
	return res, err
}

func (r *RedisCacheRepo) IncrCountPhoneCaptcha(phone string) (err error) {
	// TODO: 有bug
	if err = r.c.Incr(_phoneCaptchaKey + phone).Err(); err == nil {
		currentTime := time.Now()
		endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(),
			23, 59, 59, 0, currentTime.Location())
		err = r.c.Expire(_phoneCaptchaKey+phone, endTime.Sub(currentTime)).Err()
	}
	return
}

func (r *RedisCacheRepo) GetCountLoginErr(id uint64) (int64, error) {
	cmd := r.c.Get(_usernameLoginErrCountKey + strconv.Itoa(int(id)))
	if cmd.Err() == nil {
		return cmd.Int64()
	} else {
		currentTime := time.Now()
		endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(),
			23, 59, 59, 0, currentTime.Location())
		err := r.c.Set(_usernameLoginErrCountKey+strconv.Itoa(int(id)), 0, endTime.Sub(currentTime)).Err()
		return 0, err
	}
}

func (r *RedisCacheRepo) DelCountLoginErr(id uint64) error {
	return r.c.Del(_usernameLoginErrCountKey + strconv.Itoa(int(id))).Err()
}

func (r *RedisCacheRepo) IncrCountLoginErr(id uint64) (err error) {
	if err = r.c.Incr(_usernameLoginErrCountKey + strconv.Itoa(int(id))).Err(); err == nil {
		currentTime := time.Now()
		endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(),
			23, 59, 59, 0, currentTime.Location())
		err = r.c.Expire(_usernameLoginErrCountKey+strconv.Itoa(int(id)), endTime.Sub(currentTime)).Err()
	}
	return
}

func (r *RedisCacheRepo) GetCountWhisper(uid uint64) (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisCacheRepo) IncrCountWhisper(uid uint64) error {
	//TODO implement me
	panic("implement me")
}

func (r *RedisCacheRepo) SetRechargeStatus(tradeNo string) error {
	//TODO implement me
	panic("implement me")
}

func (r *RedisCacheRepo) DelRechargeStatus(tradeNo string) error {
	//TODO implement me
	panic("implement me")
}

func NewRedisCacheRepo(c *redis.Client) repo.IRedisCacheRepo {
	return &RedisCacheRepo{c: c}
}
