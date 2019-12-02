package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"

	OrderDeliveryEchov3 "g.ghn.vn/logistic/platform/backend/api/order/delivery/http/echov3"
)

func main() {
	e := echo.New()
	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	OrderDeliveryEchov3.Route(e)
	log.Fatal(e.Start(":9090"))

}
