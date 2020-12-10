package main

import (
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/lmolas/echo-server/docs"

	"github.com/lmolas/echo-server/routes"
)

// @title Echo API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /api/v1
func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "/static/")
	e.GET("/request", routes.GetRequest)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/swagger.json", routes.GetSwagger)

	e.Logger.Fatal(e.Start(":8080"))
}
