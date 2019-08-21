package main

import (
	"boilerplate/pkg/api"
	"boilerplate/pkg/config"
	"boilerplate/pkg/logger"
)

var commit string

func main() {
	config.Init(commit)
	logger.Init(config.Env.Log.Level)
	logger.Info("version: %v", commit)
	logger.Debug("config: %+v", config.Env)

	api.StartServer()
}
