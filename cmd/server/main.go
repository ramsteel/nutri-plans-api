package main

import (
	"nutri-plans-api/api/middlewares"
	"nutri-plans-api/api/routes"
	"nutri-plans-api/bootstraps"
	valutil "nutri-plans-api/utils/validation"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

var db *gorm.DB
var v *valutil.Validator

func init() {
	godotenv.Load()
	db = bootstraps.NewDatabase()
	v = valutil.NewValidator()
}

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(*middlewares.GetLoggerConfig()))

	routes.Init(e, db, v)

	e.Logger.Fatal(e.Start(":8080"))
}
