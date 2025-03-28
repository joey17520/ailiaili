package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joey17520/ailiaili/internal/api/v1"
	"github.com/joey17520/ailiaili/internal/middleware"
)

func CollectApiRoutes(r *gin.RouterGroup) {

	apiGroup := r.Group("api")
	apiGroup.Use(middleware.Auth())
	{
		// 获取API列表
		apiGroup.POST("getApiList", api.GetApiList)
		// 获取全部API列表
		apiGroup.GET("getAllApiList", api.GetAllApiList)
		// 新增API
		apiGroup.POST("addApi", api.AddApi)
		// 编辑API
		apiGroup.PUT("editApi", api.EditApi)
		// 删除API
		apiGroup.DELETE("deleteApi/:id", api.DeleteApi)
		// 获取角色API
		apiGroup.GET("getRoleApi", api.GetRoleApi)
		// 编辑角色API
		apiGroup.PUT("editRoleApi", api.EditRoleApi)
	}
}
