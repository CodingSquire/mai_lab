package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateOrder struct {
	UserID   int            `json:"uid"`
	Products map[string]int `json:"products"`
}

type OID struct {
	ID int `uri:"id"`
}

func (h *HandlerOrder) getOrderById(c *gin.Context) {
	var input OID

	if err := c.ShouldBindUri(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body\n")
		return
	}

	order, err := h.services.Order.Get(input.ID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"id":    order.ID,
			"uid":   order.UserID,
			"store": order.Store,
		})
	}
}

func (h *HandlerOrder) deleteOrder(c *gin.Context) {
	var input OID

	if err := c.ShouldBindUri(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body\n")
		return
	}

	id, err := h.services.Order.Delete(input.ID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})
	}

}

func (h *HandlerOrder) calcAmountOrder(c *gin.Context) {
	var input OID

	if err := c.ShouldBindUri(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body\n")
		return
	}

	amount, err := h.services.Order.Amount(input.ID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"amount": amount,
		})
	}

}

func (h *HandlerOrder) addOrder(c *gin.Context) {
	var input CreateOrder

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body\n")
		return
	}

	id, err := h.services.Order.Create(input.UserID, input.Products)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *HandlerOrder) getOrderList(c *gin.Context) {
	order_list, err := h.services.Order.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	for i := range order_list {
		c.JSON(http.StatusOK, map[string]interface{}{
			"id":    order_list[i].ID,
			"uid":   order_list[i].UserID,
			"store": order_list[i].Store,
		})
	}
}
