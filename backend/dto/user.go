/**
 ******************************************************************************
 * @file           : user.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/26
 ******************************************************************************
 */

package dto

type ChangeAvatarReq struct {
	BaseInfo `json:"-" binding:"-"`
	Avatar   string `json:"avatar" form:"avatar" binding:"required"`
}
