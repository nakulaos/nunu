/**
 ******************************************************************************
 * @file           : i18n_test.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/2
 ******************************************************************************
 */

package i18n

import (
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"log"
	"net/http"
	"testing"
)

type BaseRep struct {
	Code int
	Msg  string
	Data any
}

func TestGinI18nLocalize(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(GinI18nLocalize())
	router.GET("/city", func(c *gin.Context) {
		s, _ := ginI18n.GetMessage(c, &i18n.LocalizeConfig{MessageID: "City"})
		c.JSON(http.StatusOK, &BaseRep{
			Code: 0,
			Msg:  "",
			Data: s,
		})
	})
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
