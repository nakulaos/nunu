/**
 ******************************************************************************
 * @file           : entry.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/17
 ******************************************************************************
 */

package validation

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"nunu/backend/pkg/validation"
)

func InitValidation(e *gin.Engine) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("password", validation.PaswordValidation)
	}
}
