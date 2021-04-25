package main

import (
	"bufio"
	"log"
	"os"
	"sales-api/config"
	"strconv"
	"strings"
)

func ReadFile(path string) ([]string){
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln("%s when open %s", err,path)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
    var scannedWords []string
	for scanner.Scan() {
        scannedWords = append(scannedWords, scanner.Text())
    }
	file.Close()
	return scannedWords
}

func ConvertInt(a string)int{
	b,err := strconv.Atoi(a)
	if err!=nil{
		log.Fatal(err)
	}
	return b
}

func ConvertFloat(a string)float32{
	b,err := strconv.ParseFloat(strings.ReplaceAll(a,",","."),3)
		if err!=nil{
			log.Fatal("Error during conversion :",err)
		}
	return float32(b)
}


func InsertCustomers(){
	words := ReadFile("csvdata/Customers.csv")
	stringInserted := []string{}
	var valueParams []interface{}
	for _,word := range words[1:]{
		stringInserted = append(stringInserted, "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		a := strings.Split(word,";")
		customerID := ConvertInt(a[0])
		valueParams = append(valueParams,customerID,a[1],a[2],a[3],a[4],a[5],a[6],a[7],a[8],
		a[9],a[10],a[11],a[12],a[13],a[14],a[15],a[16])
	}

	combinedQueryVal := strings.Join(stringInserted,",")
	insertQuery := config.INSERT_CUSTOMER_QUERY + combinedQueryVal

	res := config.DB.Exec(insertQuery,valueParams...)
	if res.Error!=nil{
		log.Fatal("Error %s during insert to customer",res.Error)
	}
	log.Printf("Insert into customer success")
}

func InsertEmployee(){
	words := ReadFile("csvdata/Employees.csv")
	stringInserted := []string{}
	var valueParams []interface{}
	for _,word := range words[1:]{
		stringInserted = append(stringInserted, "(?, ?, ?, ?, ?)")
		a := strings.Split(word,";")
		employeeID := ConvertInt(a[0])
		valueParams = append(valueParams,employeeID,a[1],a[2],a[3],a[4])
	}

	combinedQueryVal := strings.Join(stringInserted,",")
	insertQuery := config.INSERT_EMPLOYEES_QUERY + combinedQueryVal

	res := config.DB.Exec(insertQuery,valueParams...)
	if res.Error!=nil{
		log.Fatal("Error %s during insert to employee",res.Error)
	}
	log.Printf("Insert into employee success")
}

func InsertShippingMethod(){
	words := ReadFile("csvdata/ShippingMethods.csv")
	stringInserted := []string{}
	var valueParams []interface{}
	for _,word := range words[1:]{
		stringInserted = append(stringInserted, "(?, ?)")
		a := strings.Split(word,";")
		shippingID := ConvertInt(a[0])
		valueParams = append(valueParams,shippingID,a[1])
	}

	combinedQueryVal := strings.Join(stringInserted,",")
	insertQuery := config.INSERT_SHIPPING_METHOD_QUERY + combinedQueryVal

	res := config.DB.Exec(insertQuery,valueParams...)
	if res.Error!=nil{
		log.Fatal("Error during insert to shipping method : ",res.Error)
	}
	log.Printf("Insert into shipping method success")
}

func InsertProducts(){
	words := ReadFile("csvdata/Products.csv")
	stringInserted := []string{}
	var valueParams []interface{}
	for _,word := range words[1:]{
		stringInserted = append(stringInserted, "(?, ?, ?, ?)")
		a := strings.Split(word,";")
		productID := ConvertInt(a[0])
		unitPrice := ConvertFloat(a[2])
		valueParams = append(valueParams,productID,a[1],unitPrice,a[3])
	}

	combinedQueryVal := strings.Join(stringInserted,",")
	insertQuery := config.INSERT_PRODUCTS_QUERY + combinedQueryVal

	res := config.DB.Exec(insertQuery,valueParams...)
	if res.Error!=nil{
		log.Fatal("Error during insert to product : ",res.Error)
	}
	log.Printf("Insert into product success")
}

func InsertOrders(){
	words := ReadFile("csvdata/Orders.csv")
	stringInserted := []string{}
	var valueParams []interface{}
	for _,word := range words[1:]{
		stringInserted = append(stringInserted, "(?, ?, ?, STR_TO_DATE(?, '%d/%m/%Y'), ?, STR_TO_DATE(?, '%d/%m/%Y'), ?, ?, ?, ?, ?)")
		a := strings.Split(word,";")
		orderID := ConvertInt(a[0])
		customerID := ConvertInt(a[1])
		employeeID:= ConvertInt(a[2])
		shippingMethodID:= ConvertInt(a[6])
		valueParams = append(valueParams,orderID,customerID,employeeID,a[3],a[4],a[5],
		shippingMethodID,a[7],a[8],a[9],a[10])
	}

	combinedQueryVal := strings.Join(stringInserted,",")
	insertQuery := config.INSERT_ORDERS_QUERY + combinedQueryVal

	res := config.DB.Exec(insertQuery,valueParams...)
	if res.Error!=nil{
		log.Fatal("Error during insert to orders : ",res.Error)
	}
	log.Printf("Insert into orders success")
}


func InsertOrderDetail(){
	words := ReadFile("csvdata/OrderDetails.csv")
	stringInserted := []string{}
	var valueParams []interface{}
	for _,word := range words[1:]{
		stringInserted = append(stringInserted, "(?, ?, ?, ?, ?, ?)")
		a := strings.Split(word,";")

		orderDetailID := ConvertInt(a[0])
		orderID := ConvertInt(a[1])
		productID := ConvertInt(a[2])
		quantity := ConvertInt(a[3])
		unitPrice := ConvertFloat(a[4])
		discount := ConvertFloat(strings.ReplaceAll(a[5],"%",""))

		valueParams = append(valueParams,orderDetailID,orderID,productID,quantity,unitPrice,discount)
	}

	combinedQueryVal := strings.Join(stringInserted,",")
	insertQuery := config.INSERT_ORDER_DETAILS_QUERY + combinedQueryVal

	res := config.DB.Exec(insertQuery,valueParams...)
	if res.Error!=nil{
		log.Fatal("Error during insert to order details : ",res.Error)
	}
	log.Printf("Insert into orders details success")
}


func main(){
	config.InitDB()
	InsertCustomers()
	InsertEmployee()
	InsertShippingMethod()
	InsertProducts()
	InsertOrders()
	InsertOrderDetail()
}

