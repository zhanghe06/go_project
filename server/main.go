package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_project/adapter/driver"
	"go_project/adapter/driver/health"
	"go_project/adapter/driver/user"
	"go_project/infra/config"
	"go_project/infra/logs"
	"go_project/infra/middleware"
	//"log"
)

type server struct {
	healthRestHandler driver.RESTHandlerInterface
	userRestHandler   driver.RESTHandlerInterface
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
			//gin.Logger(),
			middleware.RecoveryMiddleware(),
			//middleware.RequestIdMiddleware(),
		)
		engine.UseRawPath = true

		// 注册API
		s.healthRestHandler.RegisterAPI(engine)
		s.userRestHandler.RegisterAPI(engine)

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
}

func main() {
	s := &server{
		healthRestHandler: health.NewRESTHandler(),
		userRestHandler:   user.NewRESTHandler(),
	}
	s.Start()

	select {}
}
