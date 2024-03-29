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

// CreateCronJob
// @Tags CronJob
// @Summary 创建CronJob
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CronJob true "创建CronJob"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /CronJob/createCronJob [post]
func CreateCronJob(c *gin.Context) {
	var CronJob model.CronJobUser
	_ = c.ShouldBindJSON(&CronJob)
	if err := service.CreateCronJob(CronJob.Name, CronJob.UserName); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCronJob
// @Tags CronJob
// @Summary 删除CronJob
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CronJob true "删除CronJob"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /CronJob/deleteCronJob [delete]
func DeleteCronJob(c *gin.Context) {
	var CronJob model.CronJobUser
	_ = c.ShouldBindJSON(&CronJob)
	if err := service.DeleteCronJob(CronJob.Name, CronJob.Namespace, CronJob.UserName); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCronJobByIds
// @Tags CronJob
// @Summary 批量删除CronJob
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CronJob"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /CronJob/deleteCronJobByIds [delete]
func DeleteCronJobByIds(c *gin.Context) {
	var Names request.NamesReqUser //需要修改
	_ = c.ShouldBindJSON(&Names)
	if err := service.DeleteCronJobByIds(Names); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCronJob
// @Tags CronJob
// @Summary 更新CronJob
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CronJob true "更新CronJob"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /CronJob/updateCronJob [put]
func UpdateCronJob(c *gin.Context) {
	var CronJob model.CronJobUser
	_ = c.ShouldBindJSON(&CronJob)
	if err := service.UpdateCronJob(CronJob); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCronJob
// @Tags CronJob
// @Summary 用id查询CronJob
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CronJob true "用id查询CronJob"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /CronJob/findCronJob [get]
func FindCronJob(c *gin.Context) {

	var configmap model.CronJobUser
	_ = c.ShouldBindQuery(&configmap)

	if err, reCronJob := service.GetCronJob(configmap.Name,configmap.Namespace, configmap.UserName); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reCronJob": reCronJob}, c)
	}
}

// GetCronJobList
// @Tags CronJob
// @Summary 分页获取CronJob列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CronJobSearch true "分页获取CronJob列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /configmap/getCronJobList [get]
func GetCronJobList(c *gin.Context) {
	var pageInfo request.CronJobSearchUser

	_ = c.ShouldBindQuery(&pageInfo)

	if err, list, total := service.GetCronJobInfoList(pageInfo); err != nil {
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

// ApplyYamlCronJob
// @Tags CronJob
// @Summary 更新CronJob Yaml
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CronJob true "更新CronJob Yaml"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /configmap/applyYamlCronJob [put]
func ApplyYamlCronJob(c *gin.Context) {
	var configmap model.CronJobUser
	_ = c.ShouldBindJSON(&configmap)
	if err := service.ApplyYamlCronJob(configmap); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// ReadYamlCronJob
// @Tags CronJob
// @Summary 获取CronJob对应yaml
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CronJob true "获取CronJob对应yaml"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /configmap/readYamlCronJob [get]
func ReadYamlCronJob(c *gin.Context) {

	var configmap model.CronJobUser
	_ = c.ShouldBindQuery(&configmap)

	if err, reCronJobYaml := service.ReadYamlCronJob(configmap.Name, configmap.Namespace, configmap.UserName); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reCronJobYaml": reCronJobYaml}, c)
	}
}
