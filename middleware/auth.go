package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		c, err := ctx.Request.Cookie("session_token")
		if err != nil {
			if ctx.GetHeader("Content-Type") != "application/json" {
				ctx.JSON(http.StatusSeeOther, model.ErrorResponse{Error: "Content-type undefined"})
				return
			}
			if err == http.ErrNoCookie {
				ctx.JSON(http.StatusUnauthorized, model.ErrorResponse{Error: err.Error()})
				return
			}
			ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
			return
		}

		tokenString := c.Value
		claims := &model.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

			return model.JwtKey, nil
		})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
			return
		}
		if !token.Valid {
			ctx.JSON(http.StatusUnauthorized, model.ErrorResponse{Error: err.Error()})
			return
		}
		ctx.Writer.WriteHeader(200)
		ctx.Set("email", claims.Email)
		ctx.Next()
	})
}
