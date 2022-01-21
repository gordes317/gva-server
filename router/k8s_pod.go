package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitPodRouter(Router *gin.RouterGroup) {
	PodRouter := Router.Group("pod").Use(middleware.OperationRecord())
	{
		PodRouter.POST("createPod", v1.CreatePod)             // 新建Pod
		PodRouter.DELETE("deletePod", v1.DeletePod)           // 删除Pod
		PodRouter.DELETE("deletePodByIds", v1.DeletePodByIds) // 批量删除Pod
		PodRouter.PUT("updatePod", v1.UpdatePod)              // 更新Pod
		PodRouter.GET("findPod", v1.FindPod)                  // 根据ID获取Pod
		PodRouter.GET("getPodList", v1.GetPodList)            // 获取Pod列表
		PodRouter.GET("getPodLog", v1.GetPodLog)              //获取Pod log信息
		PodRouter.GET("readYamlPod", v1.ReadYamlPod)          // 获取Pod Yaml
		PodRouter.PUT("applyYamlPod", v1.ApplyYamlPod)        // 更新Pod Yaml
	}
}
