package main

import (
	"SystemEngineeringTeam/hack-teamA-2021-summer/apifunc"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// /coin : POST
	e.POST("/coin", apifunc.Coin_post)

	// /user : GET
	e.GET("/user", apifunc.User_get)
	// /user : POST
	e.POST("/user", apifunc.User_post)
	// /user : PUT
	e.PUT("/user", apifunc.User_put)

	// /user : DELETE
	e.POST("/login", apifunc.Login_post)

	e.Logger.Fatal(e.Start(":1323"))
}
