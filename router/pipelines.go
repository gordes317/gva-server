package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitPipelinesRouter(Router *gin.RouterGroup) {
	PipelinesRouter := Router.Group("pipelines").Use(middleware.OperationRecord())
	{
		PipelinesRouter.POST("startCi", v1.StartCi)             // 触发ci流水线
		PipelinesRouter.GET("getRepositoryProjectList", v1.GetRepositoryProjectList) // 获取仓库项目列表
		PipelinesRouter.GET("findRepoBranches", v1.FindRepoBranches) // 根据项目PId获取项目分支列表和项目Tags
		PipelinesRouter.GET("findRepoCommitId", v1.FindRepoCommitId) // 根据项目ID和分支名称或者tag名称获取项目CommitId
		PipelinesRouter.GET("getTektonPipelineRunsList", v1.GetTektonPipelineRunsList) // 分页获取Tekton PipelineRuns 列表
		PipelinesRouter.GET("getTaskRunList", v1.GetTaskRunList) // 分页获取Tekton TaskRuns 列表
		PipelinesRouter.GET("getPipelineTaskList", v1.GetPipelineTaskList) // 获取pipeline task列表
		PipelinesRouter.GET("getPipeline", v1.GetPipeline) // 根据pipelineName查询Tasks
		PipelinesRouter.POST("createBuildPlan", v1.CreateBuildPlan) // 新建构建计划
		PipelinesRouter.GET("getBuildPlanList", v1.GetBuildPlanList) // 分页获取构建计划列表
		PipelinesRouter.POST("deleteBuildPlan", v1.DeleteBuildPlan) // 删除构建计划
	}
}
