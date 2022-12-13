package handlars

import (
	"mymachine707/models"
	"mymachine707/protogen/eCommerce"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetOrderItemByID godoc
//
//	@Summary		GetOrderItemByID
//	@Description	get an orderItem by id
//	@Tags			orderItem
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"OrderItem id"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.OrderItem}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/orderItem/{id} [get]
func (h *handler) GetOrderItemByID(c *gin.Context) {

	idStr := c.Param("id")

	orderItem, err := h.grpcClient.OrderItem.GetOrderItemById(c.Request.Context(), &eCommerce.GetOrderItemByIDRequest{
		Id: idStr,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "OK",
		Data:    orderItem,
	})
}

// GetOrderItemList godoc
//
//	@Summary		List orderItems
//	@Description	GetOrderItemList
//	@Tags			orderItem
//	@Accept			json
//	@Produce		json
//	@Param			offset			query		int		false	"0"
//	@Param			limit			query		int		false	"100"
//	@Param			search			query		string	false	"search exapmle"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		200				{object}	models.JSONResult{data=[]models.OrderItem}
//	@Router			/v1/orderItem/ [get]
func (h *handler) GetOrderItemList(c *gin.Context) {

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

	orderItemList, err := h.grpcClient.OrderItem.GetOrderItemList(c.Request.Context(), &eCommerce.GetOrderItemListRequest{
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
		Message: "GetOrderItemList OK",
		Data:    orderItemList,
	})
}
