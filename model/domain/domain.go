package domain

import "time"

type UserDomain struct {
	Id        	uint 	`json:"user_id"`
	Username  	string 	`json:"username"`
	Email     	string 	`json:"email"`
	NamaSekolah string 	`json: "nama_sekolah"`
	Password  	string 	`json:"password"`
}

type CategoryDomain struct {
	Id          uint
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type QuizDomain struct {
	Id            uint
	CategoryId    uint
	Question      string
	CorrectAnswer string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type IncorrectAnswerDomain struct {
	Id        uint
	QuizId    uint
	OptionOne string
	OptionTwo string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AnswerAttemptDomain struct {
	Id        	uint
	Answer    	string
	QuizId    	uint
	UserId    	uint
	NamaSekolah string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ResultDomain struct {
	Id         	uint `json:"id"`
	Correct    	uint	`json:"correct"`
	Wrong      	uint	`json:"wrong"`
	Duration   	string	`json:"duration"`
	NamaSekolah string	`json:"nama_sekolah"`
	UserId    	uint	`json:"user_id"`
	CategoryId 	uint	`json:"category_id"`
	CreatedAt  time.Time	`json:"created_at"`
	UpdatedAt  time.Time	`json:"updated_at"`
}
