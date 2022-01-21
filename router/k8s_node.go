package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitNodeRouter(Router *gin.RouterGroup) {
	NodeRouter := Router.Group("node").Use(middleware.OperationRecord())
	{
		NodeRouter.POST("createNode", v1.CreateNode)             // 新建Node
		NodeRouter.DELETE("deleteNode", v1.DeleteNode)           // 删除Node
		NodeRouter.DELETE("deleteNodeByIds", v1.DeleteNodeByIds) // 批量删除Node
		NodeRouter.PUT("updateNode", v1.UpdateNode)              // 更新Node
		NodeRouter.GET("findNode", v1.FindNode)                  // 根据ID获取Node
		NodeRouter.GET("getNodeList", v1.GetNodeList)            // 获取Node列表
	}
}
