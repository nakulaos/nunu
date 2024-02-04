/**
 ******************************************************************************
 * @file           : web.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/4
 ******************************************************************************
 */

package dto

import "nunu/backend/model"

type BaseInfo struct {
	User *model.User
}

type SimpleInfo struct {
	Uid uint64
}
