package config

var INSERT_CUSTOMER_QUERY string = `
	INSERT INTO Customers(
	CustomerID,CompanyName,FirstName,LastName,BillingAddress,City,
	StateOrProvince,ZIPCode,Email,CompanyWebsite,PhoneNumber,FaxNumber
	,ShipAddress,ShipCity,ShipStateOrProvince,ShipZipCode,ShipPhoneNumber
	) VALUES `

var INSERT_EMPLOYEES_QUERY string = `
	INSERT INTO Employees(
	EmployeeID,FirstName,LastName,Title,WorkPhone
	) VALUES `

var INSERT_SHIPPING_METHOD_QUERY string = `
	INSERT INTO Shipping_Methods(
		ShippingMethodID,ShippingMethod
	) VALUES `

var INSERT_PRODUCTS_QUERY string = `
	INSERT INTO Products(
	ProductID,ProductName,UnitPrice,InStock
	) VALUES 
`

var INSERT_ORDERS_QUERY string = `
	INSERT INTO Orders(
		OrderID,CustomerID,EmployeeID,OrderDate,
		PurchaseOrderNumber,ShipDate,ShippingMethodID,
		FreightCharge,Taxes,PaymentReceived,Comment
	) VALUES
`

var INSERT_ORDER_DETAILS_QUERY string = `
	INSERT INTO Order_Details(
		OrderDetailID,OrderID,ProductID,
		Quantity,UnitPrice,Discount
	) VALUES
`