package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitConfigMapRouter(Router *gin.RouterGroup) {
	ConfigMapRouter := Router.Group("configmap").Use(middleware.OperationRecord())
	{
		ConfigMapRouter.POST("createConfigMap", v1.CreateConfigMap)             // 新建ConfigMap
		ConfigMapRouter.DELETE("deleteConfigMap", v1.DeleteConfigMap)           // 删除ConfigMap
		ConfigMapRouter.DELETE("deleteConfigMapByIds", v1.DeleteConfigMapByIds) // 批量删除ConfigMap
		ConfigMapRouter.PUT("updateConfigMap", v1.UpdateConfigMap)              // 更新ConfigMap
		ConfigMapRouter.GET("findConfigMap", v1.FindConfigMap)                  // 根据ID获取ConfigMap
		ConfigMapRouter.GET("getConfigMapList", v1.GetConfigMapList)            // 获取ConfigMap列表
		ConfigMapRouter.GET("readYamlConfigMap", v1.ReadYamlConfigMap)          // 获取ConfigMap Yaml
		ConfigMapRouter.PUT("applyYamlConfigMap", v1.ApplyYamlConfigMap)        // 更新ConfigMap Yaml
	}
}
