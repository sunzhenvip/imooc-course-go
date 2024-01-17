package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mxshop-api/user-web/api"
	"mxshop-api/user-web/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	zap.S().Info("配置用户相关都URL")
	{
		UserRouter.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		UserRouter.POST("pwd_login", api.PassWordLogin)
		UserRouter.POST("test", api.Test)
	}
}

// {
//    "expired_at": 1661835885000,
//    "id": 21,
//    "nick_name": "bobby0",
//    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MjEsIk5pY2tOYW1lIjoiYm9iYnkwIiwiQXV0aG9yaXR5SWQiOjEsImV4cCI6MTY2MTgzNTg4NSwiaXNzIjoiaW1vb2MiLCJuYmYiOjE2NTkyNDM4ODV9.q0CIQU5QoWPggO6N4Uv0dyhcZAbBSCxh_hbW00tWDUE"
// }
