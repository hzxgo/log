package log

import "github.com/sirupsen/logrus"

const (
	DEFAULT_LOG_MAX_ROLLS = 7             // 默认 log 保留最大天数
	DEFAULT_LOG_PATH      = "./log"       // 默认 log 路径
	DEFAULT_LOG_FILENAME  = "default.log" // 默认 log 文件名称
)

const (
	PANIC_LEVEL = logrus.PanicLevel
	FATAL_LEVEL = logrus.FatalLevel
	ERROR_LEVEL = logrus.ErrorLevel
	WARN_LEVEL  = logrus.WarnLevel
	INFO_LEVEL  = logrus.InfoLevel
	DEBUG_LEVEL = logrus.DebugLevel
	TRACE_LEVEL = logrus.TraceLevel
)
