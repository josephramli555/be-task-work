package config

var GET_ORDER_QUERY string = `
	SELECT CONCAT(c.FirstName,' ',c.LastName) AS CustomerName,CONCAT(e.FirstName,' ',e.LastName) AS EmployeeName,
	s.ShippingMethod as ShippingMethod
	FROM Orders AS o 
	INNER JOIN Shipping_Methods AS s ON o.ShippingMethodID = s.ShippingMethodID
	INNER JOIN Customers AS c ON o.CustomerID = c.CustomerID
	INNER JOIN Employees AS e ON o.EmployeeID = e.EmployeeID
	WHERE o.OrderID = ?
`

var GET_PRODUCTS_LIST_QUERY string = `
	SELECT p.ProductID,p.ProductName,od.Quantity,od.UnitPrice,od.Discount,
	((od.Quantity * od.UnitPrice) - od.Discount) AS Subtotal 
	FROM Order_Details AS od INNER JOIN Products AS p ON od.ProductID = p.ProductID
	WHERE od.OrderID = ?
`
