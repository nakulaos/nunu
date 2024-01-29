/**
 ******************************************************************************
 * @file           : sqlx_test.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/29
 ******************************************************************************
 */

package sqlx

import "testing"

func Test_initSqlxDB(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initSqlxDB()
		})
	}
}
