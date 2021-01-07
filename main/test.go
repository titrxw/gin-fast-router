package main

import (
	"fmt"
	"github.com/titrxw/gin-router/Src"
)

func routeHandler() {

}

func initRoute() Src.RouteGroup {
	group := Src.RouteGroup{
		"",
		"",
		make(Src.RouteMiddlewares, 0),
		&Src.RouteCollector{
			make(map[string]Src.RouteMap),
		},
		make(Src.RouteMiddlewares, 0),
	}

	return group
}

func TestBase() {
	group := initRoute()

	group.Get("test", routeHandler)
	group.Middleware(routeHandler).Group("/er", func(group *Src.RouteGroup) {
		group.Post("/erdfgf", routeHandler)
		group.Group("/fdg", func(group *Src.RouteGroup) {
			group.Delete("/fgh", routeHandler)
		})
		group.Middleware(routeHandler).Group("/hgjk", func(group *Src.RouteGroup) {
			group.Head("/he", routeHandler)
		})
		group.Post("/erdfgferw", routeHandler)
	})
	group.Name("test").Post("/test", routeHandler)

	routeDispatcher := Src.RouteDispatcher{
		group.RouteCollector,
	}

	route := routeDispatcher.Dispatcher(Src.MethodPost, "/test")

	if route.GetName() == "test" {
		fmt.Println("route /test found")
	}
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	TestBase()
}
