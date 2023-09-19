package main

import (
	"context"
	"gitlab.com/merakilab9/meracore/logger"
	"gitlab.com/merakilab9/meracrawler/kayle/conf"
	"gitlab.com/merakilab9/meracrawler/kayle/pkg/route"
	"gitlab.com/merakilab9/meracrawler/kayle/pkg/utils"

	"os"
)

const (
	APPNAME = "Kayle"
)

func main() {
	conf.SetEnv()
	logger.Init(APPNAME)
	utils.LoadMessageError()

	app := route.NewService()
	ctx := context.Background()
	err := app.Start(ctx)
	if err != nil {
		logger.Tag("main").Error(err)
	}
	os.Clearenv()
}
