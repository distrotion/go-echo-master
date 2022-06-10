package main

import (
	flow1 "echomaster/flow/001"
	flow2 "echomaster/flow/002"
	test "echomaster/flow/testflow"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "distrotion echo v0.00")
	})

	apiGroup := e.Group("FLOWTEST")
	test.UseSubroute(apiGroup)

	FLOW1Group := e.Group("FLOW001")
	flow1.UseSubroute(FLOW1Group)

	FLOW2Group := e.Group("FLOW002")
	flow2.UseSubroute(FLOW2Group)

	e.Logger.Fatal(e.Start(":12000"))
}
