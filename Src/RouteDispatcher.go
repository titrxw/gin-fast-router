package Src

import (
	"github.com/titrxw/gin-router/Src/Exception"
)

type RouteDispatcher struct {
	RouteCollector *RouteCollector
}

func (dispatcher *RouteDispatcher) Dispatcher(method string, path string) *Route {
	_, support := dispatcher.RouteCollector.RouteMethodMap[method]
	if !support {
		panic(Exception.RouteNotFoundException{
			404,
			"route " + path + " not found",
		})
	}

	_, hasRoute := dispatcher.RouteCollector.RouteMethodMap[method][path]
	if !hasRoute {
		panic(Exception.RouteNotFoundException{
			404,
			"route " + path + " not found",
		})
	}

	return dispatcher.RouteCollector.RouteMethodMap[method][path]
}
