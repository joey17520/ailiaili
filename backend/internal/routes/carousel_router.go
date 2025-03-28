package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joey17520/ailiaili/internal/api/v1"
	"github.com/joey17520/ailiaili/internal/middleware"
)

func CollectCarouselRoutes(r *gin.RouterGroup) {
	carouselGroup := r.Group("carousel")

	carouselAuth := carouselGroup.Group("")
	carouselAuth.Use(middleware.Auth())
	{
		// 新增轮播图
		carouselAuth.POST("addCarousel", api.AddCarousel)
		// 获取轮播图列表
		carouselAuth.POST("getCarouselList", api.GetCarouselList)
		// 编辑轮播图信息
		carouselAuth.PUT("editCarousel", api.EditCarousel)
		// 删除轮播图
		carouselAuth.DELETE("deleteCarousel/:id", api.DeleteCarousel)
	}

	carouselGroup.GET("getCarousel", api.GetCarousel)
}
