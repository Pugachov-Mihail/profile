package main

import (
	"profile/internal/app"
	"profile/internal/config"
)

func main() {
	cfg := config.MustLoad()
	log := config.SetupLoger(cfg.Env)
	log.Debug("Init Profile grpc")

	application := app.New(log, cfg.Grpc.Port)
	application.GRPC.MustRun()

}
