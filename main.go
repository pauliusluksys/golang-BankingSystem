package main

import (
	"bankingV2/app"
	"bankingV2/logger"
)

func main() {
	logger.Info("starting app")
	app.Start()
}
