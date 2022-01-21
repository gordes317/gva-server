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

// CreateDaemonSet
// @Tags DaemonSet
// @Summary 创建DaemonSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DaemonSet true "创建DaemonSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DaemonSet/createDaemonSet [post]
func CreateDaemonSet(c *gin.Context) {
	var DaemonSet model.DaemonSetUser
	_ = c.ShouldBindJSON(&DaemonSet)
	if err := service.CreateDaemonSet(DaemonSet.Name, DaemonSet.UserName); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDaemonSet
// @Tags DaemonSet
// @Summary 删除DaemonSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DaemonSet true "删除DaemonSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DaemonSet/deleteDaemonSet [delete]
func DeleteDaemonSet(c *gin.Context) {
	var DaemonSet model.DaemonSetUser
	_ = c.ShouldBindJSON(&DaemonSet)
	if err := service.DeleteDaemonSet(DaemonSet.Name, DaemonSet.Namespace, DaemonSet.UserName); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDaemonSetByIds
// @Tags DaemonSet
// @Summary 批量删除DaemonSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DaemonSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /DaemonSet/deleteDaemonSetByIds [delete]
func DeleteDaemonSetByIds(c *gin.Context) {
	var Names request.NamesReqUser //需要修改
	_ = c.ShouldBindJSON(&Names)
	if err := service.DeleteDaemonSetByIds(Names, Names.UserName); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDaemonSet
// @Tags DaemonSet
// @Summary 更新DaemonSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DaemonSet true "更新DaemonSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /DaemonSet/updateDaemonSet [put]
func UpdateDaemonSet(c *gin.Context) {
	var DaemonSet model.DaemonSetUser
	_ = c.ShouldBindJSON(&DaemonSet)
	if err := service.UpdateDaemonSet(DaemonSet); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDaemonSet
// @Tags DaemonSet
// @Summary 用id查询DaemonSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DaemonSet true "用id查询DaemonSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /DaemonSet/findDaemonSet [get]
func FindDaemonSet(c *gin.Context) {

	var configmap model.DaemonSetUser
	_ = c.ShouldBindQuery(&configmap)

	if err, reDaemonSet := service.GetDaemonSet(configmap.Name,configmap.Namespace, configmap.UserName); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reDaemonSet": reDaemonSet}, c)
	}
}

// GetDaemonSetList
// @Tags DaemonSet
// @Summary 分页获取DaemonSet列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DaemonSetSearch true "分页获取DaemonSet列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /configmap/getDaemonSetList [get]
func GetDaemonSetList(c *gin.Context) {
	var pageInfo request.DaemonSetSearchUser

	_ = c.ShouldBindQuery(&pageInfo)

	if err, list, total := service.GetDaemonSetInfoList(pageInfo); err != nil {
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

// ApplyYamlDaemonSet
// @Tags DaemonSet
// @Summary 更新DaemonSet Yaml
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DaemonSet true "更新DaemonSet Yaml"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /configmap/applyYamlDaemonSet [put]
func ApplyYamlDaemonSet(c *gin.Context) {
	var configmap model.DaemonSetUser
	_ = c.ShouldBindJSON(&configmap)
	if err := service.ApplyYamlDaemonSet(configmap); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// ReadYamlDaemonSet
// @Tags DaemonSet
// @Summary 获取DaemonSet对应yaml
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DaemonSet true "获取DaemonSet对应yaml"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /configmap/readYamlDaemonSet [get]
func ReadYamlDaemonSet(c *gin.Context) {

	var configmap model.DaemonSetUser
	_ = c.ShouldBindQuery(&configmap)

	if err, reDaemonSetYaml := service.ReadYamlDaemonSet(configmap.Name, configmap.Namespace, configmap.UserName); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reDaemonSetYaml": reDaemonSetYaml}, c)
	}
}
