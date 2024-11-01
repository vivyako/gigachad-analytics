package main

import (
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiKey := ctx.GetHeader("X-API-KEY")
		if apiKey == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		ctx.Next()
	}
}

func main() {
	router := gin.Default()

	authGroup := router.Group("/api")
	authGroup.Use(Auth())
	{
		authGroup.GET("/data", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "Authenticated and authorized!"})
		})
	}

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "ХОХЛЫ БУДУТ НАКАЗАНЫ")
	})

	router.Run(":8082")
}
