package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitServiceRouter(Router *gin.RouterGroup) {
	ServiceRouter := Router.Group("service").Use(middleware.OperationRecord())
	{
		ServiceRouter.POST("createService", v1.CreateService)             // 新建Service
		ServiceRouter.DELETE("deleteService", v1.DeleteService)           // 删除Service
		ServiceRouter.DELETE("deleteServiceByIds", v1.DeleteServiceByIds) // 批量删除Service
		ServiceRouter.PUT("updateService", v1.UpdateService)              // 更新Service
		ServiceRouter.GET("findService", v1.FindService)                  // 根据ID获取Service
		ServiceRouter.GET("getServiceList", v1.GetServiceList)            // 获取Service列表
		ServiceRouter.GET("readYamlService", v1.ReadYamlService)          // 获取Service Yaml
		ServiceRouter.PUT("applyYamlService", v1.ApplyYamlService)        // 更新Service Yaml
	}
}
