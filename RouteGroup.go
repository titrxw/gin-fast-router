package gin_fast_router

import (
	"path"
)

type RouteGroup struct {
	BasePath       string
	RouteName      string
	Middlewares    RouteMiddlewares
	RouteCollector *RouteCollector
	CurMiddlewares RouteMiddlewares
}

func (group *RouteGroup) Group(basePath string, handler func(group *RouteGroup)) {
	//slice map是指针类型, 在传递时虽然没有* 底层时按照指针传递 在Go语言中只存在值传递（要么是该值的副本，要么是指针的副本），不存在引用传递
	//cloneAttributes := DeepCopy(group.attributes).(RouteAttributes)
	cloneMiddlewares := make(RouteMiddlewares, len(group.Middlewares)+len(group.CurMiddlewares))
	copy(cloneMiddlewares, group.Middlewares)
	copy(cloneMiddlewares[len(group.Middlewares):], group.CurMiddlewares)

	childGroup := &RouteGroup{
		path.Join(group.BasePath, basePath),
		group.RouteName,
		cloneMiddlewares,
		group.RouteCollector,
		make(RouteMiddlewares, 0),
	}

	handler(childGroup)
	group.resetRouteGroup()
}

func (group *RouteGroup) Get(path string, handler RouteHandler) {
	group.addRoute(MethodGet, path, handler)
}

func (group *RouteGroup) Post(path string, handler RouteHandler) {
	group.addRoute(MethodPost, path, handler)
}

func (group *RouteGroup) Head(path string, handler RouteHandler) {
	group.addRoute(MethodHead, path, handler)
}

func (group *RouteGroup) Put(path string, handler RouteHandler) {
	group.addRoute(MethodPut, path, handler)
}

func (group *RouteGroup) Patch(path string, handler RouteHandler) {
	group.addRoute(MethodPatch, path, handler)
}

func (group *RouteGroup) Delete(path string, handler RouteHandler) {
	group.addRoute(MethodDelete, path, handler)
}

func (group *RouteGroup) Options(path string, handler RouteHandler) {
	group.addRoute(MethodOptions, path, handler)
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

func (group *RouteGroup) addRoute(method string, routePath string, handler RouteHandler) {
	cloneMiddlewares := make(RouteMiddlewares, len(group.Middlewares)+len(group.CurMiddlewares))
	copy(cloneMiddlewares, group.Middlewares)
	copy(cloneMiddlewares[len(group.Middlewares):], group.CurMiddlewares)

	route := &Route{
		group.RouteName,
		path.Join(group.BasePath, routePath),
		method,
		handler,
		cloneMiddlewares,
		nil,
	}

	group.RouteCollector.collect(route)
	group.resetRouteGroup()
}

func (group *RouteGroup) Name(name string) *RouteGroup {
	group.RouteName = name
	return group
}

func (group *RouteGroup) Middleware(handler RouteMiddleware) *RouteGroup {
	group.CurMiddlewares = append(group.CurMiddlewares, handler)
	return group
}

func (group *RouteGroup) resetRouteGroup() {
	group.RouteName = ""
	group.CurMiddlewares = group.CurMiddlewares[0:0]
}
