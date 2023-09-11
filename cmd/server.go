package cmd

import (
	"lido-core/v1/app"
	"lido-core/v1/pkg/configs"
	"lido-core/v1/pkg/middleware"
	"lido-core/v1/pkg/routes"
	"os"
	"os/signal"
	"syscall"
)

func StartServer() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// create a new app
	_app := app.New()
	_app.Shutdown(sigChan)

	// register tasks
	_app.Worker(Setting(configs.BrokerUrl, configs.ResultBackend))

	// register middleware
	_app.Middleware(
		middleware.FiberMiddleware,
	)

	// register route
	_app.Route(
		routes.HealthCheck,
		routes.UserRoute,
		routes.QuizRoute,
		routes.VideoRoute,
		routes.LiveRoute,
		routes.WalletRoute,
		routes.AdvertisementRoute,
		routes.NotFoundRoute,
	)
	_app.Run()
}
