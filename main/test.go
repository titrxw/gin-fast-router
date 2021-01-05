package main

import (
	"fmt"
	"github.com/titrxw/gin-router/Src"
)

func test() {

}

func main() {
	group := Src.RouteGroup{
		"",
		"",
		make(Src.RouteMiddlewares, 0),
		&Src.RouteCollector{
			make(map[string]Src.RouteMap),
		},
		make(Src.RouteMiddlewares, 0),
	}

	group.Get("test", test)
	group.Middleware(test).Group("/er", func(group *Src.RouteGroup) {
		group.Post("/erdfgf", test)
		group.Group("/fdg", func(group *Src.RouteGroup) {
			group.Delete("/fgh", test)
		})
		group.Middleware(test).Group("/hgjk", func(group *Src.RouteGroup) {
			group.Head("/he", test)
		})
		group.Post("/erdfgferw", test)
	})

	fmt.Println((group.RouteCollector.RouteMethodMap[Src.MethodHead]))

	routeDispatcher := Src.RouteDispatcher{
		group.RouteCollector,
	}
	route := routeDispatcher.Dispatcher(Src.MethodPost, "/er/erdfgferw")
	fmt.Println(route)
}
