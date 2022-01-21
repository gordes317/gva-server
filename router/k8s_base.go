package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitBaseAPIRouter(Router *gin.RouterGroup) {
	BaseAPIRouter := Router.Group("base").Use(middleware.OperationRecord())
	{

		BaseAPIRouter.GET("getSelectNamespaceList", v1.GetSelectNamespaceList) // 获取下拉Namespace列表

		BaseAPIRouter.GET("getClusterList", v1.GetClusterList) // 获取集群下拉列表

		BaseAPIRouter.GET("updateSwitchCluster", v1.UpdateSwitchCluster) // 切换集群

		BaseAPIRouter.POST("applyYamlCreate", v1.ApplyYamlCreate)                          //根据yaml文件创建对象
		BaseAPIRouter.GET("getClusterApplicationCounter", v1.GetClusterApplicationCounter) //获取集群资源计数
		BaseAPIRouter.GET("getClusterEvents", v1.GetClusterEvents)                         //获取集群Events

	}
}
