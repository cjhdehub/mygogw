package main

import (
	"flag"
	"gogw-server/config"
	"gogw-server/logger"
	"gogw-server/server"
)

var cfgFile = flag.String("c", "cfg.json", "config file")
var logLevel = flag.String("l", "info", "log level: info/debug")

func main() {
	logger.LEVEL = logger.INFO

	logger.Info("gogw start")
	flag.Parse()

	cfg, err := config.NewConfigFromFile(*cfgFile)
	if err != nil {
		logger.Error(err)
		return
	}

	if *logLevel == "debug" {
		logger.LEVEL = logger.DEBUG
	}

	server := server.NewServer(cfg.Server.ServerAddr, cfg.Server.TimeoutSecond)
	server.Start()
}
