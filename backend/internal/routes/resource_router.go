package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joey17520/ailiaili/internal/api/v1"
	"github.com/joey17520/ailiaili/internal/middleware"
)

func CollectResourceRoutes(r *gin.RouterGroup) {
	resourceGroup := r.Group("resource")

	resourceAuth := resourceGroup.Group("")
	resourceAuth.Use(middleware.Auth())
	{
		resourceAuth.PUT("modifyTitle", api.ModifyResourceTitle)
		resourceAuth.DELETE("deleteResource/:id", api.DeleteResource)
	}
}
