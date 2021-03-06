package generic

import (
	"github.com/labstack/echo/v4"
)

func InitModule(e *echo.Echo) {
	genericPresenter := newGenericEcho()

	// Bind generic routes
	generic := e.Group("")
	generic.GET("/version", genericPresenter.GetVersion)
}
