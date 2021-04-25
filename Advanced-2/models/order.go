package models

import(
	"time"
)

type Orders struct{
	OrderID int
	CustomerID int
	EmployeeID int
	OrderDate time.Time
	PurchaseOrderNumber string
	ShipDate time.Time
	ShippingMethodID int
	FreightCharge int
	Taxes int
	PaymentReceived string
	Comment string
}


type Order_Result struct{
	EmployeeName string `json:"employeename"`
	CustomerName string `json:"customername"`
	ShippingMethod string `json:"shippingmethod"`
}