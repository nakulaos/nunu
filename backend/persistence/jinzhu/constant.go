/**
 ******************************************************************************
 * @file           : constant.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/13
 ******************************************************************************
 */

package jinzhu

import "nunu/backend/constant"

var (
	_user_       string
	_userMetric_ string
)

func InitTableName(prefix string) {
	// todo : wait to finish
	_user_ = prefix + constant.TableUser
	_userMetric_ = prefix + constant.TableUserMetric

}
