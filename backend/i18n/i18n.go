/**
 ******************************************************************************
 * @file           : i18n.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/24
 ******************************************************************************
 */

package i18n

import (
	"embed"
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
	"strings"
)

//go:embed *.yaml
var fs embed.FS

func GetMsgWithMap(key string, maps map[string]interface{}) string {
	content := ""
	if maps == nil {
		content = ginI18n.MustGetMessage(&i18n.LocalizeConfig{
			MessageID: key,
		})
	} else {
		content = ginI18n.MustGetMessage(&i18n.LocalizeConfig{
			MessageID:    key,
			TemplateData: maps,
		})
	}

	content = strings.ReplaceAll(content, ": <no value>", "")
	if content == "" {
		return key
	} else {
		return content
	}
}

func GetErrMsg(key string, maps map[string]interface{}) string {
	content := ""
	if maps == nil {
		content = ginI18n.MustGetMessage(&i18n.LocalizeConfig{
			MessageID: key,
		})
	} else {
		content = ginI18n.MustGetMessage(&i18n.LocalizeConfig{
			MessageID:    key,
			TemplateData: maps,
		})
	}
	return content
}

func GetMsgByKey(key string) string {
	content := ginI18n.MustGetMessage(&i18n.LocalizeConfig{
		MessageID: key,
	})
	return content
}

// GinI18nLocalize 应在初始化router时进行全局注册
func GinI18nLocalize() gin.HandlerFunc {
	return ginI18n.Localize(
		ginI18n.WithBundle(&ginI18n.BundleCfg{
			RootPath:         ".",
			AcceptLanguage:   []language.Tag{language.Chinese, language.English, language.TraditionalChinese, language.German, language.French}, //nolint:lll
			DefaultLanguage:  language.Chinese,
			FormatBundleFile: "yaml",
			UnmarshalFunc:    yaml.Unmarshal,
			Loader:           &ginI18n.EmbedLoader{FS: fs},
		}),
		ginI18n.WithGetLngHandle(
			func(context *gin.Context, defaultLng string) string {
				lng := context.GetHeader("Accept-Language")
				if lng == "" {
					return defaultLng
				}
				return lng
			},
		))
}
