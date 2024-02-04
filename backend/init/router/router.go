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
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"nunu/backend/i18n"
	implApi "nunu/backend/init/api/v1"
	v1 "nunu/backend/router/v1"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.HandleMethodNotAllowed = true
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// 添加全局中间件
	// 跨域配置
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	router.Use(cors.New(corsConfig))

	// 国际化中间件
	router.Use(i18n.GinI18nLocalize())

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
		systemRouter.InitAdminRouter(PrivateGroup, implApi.ImplAdminApi)
	}

	return router

}
