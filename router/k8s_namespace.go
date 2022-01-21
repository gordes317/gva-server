package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitNamespaceRouter(Router *gin.RouterGroup) {
	NamespaceRouter := Router.Group("namespace").Use(middleware.OperationRecord())
	{
		NamespaceRouter.POST("createNamespace", v1.CreateNamespace)             // 新建Namespace
		NamespaceRouter.DELETE("deleteNamespace", v1.DeleteNamespace)           // 删除Namespace
		NamespaceRouter.DELETE("deleteNamespaceByIds", v1.DeleteNamespaceByIds) // 批量删除Namespace
		NamespaceRouter.PUT("updateNamespace", v1.UpdateNamespace)              // 更新Namespace
		NamespaceRouter.GET("findNamespace", v1.FindNamespace)                  // 根据ID获取Namespace
		NamespaceRouter.GET("getNamespaceList", v1.GetNamespaceList)            // 获取Namespace列表
	}
}
