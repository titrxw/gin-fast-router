package Src

type RouteHandler func()
type RouteAttributes map[string]interface{}
type RouteMiddleware RouteHandler
type RouteMiddlewares []RouteMiddleware
type RouteMap map[string]*Route
