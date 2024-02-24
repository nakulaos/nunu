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
	"fmt"
	"gorm.io/gorm"
	"nunu/backend/model"
	"nunu/backend/repo"
	"strings"
)

var (
	_ repo.IUserManageRepo = (*UserManageRepo)(nil)
)

type UserManageRepo struct {
	db  *gorm.DB
	ums repo.IUserMetricsRepo
}

func NewUserManageRepo(db *gorm.DB, ums repo.IUserMetricsRepo) repo.IUserManageRepo {
	return &UserManageRepo{db: db, ums: ums}
}

func (u *UserManageRepo) GetUserByID(id uint64) (*model.User, error) {
	var user model.User
	var err error

	if id > 0 {
		err = u.db.Where("id = ? AND is_del = ?", id, 0).First(&user).Error
	}

	return &user, err
}

func (u *UserManageRepo) GetUserByPhone(phone string) (*model.User, error) {
	var user model.User
	var err error

	if phone != "" {
		err = u.db.Where("phone = ? AND is_del = ?", phone, 0).First(&user).Error
	}

	return &user, err
}

func (u *UserManageRepo) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	var err error

	if username != "" {
		err = u.db.Where("username = ? AND is_del = ?", username, 0).First(&user).Error
	}

	return &user, err
}

func (u *UserManageRepo) GetUserByMail(mail string) (*model.User, error) {
	var user model.User
	var err error

	if mail != "" {
		err = u.db.Where("mail = ? AND is_del = ?", mail, 0).First(&user).Error
	}

	return &user, err
}

func (u *UserManageRepo) GetUsersByIDS(ids []uint64) ([]*model.User, error) {
	var users []*model.User
	var err error

	err = u.db.Where("id in ?", ids).Find(&users).Error

	return users, err
}

func (u *UserManageRepo) GetUserByKeyword(keyword string) ([]*model.User, error) {
	var user []*model.User
	var err error //nolint:wsl

	keyword = strings.Trim(keyword, " ") + "%"
	if keyword == "%" {
		err = u.db.Offset(0).Limit(10).Where("id ASC").Find(&user).Error //nolint:gomnd
	} else {
		err = u.db.Offset(0).Limit(10).Where("username LIKE ?", keyword).Find(&user).Error
	}

	return user, err
}

func (u *UserManageRepo) CreateUser(user *model.User) (*model.User, error) {
	err := u.db.Create(&user).Error

	if err == nil {
		// 忽略错误
		_ = u.ums.AddUserMetrics(user.ID)
	}
	//nolint:gci
	return user, err
}

func (u *UserManageRepo) UpdateUser(user *model.User) error {
	err := u.db.Model(&model.User{}).Where("id = ? AND is_del = ?", user.ID, 0).Save(user).Error
	return err

}

func (u *UserManageRepo) GetRegisterUserCount() (res int64, err error) {
	err = u.db.Model(&model.User{}).Count(&res).Error
	return
}

func (u *UserManageRepo) GetUserProfileByName(username string) (*model.UserProfile, error) {
	var res model.UserProfile
	err := u.db.Table(_user_).Joins(fmt.Sprintf("LEFT JOIN %s m ON %s.id == m.id", _userMetric_, _user_)).
		Where(fmt.Sprintf("%s.username = ? AND %s.is_del = 0", _user_, _user_), username).
		Select([]string{
			fmt.Sprintf("%s.id", _user_),
			fmt.Sprintf("%s.username", _user_),
			fmt.Sprintf("%s.nickname", _user_),
			fmt.Sprintf("%s.phone", _user_),
			fmt.Sprintf("%s.status", _user_),
			fmt.Sprintf("%s.avatar", _user_),
			fmt.Sprintf("%s.balance", _user_),
			fmt.Sprintf("%s.is_admin", _user_),
			fmt.Sprintf("%s.created_on", _user_),
			"m.tweets_count",
		}).First(&res).Error
	return &res, err
}
