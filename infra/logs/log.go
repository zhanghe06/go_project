package logs

import (
	"fmt"
	"log"
	"os"
	"path"
	"sap_cert_mgt/infra/config"
	"sap_cert_mgt/infra/utils"
	"sync"

	"github.com/sirupsen/logrus"
)

// Logger 服务日志服务，可适配其他日志组件
type Logger interface {
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})
	Warnf(format string, args ...interface{})
	Warnln(args ...interface{})
	Tracef(format string, args ...interface{})
	Traceln(args ...interface{})
	Panicf(format string, args ...interface{})
	Panicln(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})
}

var (
	logOnce sync.Once
	l       Logger
	logDir  string
	logName string
)

// NewLogger 获取日志句柄
func NewLogger() Logger {
	initLogger()

	return l
}

type serverLog struct {
	logger *logrus.Logger
}

// Infof 普通信息
func (l *serverLog) Infof(format string, args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Infof(format, args...)
}

// Infoln 普通信息
func (l *serverLog) Infoln(args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Infoln(args...)
}

// Warnf 警告信息
func (l *serverLog) Warnf(format string, args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Warnf(format, args...)
}

// Warnln 警告信息
func (l *serverLog) Warnln(args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Warnln(args...)
}

// Errorf 错误信息
func (l *serverLog) Errorf(format string, args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Errorf(format, args...)
}

// Errorln 错误信息
func (l *serverLog) Errorln(args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Errorln(args...)
}

// Debugf 调试信息
func (l *serverLog) Debugf(format string, args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Debugf(format, args...)
}

// Debugln 调试信息
func (l *serverLog) Debugln(args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Debugln(args...)
}

// Tracef 跟踪信息
func (l *serverLog) Tracef(format string, args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Tracef(format, args...)
}

// Traceln 跟踪信息
func (l *serverLog) Traceln(args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Traceln(args...)
}

// Fatalf 致命错误
func (l *serverLog) Fatalf(format string, args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Fatalf(format, args...)
}

// Fatalln 致命错误
func (l *serverLog) Fatalln(args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Fatalln(args...)
}

// Panicf 恐慌错误
func (l *serverLog) Panicf(format string, args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Panicf(format, args...)
}

// Panicln 恐慌错误
func (l *serverLog) Panicln(args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Panicln(args...)
}

func initLogger() {
	logOnce.Do(func() {
		conf := config.NewConfig()
		log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
		logHandle := &serverLog{}
		logHandle.logger = logrus.New()
		logHandle.logger.SetFormatter(&logrus.JSONFormatter{})
		logout := os.Getenv("LOGOUT")
		if len(logout) > 0 {
			logHandle.logger.SetOutput(os.Stdout)
		} else {
			logDir = conf.Log.LogDir
			logName = conf.Log.LogName
			if logDir == "" {
				logDir = utils.GetProjectAbPathByCaller()
			} else {
				err := os.MkdirAll(logDir, 0750)
				if err != nil {
					fmt.Println("mkdir err", err)
				}
			}
			if logName == "" {
				logName = "server.log"
			}
			logFileName := path.Join(logDir, logName)
			logFile, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0)
			if err != nil {
				fmt.Println("open log file err", err)
			}
			logHandle.logger.SetOutput(logFile)
			l = logHandle
		}
	})
}
