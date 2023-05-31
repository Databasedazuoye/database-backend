package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	utils "goodsManagement/utils"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 验证请求
		token := c.GetHeader("Authorization")
		fmt.Println(token)
		flag, username, id := utils.ParseToken(token)
		fmt.Println("Authorization")
		fmt.Println(flag, username, id)

		if !flag {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "unauthorized",
			})
			return
		}

		c.Set("username", username)
		c.Set("id", id)
		// 继续处理请求
		c.Next()
	}
}
