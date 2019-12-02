package nethttp

import (
	"encoding/json"
	"net/http"

	"g.ghn.vn/logistic/platform/backend/api/order"
	orderUsecase "g.ghn.vn/logistic/platform/backend/api/order/usecase"
)

func Route() {

	uc := orderUsecase.New()
	h := Handler{uc}
	// Routes
	http.HandleFunc("/order", h.GetOrder)
	http.HandleFunc("/health", h.CheckHealth)
}

type Handler struct {
	orderUC order.Usecase
}

func (h *Handler) GetOrder(w http.ResponseWriter, r *http.Request) {
	type req struct {
		ID int
	}

	var reqObj req
	// if err := c.Bind(&reqObj); err != nil {
	// 	return c.JSON(http.StatusBadRequest, nil)
	// }

	// c.JSON(http.StatusOK, h.orderUC.GetOrder(reqObj.ID))
	order := h.orderUC.GetOrder(reqObj.ID)
	bs, _ := json.Marshal(order)
	w.WriteHeader(http.StatusOK)
	w.Write(bs)
}

func (h *Handler) CheckHealth(w http.ResponseWriter, r *http.Request) {
	// c.String(http.StatusOK, "OK")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
