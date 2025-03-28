package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joey17520/ailiaili/internal/api/v1"
	"github.com/joey17520/ailiaili/internal/middleware"
)

func CollectRelationRoutes(r *gin.RouterGroup) {
	relationGroup := r.Group("relation")

	relationAuth := relationGroup.Group("")
	relationAuth.Use(middleware.Auth())
	{
		// 关注
		relationAuth.POST("follow", api.Follow)
		// 取消关注
		relationAuth.POST("unfollow", api.Unfollow)
		// 获取用户关系
		relationAuth.GET("getUserRelation", api.GetUserRelation)
	}

	// 获取关注列表
	relationGroup.GET("getFollowings", api.GetFollowings)
	// 获取粉丝列表
	relationGroup.GET("getFollowers", api.GetFollowers)
	// 获取关注和粉丝数
	relationGroup.GET("getFollowCount", api.GetFollowCount)
}
