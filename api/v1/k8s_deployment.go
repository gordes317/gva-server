package v1

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateDeployment
// @Tags Deployment
// @Summary 创建Deployment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Deployment true "创建Deployment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deployment/createDeployment [post]
func CreateDeployment(c *gin.Context) {
	var deployment model.DeploymentUser
	_ = c.ShouldBindJSON(&deployment)
	if err := service.CreateDeployment(deployment.Name, deployment.UserName); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDeployment
// @Tags Deployment
// @Summary 删除Deployment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Deployment true "删除Deployment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /deployment/deleteDeployment [delete]
func DeleteDeployment(c *gin.Context) {
	var deployment model.DeploymentUser
	_ = c.ShouldBindJSON(&deployment)
	fmt.Println("deploy:", deployment)
	if err := service.DeleteDeployment(deployment.Name, deployment.Namespace, deployment.UserName); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDeploymentByIds
// @Tags Deployment
// @Summary 批量删除Deployment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Deployment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /deployment/deleteDeploymentByIds [delete]
func DeleteDeploymentByIds(c *gin.Context) {
	var Names request.NamesReqUser //需要修改
	_ = c.ShouldBindJSON(&Names)
	if err := service.DeleteDeploymentByIds(Names); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDeployment
// @Tags Deployment
// @Summary 更新Deployment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Deployment true "更新Deployment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /deployment/updateDeployment [put]
func UpdateDeployment(c *gin.Context) {
	var deployment model.DeploymentUser
	_ = c.ShouldBindJSON(&deployment)
	if err := service.UpdateDeployment(deployment); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDeployment
// @Tags Deployment
// @Summary 用id查询Deployment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Deployment true "用id查询Deployment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /deployment/findDeployment [get]
func FindDeployment(c *gin.Context) {

	var deployment model.DeploymentUser
	_ = c.ShouldBindQuery(&deployment)

	if err, reDeployment := service.GetDeployment(deployment.Name, deployment.Namespace, deployment.UserName); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reDeployment": reDeployment}, c)
	}
}

// GetDeploymentList
// @Tags Deployment
// @Summary 分页获取Deployment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DeploymentSearch true "分页获取Deployment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deployment/getDeploymentList [get]
func GetDeploymentList(c *gin.Context) {
	var pageInfo request.DeploymentSearchUser

	_ = c.ShouldBindQuery(&pageInfo)

	if err, list, total := service.GetDeploymentInfoList(pageInfo); err != nil {
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

// ApplyYamlDeployment
// @Tags Deployment
// @Summary 更新Deployment Yaml
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Deployment true "更新Deployment Yaml"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /deployment/applyYamlDeployment [put]
func ApplyYamlDeployment(c *gin.Context) {
	var deployment model.DeploymentUser
	_ = c.ShouldBindJSON(&deployment)
	if err := service.ApplyYamlDeployment(deployment); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// ReadYamlDeployment
// @Tags Deployment
// @Summary 获取Deployment对应yaml
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Deployment true "获取Deployment对应yaml"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /deployment/readYamlDeployment [get]
func ReadYamlDeployment(c *gin.Context) {

	var deployment model.DeploymentUser
	_ = c.ShouldBindQuery(&deployment)

	if err, reDeploymentYaml := service.ReadYamlDeployment(deployment.Name, deployment.Namespace, deployment.UserName); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reDeploymentYaml": reDeploymentYaml}, c)
	}
}

// RestartDeployment
// @Tags Deployment
// @Summary 更新Deployment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Deployment true "更新Deployment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /deployment/restartDeployment [put]
func RestartDeployment(c *gin.Context) {
	var deployment model.DeploymentUser
	_ = c.ShouldBindJSON(&deployment)
	if err := service.RestartDeployment(deployment); err != nil {
		global.GVA_LOG.Error("重启失败!", zap.Any("err", err))
		response.FailWithMessage("重启失败", c)
	} else {
		response.OkWithMessage("重启成功", c)
	}
}
