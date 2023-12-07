package router

import (
	"github.com/labstack/echo/v4"
	"server/internal/onlend/rest"
)

type Router struct {
	router *echo.Echo
}

func NewRouter() *Router {
	return &Router{
		router: echo.New(),
	}
}

func (r *Router) InitRouter(handler *rest.UserHandler) {
	r.router.POST("/api/v1/signup", handler.CreateUser)
}

func (r *Router) Start(addr string) error {
	return r.router.Start(addr)
}
