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
	r.router.POST("/api/v1/login", handler.Login)
	r.router.GET("/api/v1/logout", handler.Logout)
	r.router.GET("/api/v1/users", handler.GetAllUsers)
}

func (r *Router) Start(addr string) error {
	return r.router.Start(addr)
}
