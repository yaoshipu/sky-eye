package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// AddRouters ...
func AddRouters(e *echo.Echo) {

	// ===== PUblic APIs =====
	public := e.Group("/api")
	public.POST("/users/login", login)
	public.GET("/hello", helloworld)

	// ===== Authorized APIs =====
	api := e.Group("/api")
	api.Use(middleware.JWT([]byte("secret")))
}
