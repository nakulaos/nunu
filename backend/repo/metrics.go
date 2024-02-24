/**
 ******************************************************************************
 * @file           : metrics.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/13
 ******************************************************************************
 */

package repo

const (
	MetricActionCreateTweet uint8 = iota
	MetricActionDeleteTweet
)

type IUserMetricsRepo interface {
	UpdateUserMetrics(UserID uint64, action uint8) error
	AddUserMetrics(UserID uint64) error
	DeleteUserMetric(userId uint64) error
}
