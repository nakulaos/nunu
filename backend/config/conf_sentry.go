/**
 ******************************************************************************
 * @file           : conf_sentry.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/30
 ******************************************************************************
 */

package config

type SentryConf struct {
	Dsn string
	// In debug mode, the debug information is printed to stdout to help you
	// understand what Sentry is doing.
	Debug bool
	// Configures whether SDK should generate and attach stack traces to pure
	// capture message calls.
	AttachStacktrace bool
	// The sample rate for sampling traces in the range [0.0, 1.0].
	TracesSampleRate float64
	AttachLogrus     bool
	AttachGin        bool
	Enable           bool
}

func (s *SentryConf) UseSentryGin() bool {
	return s.Enable && s.AttachGin
}
