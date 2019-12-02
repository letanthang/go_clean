package main

import (
	"log"
	"net/http"

	OrderDeliveryNetHttp "g.ghn.vn/logistic/platform/backend/api/order/delivery/http/nethttp"
)

func main() {

	// Middleware
	OrderDeliveryNetHttp.Route()

	log.Fatal(http.ListenAndServe(":9090", nil))

}
