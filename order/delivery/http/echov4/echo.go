package echov4

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"g.ghn.vn/logistic/platform/backend/api/order"
	orderUsecase "g.ghn.vn/logistic/platform/backend/api/order/usecase"
)

func Route(e *echo.Echo) {

	uc := orderUsecase.New()
	h := Handler{uc}
	// Routes
	e.PATCH("/order", h.GetOrder)
	e.GET("/health", h.CheckHealth)
}

type Handler struct {
	orderUC order.Usecase
}

func (h *Handler) GetOrder(c echo.Context) error {
	type req struct {
		ID int
	}

	var reqObj req
	if err := c.Bind(&reqObj); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	return c.JSON(http.StatusOK, h.orderUC.GetOrder(reqObj.ID))
}

func (h *Handler) CheckHealth(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
