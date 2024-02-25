/**
 ******************************************************************************
 * @file           : wire_injector.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/18
 ******************************************************************************
 */

package cache

import "github.com/google/wire"

var Set = wire.NewSet(
	NewRedisCacheRepo,
)
