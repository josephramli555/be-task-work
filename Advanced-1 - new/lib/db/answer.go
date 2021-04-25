package db

import (
	"surveyapp/models"
	"surveyapp/config"
	"log"
)

func InsertAnswers(answers []models.Answer){

	res :=config.DB.Create(answers)
	if res.Error != nil{
		panic(res.Error)
	}
	if res.RowsAffected >0{
		log.Println("Insert User Answer success")
	}
}

func GetAnswer(searchId string) models.Answer{
	answer := models.Answer{}
	res := config.DB.First(&answer,searchId)
	if res.Error != nil{
		panic(res.Error)
	}
	return answer
}

func GetAllAnswers()(answers []models.Answer){
	res := config.DB.Order("answer_id").Find(&answers)
	if res.Error != nil{
		panic(res.Error)
	}
	return answers
}

func DeleteAnswer(id string){
	res := config.DB.Delete(&models.Answer{},id)
	if res.Error != nil{
		panic(res.Error)
	}
}

func UpdateAnswer(id string,newAnswer string){
	res := config.DB.Model(&models.Answer{}).Where("answer_id = ?",id).Update("answer",newAnswer)
	if res.Error != nil{
		panic(res.Error)
	}
}