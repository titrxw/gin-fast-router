package gin_fast_router

import "path"

type RouteGroup struct {
	basePath       string
	middlewares    RouteMiddlewares
	attributes     RouteAttributes
	routeCollector *RouteCollector
}

func (group *RouteGroup) Group(basePath string, handler func(group *RouteGroup)) {
	//slice map是指针类型, 在传递时虽然没有* 底层时按照指针传递 在Go语言中只存在值传递（要么是该值的副本，要么是指针的副本），不存在引用传递
	cloneMiddlewares := make(RouteMiddlewares, len(group.middlewares))
	copy(cloneMiddlewares, group.middlewares)
	cloneAttributes := DeepCopy(group.attributes).(RouteAttributes)

	childGroup := &RouteGroup{
		path.Join(group.basePath, basePath),
		cloneMiddlewares,
		cloneAttributes,
		group.routeCollector,
	}

	handler(childGroup)
}

func (group *RouteGroup) Get(path string, handler RouteHandler) *RouteGroup {
	return group.AddRoute(MethodGet, path, handler)
}

func (group *RouteGroup) Post(path string, handler RouteHandler) *RouteGroup {
	return group.AddRoute(MethodPost, path, handler)
}

func (group *RouteGroup) Head(path string, handler RouteHandler) *RouteGroup {
	return group.AddRoute(MethodHead, path, handler)
}

func (group *RouteGroup) Put(path string, handler RouteHandler) *RouteGroup {
	return group.AddRoute(MethodPut, path, handler)
}

func (group *RouteGroup) Patch(path string, handler RouteHandler) *RouteGroup {
	return group.AddRoute(MethodPatch, path, handler)
}

func (group *RouteGroup) Delete(path string, handler RouteHandler) *RouteGroup {
	return group.AddRoute(MethodDelete, path, handler)
}

func (group *RouteGroup) Options(path string, handler RouteHandler) *RouteGroup {
	return group.AddRoute(MethodOptions, path, handler)
}

func (group *RouteGroup) AddRoute(method string, path string, handler RouteHandler) *RouteGroup {
	return group
}

func (group *RouteGroup) Middleware(handler RouteMiddleware) *RouteGroup {
	group.middlewares = append(group.middlewares, handler)
	return group
}

func (group *RouteGroup) Attribute(name string, value interface{}) *RouteGroup {
	group.attributes[name] = value
	return group
}
