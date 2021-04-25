package models

type Answer struct {
	Answer_id   int       `gorm:"primarykey;not null;AUTO_INCREMENT"`
	Answer string    `gorm:"type:varchar(255);not null"`
	Question string   `gorm:"type:varchar(255);not null"`
	User   string       `gorm:"type:varchar(255);not null"`
}

func NewAnswer(user string,question string,answer string) Answer{
	return Answer{
		Answer: answer,
		Question: question,
		User : user,
	}
}