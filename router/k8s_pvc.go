package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitPVCRouter(Router *gin.RouterGroup) {
	PVCRouter := Router.Group("pvc").Use(middleware.OperationRecord())
	{
		PVCRouter.POST("createPVC", v1.CreatePVC)             // 新建PVC
		PVCRouter.DELETE("deletePVC", v1.DeletePVC)           // 删除PVC
		PVCRouter.DELETE("deletePVCByIds", v1.DeletePVCByIds) // 批量删除PVC
		PVCRouter.PUT("updatePVC", v1.UpdatePVC)              // 更新PVC
		PVCRouter.GET("findPVC", v1.FindPVC)                  // 根据ID获取PVC
		PVCRouter.GET("getPVCList", v1.GetPVCList)            // 获取PVC列表
	}
}
