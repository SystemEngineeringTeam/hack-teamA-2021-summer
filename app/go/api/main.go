package main

import (
	"SystemEngineeringTeam/hack-teamA-2021-summer/apifunc"
	"fmt"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Static("/static/img", "./static/img")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		fmt.Fprintf(os.Stderr, "Request: %v\n", string(reqBody))
	}))

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

	// http://localhost:8080/setting : GET apifunc->setting.go->SettingGet()
	e.GET("/setting", apifunc.SettingGet)
	// http://localhost:8080/setting : POST apifunc->setting.go->SettingPost()
	e.POST("/setting", apifunc.SettingPost)

	// http://localhost:8080/images : GET apifunc->images.go->imagesGet()
	e.GET("/images", apifunc.ImageGet)

	// ChartData : GET apifunc->ChartData.go->ChartDataGet()
	e.GET("/chartdata", apifunc.ChartDataGet)
	// http://localhost:8080/chartdata : POST apifunc->ChartData.go->ChartDataPost()
	e.POST("/chartdata", apifunc.ChartDataPost)

	// 8080番ポートで待ち受け
	e.Logger.Fatal(e.Start(":8080"))
}
