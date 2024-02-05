/**
 ******************************************************************************
 * @file           : wire_injector.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/5
 ******************************************************************************
 */

package jinzhu

import "github.com/google/wire"

var Set = wire.NewSet(
	NewUserManageRepo,
)
