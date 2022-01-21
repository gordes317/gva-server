package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitStatefulSetRouter(Router *gin.RouterGroup) {
	StatefulSetRouter := Router.Group("statefulset").Use(middleware.OperationRecord())
	{
		StatefulSetRouter.POST("createStatefulSet", v1.CreateStatefulSet)             // 新建StatefulSet
		StatefulSetRouter.DELETE("deleteStatefulSet", v1.DeleteStatefulSet)           // 删除StatefulSet
		StatefulSetRouter.DELETE("deleteStatefulSetByIds", v1.DeleteStatefulSetByIds) // 批量删除StatefulSet
		StatefulSetRouter.PUT("updateStatefulSet", v1.UpdateStatefulSet)              // 更新StatefulSet
		StatefulSetRouter.GET("findStatefulSet", v1.FindStatefulSet)                  // 根据ID获取StatefulSet
		StatefulSetRouter.GET("getStatefulSetList", v1.GetStatefulSetList)            // 获取StatefulSet列表
		StatefulSetRouter.GET("readYamlStatefulSet", v1.ReadYamlStatefulSet)          // 获取StatefulSet Yaml
		StatefulSetRouter.PUT("applyYamlStatefulSet", v1.ApplyYamlStatefulSet)        // 更新StatefulSet Yaml
	}
}
