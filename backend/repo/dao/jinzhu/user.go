/**
 ******************************************************************************
 * @file           : user.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/1
 ******************************************************************************
 */

package jinzhu

import (
	"gorm.io/gorm"
	"nunu/backend/repo/core"
	"nunu/backend/repo/model"
)

var (
	_ core.IUserManageRepo = (*UserManageRepo)(nil)
)

type UserManageRepo struct {
	db *gorm.DB
}

func NewUserManageRepo(db *gorm.DB) core.IUserManageRepo {
	return &UserManageRepo{db: db}
}

func (u *UserManageRepo) GetUserByID(id uint64) (*model.User, error) {
	//TODO implement me
	panic("implement me")

}

func (u *UserManageRepo) GetUserByPhone(phone string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserManageRepo) GetUserByUsername(username string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserManageRepo) GetUserByMail(mail string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserManageRepo) GetUserByIDS(ids []uint64) ([]*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserManageRepo) GetUserByKeyword(keyword string) ([]*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserManageRepo) CreateUser(user *model.User) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserManageRepo) UpdateUser(user *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserManageRepo) GetRegisterUserCount() (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserManageRepo) GetUserProfileByName(username string) (*model.UserProfile, error) {
	//TODO implement me
	panic("implement me")
}
