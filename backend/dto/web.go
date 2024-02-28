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

func (b *BaseInfo) SetUser(user *model.User) {
	b.User = user
}

type SimpleInfo struct {
	Uid uint64
}

func (s *SimpleInfo) SetUserUID(uid uint64) {
	s.Uid = uid
}
