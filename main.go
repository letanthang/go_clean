package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"

	OrderDeliveryEchov4 "g.ghn.vn/logistic/platform/backend/api/order/delivery/http/echov4"
)

func main() {
	e := echo.New()
	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	OrderDeliveryEchov4.Route(e)
	log.Fatal(e.Start(":9090"))

}
