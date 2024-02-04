/**
 ******************************************************************************
 * @file           : redis_test.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/31
 ******************************************************************************
 */

package redis

import "testing"

func Test_initRedisClient(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initRedisClient()
		})
	}
}
