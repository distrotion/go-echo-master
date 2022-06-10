package flow2

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func UseSubroute(group *echo.Group) {
	group.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Flow02")
	})

}
