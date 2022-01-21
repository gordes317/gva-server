package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitDaemonSetRouter(Router *gin.RouterGroup) {
	DaemonSetRouter := Router.Group("daemonset").Use(middleware.OperationRecord())
	{
		DaemonSetRouter.POST("createDaemonSet", v1.CreateDaemonSet)             // 新建DaemonSet
		DaemonSetRouter.DELETE("deleteDaemonSet", v1.DeleteDaemonSet)           // 删除DaemonSet
		DaemonSetRouter.DELETE("deleteDaemonSetByIds", v1.DeleteDaemonSetByIds) // 批量删除DaemonSet
		DaemonSetRouter.PUT("updateDaemonSet", v1.UpdateDaemonSet)              // 更新DaemonSet
		DaemonSetRouter.GET("findDaemonSet", v1.FindDaemonSet)                  // 根据ID获取DaemonSet
		DaemonSetRouter.GET("getDaemonSetList", v1.GetDaemonSetList)            // 获取DaemonSet列表
		DaemonSetRouter.GET("readYamlDaemonSet", v1.ReadYamlDaemonSet)          // 获取DaemonSet Yaml
		DaemonSetRouter.PUT("applyYamlDaemonSet", v1.ApplyYamlDaemonSet)        // 更新DaemonSet Yaml
	}
}
