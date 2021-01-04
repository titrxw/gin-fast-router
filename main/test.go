package main

import (
	"fmt"
	"github.com/titrxw/gin-router"
)

func test() {

}

func main() {
	group := gin_fast_router.RouteGroup{
		"",
		"",
		make(gin_fast_router.RouteMiddlewares, 0),
		&gin_fast_router.RouteCollector{
			make(map[string]gin_fast_router.RouteMap),
		},
		make(gin_fast_router.RouteMiddlewares, 0),
	}

	group.Get("test", test)
	group.Middleware(test).Group("/er", func(group *gin_fast_router.RouteGroup) {
		group.Post("/erdfgf", test)
		group.Group("/fdg", func(group *gin_fast_router.RouteGroup) {
			group.Delete("/fgh", test)
		})
		group.Middleware(test).Group("/hgjk", func(group *gin_fast_router.RouteGroup) {
			group.Head("/he", test)
		})
		group.Post("/erdfgferw", test)
	})

	fmt.Println((group.RouteCollector.RouteMethodMap[gin_fast_router.MethodHead][0]))

}
