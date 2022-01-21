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

// CreateConfigMap
// @Tags ConfigMap
// @Summary 创建ConfigMap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ConfigMap true "创建ConfigMap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ConfigMap/createConfigMap [post]
func CreateConfigMap(c *gin.Context) {
	var ConfigMap model.ConfigMapUser
	_ = c.ShouldBindJSON(&ConfigMap)
	if err := service.CreateConfigMap(ConfigMap.Name, ConfigMap.UserName); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteConfigMap
// @Tags ConfigMap
// @Summary 删除ConfigMap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ConfigMap true "删除ConfigMap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ConfigMap/deleteConfigMap [delete]
func DeleteConfigMap(c *gin.Context) {
	var ConfigMap model.ConfigMapUser
	_ = c.ShouldBindJSON(&ConfigMap)
	if err := service.DeleteConfigMap(ConfigMap.Name, ConfigMap.UserName); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteConfigMapByIds
// @Tags ConfigMap
// @Summary 批量删除ConfigMap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ConfigMap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ConfigMap/deleteConfigMapByIds [delete]
func DeleteConfigMapByIds(c *gin.Context) {
	var Names request.NamesReq //需要修改
	var NamesUser request.NamesReqUser //需要修改
	_ = c.ShouldBindJSON(&Names)
	if err := service.DeleteConfigMapByIds(Names, NamesUser.UserName); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateConfigMap
// @Tags ConfigMap
// @Summary 更新ConfigMap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ConfigMap true "更新ConfigMap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ConfigMap/updateConfigMap [put]
func UpdateConfigMap(c *gin.Context) {
	var ConfigMap model.ConfigMap
	var ConfigMapUser model.ConfigMapUser
	_ = c.ShouldBindJSON(&ConfigMap)
	if err := service.UpdateConfigMap(ConfigMap, ConfigMapUser.UserName); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindConfigMap
// @Tags ConfigMap
// @Summary 用id查询ConfigMap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ConfigMap true "用id查询ConfigMap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ConfigMap/findConfigMap [get]
func FindConfigMap(c *gin.Context) {

	var configmap model.ConfigMapUser
	_ = c.ShouldBindQuery(&configmap)

	if err, reConfigMap := service.GetConfigMap(configmap.Name,configmap.Namespace, configmap.UserName); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reConfigMap": reConfigMap}, c)
	}
}

// GetConfigMapList
// @Tags ConfigMap
// @Summary 分页获取ConfigMap列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ConfigMapSearch true "分页获取ConfigMap列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /configmap/getConfigMapList [get]
func GetConfigMapList(c *gin.Context) {
	var pageInfo request.ConfigMapSearchUser

	_ = c.ShouldBindQuery(&pageInfo)

	if err, list, total := service.GetConfigMapInfoList(pageInfo); err != nil {
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

// ApplyYamlConfigMap
// @Tags ConfigMap
// @Summary 更新ConfigMap Yaml
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ConfigMap true "更新ConfigMap Yaml"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /configmap/applyYamlConfigMap [put]
func ApplyYamlConfigMap(c *gin.Context) {
	var configmap model.ConfigMap
	var configmapUser model.ConfigMapUser
	_ = c.ShouldBindJSON(&configmap)
	_ = c.ShouldBindJSON(&configmapUser)
	if err := service.ApplyYamlConfigMap(configmap, configmapUser.UserName); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// ReadYamlConfigMap
// @Tags ConfigMap
// @Summary 获取ConfigMap对应yaml
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ConfigMap true "获取ConfigMap对应yaml"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /configmap/readYamlConfigMap [get]
func ReadYamlConfigMap(c *gin.Context) {

	var configmap model.ConfigMapUser
	_ = c.ShouldBindQuery(&configmap)

	if err, reConfigMapYaml := service.ReadYamlConfigMap(configmap.Name, configmap.Namespace, configmap.UserName); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reConfigMapYaml": reConfigMapYaml}, c)
	}
}
