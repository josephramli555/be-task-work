package models

type Products struct{
	ProductID int
	ProductName string
	UnitPrice float32
	InStock string
}


type Product_Result struct{
	ProductID int `json:"productid"`
	ProductName string `json:"productname"`
	UnitPrice float32 `json:"unitprice"`
	Quantity int `json:"quantity"`
	Discount float32 `json:"discount"`
	Subtotal float32 `json:"subtotal"`
}