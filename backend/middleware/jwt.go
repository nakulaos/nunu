/**
 ******************************************************************************
 * @file           : jwt.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/26
 ******************************************************************************
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"nunu/backend/api/v1/base"
	"nunu/backend/constant"
	"nunu/backend/pkg/utils"
	"nunu/backend/pkg/xerror"
)

// JWTAuth jwt 认证中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.Request.Header.Get("token")
		}
		if token == "" {
			base.Render(c, nil, xerror.NewErrorWithMapAndXError(constant.ErrUnauthorizedTokenError, nil))
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(token)
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				base.Render(c, nil, xerror.NewErrorWithMapAndXError(constant.ErrUnauthorizedTokenTimeout, nil))
			} else {
				base.Render(c, nil, xerror.NewErrorWithMapAndXError(constant.ErrUnauthorizedTokenError, nil))
			}
			c.Abort()
			return
		}

		user, err := _userManageRepo.GetUserByID(claims.UID)
		if err != nil {
			base.Render(c, nil, xerror.NewErrorWithMapAndXError(constant.ErrInternalServer, nil))
		}

		if utils.IssuerFrom(user.Salt) != claims.Issuer {
			base.Render(c, nil, xerror.NewErrorWithMapAndXError(constant.ErrUnauthorizedTokenError, nil))
		}

		c.Set("claims", claims)
		c.Set("uid", claims.UID)
		c.Set("username", claims.Username)
		c.Set("authMethod", constant.AuthMethodJWT)
		c.Set("user", user)
		c.Next()

	}
}
