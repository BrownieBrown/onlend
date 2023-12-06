package router

import (
	"github.com/labstack/echo/v4"
	"server/internal/onlend/rest"
)

var router *echo.Echo

func InitRouter(handler *rest.UserHandler) {
	router = echo.New()

	router.POST("/api/v1/signup", handler.CreateUser)
}

func Start(addr string) error {
	return router.Start(addr)
}
