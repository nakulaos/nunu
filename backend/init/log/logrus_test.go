/**
 ******************************************************************************
 * @file           : logrus_test.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/26
 ******************************************************************************
 */

package log

import (
	"nunu/backend/init/gorm"
	"testing"
)

func TestInitLogger(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitLogger()
			_ = gorm.GetGormDB()
		})
	}
}
