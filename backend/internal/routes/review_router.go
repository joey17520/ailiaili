package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joey17520/ailiaili/internal/api/v1"
	"github.com/joey17520/ailiaili/internal/middleware"
)

func CollectReviewRoutes(r *gin.RouterGroup) {
	reviewGroup := r.Group("review")

	reviewAuth := reviewGroup.Group("")
	reviewAuth.Use(middleware.Auth())
	{
		// 视频审核通过
		reviewAuth.POST("reviewVideoApproved", api.ReviewVideoApproved)
		// 视频审核不通过
		reviewAuth.POST("reviewVideoFailed", api.ReviewVideoFailed)
		// 获取视频审核记录
		reviewAuth.GET("getVideoReviewRecord", api.GetVideoReviewRecord)

		// 文章审核通过
		reviewAuth.POST("reviewArticleApproved", api.ReviewArticleApproved)
		// 文章审核不通过
		reviewAuth.POST("reviewArticleFailed", api.ReviewArticleFailed)
		// 获取文章审核记录
		reviewAuth.GET("getArticleReviewRecord", api.GetArticleReviewRecord)
	}
}
