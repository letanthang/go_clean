package postgres

import (
	"g.ghn.vn/logistic/platform/backend/api/model"
)

type Repository struct{}

func (r Repository) GetOne(ID int) *model.Order {
	return &model.Order{ID: 1, Receiver: "Dat", Sender: "Tien"}
}

func (r Repository) GetByStatus(OrderStatus string) *[]model.Order {
	return &[]model.Order{{ID: 1, Receiver: "Dat", Sender: "Tien"}, {ID: 2, Receiver: "Thang", Sender: "Tien"}}
}
