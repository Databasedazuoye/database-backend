package middleware

import (
	"github.com/gin-gonic/gin"
	utils "goodsManagement/util"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 验证请求
		token := c.GetHeader("Authorization")
		if utils.ParseToken(token) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		// 继续处理请求
		c.Next()
	}
}
