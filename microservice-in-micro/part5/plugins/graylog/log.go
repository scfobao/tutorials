package graylog

import (
	"sync"

	"github.com/micro-in-cn/tutorials/microservice-in-micro/part5/basic"
	"github.com/micro/go-micro/util/log"
)

var (
	inited bool
	myLog  *Graylog
	m      sync.RWMutex
)

func init() {
	basic.Register(initLOG)
}

func initLOG() {
	m.Lock()
	defer m.Unlock()

	if inited {
		log.Logf("[initLog] 日志插件已经初始化")
		return
	}

	initGraylog()

	inited = true

}

func GetLog(appinfo interface{}) *Graylog {
	if appinfo != nil {
		myLog.Appinfo = appinfo
	}

	return myLog
}
