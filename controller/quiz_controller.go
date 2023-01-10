package controller

import (
	"net/http"
	"strconv"
	"exacta/backend/model/domain"
	"exacta/backend/model/web"	
	"github.com/gin-gonic/gin"
	"fmt"

)

func (api *API) GetCategories(c *gin.Context) {
	//go api.AllowOrigin(c)
	categories, err := api.quizRepo.FindCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err})
		return
	}

	var categoriesResponse []web.CategoryResponse

	for _, v := range categories {
		categorieResponse := convertToCategorieResponse(v)

		categoriesResponse = append(categoriesResponse, categorieResponse)
	}
	c.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    categoriesResponse,
	})
}

func (api *API) GetQuizByCategoryIdWithPagination(c *gin.Context) {
	//go api.AllowOrigin(c)
	categoryId, _ := strconv.Atoi(c.Query("category_id"))
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _:= strconv.Atoi(c.Query("limit"))
	questions, err := api.quizRepo.FindQuizByCategoryIdWithPagination(
		uint(categoryId), uint(page), uint(limit),
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err})
		return
		}
	
	category, err := api.quizRepo.FindCategoryById(uint(categoryId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err})
		return
		}
	
	var incorrectAnswers domain.IncorrectAnswerDomain
	for _, question := range questions {
		incorrectAnswerDomain, err := api.quizRepo.FindIncorrectAnswersByQuizId(question.Id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": err})
			return
			}
	
		incorrectAnswers = incorrectAnswerDomain
	}

	QuizResponse := convertToQuizResponses(questions, incorrectAnswers, category)

	c.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    QuizResponse,
	})
}	

func (api *API) SubmitAnswersAttempts(c *gin.Context) {
	//go api.AllowOrigin(c)
	var namaSekolah string
	userId, _ := GetUserId()
	users, _ := api.usersRepo.FetchUsers()
	for _, user := range users{
		if user.Id == userId{
			namaSekolah = user.NamaSekolah
		}
	}
	
	var answerAttemptReq web.AnswerAttemptRequest

	if err := c.ShouldBindJSON(&answerAttemptReq); err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Data:    err.Error(),
		})
		return
	}

	err := answerAttemptReq.ValidateAnswerAttempt()
	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, web.WebResponse{
				Code:    http.StatusUnprocessableEntity,
				Message: http.StatusText(http.StatusUnprocessableEntity),
				Data:    err.Error(),
			})
			return
		}
	}

	var answersAttempt []domain.AnswerAttemptDomain

	for _, answer := range answerAttemptReq.Answers {
		answersAttempt = append(answersAttempt, domain.AnswerAttemptDomain{
			Answer: answer.Answer,
			QuizId: answer.QuizId, 
			UserId: userId,
			NamaSekolah: namaSekolah, 
		})
	}

	_, err = api.quizRepo.SaveAnswerAttempt(userId, answersAttempt, namaSekolah)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = api.quizRepo.SaveResult(
		answerAttemptReq.Duration, 
		userId, 
		answerAttemptReq.CategoryId, 
		namaSekolah,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}


	result, err := api.quizRepo.FindResultByCategoryId(answerAttemptReq.CategoryId, userId, namaSekolah)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println("result", result)

	resultResp := convertToAnswersAttemptResponse(result)
	fmt.Println("resultResp", resultResp)


	c.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    resultResp,
	})
}

func (api *API) GetScoresBoardByCategoryId(c *gin.Context) {
	//go api.AllowOrigin(c)
	categoryId, _ := strconv.Atoi(c.Query("category_id"))

	usersResp,err := api.usersRepo.FetchUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    err.Error(),
		})
		return
	}
	var users []domain.UserDomain
	for _, userResp := range usersResp{
		user, err := api.usersRepo.FetchUserByID(userResp.Id)
		if err != nil {			
			c.JSON(http.StatusInternalServerError, web.WebResponse{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
				Data:    err.Error(),
			})
			return
		}

		users = append(users, user)
	}
	scoreBoardResponse, err := api.quizRepo.FindScoresBoardByCategoryId(uint(categoryId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    err.Error(),
		})
		return
	}
	scoresBoard := convertToScoreBoardResponses(users, scoreBoardResponse)

	c.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    scoresBoard,
	})
}	



func convertToCategorieResponse(c domain.CategoryDomain) web.CategoryResponse {
	return web.CategoryResponse{
		Id 			: c.Id,
		Name 		: c.Name,
		Description : c.Description,
	}
}

func convertToQuizResponses(questions []domain.QuizDomain, incorrectAnswer domain.IncorrectAnswerDomain, category domain.CategoryDomain) []web.QuizResponse {
	var incorrectAnswerResponses []string
	incorrectAnswerResponses = append(
		incorrectAnswerResponses, incorrectAnswer.OptionOne, incorrectAnswer.OptionTwo,
	)

	var questionResponses []web.QuizResponse
	for _, question := range questions {
		questionResponses = append(questionResponses, web.QuizResponse{
			Id:               question.Id,
			Category:         category.Name,
			Question:         question.Question,
			CorrectAnswer:    question.CorrectAnswer,
			IncorrectAnswers: incorrectAnswerResponses,
		})
	}

	return questionResponses
}

func convertToAnswersAttemptResponse(result domain.ResultDomain) web.AnswerAttemptResponse {
	return web.AnswerAttemptResponse{
		Correct:  result.Correct,
		Wrong:    result.Wrong,
		Duration: result.Duration,
	}
}

func convertToScoreBoardResponses(users []domain.UserDomain, scoresBoard []domain.ResultDomain) []web.ScoreBoardResponse {
	var scoreBoardResponses []web.ScoreBoardResponse

	for _, scoreBoard := range scoresBoard {
		var username string
		var namaSekolah string
		for _, user := range users {
			if user.Id == scoreBoard.UserId {
				username = user.Username
				namaSekolah = user.NamaSekolah
			}
		}

		scoreBoardResponses = append(scoreBoardResponses, web.ScoreBoardResponse{
			Username: username,
			Score:    scoreBoard.Correct * 10,
			Duration: scoreBoard.Duration,
			NamaSekolah: namaSekolah,

		})
	}

	return scoreBoardResponses
}
