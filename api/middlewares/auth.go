package middlewares

import (
	"errors"
	"net/http"
	"strings"
	"uploader/api/auth"

	"github.com/gin-gonic/gin"
)

// CheckAuth get token from payload and
// check is it valid or not.
func CheckAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtService := auth.JWTAuthService()
		header := ctx.Request.Header
		if authToken, ok := header["Authorization"]; ok {
			token := strings.Split(authToken[0], " ")[1]
			if t, err := jwtService.ValidateToken(token); err != nil {
				ctx.AbortWithError(http.StatusUnauthorized, err)
				return
			} else {
				claims := jwtService.GetMapClaims(t)
				ctx.Set("email", claims["email"])
				ctx.Next()
			}
		} else {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
				"error":   errors.New("request does not contain any token"),
				"message": "401 UnAuthorized",
			})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

func IsSuperUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtService := auth.JWTAuthService()
		header := ctx.Request.Header
		token := strings.Split(header["Authorization"][0], " ")[1]
		validatedToken, _ := jwtService.ValidateToken(token)
		jwtClaims := jwtService.GetMapClaims(validatedToken)
		if jwtClaims["isAdmin"] == true {
			ctx.Set("isAdmin", true)
		} else {
			ctx.Set("isAdmin", false)
			ctx.AbortWithError(http.StatusUnauthorized, http.ErrBodyNotAllowed)
		}
		ctx.Next()
	}
}
