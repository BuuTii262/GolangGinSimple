package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"MyGO.com/m/helper"
	"MyGO.com/m/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(jwtService service.JwtService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			response := helper.ResponseErrorData(401, "No token found")
			ctx.JSON(http.StatusOK, response)
			ctx.Abort()
			return
		}

		splitToken := strings.Split(authHeader, "Bearer ")
		authHeader = splitToken[1]
		token, err := jwtService.ValidateToken(authHeader)
		if err != nil {
			fmt.Println("Here have error in Middle warre", err.Error())
			response := helper.ResponseErrorData(401, err.Error())
			ctx.JSON(http.StatusOK, response)
			ctx.Abort()
			return
		}

		if !token.Valid {
			response := helper.ResponseErrorData(401, "Token is not valid")
			ctx.JSON(http.StatusOK, response)
			ctx.Abort()
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		log.Println("Claim[user_id]: ", claims["user_id"])

		ctx.Next()

	}
}
