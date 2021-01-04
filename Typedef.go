package gin_fast_router

type RouteHandler func()
type RouteAttributes map[string]interface{}
type RouteMiddleware RouteHandler
type RouteMiddlewares []RouteMiddleware
type RouteMap []*Route
