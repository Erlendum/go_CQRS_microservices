package main

import (
	"flag"
	"github.com/Erlendum/go_CQRS_microservices/api_gateway_service/config"
	"github.com/Erlendum/go_CQRS_microservices/pkg/logger"
	"log"
)

func main() {
	flag.Parse()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	appLogger := logger.NewAppLogger(cfg.Logger)
	appLogger.InitLogger()
	appLogger.WithName("ApiGateway")

	//s := server.NewServer(appLogger, cfg)
	//appLogger.Fatal(s.Run())
}
