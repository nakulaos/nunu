/**
 ******************************************************************************
 * @file           : router.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/2
 ******************************************************************************
 */

package router

import (
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"nunu/backend/i18n"
	"nunu/backend/init/conf"
	"nunu/backend/init/validation"
	v1 "nunu/backend/router/v1"
	"nunu/script/wire"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.HandleMethodNotAllowed = true

	// 初始化校验器
	validation.InitValidation(router)

	// 添加全局中间件
	// 跨域配置
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	router.Use(cors.New(corsConfig))

	// 国际化中间件
	router.Use(i18n.GinI18nLocalize())

	senryConf := conf.GetConfig().SentryConf
	if senryConf.UseSentryGin() {
		router.Use(sentrygin.New(sentrygin.Options{
			// repanic配置Sentry恢复后是否应该重新panic，在大多数情况下应该设置为true，
			// as gin。默认包括它自己的恢复中间件处理http响应。
			Repanic: true,
		}))
	}

	// 默认404
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "Not Found",
		})
	})

	// 默认405
	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code": 405,
			"msg":  "Method Not Allowed",
		})
	})

	var systemRouter = new(v1.RouterGroup)

	PrivateGroup := router.Group("/api/v1")
	{
		systemRouter.InitAdminRouter(PrivateGroup, wire.CreateAdminApi())
		systemRouter.InitPubRouter(PrivateGroup, wire.CreatePubApi())
	}

	return router

}
