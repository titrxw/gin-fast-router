package gin_fast_router

import "github.com/gin-gonic/gin"

type RouteHandler func(ctx *gin.Context)
type RouteAttributes map[string]interface{}
type RouteMiddleware RouteHandler
type RouteMiddlewares []RouteMiddleware
