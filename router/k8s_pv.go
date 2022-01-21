package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitPVRouter(Router *gin.RouterGroup) {
	PVRouter := Router.Group("pv").Use(middleware.OperationRecord())
	{
		PVRouter.POST("createPV", v1.CreatePV)             // 新建PV
		PVRouter.DELETE("deletePV", v1.DeletePV)           // 删除PV
		PVRouter.DELETE("deletePVByIds", v1.DeletePVByIds) // 批量删除PV
		PVRouter.PUT("updatePV", v1.UpdatePV)              // 更新PV
		PVRouter.GET("findPV", v1.FindPV)                  // 根据ID获取PV
		PVRouter.GET("getPVList", v1.GetPVList)            // 获取PV列表
	}
}
