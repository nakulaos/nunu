/**
 ******************************************************************************
 * @file           : constant.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/13
 ******************************************************************************
 */

package service

const (
	userStatusNormal = iota
	userStatusClosed
)

var (
	_allowNewUserRegister_ bool
)
