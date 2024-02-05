/**
 ******************************************************************************
 * @file           : wire_injector.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/5
 ******************************************************************************
 */

package service

import "github.com/google/wire"

var Set = wire.NewSet(
	NewAdminService,
)
