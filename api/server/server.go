package server

import (
	"webapi/api/handlers"
	"webapi/configs"
	"webapi/db"
	"webapi/logger"
	"webapi/repositories"
	"webapi/services"

	"github.com/go-chi/chi/middleware"
	"github.com/rafiulgits/gotnet"
)

func Start() {
	builder := gotnet.NewAppBuilder()
	injectDependencies(builder.Service)

	app := builder.Build()

	// middlewares
	configureLogger(app)

	// routing
	mapHandler(app)

	// migrations and background services
	app.BeforeRun(func() {
		dbMigration(app.Service.Container())
	})

	// run applications
	app.Run(&gotnet.AppRunConfig{
		Port: configs.LoadConfig().ListenPort,
	})
}

func configureLogger(app *gotnet.App) {
	app.Service.Container().Invoke(func(cfg *configs.AppConfig) {
		logger.NewLogger(&cfg.LogConfig)
		if cfg.DebugEnv {
			app.Router.Use(middleware.Logger)
		} else {
			app.Router.Use(logger.ZapFileLogging(logger.Log))
		}
	})
}

func dbMigration(container *gotnet.Container) {
	container.Invoke(func(database *db.DB, cfg *configs.AppConfig) {
		if cfg.DBMigration {
			database.Migration()
		}
	})
}

func injectDependencies(dependency *gotnet.Service) {
	dependency.AddSingleton(configs.LoadConfig)
	dependency.AddSingleton(services.NewDemoService)
	dependency.AddSingleton(repositories.NewDemoRepository)
	dependency.AddSingleton(db.OpenConnection)
}

func mapHandler(app *gotnet.App) {
	app.RegisterHandler(handlers.NewDemoHandler, func(handler handlers.IDemoHandler) {
		app.Router.Mount("/demos", handler.Handle())
	})
}
