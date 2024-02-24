/**
 ******************************************************************************
 * @file           : jinzhu.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/1
 ******************************************************************************
 */

package jinzhu

import "nunu/backend/init/conf"

func init() {
	// 顺便初始化表名
	sysConfig := conf.GetConfig()
	InitTableName(sysConfig.DatabaseConf.TablePrefix)
}
