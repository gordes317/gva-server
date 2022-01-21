package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitIngressRouter(Router *gin.RouterGroup) {
	IngressRouter := Router.Group("ingress").Use(middleware.OperationRecord())
	{
		IngressRouter.POST("createIngress", v1.CreateIngress)             // 新建Ingress
		IngressRouter.DELETE("deleteIngress", v1.DeleteIngress)           // 删除Ingress
		IngressRouter.DELETE("deleteIngressByIds", v1.DeleteIngressByIds) // 批量删除Ingress
		IngressRouter.PUT("updateIngress", v1.UpdateIngress)              // 更新Ingress
		IngressRouter.GET("findIngress", v1.FindIngress)                  // 根据ID获取Ingress
		IngressRouter.GET("getIngressList", v1.GetIngressList)            // 获取Ingress列表
		IngressRouter.GET("readYamlIngress", v1.ReadYamlIngress)          // 获取Ingress Yaml
		IngressRouter.PUT("applyYamlIngress", v1.ApplyYamlIngress)        // 更新Ingress Yaml
	}
}
