package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitTemplateRouter(Router *gin.RouterGroup) {
	TemplateRouter := Router.Group("template").Use(middleware.OperationRecord())
	{
		TemplateRouter.POST("createTemplate", v1.CreateTemplate)             // 新建Template
		TemplateRouter.DELETE("deleteTemplate", v1.DeleteTemplate)           // 删除Template
		TemplateRouter.DELETE("deleteTemplateByIds", v1.DeleteTemplateByIds) // 批量删除Template
		TemplateRouter.PUT("updateTemplate", v1.UpdateTemplate)              // 更新Template
		TemplateRouter.GET("findTemplate", v1.FindTemplate)                  // 根据ID获取Template
		TemplateRouter.GET("getTemplateList", v1.GetTemplateList)            // 获取Template列表
	}
}
