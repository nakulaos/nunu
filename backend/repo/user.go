/**
 ******************************************************************************
 * @file           : user.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/24
 ******************************************************************************
 */

package repo

import "nunu/backend/model"

type IUserManageRepo interface {
	GetUserByID(id uint64) (*model.User, error)
	GetUserByPhone(phone string) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	GetUserByMail(mail string) (*model.User, error)
	GetUserByIDS(ids []uint64) ([]*model.User, error)
	// GetUserByKeyword 通过搜索用户名找账号
	GetUserByKeyword(keyword string) ([]*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) error
	GetRegisterUserCount() (int64, error)
	GetUserProfileByName(username string) (*model.UserProfile, error)
}
