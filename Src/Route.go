package Src

type RouteInterface interface {
	GetName() string
	GetPath() string
	GetMethod() string
	GetHandler() RouteHandler
	GetMiddlewares() RouteMiddlewares
	GetAttributes() RouteAttributes
}

type Route struct {
	name        string
	path        string
	method      string
	handler     RouteHandler
	middlewares RouteMiddlewares
	attributes  RouteAttributes
}

var _ RouteInterface = &Route{}

func (route *Route) GetName() string {
	return route.name
}

func (route Route) GetMethod() string {
	return route.method
}

func (route Route) GetPath() string {
	return route.path
}

func (route Route) GetHandler() RouteHandler {
	return route.handler
}

func (route Route) GetMiddlewares() RouteMiddlewares {
	return route.middlewares
}

func (route Route) GetAttributes() RouteAttributes {
	return route.attributes
}
