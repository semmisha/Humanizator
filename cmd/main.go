package main

import (
	"Humanizator/config"
	"Humanizator/http/httpRouter"
	"Humanizator/logging"
)

// TODO ----- Add additional checking file with vrd_ and file without
const (
	envPath = "data/env.conf"
)

var (
	logger = logging.Logger()
)

func main() {
	EnvMap := config.GetEnv(envPath, logger)
	router := httpRouter.NewRouterVar("", EnvMap, logger)
	router.Setup()

}
