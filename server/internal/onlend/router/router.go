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

func (r *Router) InitRouter(userHandler *rest.UserHandler, accountHandler *rest.AccountHandler, transactionHandler *rest.TransactionHandler) {
	r.router.POST("/api/v1/signup", userHandler.CreateUser)
	r.router.POST("/api/v1/login", userHandler.Login)
	r.router.GET("/api/v1/logout", userHandler.Logout)
	r.router.GET("/api/v1/users", userHandler.GetAllUsers)

	r.router.GET("/api/v1/accounts/:id", accountHandler.GetAccount)
	r.router.GET("/api/v1/accounts/user/:id", accountHandler.GetAccountByUserId)
	r.router.GET("/api/v1/accounts", accountHandler.GetAllAccounts)
	r.router.PUT("/api/v1/accounts", accountHandler.UpdateAccount)

	r.router.GET("/api/v1/transactions/:id", transactionHandler.GetTransaction)
	r.router.GET("/api/v1/transactions", transactionHandler.GetAllTransactions)
	r.router.POST("/api/v1/transactions/transfer", transactionHandler.TransferFunds)
	r.router.POST("/api/v1/transactions", transactionHandler.CreateTransaction)

}

func (r *Router) Start(addr string) error {
	return r.router.Start(addr)
}
