/**
 ******************************************************************************
 * @file           : sentry.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/26
 ******************************************************************************
 */

package sentry

import (
	"github.com/getsentry/sentry-go"
	sentrylogrus "github.com/getsentry/sentry-go/logrus"
	"github.com/sirupsen/logrus"
	"nunu/backend/init/conf"
	"nunu/backend/pkg/version"
	"time"
)

func setupSentryLogrus(opts sentry.ClientOptions) {
	// Send only ERROR and higher level logs to Sentry
	sentryLevels := []logrus.Level{logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}
	sentryHook, err := sentrylogrus.New(sentryLevels, opts)
	if err != nil {
		panic(err)
	}
	logrus.AddHook(sentryHook)

	// Flushes before calling os.Exit(1) when using logger.Fatal
	// (else all defers are not called, and Sentry does not have time to send the event)
	logrus.RegisterExitHandler(func() { sentryHook.Flush(5 * time.Second) })
}

func InitSentry() {
	sysConf := conf.GetConfig()

	if sysConf.SentryConf.Enable {
		opts := sentry.ClientOptions{
			Dsn:              sysConf.SentryConf.Dsn,
			Debug:            sysConf.SentryConf.Debug,
			AttachStacktrace: sysConf.SentryConf.AttachStacktrace,
			TracesSampleRate: sysConf.SentryConf.TracesSampleRate,
		}

		err := sentry.Init(opts)
		if err != nil {
			logrus.Errorln("sentry init failed!")
		}

		if sysConf.SentryConf.AttachLogrus {
			setupSentryLogrus(opts)
		}

		sentry.WithScope(func(scope *sentry.Scope) {
			scope.SetExtras(map[string]interface{}{
				"version": version.GetVersionInfo(),
				"time":    time.Now().Local(),
			})

			sentry.CaptureMessage("nunu sentry works!")
			logrus.Debugln("nunu sentry works!")
		})

	}
}
