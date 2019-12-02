package main

import (
	"github.com/gin-gonic/gin"
	"log"

	OrderDeliveryGin "g.ghn.vn/logistic/platform/backend/api/order/delivery/http/gin"
)

func main() {

	// g := gin.Default()
	g := gin.New()
	g.Use(gin.Recovery())

	// Middleware
	OrderDeliveryGin.Route(g)
	log.Fatal(g.Run(":9090"))

}
