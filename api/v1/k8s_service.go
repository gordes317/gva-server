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

// CreateService
// @Tags Service
// @Summary 创建Service
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Service true "创建Service"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /service/createService [post]
func CreateService(c *gin.Context) {
	var Service model.ServiceUser
	_ = c.ShouldBindJSON(&Service)
	if err := service.CreateService(Service.Name, Service.UserName); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteService
// @Tags Service
// @Summary 删除Service
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Service true "删除Service"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /service/deleteService [delete]
func DeleteService(c *gin.Context) {
	var Service model.ServiceUser
	_ = c.ShouldBindJSON(&Service)
	if err := service.DeleteService(Service.Name, Service.Namespace, Service.UserName); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteServiceByIds
// @Tags Service
// @Summary 批量删除Service
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Service"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /service/deleteServiceByIds [delete]
func DeleteServiceByIds(c *gin.Context) {
	var Names request.NamesReqUser //需要修改
	_ = c.ShouldBindJSON(&Names)
	if err := service.DeleteServiceByIds(Names); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateService
// @Tags Service
// @Summary 更新Service
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Service true "更新Service"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /service/updateService [put]
func UpdateService(c *gin.Context) {
	var Service model.ServiceUser
	_ = c.ShouldBindJSON(&Service)
	if err := service.UpdateService(Service); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindService
// @Tags Service
// @Summary 用id查询Service
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Service true "用id查询Service"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /service/findService [get]
func FindService(c *gin.Context) {

	var Service model.ServiceUser
	_ = c.ShouldBindQuery(&Service)

	if err, reService := service.GetService(Service.Name, Service.UserName); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reService": reService}, c)
	}
}

// GetServiceList
// @Tags Service
// @Summary 分页获取Service列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ServiceSearch true "分页获取Service列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /service/getServiceList [get]
func GetServiceList(c *gin.Context) {
	var pageInfo request.ServiceSearchUser

	_ = c.ShouldBindQuery(&pageInfo)

	if err, list, total := service.GetServiceInfoList(pageInfo); err != nil {
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

// ApplyYamlService
// @Tags Service
// @Summary 更新Service Yaml
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Service true "更新Service Yaml"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /service/applyYamlService [put]
func ApplyYamlService(c *gin.Context) {
	var svc model.ServiceUser
	_ = c.ShouldBindJSON(&svc)
	if err := service.ApplyYamlService(svc); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// ReadYamlService
// @Tags Service
// @Summary 获取Service对应yaml
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Service true "获取Service对应yaml"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /service/readYamlService [get]
func ReadYamlService(c *gin.Context) {

	var svc model.ServiceUser
	_ = c.ShouldBindQuery(&svc)

	if err, reServiceYaml := service.ReadYamlService(svc.Name, svc.Namespace, svc.UserName); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reServiceYaml": reServiceYaml}, c)
	}
}
