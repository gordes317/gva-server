package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	_ "gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// StartCi
// @Tags Pipelines
// @Summary 开始Ci
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Pod true "开始Ci流水线"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Pipelines/StartCi [Post]
func StartCi(c *gin.Context) {
	var repository model.Repository
	_ = c.ShouldBindJSON(&repository)
	if err, rePipelines := service.StartCi(repository); err != nil {
		global.GVA_LOG.Error("触发Ci流水线成功!", zap.Any("err", err))
		response.FailWithMessage("触发Ci流水线成功", c)
	} else {
		response.OkWithData(gin.H{"rePipelines": rePipelines}, c)
	}
}


// GetRepositoryProjectList
// @Tags Pipelines
// @Summary 获取仓库项目列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body true "获取仓库项目列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pipelines/getRepositoryList [GET]
func GetRepositoryProjectList(c *gin.Context) {
	//var pageInfo request.PodSearchUser
	//
	//_ = c.ShouldBindQuery(&pageInfo)

	if err, reRepositoryProject := service.GetRepositoryProjectList(); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"rePipelines": reRepositoryProject}, c)
	}
}


// FindRepoBranches
// @Tags Pipelines
// @Summary 根据项目ID获取项目分支列表和项目Tags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Repository true "根据项目PId获取项目分支列表和项目Tags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pipelines/FindRepoBranches [get]
func FindRepoBranches(c *gin.Context) {
	var Repo model.Repository
	_ = c.ShouldBindQuery(&Repo)

	if err, reBranches := service.FindRepoBranches(Repo.PId); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"reBranches": reBranches}, c)
	}
}

// FindRepoCommitId
// @Tags Pipelines
// @Summary 根据项目ID和分支名称或者tag名称获取项目CommitId
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Branch true "根据项目ID和分支名称或者tag名称获取项目CommitId"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pipelines/FindRepoCommitId [get]
func FindRepoCommitId(c *gin.Context) {
	var branch model.Branch
	_ = c.ShouldBindQuery(&branch)

	if err, reBranches := service.FindRepoCommitId(branch); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"rePipelines": reBranches}, c)
	}
}


// GetTektonPipelineRunsList
// @Tags Pipelines
// @Summary 分页获取Tekton PipelineRuns 列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TektonPipelineRunsSearchUser true "分页获取Tekton PipelineRuns 列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pipelines/getTektonPipelineRunsList [get]
func GetTektonPipelineRunsList(c *gin.Context) {
	var tkn request.TektonPipelineRunsListSearchUser
	_ = c.ShouldBindQuery(&tkn)
	tkn.Watch = false
	if err, list, total := service.GetTektonPipelineRunsList(tkn); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     tkn.Page,
			PageSize: tkn.PageSize,
		}, "获取成功", c)
	}
}


// GetPipeline
// @Tags Pipelines
// @Summary 用pipelineName查询Tasks
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TektonPipelineSearchUser true "用pipelineName查询Tasks"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Pipelines/getPipeline [get]
func GetPipeline(c *gin.Context) {

	var pipeline request.TektonPipelineSearchUser
	_ = c.ShouldBindQuery(&pipeline)

	if err, rePipeline := service.GetPipeline(pipeline.Name, pipeline.Namespace, pipeline.UserName); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rePipeline": rePipeline}, c)
	}
}


// GetPipelineTaskList
// @Tags Pipelines
// @Summary 获取pipeline task列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TektonPipelineUser true "获取pipeline task列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Pipelines/getPipelineTaskList [get]
func GetPipelineTaskList(c *gin.Context) {

	var _task request.TektonPipelineTaskUser
	_ = c.ShouldBindQuery(&_task)
	page, pageSize := 1, 100000
	if err, list, total := service.GetPipelineTaskList(_task.Name, _task.Namespace, _task.UserName, page, pageSize); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		}, "获取成功", c)
	}
}


// GetTaskRunList
// @Tags Pipelines
// @Summary 根据pipelineRun名称获取 taskRun
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TektonTaskRunListUser true "根据pipelineRun名称获取 taskRun"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Pipelines/GetTaskRunList [get]
func GetTaskRunList(c *gin.Context) {

	var taskRun request.TektonTaskRunListUser
	_ = c.ShouldBindQuery(&taskRun)
	page, pageSize := 1, 100000
	if err, list, total := service.GetTaskRunList(taskRun, page, pageSize); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		}, "获取成功", c)
	}
}

// CreateBuildPlan
// @Tags CreateBuildPlan
// @Summary 新建构建计划
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BuildPlan true
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /api/createBuildPlan [post]
func CreateBuildPlan(c *gin.Context) {
	var plan model.SysBuildPlan
	_ = c.ShouldBindJSON(&plan)
	if err := utils.Verify(plan, utils.BuildPlanVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.CreateBuildPlan(plan); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}



// GetBuildPlanList
// @Tags Pipelines
// @Summary 获取构建计划列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysBuildPlan true "获取构建计划列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Pipelines/getBuildPlanList [get]
func GetBuildPlanList(c *gin.Context) {

	//var bpl model.SysBuildPlan
	//_ = c.ShouldBindQuery(&bpl)
	var pageInfo request.SearchBuildPlanParams
	_ = c.ShouldBindQuery(&pageInfo)
	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := service.GetBuildPlanList(pageInfo.SysBuildPlan, pageInfo.PageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// DeleteBuildPlan
// @Tags DeleteBuildPlan
// @Summary 删除构建计划
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysBuildPlan true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /api/deleteBuildPlan [post]
func DeleteBuildPlan(c *gin.Context) {
	var buildPlan model.SysBuildPlan
	_ = c.ShouldBindJSON(&buildPlan)
	if err := utils.Verify(buildPlan.GVA_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.DeleteBuildPlan(buildPlan); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

