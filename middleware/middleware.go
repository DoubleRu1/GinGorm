package middleware

import (
	"GinGormCRUD/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing autherized token"})
			ctx.Abort()
			return
		}
		name, err := utils.ParseJWT(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "JWt token is not right"})
			ctx.Abort()
			return
		}
		ctx.Set("name", name)
		ctx.Next()
	}
}
