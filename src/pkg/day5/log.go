package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()
func init() {
	//log 格式化方式
	log.Formatter = &logrus.JSONFormatter{}
	f,_ := os.Create("/Users/abcd/go_work/go_study/src/pkg/day5/log/logerr.log")
	log.Out = f
	log.Level = logrus.InfoLevel
}
func main() {
	log.WithFields(logrus.Fields{"name":"zzq"}).Info("这是一条log")
	log.WithFields(logrus.Fields{"name":"zzq"}).Error("这是一条ERROR log")
}
