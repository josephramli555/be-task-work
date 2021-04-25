package config

var CREATE_CUSTOMER_TABLE = `
	CREATE TABLE IF NOT EXISTS Customers(
		CustomerID INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
		CompanyName VARCHAR(50),
		FirstName VARCHAR(30),
		LastName VARCHAR(50) ,
		BillingAddress VARCHAR(255),
		City VARCHAR(50) ,
		StateOrProvince VARCHAR(20),
		ZIPCode VARCHAR(20),
		Email VARCHAR(75) ,
		CompanyWebsite VARCHAR(200) ,
		PhoneNumber VARCHAR(30), 
		FaxNumber VARCHAR(30),
		ShipAddress VARCHAR(255),
		ShipCity VARCHAR(50) ,
		ShipStateOrProvince	VARCHAR(50) ,
		ShipZIPCode VARCHAR(20) ,
		ShipPhoneNumber VARCHAR(30) 
	)
`

var CREATE_EMPLOYEES_TABLE = `
	CREATE TABLE IF NOT EXISTS Employees(
		EmployeeID INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
		FirstName VARCHAR(50),
		LastName VARCHAR(50),
		Title VARCHAR(50),
		WorkPhone VARCHAR(30)
	)
`

var CREATE_SHIPPING_METHODS_TABLE = `
	CREATE TABLE IF NOT EXISTS Shipping_Methods(
		ShippingMethodID INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
		ShippingMethod VARCHAR(20)
	)
`

var CREATE_PRODUCTS_TABLE = `
	CREATE TABLE IF NOT EXISTS Products(
		ProductID INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
		ProductName VARCHAR(50),
		UnitPrice Decimal(15,2),
		InStock CHAR(1)
	)
`

var CREATE_ORDER_DETAILS_TABLE = `
	CREATE TABLE IF NOT EXISTS Order_Details(
		OrderDetailID INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
		OrderID INT,
		ProductID INT,
		Quantity INT,
		UnitPrice Decimal(15,2),
		Discount Decimal(5,2),
		FOREIGN KEY (OrderID) REFERENCES Orders(OrderID),
		FOREIGN KEY (ProductID) REFERENCES Products(ProductID)
	)
`

var CREATE_ORDERS_TABLE = `
	CREATE TABLE IF NOT EXISTS Orders(
		OrderID INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
		CustomerID INT,
		EmployeeID INT,
		OrderDate DATE,
		PurchaseOrderNumber VARCHAR(30),
		ShipDate DATE,
		ShippingMethodID INT,
		FreightCharge INT,
		Taxes INT,
		PaymentReceived CHAR(1),
		Comment VARCHAR(150),
		FOREIGN KEY (CustomerID) REFERENCES Customers(CustomerID),
		FOREIGN KEY (EmployeeID) REFERENCES Employees(EmployeeID),
		FOREIGN KEY (ShippingMethodID) REFERENCES Shipping_Methods(ShippingMethodID)
	)
`