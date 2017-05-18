package router

import (
	"net/http"

	"github.com/labstack/echo"
)

func helloworld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
