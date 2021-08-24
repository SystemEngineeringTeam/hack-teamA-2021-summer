package main

import (
	"SystemEngineeringTeam/hack-teamA-2021-summer/apifunc"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.POST("/coin", apifunc.Coin_post)
	e.GET("/user", apifunc.User_get)
	e.POST("/user", apifunc.User_post)
	e.PUT("/user", apifunc.User_put)
	e.POST("/login", apifunc.Login_post)
	e.Logger.Fatal(e.Start(":1323"))
}
