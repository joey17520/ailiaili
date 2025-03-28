package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joey17520/ailiaili/internal/api/v1"
	"github.com/joey17520/ailiaili/internal/middleware"
)

func CollectCollectionRoutes(r *gin.RouterGroup) {
	collectionGroup := r.Group("collection")

	collectionAuth := collectionGroup.Group("")
	collectionAuth.Use(middleware.Auth())
	{
		// 获取收藏夹列表
		collectionAuth.GET("getCollectionList", api.GetCollectionList)
		// 获取收藏夹信息
		collectionAuth.GET("getCollectionInfo", api.GetCollectionInfo)
		// 添加收藏夹
		collectionAuth.POST("addCollection", api.AddCollection)
		// 获取收藏夹的视频列表
		collectionAuth.GET("getVideoList", api.GetCollectVideo)
		// 修改收藏夹
		collectionAuth.PUT("editCollection", api.EditCollection)
		// 删除收藏夹
		collectionAuth.DELETE("deleteCollection/:id", api.DeleteCollection)
	}
}
