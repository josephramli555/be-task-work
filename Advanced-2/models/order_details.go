package models

type Order_Details struct{
	OrderDetailID int
	OrderID int
	ProductID int
	Quantity int
	UnitPrice float32
	Discount float32
}

type Order_Details_Result struct{
	Code int `json:"code"`
	Status string `json:"status"`
	Message string `json:"message"`
	Data map[string]interface{} `json:"data"`
}