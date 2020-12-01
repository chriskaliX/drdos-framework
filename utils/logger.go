package utils

import (
	"fmt"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

func Initlog() {
	path := Dir + "/data/logs/drdos.log"
	writer, err := rotatelogs.New(
		path + ".%Y%m%d",
	)
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(writer)
	log.SetLevel(log.InfoLevel)
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)
}
