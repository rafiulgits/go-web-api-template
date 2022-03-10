package server

import (
	"webapi/api/handlers"

	"github.com/rafiulgits/gotnet"
)

func MapAppRoutes(app *gotnet.App) {
	app.RegisterHandler(handlers.NewDemoHandler, func(handler handlers.IDemoHandler) {
		app.Router.Mount("/demos", handler.Handle())
	})
}
