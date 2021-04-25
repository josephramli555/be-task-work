package controllers

import (
	"net/http"
	"sales-api/config"
	"sales-api/models"
	"github.com/labstack/echo/v4"
	"fmt"
	"strconv"
)




func GetOrderController(c echo.Context) error {
	orderID := c.Param("id")

	if _,err := strconv.Atoi(orderID);err!=nil{
		return c.JSON(http.StatusInternalServerError,map[string]interface{}{
			"code":    400,
			"status": "fail",
			"message":  "invalid id supplied",
		})
	}

	var order models.Order_Result
	var products []models.Product_Result
	res :=config.DB.Raw(config.GET_ORDER_QUERY,orderID).Scan(&order)
	if res.Error!=nil{
		return c.JSON(http.StatusInternalServerError,map[string]interface{}{
			"code":    500,
			"status": "fail",
			"message":  res.Error,
		})
	}else{
		if res:=config.DB.Raw(config.GET_PRODUCTS_LIST_QUERY,orderID).Scan(&products);res.Error!=nil{
			return c.JSON(http.StatusInternalServerError,map[string]interface{}{
				"code":    500,
				"status": "fail",
				"message":  res.Error,
			})
		}
	}
	total := float32(0)
	for _,product := range products{
		total += product.Subtotal
	}
	totalPrice := fmt.Sprintf("$ %.2f",total)
	return c.JSON(http.StatusOK, models.Order_Details_Result{
		Code : 200,
		Status : "success",
		Message: "success get order details",
		Data:  map[string]interface{}{
			"order"  : order,
			"products" : products,
			"totalprice": totalPrice,
		},
	})

}