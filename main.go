package main

import (
	"github.com/AliSayyah/banking/app"
	"github.com/AliSayyah/banking/logger"
)

func main() {
	logger.Info("Starting the application")
	app.Start()
}
