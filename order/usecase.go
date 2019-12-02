package order

import (
	"g.ghn.vn/logistic/platform/backend/api/model"
)

type Usecase interface {
	GetOrder(ID int) *model.Order
	GetTripOrders(tripID int) *[]model.Order
}
