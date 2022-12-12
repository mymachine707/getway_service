package handlars

import (
	"mymachine707/models"
	"mymachine707/protogen/eCommerce"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreatOrder godoc
//
//	@Summary		Creat Order
//	@Description	Creat a new order
//	@Tags			order
//	@Accept			json
//	@Produce		json
//	@Param			order		body		models.CreateOrderModul	true	"Order body"
//	@Param			Authorization	header		string						false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Order}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/order [post]
func (h *handler) CreatOrder(c *gin.Context) {

	var body models.CreateOrderModul

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	// Product bizda bor yo'ligini tekshiramiz!!!
	product, err := h.grpcClient.Product.GetProductById(c.Request.Context(), &eCommerce.GetProductByIDRequest{
		Id: body.Product_id,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: "We do not have this product!",
		})
		return
	}
	// Tekshiruvdan o'tsa keyingi bosqichga boradi. O'tmasa chiqib ketadi.

	// Client bizda bor yo'ligini tekshiramiz!!!
	client, err := h.grpcClient.Client.GetClientById(c.Request.Context(), &eCommerce.GetClientByIDRequest{
		Id: body.Client_id,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: "You are not registered! Please register!",
		})
		return
	}
	// Tekshiruvdan o'tsa keyingi bosqichga boradi. O'tmasa chiqib ketadi.

	order, err := h.grpcClient.Order.CreateOrder(c.Request.Context(), &eCommerce.CreateOrderRequest{
		ProductId: product.Id,
		ClientId:  client.Id,
		Count:     body.Count,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	// Todo calculate total price
	price, err := strconv.Atoi(product.Price)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	count, err := strconv.Atoi(body.Count)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	totalPriceProduct := price * count
	total := strconv.Itoa(totalPriceProduct)
	// finish

	// Orderni OrderItemga kirgazish

	_, err = h.grpcClient.OrderItem.CreateOrderItem(c.Request.Context(), &eCommerce.CreateOrderItemRequest{
		OredrId:    order.Id,
		TotalPrice: total,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "CreatOrder",
		Data:    order,
	})

}

// GetOrderByID godoc
//
//	@Summary		GetOrderByID
//	@Description	get an order by id
//	@Tags			order
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"Order id"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Order}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/order/{id} [get]
func (h *handler) GetOrderByID(c *gin.Context) {

	idStr := c.Param("id")

	order, err := h.grpcClient.Order.GetOrderById(c.Request.Context(), &eCommerce.GetOrderByIDRequest{
		Id: idStr,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	// get order with product info

	product, err := h.grpcClient.Product.GetProductById(c.Request.Context(), &eCommerce.GetProductByIDRequest{
		Id: order.ProductId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	//---------------------------------

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "GetOrderById OK",
		Data: eCommerce.GetOrderByIDResponse{
			Id:        order.Id,
			ProductId: order.ProductId,
			Product: &eCommerce.Product{
				Id:          product.Id,
				CategoryId:  product.Category.Id,
				ProductName: product.ProductName,
				Description: product.Description,
				Price:       product.Price,
				CreatedAt:   product.CreatedAt,
				UpdatedAt:   product.UpdatedAt,
			},
			ClientId:       order.ClientId,
			ClientUsername: order.ClientUsername,
			CreatedAt:      order.CreatedAt,
			UpdatedAt:      order.UpdatedAt,
		},
	})
}

// GetOrderList godoc
//
//	@Summary		List orders
//	@Description	GetOrderList
//	@Tags			order
//	@Accept			json
//	@Produce		json
//	@Param			offset			query		int		false	"0"
//	@Param			limit			query		int		false	"100"
//	@Param			search			query		string	false	"search exapmle"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		200				{object}	models.JSONResult{data=[]models.Order}
//	@Router			/v1/order/ [get]
func (h *handler) GetOrderList(c *gin.Context) {

	offsetStr := c.DefaultQuery("offset", h.cfg.Default_Offset)
	limitStr := c.DefaultQuery("limit", h.cfg.Default_Limit)

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	orderList, err := h.grpcClient.Order.GetOrderList(c.Request.Context(), &eCommerce.GetOrderListRequest{
		Offset: int32(offset),
		Limit:  int32(limit),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "GetOrderList OK",
		Data:    orderList,
	})
}

// OrderUpdate godoc
//
//	@Summary		Update Order
//	@Description	Update Order
//	@Tags			order
//	@Accept			json
//	@Produce		json
//	@Param			order		body		models.UpdateOrderModul	true	"Order body"
//	@Param			Authorization	header		string						false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=[]models.Order}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/order/ [put]
func (h *handler) OrderUpdate(c *gin.Context) {

	var body models.UpdateOrderModul

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	order, err := h.grpcClient.Order.UpdateOrder(c.Request.Context(), &eCommerce.UpdateOrderRequest{
		Id:    body.ID,
		Count: body.Count,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order Updated",
		"data":    order,
	})

}

// DeleteOrder godoc
//
//	@Summary		Delete Order
//	@Description	get element by id and delete this order
//	@Tags			order
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"Order id"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Order}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/order/{id} [delete]
func (h *handler) DeleteOrder(c *gin.Context) {
	idStr := c.Param("id")

	order, err := h.grpcClient.Order.DeleteOrder(c.Request.Context(), &eCommerce.DeleteOrderRequest{
		Id: idStr,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order Deleted",
		"data":    order,
	})

}
