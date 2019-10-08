package server

import (
	"net/http"
)

type Router struct {
	routes map[Path]Route
}

type Path string
type Method string

type Route map[Method]http.Handler

func NewRouter() *Router {
	return &Router{
		routes: map[Path]Route{},
	}
}

func (r *Router) Route(p Path) Route {
	if r.routes[p] == nil {
		r.routes[p] = Route{}
	}
	return r.routes[p]
}
