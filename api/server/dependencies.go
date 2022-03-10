package server

import (
	"webapi/configs"
	"webapi/db"
	"webapi/repositories"
	"webapi/services"

	"github.com/rafiulgits/gotnet"
)

func injectDependencies(dependency *gotnet.Service) {
	dependency.AddSingleton(configs.LoadConfig)
	dependency.AddSingleton(services.NewDemoService)
	dependency.AddSingleton(repositories.NewDemoRepository)
	dependency.AddSingleton(db.OpenConnection)
}
