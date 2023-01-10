package web

type CategoryResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type QuizResponse struct {
	Id               uint     `json:"id"`
	Category         string   `json:"category"`
	Question         string   `json:"question"`
	CorrectAnswer    string   `json:"correct_answer"`
	IncorrectAnswers []string `json:"incorrect_answers"`
}

type AnswerAttemptResponse struct {
	Correct  uint   `json:"correct"`
	Wrong    uint   `json:"wrong"`
	Duration string `json:"duration"`
}

type ScoreBoardResponse struct {
	Username 	string 	`json:"username"`
	Score    	uint   	`json:"score"`
	Duration 	string 	`json:"duration"`
	NamaSekolah string 	`json:"nama_sekolah"`
}
