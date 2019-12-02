package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"g.ghn.vn/logistic/platform/backend/api/order"
	orderUsecase "g.ghn.vn/logistic/platform/backend/api/order/usecase"
)

func Route(g *gin.Engine) {

	uc := orderUsecase.New()
	h := Handler{uc}
	// Routes
	g.PATCH("/order", h.GetOrder)
	g.GET("/health", h.CheckHealth)
}

type Handler struct {
	orderUC order.Usecase
}

func (h *Handler) GetOrder(c *gin.Context) {
	type req struct {
		ID int
	}

	var reqObj req
	// if err := c.Bind(&reqObj); err != nil {
	// 	return c.JSON(http.StatusBadRequest, nil)
	// }

	c.JSON(http.StatusOK, h.orderUC.GetOrder(reqObj.ID))
}

func (h *Handler) CheckHealth(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
