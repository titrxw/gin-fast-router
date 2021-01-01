package gin_fast_router

import "path"

type RouteGroup struct {
	basePath       string
	name           string
	middlewares    RouteMiddlewares
	routeCollector *RouteCollector
	parentGroup    *RouteGroup
}

func (group *RouteGroup) Group(basePath string, handler func(group *RouteGroup)) {
	//slice map是指针类型, 在传递时虽然没有* 底层时按照指针传递 在Go语言中只存在值传递（要么是该值的副本，要么是指针的副本），不存在引用传递
	cloneMiddlewares := make(RouteMiddlewares, len(group.middlewares))
	copy(cloneMiddlewares, group.middlewares)
	//cloneAttributes := DeepCopy(group.attributes).(RouteAttributes)

	childGroup := &RouteGroup{
		path.Join(group.basePath, basePath),
		group.name,
		cloneMiddlewares,
		group.routeCollector,
		group,
	}

	handler(childGroup)
	group.resetRouteGroup()
}

func (group *RouteGroup) Get(path string, handler RouteHandler) {
	group.AddRoute(MethodGet, path, handler)
}

func (group *RouteGroup) Post(path string, handler RouteHandler) {
	group.AddRoute(MethodPost, path, handler)
}

func (group *RouteGroup) Head(path string, handler RouteHandler) {
	group.AddRoute(MethodHead, path, handler)
}

func (group *RouteGroup) Put(path string, handler RouteHandler) {
	group.AddRoute(MethodPut, path, handler)
}

func (group *RouteGroup) Patch(path string, handler RouteHandler) {
	group.AddRoute(MethodPatch, path, handler)
}

func (group *RouteGroup) Delete(path string, handler RouteHandler) {
	group.AddRoute(MethodDelete, path, handler)
}

func (group *RouteGroup) Options(path string, handler RouteHandler) {
	group.AddRoute(MethodOptions, path, handler)
}

func (group *RouteGroup) any(path string, handler RouteHandler) {
	group.Get(path, handler)
	group.Post(path, handler)
	group.Options(path, handler)
	group.Head(path, handler)
	group.Put(path, handler)
	group.Delete(path, handler)
	group.Patch(path, handler)
}

func (group *RouteGroup) AddRoute(method string, routePath string, handler RouteHandler) {
	var cloneMiddlewares RouteMiddlewares

	if group.parentGroup != nil {
		cloneMiddlewares := make(RouteMiddlewares, len(group.middlewares)+len(group.parentGroup.middlewares))
		copy(cloneMiddlewares, group.middlewares)
		copy(cloneMiddlewares[len(group.middlewares):], group.parentGroup.middlewares)
	} else {
		cloneMiddlewares := make(RouteMiddlewares, len(group.middlewares))
		copy(cloneMiddlewares, group.middlewares)
	}

	route := &Route{
		group.name,
		path.Join(group.basePath, routePath),
		method,
		handler,
		cloneMiddlewares,
		nil,
	}

	group.routeCollector.collect(route)
	group.resetRouteGroup()
}

func (group *RouteGroup) Name(name string) *RouteGroup {
	group.name = name
	return group
}

func (group *RouteGroup) Middleware(handler RouteMiddleware) *RouteGroup {
	group.middlewares = append(group.middlewares, handler)
	return group
}

func (group *RouteGroup) resetRouteGroup() {
	group.name = ""
	group.middlewares = group.middlewares[0:0]
}
