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

// CreateStatefulSet
// @Tags StatefulSet
// @Summary 创建StatefulSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StatefulSet true "创建StatefulSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /StatefulSet/createStatefulSet [post]
func CreateStatefulSet(c *gin.Context) {
	var StatefulSet model.StatefulSetUser
	_ = c.ShouldBindJSON(&StatefulSet)
	if err := service.CreateStatefulSet(StatefulSet.Name, StatefulSet.UserName); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteStatefulSet
// @Tags StatefulSet
// @Summary 删除StatefulSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StatefulSet true "删除StatefulSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /StatefulSet/deleteStatefulSet [delete]
func DeleteStatefulSet(c *gin.Context) {
	var StatefulSet model.StatefulSetUser
	_ = c.ShouldBindJSON(&StatefulSet)
	if err := service.DeleteStatefulSet(StatefulSet.Name, StatefulSet.Namespace, StatefulSet.UserName); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteStatefulSetByIds
// @Tags StatefulSet
// @Summary 批量删除StatefulSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除StatefulSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /StatefulSet/deleteStatefulSetByIds [delete]
func DeleteStatefulSetByIds(c *gin.Context) {
	var Names request.NamesReqUser//需要修改
	_ = c.ShouldBindJSON(&Names)
	if err := service.DeleteStatefulSetByIds(Names); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateStatefulSet
// @Tags StatefulSet
// @Summary 更新StatefulSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StatefulSet true "更新StatefulSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /StatefulSet/updateStatefulSet [put]
func UpdateStatefulSet(c *gin.Context) {
	var StatefulSet model.StatefulSetUser
	_ = c.ShouldBindJSON(&StatefulSet)
	if err := service.UpdateStatefulSet(StatefulSet); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindStatefulSet
// @Tags StatefulSet
// @Summary 用id查询StatefulSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StatefulSet true "用id查询StatefulSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /StatefulSet/findStatefulSet [get]
func FindStatefulSet(c *gin.Context) {

	var configmap model.StatefulSetUser
	_ = c.ShouldBindQuery(&configmap)

	if err, reStatefulSet := service.GetStatefulSet(configmap.Name,configmap.Namespace, configmap.UserName); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reStatefulSet": reStatefulSet}, c)
	}
}

// GetStatefulSetList
// @Tags StatefulSet
// @Summary 分页获取StatefulSet列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.StatefulSetSearch true "分页获取StatefulSet列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /configmap/getStatefulSetList [get]
func GetStatefulSetList(c *gin.Context) {
	var pageInfo request.StatefulSetSearchUser

	_ = c.ShouldBindQuery(&pageInfo)

	if err, list, total := service.GetStatefulSetInfoList(pageInfo); err != nil {
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

// ApplyYamlStatefulSet
// @Tags StatefulSet
// @Summary 更新StatefulSet Yaml
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StatefulSet true "更新StatefulSet Yaml"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /configmap/applyYamlStatefulSet [put]
func ApplyYamlStatefulSet(c *gin.Context) {
	var configmap model.StatefulSetUser
	_ = c.ShouldBindJSON(&configmap)
	if err := service.ApplyYamlStatefulSet(configmap); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// ReadYamlStatefulSet
// @Tags StatefulSet
// @Summary 获取StatefulSet对应yaml
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StatefulSet true "获取StatefulSet对应yaml"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /configmap/readYamlStatefulSet [get]
func ReadYamlStatefulSet(c *gin.Context) {

	var configmap model.StatefulSetUser
	_ = c.ShouldBindQuery(&configmap)

	if err, reStatefulSetYaml := service.ReadYamlStatefulSet(configmap.Name, configmap.Namespace, configmap.UserName); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reStatefulSetYaml": reStatefulSetYaml}, c)
	}
}
