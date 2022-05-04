package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"sap_cert_mgt/adapter/driver"
	"sap_cert_mgt/adapter/driver/cert"
	"sap_cert_mgt/adapter/driver/event"
	"sap_cert_mgt/adapter/driver/health"
	"sap_cert_mgt/adapter/driver/notice_conf"
	"sap_cert_mgt/adapter/driver/notice_strategy"
	"sap_cert_mgt/adapter/driver/operation_log"
	"sap_cert_mgt/infra/config"
	"sap_cert_mgt/infra/logs"
	"sap_cert_mgt/infra/middleware"
	"time"
)

type server struct {
	// Rest API
	healthRestHandler         driver.RESTHandlerInterface
	certRestHandler           driver.RESTHandlerInterface
	noticeConfRestHandler     driver.RESTHandlerInterface
	noticeStrategyRestHandler driver.RESTHandlerInterface
	operationLogRestHandler   driver.RESTHandlerInterface
	// 邮件通知
	emailNotice event.NoticeInterface
	// TODO
}

func (s *server) Start() {
	log := logs.NewLogger()
	log.Infoln("start server.....")
	conf := config.NewConfig()
	go func() {
		engine := gin.New()
		//engine.Use(gin.Recovery())

		//formatter := func(p gin.LogFormatterParams) string {
		//	return fmt.Sprintf("[logger] %s %s %s %d %s\n",
		//		p.TimeStamp.Format("2006-01-02_15:04:05"),
		//		p.Path,
		//		p.Method,
		//		p.StatusCode,
		//		p.ClientIP,
		//	)
		//}
		logConf := gin.LoggerConfig{
			SkipPaths: []string{"/heartbeat"},
			Output:    os.Stderr,
			//Formatter: formatter,
		}
		engine.Use(
			//gin.Logger(),
			gin.LoggerWithConfig(logConf),
			middleware.RecoveryMiddleware(),
			//middleware.RequestIdMiddleware(),
			middleware.CORSMiddleware(), // 跨域中间件
		)
		engine.UseRawPath = true

		// 注册API
		s.healthRestHandler.RegisterAPI(engine)
		s.certRestHandler.RegisterAPI(engine)
		s.noticeConfRestHandler.RegisterAPI(engine)
		s.noticeStrategyRestHandler.RegisterAPI(engine)
		s.operationLogRestHandler.RegisterAPI(engine)

		url := fmt.Sprintf(
			"%s:%s",
			conf.Service.Host,
			conf.Service.PublicPort,
		)
		if err := engine.Run(url); err != nil {
			log.Errorln(err)
		}
		log.Infof("api server run in %s", url)
	}()

	// 过期提醒（扫描频次：每天一次）
	go func() {
		d := time.Duration(time.Hour * 24)
		t := time.NewTicker(d)
		defer t.Stop()

		for {
			<-t.C
			// 检查临期证书
			log.Infoln("Scan cert timeout, start")
			s.emailNotice.Scan()
			log.Infoln("Scan cert timeout, end")
			// 临期邮件通知
			log.Infoln("Send cert timeout, start")
			s.emailNotice.Send()
			log.Infoln("Send cert timeout, end")
		}
	}()

}

func main() {
	s := &server{
		healthRestHandler:         health.NewRESTHandler(),
		certRestHandler:           cert.NewRESTHandler(),
		noticeConfRestHandler:     noticeConf.NewRESTHandler(),
		noticeStrategyRestHandler: noticeStrategy.NewRESTHandler(),
		operationLogRestHandler:   operationLog.NewRESTHandler(),
		emailNotice:               event.NewEmailNotice(),
	}
	s.Start()

	select {}
}
