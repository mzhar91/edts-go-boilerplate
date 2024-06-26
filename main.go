package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	echoSwagger "github.com/swaggo/echo-swagger"

	_config "sg-edts.com/edts-go-boilerplate/config"
	_ "sg-edts.com/edts-go-boilerplate/docs"
	_load "sg-edts.com/edts-go-boilerplate/load"
	_auth "sg-edts.com/edts-go-boilerplate/pkg/auth"
	_middleware "sg-edts.com/edts-go-boilerplate/pkg/middleware"
)

func init() {
	log.SetFlags(log.Flags() | log.Llongfile)
	log.SetOutput(os.Stdout)

	_config.LoadEnv()

	if _config.Cfg.Debug {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

// @title Swagger EDTS go boilerplate API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name Harry K
// @contact.email k.harry791@gmail.com

// @host localhost:3001
// @BasePath /
func main() {
	dbConn := _config.InitDB()
	defer func(dbConn *sql.DB) {
		err := dbConn.Close()
		if err != nil {
			fmt.Println("DB connection close failure")

			if _config.Cfg.Debug {
				fmt.Println(err.Error())
			}
		}
	}(dbConn)

	connection := _config.Connection{Database: dbConn}

	e := echo.New()
	middL := _middleware.InitMiddleware()
	e.Use(middL.CORS)

	claims := _auth.InitClaims(_config.Cfg.Debug)
	e.Use(claims.ClaimsContext)

	// Get timeoutcontext
	timeoutContext := _config.GetTimeoutContext()

	_load.Load(e, &connection, timeoutContext)

	_config.ApiSetup()

	// swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	err := e.Start(":" + _config.Cfg.Port)
	if err != nil {
		fmt.Println("Application start up failure")

		if _config.Cfg.Debug {
			fmt.Println(err.Error())
		}
	}
}
