package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitCronJobRouter(Router *gin.RouterGroup) {
	CronJobRouter := Router.Group("cronjob").Use(middleware.OperationRecord())
	{
		CronJobRouter.POST("createCronJob", v1.CreateCronJob)             // 新建CronJob
		CronJobRouter.DELETE("deleteCronJob", v1.DeleteCronJob)           // 删除CronJob
		CronJobRouter.DELETE("deleteCronJobByIds", v1.DeleteCronJobByIds) // 批量删除CronJob
		CronJobRouter.PUT("updateCronJob", v1.UpdateCronJob)              // 更新CronJob
		CronJobRouter.GET("findCronJob", v1.FindCronJob)                  // 根据ID获取CronJob
		CronJobRouter.GET("getCronJobList", v1.GetCronJobList)            // 获取CronJob列表
		CronJobRouter.GET("readYamlCronJob", v1.ReadYamlCronJob)          // 获取CronJob Yaml
		CronJobRouter.PUT("applyYamlCronJob", v1.ApplyYamlCronJob)        // 更新CronJob Yaml
	}
}
