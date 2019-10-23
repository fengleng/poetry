/*
@Time : 2019/9/17 14:52
@Author : zxr
@File : logger
@Software: GoLand
*/
package bootstrap

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
	"io"
	"os"
	"poetry/config/define"
	"time"
)

func InitLogger() {
	formatter := &prefixed.TextFormatter{}
	log.SetFormatter(formatter)
	log.SetLevel(log.InfoLevel)
	log.SetReportCaller(true)
	point := getLogWriter()
	if point != nil {
		log.SetOutput(point)
	}
	formatter.ForceFormatting = true
	formatter.DisableColors = false
	formatter.ForceColors = true
	formatter.FullTimestamp = true
	formatter.TimestampFormat = "2006-01-02 15:04:05 000000"
}

//logrus记录到文件中
func getLogWriter() (point io.Writer) {
	logFile := fmt.Sprintf(define.BaseDir+"/logs/poetry-%d-%d-%d.log", time.Now().Year(), time.Now().Month(), time.Now().Day())
	if file, e := os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666); e == nil {
		return file
	}
	return nil
}
