package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitJobRouter(Router *gin.RouterGroup) {
	JobRouter := Router.Group("job").Use(middleware.OperationRecord())
	{
		JobRouter.POST("createJob", v1.CreateJob)             // 新建Job
		JobRouter.DELETE("deleteJob", v1.DeleteJob)           // 删除Job
		JobRouter.DELETE("deleteJobByIds", v1.DeleteJobByIds) // 批量删除Job
		JobRouter.PUT("updateJob", v1.UpdateJob)              // 更新Job
		JobRouter.GET("findJob", v1.FindJob)                  // 根据ID获取Job
		JobRouter.GET("getJobList", v1.GetJobList)            // 获取Job列表
		JobRouter.GET("readYamlJob", v1.ReadYamlJob)          // 获取Job Yaml
		JobRouter.PUT("applyYamlJob", v1.ApplyYamlJob)        // 更新Job Yaml
	}
}
