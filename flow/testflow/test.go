package test

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
)

func TESTPOST(c echo.Context) error {
	fmt.Println(`--TESTPOST--`)

	input := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&input)
	if err != nil {
		log.Error("empty json body")
		return nil
	}
	return c.JSON(200, input)
}

func TESTPOSTBSON(c echo.Context) error {
	fmt.Println(`--TESTPOSTBSON--`)

	input := make(bson.M)
	err := json.NewDecoder(c.Request().Body).Decode(&input)
	if err != nil {
		log.Error("empty json body")
		return nil
	}
	return c.JSON(200, input)
}

func UseSubroute(group *echo.Group) {
	group.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "TESTFLOW")
	})
	group.POST("/test", TESTPOST)
	group.POST("/testBSON", TESTPOSTBSON)
}
