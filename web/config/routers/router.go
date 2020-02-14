package routers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var Echo *echo.Echo

func init() {
	Echo = echo.New()

	// middleware setting
	Echo.Pre(middleware.NonWWWRedirect())
	Echo.Use(middleware.Logger())
}
