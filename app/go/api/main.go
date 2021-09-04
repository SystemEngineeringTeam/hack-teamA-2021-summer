package main

import (
	"SystemEngineeringTeam/hack-teamA-2021-summer/apifunc"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// http://localhost:8080/coin : GET apifunc->coin.go->CoinPost()
	e.POST("/coin", apifunc.CoinPost)

	// http://localhost:8080/user : GET apifunc->user.go->UserGet()
	e.GET("/user", apifunc.UserGet)
	// http://localhost:8080/user : POST apifunc->user.go->UserPost()
	e.POST("/user", apifunc.UserPost)
	// http://localhost:8080/user : PUT apifunc->user.go->UserPut()
	e.PUT("/user", apifunc.UserPut)

	// http://localhost:8080/login : POST apifunc->login.go->LoginPost()
	e.POST("/login", apifunc.LoginPost)

	// 8080番ポートで待ち受け
	e.Logger.Fatal(e.Start(":8080"))
}
