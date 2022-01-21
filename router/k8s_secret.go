package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitSecretRouter(Router *gin.RouterGroup) {
	SecretRouter := Router.Group("secret").Use(middleware.OperationRecord())
	{
		SecretRouter.POST("createSecret", v1.CreateSecret)             // 新建Secret
		SecretRouter.DELETE("deleteSecret", v1.DeleteSecret)           // 删除Secret
		SecretRouter.DELETE("deleteSecretByIds", v1.DeleteSecretByIds) // 批量删除Secret
		SecretRouter.PUT("updateSecret", v1.UpdateSecret)              // 更新Secret
		SecretRouter.GET("findSecret", v1.FindSecret)                  // 根据ID获取Secret
		SecretRouter.GET("getSecretList", v1.GetSecretList)            // 获取Secret列表
		SecretRouter.GET("readYamlSecret", v1.ReadYamlSecret)          // 获取Secret Yaml
		SecretRouter.PUT("applyYamlSecret", v1.ApplyYamlSecret)        // 更新Secret Yaml
	}
}
