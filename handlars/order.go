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
//	@Description	Creat a new category
//	@Tags			category
//	@Accept			json
//	@Produce		json
//	@Param			category		body		models.CreateOrderModul	true	"Order body"
//	@Param			Authorization	header		string						false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Order}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/category [post]
func (h *handler) CreatOrder(c *gin.Context) {

	var body models.CreateOrderModul

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	category, err := h.grpcClient.Order.CreateOrder(c.Request.Context(), &eCommerce.CreateOrderRequest{
		ProductId: body.Product_id,
		ClientId:  body.Client_id,
		Count:     body.Count,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "CreatOrder",
		Data:    category,
	})
}

// GetOrderByID godoc
//
//	@Summary		GetOrderByID
//	@Description	get an category by id
//	@Tags			category
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"Order id"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Order}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/category/{id} [get]
func (h *handler) GetOrderByID(c *gin.Context) {

	idStr := c.Param("id")

	category, err := h.grpcClient.Order.GetOrderById(c.Request.Context(), &eCommerce.GetOrderByIDRequest{
		Id: idStr,
	})

	// get order with product info

	product, err := h.grpcClient.Product.GetProductById(c.Request.Context(), &eCommerce.GetProductByIDRequest{
		Id: idStr,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	//---------------------------------

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "GetOrderById OK",
		Data:    category,
	})
}

// GetOrderList godoc
//
//	@Summary		List categorys
//	@Description	GetOrderList
//	@Tags			category
//	@Accept			json
//	@Produce		json
//	@Param			offset			query		int		false	"0"
//	@Param			limit			query		int		false	"100"
//	@Param			search			query		string	false	"search exapmle"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		200				{object}	models.JSONResult{data=[]models.Order}
//	@Router			/v1/category/ [get]
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

	categoryList, err := h.grpcClient.Order.GetOrderList(c.Request.Context(), &eCommerce.GetOrderListRequest{
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
		Data:    categoryList,
	})
}

// OrderUpdate godoc
//
//	@Summary		Update Order
//	@Description	Update Order
//	@Tags			category
//	@Accept			json
//	@Produce		json
//	@Param			category		body		models.UpdateOrderModul	true	"Order body"
//	@Param			Authorization	header		string						false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=[]models.Order}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/category/ [put]
func (h *handler) OrderUpdate(c *gin.Context) {

	var body models.UpdateOrderModul

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	category, err := h.grpcClient.Order.UpdateOrder(c.Request.Context(), &eCommerce.UpdateOrderRequest{
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
		"data":    category,
	})

}

// DeleteOrder godoc
//
//	@Summary		Delete Order
//	@Description	get element by id and delete this category
//	@Tags			category
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"Order id"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Order}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/category/{id} [delete]
func (h *handler) DeleteOrder(c *gin.Context) {
	idStr := c.Param("id")

	category, err := h.grpcClient.Order.DeleteOrder(c.Request.Context(), &eCommerce.DeleteOrderRequest{
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
		"data":    category,
	})

}
