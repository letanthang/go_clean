package order

import (
	"g.ghn.vn/logistic/platform/backend/api/model"
)

type Repository interface {
	GetOne(ID int) *model.Order
	GetByStatus(OrderStatus string) *[]model.Order
}
