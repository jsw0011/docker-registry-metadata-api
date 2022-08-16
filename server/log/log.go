package log

import (
	easyinitgolog "github.com/jsw0011/easy-init-go-log"
	l "github.com/op/go-logging"
)

var Logger *l.Logger

func InitLogger(logFile *string, logLevel string) {

	appname := "docker_image_metadata_api"
	easyinitgolog.InitLogger(logFile, &logLevel, appname)
	Logger = easyinitgolog.GetLoggerByName(appname)
}
