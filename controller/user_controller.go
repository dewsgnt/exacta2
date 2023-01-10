package controller

import(
	_"exacta/backend/repository"
	"exacta/backend/model/web"
	"net/http"
	"fmt"
	"time"

	validator"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mattn/go-sqlite3"	
	"github.com/gin-gonic/gin"

)

type User struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Id     uint    `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

type Token struct {
	UserId    uint       `json:"user_id" binding:"required"`
	Token     string    `json:"token" binding:"required"`
	ExpiresAt time.Time `json:"expires_at" binding:"required"`
}

type AuthErrorResponse struct {
	Error string `json:"error"`
}

//jwt key for signature
var jwtKey = []byte("secret")

type Claims struct {
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

func (api *API) PostUserRegist(c *gin.Context) {
	 //go api.AllowOrigin(c)
	var user web.RegisterRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	err := user.ValidateRegister()
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: http.StatusText(http.StatusUnprocessableEntity),
			Data:    err.Error(),
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)

	res, err := api.usersRepo.InsertUser(user.Username, user.Email, string(hashedPassword), user.NamaSekolah)
	if res == nil && err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusBadRequest,
			"message":     "Email Has Registered!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":     "registered!",
	})

}

func (api *API) LoginUser(c *gin.Context) {
	//go api.AllowOrigin(c)
	var user web.LoginRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := user.ValidateLogin()
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: http.StatusText(http.StatusUnprocessableEntity),
			Data:    err.Error(),
		})
		return
	}

	pass, err := api.usersRepo.GetPasswordCompare(user.Email)

	if err != nil {
		if err == sqlite3.ErrConstraintRowID {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "e-mail is not registered!"})
		return
	}

	compareVal := bcrypt.CompareHashAndPassword([]byte(*pass), []byte(user.Password))
	if compareVal != nil {		
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "wrong password/password too short"})
		return
	}

	res, err := api.usersRepo.LoginUser(user.Email, *pass)

	c.Header("Content-Type", "application/json")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		UserId: *res,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	//encode claim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	//push token to db auth when user logged in
	var pushtoken Token

	//fetchUserId, _ := api.usersRepo.FetchUserIdByEmail(*res)

	pushtoken = Token{
		UserId:    *res,
		Token:     tokenString,
		ExpiresAt: expirationTime,
	}

	tknToDb, err := api.usersRepo.PushToken(pushtoken.UserId, pushtoken.Token, pushtoken.ExpiresAt)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
		Path:    ("/"),
	})

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":      "success",
		"data": LoginResponse{
			Id:     *res,
			Token:     *tknToDb,
			ExpiresAt: expirationTime,
		},
	})
}

func (api *API) LogoutUser(c *gin.Context) {
	// go api.AllowOrigin(c)
	userId, err := GetUserId()
	if err != nil{
		c.JSON(http.StatusUnauthorized, web.WebResponse{
			Code : http.StatusUnauthorized,
			Message : http.StatusText(http.StatusUnauthorized),
			Data: err.Error(),
		})
		return
	}
	api.usersRepo.DeleteToken(userId)

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Expires: time.Unix(0, 0),
		MaxAge:  -1,
	})

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "logged out!",
	})
}


