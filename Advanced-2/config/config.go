package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sales-api/modules"
)

var DB *gorm.DB


func InitDB(){
	var err error
	new(modules.Configuration).Init()
	conf := new(modules.Configuration).GetConfig()
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.GetString("DB_USER"), conf.GetString("DB_PASS"), conf.GetString("DB_HOST"), conf.GetString("DB_PORT"), conf.GetString("DATABASE"))
	
	DB,err = gorm.Open(mysql.Open(connection),&gorm.Config{})
	
	if err!=nil{
		log.Fatal(err)
	}

	queries := []string{
		CREATE_CUSTOMER_TABLE,
		CREATE_EMPLOYEES_TABLE,
		CREATE_SHIPPING_METHODS_TABLE,
		CREATE_PRODUCTS_TABLE,
		CREATE_ORDERS_TABLE,
		CREATE_ORDER_DETAILS_TABLE,
	}

	for _,query := range queries{
		res := DB.Exec(query)
		if res.Error != nil {
			panic(err)
		}
	}

}