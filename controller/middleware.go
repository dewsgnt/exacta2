package controller
import (
	"context"
	"log"
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"exacta/backend/model/web"
	"strings"

)
const(
	authorizationHeader = "Authorization"
	userIdKey = "user_id"
)

var userIdVal uint

func (h *API) AllowOrigin(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}

func (m *API) AuthMiddleware(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		//m.AllowOrigin(c)
		header := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
		if header == " "{
			c.JSON(http.StatusUnauthorized, web.WebResponse{
				Code : http.StatusUnauthorized,
				Message : "Unauthorized",
				Data : "Need Authorization header",
			})
			return
		}
		userId, err := ParseToken(c, header)
		if err != nil {
			c.JSON(http.StatusUnauthorized, web.WebResponse{
				Code : http.StatusUnauthorized,
				Message : "Unauthorized",
				Data : "Need Authorization header",
			})
			return
		}
		ctx := context.WithValue(c.Request.Context(), userIdKey, userId)
		userIdVal = ctx.Value(userIdKey).(uint)


		// token, err := c.Request.Cookie("token")
		// fmt.Println("isi tkn", token,"ini gaada?")
		// if err != nil {
		// 	if err == http.ErrNoCookie {
		// 		// If no cookie is present, return unauthorized
		// 		c.JSON(http.StatusUnauthorized, gin.H{"Error4": err.Error()})
		// 		return
		// 	}
		// 	//no token field, return bad request
		// 	c.JSON(http.StatusBadRequest, gin.H{"Error5": err.Error()})
		// 	return
		// }

		// tokenString := token.Value

		// claims := &Claims{}

		// // Parse the JWT string and store the result in `claims`.
		// tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// 	return jwtKey, nil
		// })

		// if err != nil {
		// 	if err == jwt.ErrSignatureInvalid {
		// 		//signature invalid
		// 		c.JSON(http.StatusUnauthorized, gin.H{"Error1": err.Error()})
		// 		return
		// 	}
		// 	c.JSON(http.StatusBadRequest, gin.H{"Error2": err.Error()})
		// 	return
		// }

		// if !tkn.Valid {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"Error3": err.Error()})
		// 	c.Abort()
		// 	return
		// }
		// ctx := context.WithValue(c.Request.Context(), "email", claims.Email)
		// ctx = context.WithValue(ctx, "props", claims)
		// c.Request = c.Request.WithContext(ctx)

		next(c)

	}
}
func ParseToken(ctx context.Context, token string) (uint, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := tkn.Claims.(*Claims)
	if !ok || !tkn.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	newCtx := context.WithValue(ctx, "user_id", claims.UserId)
	_ = context.WithValue(newCtx, "props", claims)

	return claims.UserId, nil
}


func (api *API) GET(next gin.HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		//api.AllowOrigin(ctx)
		log.Print("GET: ", http.MethodGet)
		log.Print("Request method: ", ctx.Request.Method)
		if ctx.Request.Method != http.MethodGet {
			ctx.JSON(http.StatusMethodNotAllowed, gin.H{"Error": "Need GET Method!"})
			return
		}
		next(ctx)
	})
}

func (api *API) POST(next gin.HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		//api.AllowOrigin(ctx)
		log.Print("POST: ", http.MethodPost)
		log.Print("Request method: ", ctx.Request.Method)
		if ctx.Request.Method != http.MethodPost {
			ctx.JSON(http.StatusMethodNotAllowed, gin.H{"Error": "Need POST Method!"})
			return
		}
		next(ctx)
	})
}

func (api *API) DELETE(next gin.HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		//api.AllowOrigin(ctx)
		log.Print("Delete: ", http.MethodDelete)
		log.Print("Request method: ", ctx.Request.Method)
		if ctx.Request.Method != http.MethodDelete {
			ctx.JSON(http.StatusMethodNotAllowed, gin.H{"Error": "Need DELETE Method!"})
			return
		}
		next(ctx)
	})
}

func (api *API) PATCH(next gin.HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		//api.AllowOrigin(ctx)
		log.Print("PATCH: ", http.MethodPatch)
		log.Print("Request method: ", ctx.Request.Method)
		if ctx.Request.Method != http.MethodPatch {
			ctx.JSON(http.StatusMethodNotAllowed, gin.H{"Error": "Need PATCH Method!"})
			return
		}
		next(ctx)
	})
}

func GetUserId() (uint, error) {
	if userIdVal == 0 {
		return 0, fmt.Errorf("user id is not set")
	}

	return userIdVal, nil
}