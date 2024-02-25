/**
 ******************************************************************************
 * @file           : metrics.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/13
 ******************************************************************************
 */

package jinzhu

import (
	"gorm.io/gorm"
	"nunu/backend/model"
	"nunu/backend/repo"
	"time"
)

var _ repo.IUserMetricsRepo = (*UserMetricsRepo)(nil)

type UserMetricsRepo struct {
	db *gorm.DB
}

func (u *UserMetricsRepo) UpdateUserMetrics(UserID uint64, action uint8) error {
	metric := model.UserMetric{UserId: UserID}
	u.db.Model(metric).Where("user_id = ? ", UserID).First(metric)
	metric.LatestTrendsOn = time.Now().Unix()
	switch action {
	case repo.MetricActionCreateTweet:
		metric.TweetsCount++
	case repo.MetricActionDeleteTweet:
		if metric.TweetsCount > 0 {
			metric.TweetsCount--
		}
	}
	return u.db.Save(metric).Error
}

func (u *UserMetricsRepo) AddUserMetrics(userID uint64) error {
	return u.db.Create(&model.UserMetric{UserId: userID}).Error
}

func (u *UserMetricsRepo) DeleteUserMetric(userId uint64) error {
	return u.db.Model(&model.UserMetric{UserId: userId}).Updates(map[string]any{
		"deleted_on": time.Now().Unix(),
		"is_del":     1,
	}).Error
}

func NewUserMetricsRepo(db *gorm.DB) repo.IUserMetricsRepo {
	return &UserMetricsRepo{db: db}
}
