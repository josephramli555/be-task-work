package config

import (
	"fmt"
	"surveyapp/models"
	"surveyapp/modules"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var QuestionList = []string{
	"Q1: What is your favourite food?",
	"Q2: What is your favourite hero?",
	"Q3: What is your No.1 holiday destination?",
}

const (
	//redeclared because there is some bug, when testing that program cannot take
	//configuration from conf folder
	DB_HOST       = "localhost"
	DB_PORT       = "3306"
	DB_USER       = "root"
	DB_PASS       = ""
	DATABASE      = "db_surveyapp"
	DATABASE_TEST = "db_testsurvey"
)

func InitDB() {
	var err error
	new(modules.Configuration).Init()
	conf := new(modules.Configuration).GetConfig()
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.GetString("DB_USER"), conf.GetString("DB_PASS"), conf.GetString("DB_HOST"), conf.GetString("DB_PORT"), conf.GetString("DATABASE"))

	DB, err = gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&models.Answer{})
}

func InitDBTest() {
	var err error
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DB_USER, DB_PASS, DB_HOST, DB_PORT, DATABASE_TEST)

	DB, err = gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.Answer{})
	DB.Where("1 = 1").Delete(&models.Answer{})
}
