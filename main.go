package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/thanhlam/home-sso-service/service"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "SSO SERVICE")
	})
	//<---------------------------SSO------------------------------>
	e.POST("/api/sso/requestToken", service.RequestSSOTokenv2)
	e.POST("/api/sso/parseToken", service.ParseSSOToken)
	e.Logger.Fatal(e.Start(":1323"))
}
