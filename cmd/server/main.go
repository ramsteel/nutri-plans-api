package main

import (
	"nutri-plans-api/api/routes"
	"nutri-plans-api/bootstraps"
	logutil "nutri-plans-api/utils/logger"
	valutil "nutri-plans-api/utils/validation"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

var db *gorm.DB
var v *valutil.Validator

func init() {
	db = bootstraps.NewDatabase()
	v = valutil.NewValidator()
}

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(*logutil.GetLoggerConfig()))

	routes.Init(e, db, v)

	e.Logger.Fatal(e.Start(":8080"))
}
