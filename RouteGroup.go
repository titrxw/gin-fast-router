package gin_fast_router

type RouteGroup struct {
	basePath       string
	middlewares    RouteMiddlewares
	attributes     RouteAttributes
	routeCollector RouteCollector
}

func (group *RouteGroup) Group(basePath string, handler func(group *RouteGroup)) {
	childGroup := &RouteGroup{
		basePath,
		group.middlewares,
		group.attributes,
		group.routeCollector,
	}

	handler(childGroup)
}

func (group *RouteGroup) get(path string, handler RouteHandler) *RouteGroup {
	return group.addRoute(MethodGet, path, handler)
}

func (group *RouteGroup) post(path string, handler RouteHandler) *RouteGroup {
	return group.addRoute(MethodPost, path, handler)
}

func (group *RouteGroup) head(path string, handler RouteHandler) *RouteGroup {
	return group.addRoute(MethodHead, path, handler)
}

func (group *RouteGroup) put(path string, handler RouteHandler) *RouteGroup {
	return group.addRoute(MethodPut, path, handler)
}

func (group *RouteGroup) patch(path string, handler RouteHandler) *RouteGroup {
	return group.addRoute(MethodPatch, path, handler)
}

func (group *RouteGroup) delete(path string, handler RouteHandler) *RouteGroup {
	return group.addRoute(MethodDelete, path, handler)
}

func (group *RouteGroup) options(path string, handler RouteHandler) *RouteGroup {
	return group.addRoute(MethodOptions, path, handler)
}

func (group *RouteGroup) addRoute(method string, path string, handler RouteHandler) *RouteGroup {
	return group
}

func (group *RouteGroup) middleware(handler RouteMiddleware) *RouteGroup {
	group.middlewares = append(group.middlewares, handler)
	return group
}

func (group *RouteGroup) attribute(name string, value interface{}) *RouteGroup {
	group.attributes[name] = value
	return group
}
