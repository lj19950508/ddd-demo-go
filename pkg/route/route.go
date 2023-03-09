package route

type Routeable interface{
	Route() *HttpRoutes
}

type HttpRoute struct {
	Handler any
	Pattern string
}

type HttpRoutes []HttpRoute

// type Route interface {
// 	any
// 	Pattern() string
// }