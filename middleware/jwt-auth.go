package middleware

import (
	"net/http"

	"MyGO.com/m/helper"
	"MyGO.com/m/service"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(jwtService service.JwtService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			response := helper.ResponseErrorData(401, "No token found")
			ctx.JSON(http.StatusOK, response)
			return
		}
	}
}
