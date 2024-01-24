/**
 ******************************************************************************
 * @file           : config_test.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/24
 ******************************************************************************
 */

package config

import (
	"fmt"
	viper2 "github.com/spf13/viper"
	"reflect"
	"testing"
)

func Test_initConf(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{},
	}
	fmt.Println(NewConfig())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initConf()
		})
	}
}

func TestGetViper(t *testing.T) {
	tests := []struct {
		name string
		want *viper2.Viper
	}{
		// TODO: Add test cases.
		{},
	}
	fmt.Println(GetViper().Get("MySQL"))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetViper(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetViper() = %v, want %v", got, tt.want)
			}
		})
	}
}
