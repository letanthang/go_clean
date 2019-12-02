package usecase

import (
	"g.ghn.vn/logistic/platform/backend/api/model"
	"g.ghn.vn/logistic/platform/backend/api/order"
	orderRepo "g.ghn.vn/logistic/platform/backend/api/order/repository"
)

func New() order.Usecase {
	return &OrderUseCase{&orderRepo.Repository{}}
}

type OrderUseCase struct {
	orderRepo order.Repository
}

func (o *OrderUseCase) GetOrder(ID int) *model.Order {
	return o.orderRepo.GetOne(ID)
}
func (o *OrderUseCase) GetTripOrders(tripID int) *[]model.Order {
	return nil
}
