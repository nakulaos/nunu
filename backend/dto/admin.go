/**
 ******************************************************************************
 * @file           : ro_admin.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/4
 ******************************************************************************
 */

package dto

type ChangeUserStatusReq struct {
	BaseInfo `json:"-" binding:"-"`
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Status   int    `json:"status" form:"status" binding:"required,oneof=1 2"`
}

type SiteInfoReq struct {
	SimpleInfo `json:"-" binding:"-"`
}

type SiteInfoResp struct {
	RegisterUserCount int64 `json:"register_user_count"`
	OnlineUserCount   int   `json:"online_user_count"`
	HistoryMaxOnline  int   `json:"history_max_online"`
	ServerUpTime      int64 `json:"server_up_time"`
}
