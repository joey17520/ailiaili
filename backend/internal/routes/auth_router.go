package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joey17520/ailiaili/internal/api/v1"
	"github.com/joey17520/ailiaili/internal/middleware"
)

func CollectAuthRoutes(r *gin.RouterGroup) {

	authGroup := r.Group("auth")

	authAuth := authGroup.Group("")
	authAuth.Use(middleware.Auth())
	{
		authAuth.POST("logout", api.Logout)
	}

	// 用户注册
	authGroup.POST("register", api.Register)
	// 用户登录(密码)
	authGroup.POST("login", api.Login)
	// 用户登录(邮箱)
	authGroup.POST("login/email", api.EmailLogin)
	// 更新token
	authGroup.POST("updateToken", api.UpdateToken)
	// 清除Cookie
	authGroup.POST("clearCookie", api.ClearCookie)
	// 修改密码检查
	authGroup.POST("resetpwdCheck", api.ResetPwdCheck)
	// 修改密码
	authGroup.POST("modifyPwd", api.ModifyPwd)
}
