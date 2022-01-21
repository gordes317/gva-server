package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateJob
// @Tags Job
// @Summary 创建Job
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Job true "创建Job"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Job/createJob [post]
func CreateJob(c *gin.Context) {
	var Job model.JobUser
	_ = c.ShouldBindJSON(&Job)
	if err := service.CreateJob(Job.Name, Job.Namespace, Job.UserName); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteJob
// @Tags Job
// @Summary 删除Job
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Job true "删除Job"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Job/deleteJob [delete]
func DeleteJob(c *gin.Context) {
	var Job model.JobUser
	_ = c.ShouldBindJSON(&Job)
	if err := service.DeleteJob(Job.Name, Job.Namespace, Job.UserName); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteJobByIds
// @Tags Job
// @Summary 批量删除Job
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Job"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Job/deleteJobByIds [delete]
func DeleteJobByIds(c *gin.Context) {
	var Names request.NamesReqUser //需要修改
	_ = c.ShouldBindJSON(&Names)
	if err := service.DeleteJobByIds(Names); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateJob
// @Tags Job
// @Summary 更新Job
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Job true "更新Job"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Job/updateJob [put]
func UpdateJob(c *gin.Context) {
	var Job model.JobUser
	_ = c.ShouldBindJSON(&Job)
	if err := service.UpdateJob(Job); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindJob
// @Tags Job
// @Summary 用id查询Job
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Job true "用id查询Job"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Job/findJob [get]
func FindJob(c *gin.Context) {

	var configmap model.JobUser
	_ = c.ShouldBindQuery(&configmap)

	if err, reJob := service.GetJob(configmap.Name,configmap.Namespace, configmap.UserName); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reJob": reJob}, c)
	}
}

// GetJobList
// @Tags Job
// @Summary 分页获取Job列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.JobSearch true "分页获取Job列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /configmap/getJobList [get]
func GetJobList(c *gin.Context) {
	var pageInfo request.JobSearchUser

	_ = c.ShouldBindQuery(&pageInfo)

	if err, list, total := service.GetJobInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
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

// ApplyYamlJob
// @Tags Job
// @Summary 更新Job Yaml
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Job true "更新Job Yaml"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /configmap/applyYamlJob [put]
func ApplyYamlJob(c *gin.Context) {
	var configmap model.JobUser
	_ = c.ShouldBindJSON(&configmap)
	if err := service.ApplyYamlJob(configmap); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// ReadYamlJob
// @Tags Job
// @Summary 获取Job对应yaml
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Job true "获取Job对应yaml"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /configmap/readYamlJob [get]
func ReadYamlJob(c *gin.Context) {

	var configmap model.JobUser
	_ = c.ShouldBindQuery(&configmap)

	if err, reJobYaml := service.ReadYamlJob(configmap.Name, configmap.Namespace, configmap.UserName); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reJobYaml": reJobYaml}, c)
	}
}
