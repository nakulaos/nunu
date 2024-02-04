/**
 ******************************************************************************
 * @file           : user.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/24
 ******************************************************************************
 */

package model

type User struct {
	*Model
	Nickname string `json:"nickname"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	Status   int    `json:"status"`
	Avatar   string `json:"avatar"`
	Balance  int64  `json:"balance"`
	IsAdmin  bool   `json:"is_admin"`
	Mail     string `json:"mail"`
}

// UserProfile 用户概要
type UserProfile struct {
	ID          int64  `json:"id" db:"id"`
	Nickname    string `json:"nickname"`
	Username    string `json:"username"`
	Phone       string `json:"phone"`
	Status      int    `json:"status"`
	Avatar      string `json:"avatar"`
	Balance     int64  `json:"balance"`
	IsAdmin     bool   `json:"is_admin"`
	CreatedOn   int64  `json:"created_on"`
	TweetsCount int    `json:"tweets_count"`
	Mail        string `json:"mail"`
}
