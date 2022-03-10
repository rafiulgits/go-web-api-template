package server

import (
	"webapi/configs"

	"github.com/rafiulgits/gotnet"
)

func Start() {
	builder := gotnet.NewAppBuilder()

	injectDependencies(builder.Service)

	app := builder.Build()

	ConfigureMiddlewares(app)

	//routes must set after middlewares
	MapAppRoutes(app)

	app.BeforeRun(func() {
		BeforeAppRun(app)
	})

	app.Run(&gotnet.AppRunConfig{
		Port: configs.GetAppConfig().ListenPort,
	})
}
