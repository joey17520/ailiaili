package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joey17520/ailiaili/internal/api/v1"
	"github.com/joey17520/ailiaili/internal/middleware"
)

func CollectPartitionRoutes(route *gin.RouterGroup) {
	partitionGroup := route.Group("partition")

	//获取分区列表
	partitionGroup.GET("getPartitionList", api.GetPartitionList)

	partitionAuth := partitionGroup.Use(middleware.Auth())
	{
		//添加分区
		partitionAuth.POST("addPartition", api.AddPartition)
		//删除分区
		partitionAuth.DELETE("deletePartition/:id", api.DeletePartition)
	}

}
