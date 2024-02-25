/**
 ******************************************************************************
 * @file           : entry.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/4
 ******************************************************************************
 */

package service

import "nunu/backend/init/conf"

func init() {
	customProfile := conf.GetConfig().CustomProfile
	userProfile := customProfile.UserProfile

	// 初始化service层的一些配置
	_allowNewUserRegister_ = userProfile.AllowNewUserRegister

}
