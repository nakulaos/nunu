/**
 ******************************************************************************
 * @file           : sentry_test.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/30
 ******************************************************************************
 */

package sentry

import (
	"nunu/backend/init/log"
	"testing"
)

func TestInitSentry(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.InitLogger()
			InitSentry()
		})
	}
}
