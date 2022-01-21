package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitStorageClassRouter(Router *gin.RouterGroup) {
	StorageClassRouter := Router.Group("storageclass").Use(middleware.OperationRecord())
	{
		StorageClassRouter.POST("createStorageClass", v1.CreateStorageClass)             // 新建StorageClass
		StorageClassRouter.DELETE("deleteStorageClass", v1.DeleteStorageClass)           // 删除StorageClass
		StorageClassRouter.DELETE("deleteStorageClassByIds", v1.DeleteStorageClassByIds) // 批量删除StorageClass
		StorageClassRouter.PUT("updateStorageClass", v1.UpdateStorageClass)              // 更新StorageClass
		StorageClassRouter.GET("findStorageClass", v1.FindStorageClass)                  // 根据ID获取StorageClass
		StorageClassRouter.GET("getStorageClassList", v1.GetStorageClassList)            // 获取StorageClass列表
	}
}
