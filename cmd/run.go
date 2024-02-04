/**
 ******************************************************************************
 * @file           : run.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/24
 ******************************************************************************
 */

package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"nunu/backend/init/conf"
	"nunu/backend/init/log"
	"nunu/backend/init/router"
)

func commandRunHandle(_cmd *cobra.Command, _args []string) {
	webServerConf := conf.GetConfig().WebServerConf
	log.InitLogger()
	router := router.InitRouter()
	err := router.Run(webServerConf.HttpIp + ":" + webServerConf.HttpPort)
	if err != nil {
		logrus.Errorln(err)
	}
}

func initRunCommand() {
	_commandRun := &cobra.Command{
		Use:   "run",
		Short: "run nunu backend",
		Long:  "run nunu backend",
		Run:   commandRunHandle,
	}

	rootCmd.AddCommand(_commandRun)
}
