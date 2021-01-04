package gin_fast_router

type RouteCollector struct {
	RouteMethodMap map[string]RouteMap
}

func (collector *RouteCollector) collect(route *Route) {
	_, has := collector.RouteMethodMap[route.method]
	if !has {
		collector.RouteMethodMap[route.method] = make(RouteMap, 0)
	}
	collector.RouteMethodMap[route.method] = append(collector.RouteMethodMap[route.method], route)
}
