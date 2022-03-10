package handlers

import (
	"net/http"
	"webapi/services"

	"github.com/rafiulgits/gotnet"
)

type IDemoHandler interface {
	gotnet.IHandler
}

type DemoHandler struct {
	demoService services.IDemoService
}

func NewDemoHandler(demoService services.IDemoService) IDemoHandler {
	return &DemoHandler{
		demoService: demoService,
	}
}

func (handler *DemoHandler) Handle() http.Handler {
	router := gotnet.NewRouter()
	router.Post("/", handler.createHandler)
	router.Get("/", handler.getAllDemos)
	return router
}

func (handler *DemoHandler) createHandler(w http.ResponseWriter, r *http.Request) {
	createdDemo, err := handler.demoService.CreateDemo()
	if err != nil {
		gotnet.BadRequest(w, err)
		return
	}
	gotnet.Ok(w, createdDemo)
}

func (handler *DemoHandler) getAllDemos(w http.ResponseWriter, r *http.Request) {
	demos, err := handler.demoService.GetAllDemos()
	if err != nil {
		gotnet.NotFound(w, err)
		return
	}
	gotnet.Ok(w, demos)
}
