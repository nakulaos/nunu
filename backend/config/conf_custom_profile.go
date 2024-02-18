/**
 ******************************************************************************
 * @file           : conf_custom_profile.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/13
 ******************************************************************************
 */

package config

type CustomProfile struct {
	UserProfile UserProfile
}

type UserProfile struct {
	AllowNewUserRegister bool
}
