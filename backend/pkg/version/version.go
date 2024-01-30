/**
 ******************************************************************************
 * @file           : version.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/30
 ******************************************************************************
 */

package version

import "fmt"

var (
	version   = "unknown"
	commitID  = "unknown"
	buildDate = "unknown"
	buildTags = "unknown"
)

func GetVersionInfo() string {
	return fmt.Sprintf("nunu %s (build:%s %s)", version, commitID, buildDate)
}
