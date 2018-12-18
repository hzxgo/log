package log

import (
	"os"
	"path/filepath"
	"time"

	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var (
	log *logrus.Logger
)

func init() {
	log = logrus.New()

	// set log formatter
	formatter := new(logrus.TextFormatter)
	formatter.TimestampFormat = "2006-01-02 15:04:05"
	formatter.FieldMap = logrus.FieldMap{
		logrus.FieldKeyTime:  "[T]",
		logrus.FieldKeyLevel: "[L]",
		logrus.FieldKeyMsg:   "[Msg]",
	}
	log.SetFormatter(formatter)
	log.AddHook(NewLineHook(true))
	log.SetLevel(logrus.DebugLevel)

	// run as debug?
	var debug bool
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			if arg == "-debug" {
				debug = true
				break
			}
		}
	}
	if !debug {
		logFilename := getLogFilename()
		rotateLog, err := rotatelogs.New(
			getLogFilename()+".%Y%m%d",
			rotatelogs.WithLinkName(logFilename),
			rotatelogs.WithMaxAge(24*time.Duration(DEFAULT_LOG_MAX_ROLLS)*time.Hour),
			rotatelogs.WithRotationTime(24*time.Hour),
		)
		if err != nil {
			log.Errorf("new rotate log failed | %v", err)
		} else {
			log.Out = rotateLog
		}
	}
}

func Init(isDebug bool, maxRolls uint32, filename string) {

	if isDebug {
		log.SetLevel(DEBUG_LEVEL)
	} else {
		log.SetLevel(INFO_LEVEL)
	}

	if len(filename) > 0 {
		logFilename := getLogFilename(filename)
		rotateLog, _ := rotatelogs.New(
			logFilename+".%Y%m%d",
			rotatelogs.WithLinkName(logFilename),
			rotatelogs.WithMaxAge(24*time.Duration(maxRolls)*time.Hour),
			rotatelogs.WithRotationTime(24*time.Hour),
		)
		log.Out = rotateLog
	}
}

func SetLevel(level logrus.Level) {
	log.SetLevel(level)
}

// ---------------------------------------------------------------------------------------------------------------------

func isDirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}
}

func getLogFilename(filename ...string) string {
	var err error
	var logFilename string

	if len(filename) == 0 {
		logFilename = DEFAULT_LOG_PATH + "/" + DEFAULT_LOG_FILENAME
	} else {
		logFilename = filename[0]
	}

	if !filepath.IsAbs(logFilename) {
		logFilename, err = filepath.Abs(logFilename)
		if err != nil {
			log.Errorf("get log filename abs failed | %v", err)
		}
	}

	if path := filepath.Dir(logFilename); !isDirExists(path) {
		if err := os.MkdirAll(path, 0744); err != nil {
			log.Errorf("make dir all failed | %v", path, err)
		}
	}

	return logFilename
}
