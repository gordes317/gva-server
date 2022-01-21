package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitMultiClusterRouter(Router *gin.RouterGroup) {
	MultiClusterRouter := Router.Group("multicluster").Use(middleware.OperationRecord())
	{
		MultiClusterRouter.POST("createMultiCluster", v1.CreateMultiCluster)             // 新建MultiCluster
		MultiClusterRouter.DELETE("deleteMultiCluster", v1.DeleteMultiCluster)           // 删除MultiCluster
		MultiClusterRouter.DELETE("deleteMultiClusterByIds", v1.DeleteMultiClusterByIds) // 批量删除MultiCluster
		MultiClusterRouter.PUT("updateMultiCluster", v1.UpdateMultiCluster)              // 更新MultiCluster
		MultiClusterRouter.GET("findMultiCluster", v1.FindMultiCluster)                  // 根据ID获取MultiCluster
		MultiClusterRouter.GET("getMultiClusterList", v1.GetMultiClusterList)            // 获取MultiCluster列表
	}
}
