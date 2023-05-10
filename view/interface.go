package view

import "github.com/gorilla/mux"

type HttpModule interface {
	Setup(router *mux.Router)
}

type HttpModulew interface {
	Setup(router *mux.Router)
}
type HttpModuleww interface {
	Setup(router *mux.Router)
}
