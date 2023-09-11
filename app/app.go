package app

import (
	"fmt"
	"lido-core/v1/platform/cache"
	"log"
	"os"

	"lido-core/v1/pkg/configs"
	"lido-core/v1/pkg/utils"
	"lido-core/v1/pkg/workers"
	"lido-core/v1/platform/database"

	"github.com/gofiber/fiber/v2"
)

type _App struct {
	engine *fiber.App
}

type IApp interface {
	Worker(*configs.WorkerConfig)
	Middleware(...middleware) IApp
	Route(...route) IApp
	Run()
	Shutdown(<-chan os.Signal)
	BackgroundTask(...backgroundTask)
}

type middleware func(*fiber.App)
type route func(*fiber.App)
type backgroundTask func()

func New() IApp {
	config := configs.FiberConfig()
	return &_App{
		engine: fiber.New(config),
	}
}

func (app *_App) BackgroundTask(tasks ...backgroundTask) {
	for _, task := range tasks {
		go task()
	}
}

func (app *_App) Worker(wcf *configs.WorkerConfig) {
	workers.WorkerConfig = wcf
}

func (app *_App) Middleware(middlewares ...middleware) IApp {
	for _, middleware := range middlewares {
		middleware(app.engine)
	}
	return app
}

func (app *_App) Route(routes ...route) IApp {
	for _, route := range routes {
		route(app.engine)
	}
	return app
}

func (app *_App) Shutdown(sig <-chan os.Signal) {
	go func() {
		<-sig
		fmt.Println()
		database.Shutdown()
		cache.Shutdown()
		if configs.StageStatus == "prod" {
			log.Println("[SERVER] Server is shutting down ..")
			if err := app.engine.Shutdown(); err != nil {
				log.Printf("Oops... Server is not shutting down! Reason: %v", err)
			}
		} else {
			os.Exit(0)
		}
	}()
}

func (app *_App) Run() {
	utils.StartServer(app.engine)
}
