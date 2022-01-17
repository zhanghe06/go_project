package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_project/adapter/driver"
	"go_project/adapter/driver/cert"
	"go_project/adapter/driver/event"
	"go_project/adapter/driver/health"
	"go_project/adapter/driver/notice_conf"
	"go_project/adapter/driver/notice_strategy"
	"go_project/adapter/driver/operation_log"
	"go_project/infra/config"
	"go_project/infra/logs"
	"go_project/infra/middleware"
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
		engine.Use(
			gin.Logger(),
			middleware.RecoveryMiddleware(),
			//middleware.RequestIdMiddleware(),
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
			conf.Service.Port,
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
