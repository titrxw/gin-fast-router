package Src

import "github.com/titrxw/gin-router/Src/Exception"

type RouteCollector struct {
	RouteMethodMap map[string]RouteMap
}

func (collector *RouteCollector) collect(route *Route) {
	_, has := collector.RouteMethodMap[route.method]
	if !has {
		collector.RouteMethodMap[route.method] = make(RouteMap, 0)
	}

	_, hasRegister := collector.RouteMethodMap[route.method][route.path]
	if hasRegister {
		panic(Exception.RouteHasRegisterException{
			500,
			"route " + route.path + ", method " + route.method + " has register",
		})
	}

	collector.RouteMethodMap[route.method][route.path] = route
}
