package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"goodsManagement/utils"
)

func Test(c *gin.Context) {
	get := utils.GetRedisHelper().Get(context.Background(), "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwidXNlcklkIjoxLCJleHAiOjE2ODU2MjQ4OTV9.ju555laSDLr907aNMChxz6sYQImRy2eWfqFBjzy3O-s")
	fmt.Println(get)
}
