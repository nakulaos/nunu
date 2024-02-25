/**
 ******************************************************************************
 * @file           : metrics.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/13
 ******************************************************************************
 */

package model

type UserMetric struct {
	*Model
	UserId         uint64
	TweetsCount    int
	LatestTrendsOn int64
}
