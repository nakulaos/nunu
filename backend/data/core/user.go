/**
 ******************************************************************************
 * @file           : user.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/24
 ******************************************************************************
 */

package core

import "nunu/backend/data/model"

type UserManageStore interface {
	GetUserByID(id uint64) (*model.User, error)
}
