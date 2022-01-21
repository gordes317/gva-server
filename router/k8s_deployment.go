package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitDeploymentRouter(Router *gin.RouterGroup) {
	DeploymentRouter := Router.Group("deployment").Use(middleware.OperationRecord())
	{
		DeploymentRouter.POST("createDeployment", v1.CreateDeployment)             // 新建Deployment
		DeploymentRouter.DELETE("deleteDeployment", v1.DeleteDeployment)           // 删除Deployment
		DeploymentRouter.DELETE("deleteDeploymentByIds", v1.DeleteDeploymentByIds) // 批量删除Deployment
		DeploymentRouter.PUT("updateDeployment", v1.UpdateDeployment)              // 更新Deployment 副本
		DeploymentRouter.GET("findDeployment", v1.FindDeployment)                  // 根据ID获取Deployment
		DeploymentRouter.GET("getDeploymentList", v1.GetDeploymentList)            // 获取Deployment列表
		DeploymentRouter.GET("readYamlDeployment", v1.ReadYamlDeployment)          // 获取Deployment Yaml
		DeploymentRouter.PUT("applyYamlDeployment", v1.ApplyYamlDeployment)        // 更新Deployment Yaml
		DeploymentRouter.PUT("restartDeployment", v1.RestartDeployment)            // 重启Deployment
	}
}
