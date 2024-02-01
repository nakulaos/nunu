/**
 ******************************************************************************
 * @file           : cmd.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/1
 ******************************************************************************
 */

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// root.go 定义 根命令

var (
	rootCmd = &cobra.Command{
		Use:   "nunu",
		Short: "a community like 'weibo'",
		Long:  "a community like 'weibo'",
	}
)

// SetupRootCmd set root command name,short-describe, long-describe
// return &cobra.Command to custom other options
// SetupRootCmd 函数 设置root命令名和描述
func SetupRootCmd(use, short, long string) {
	rootCmd.Use = use
	rootCmd.Short = short
	rootCmd.Long = long
}

// RegisterCmd add sub-command
// RegisterCmd 添加子命令
func RegisterCmd(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

// Execute start application
// Execute 启动应用
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	initMigrationCommand()
}
