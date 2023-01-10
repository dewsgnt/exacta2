package controller

import(
	"exacta/backend/repository"
	"github.com/gin-gonic/gin"
	"fmt"
	"time"
	"net/http"

	"github.com/gin-contrib/cors"

)

type API struct{
	usersRepo repository.UserRepositoryImpl
	quizRepo repository.QuizRepositoryImpl
	gin *gin.Engine
}


func NewAPI(usersRepo repository.UserRepositoryImpl, quizRepo repository.QuizRepositoryImpl) API{
	gin := gin.Default()
	api := API{
		usersRepo,
		quizRepo,
		gin,
	}
	//gin.Use(cors.Default())
	gin.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
  		AllowMethods:     []string{"PUT", "PATCH","GET","POST", "OPTIONS"},
		AllowHeaders:     []string{"origin","Authorization","Cookie", "Content-Type", "X-CSRF-Token"},
  		ExposeHeaders:    []string{"Content-Length","Authorization"},
  		 AllowCredentials: true,
  		AllowOriginFunc: func(origin string) bool {
   		return origin == "https://github.com"
  		},
  		MaxAge: 12 * time.Hour,
 	}))

	
	v1 := gin.Group("/api/v1")

	//users
	v1.POST("/users/regist", api.POST(api.PostUserRegist))
	v1.POST("/users/login", api.POST(api.LoginUser))
	v1.POST("/users/logout", api.POST(api.AuthMiddleware(api.LogoutUser)))

	//quiz
	v1.GET("/home/categories", api.GET(api.AuthMiddleware(api.GetCategories)))
	// v1.GET("/home/categories", api.GET((api.GetCategories)))

	v1.GET("/home/quizzes", api.GET(api.AuthMiddleware(api.GetQuizByCategoryIdWithPagination)))
	v1.POST("/home/submitanswer", api.POST(api.AuthMiddleware(api.SubmitAnswersAttempts)))
	v1.GET("/home/score-boards", api.GET(api.GetScoresBoardByCategoryId))

	return api
}

func (api *API) Handler() *gin.Engine {
	return api.gin
}
func (api *API) Start() {
	fmt.Println("starting web server at https://localhost:8080/")
	http.ListenAndServe(":8080", api.Handler())
}