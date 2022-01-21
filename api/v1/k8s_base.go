package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Namespace
// @Summary 分页获取Namespace列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.NamespaceSearch true "分页获取Namespace列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /base/getSelectNamespaceList [get]
func GetSelectNamespaceList(c *gin.Context) {
	var pageInfo request.SelectNamespacesSearchUser
	_ = c.ShouldBindQuery(&pageInfo)

	if err, list := service.GetSelectNamespaceList(pageInfo.UserName); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: list,
			//Total:    0,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// @Tags Cluster
// @Summary 获取集群列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.NamespaceSearch true "获取集群列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /base/getClusterList [get]
func GetClusterList(c *gin.Context) {

	if err, list := service.GetSelectClusterList(); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: list,
		}, "获取成功", c)
	}
}

//
// @Tags SwitchCluster
// @Summary 切换SwitchCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ConfigMap true "切换SwitchCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /base/updateSwitchCluster [put]
func UpdateSwitchCluster(c *gin.Context) {

	name := c.Query("name")
	user := c.Query("user")

	if err := service.SwitchCluster(name, user); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags DaemonSet
// @Summary 更新DaemonSet Yaml
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DaemonSet true "更新DaemonSet Yaml"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /configmap/applyYamlDaemonSet [put]
func ApplyYamlCreate(c *gin.Context) {

	json := make(map[string]string) //注意该结构接受的内容
	_ = c.BindJSON(&json)
	yamldata := json["yamldata"]
	userName := json["UserName"]
	if err := service.ApplyYamlCreate(yamldata, userName); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// GetClusterApplicationCounter
// @Tags Cluster
// @Summary 获取集群资源计数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfoUser true "分页获取Namespace列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /base/GetClusterApplicationCounter [get]
func GetClusterApplicationCounter(c *gin.Context) {
	var pageInfo request.PageInfoUser
	_ = c.ShouldBindQuery(&pageInfo)

	if err, list := service.GetClusterApplicationCounter(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: list,
			//Total:    0,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetClusterEvents
// @Tags Cluster
// @Summary 获取集群Events
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfoUser true "获取集群Events"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /base/GetClusterApplicationCounter [get]
func GetClusterEvents(c *gin.Context) {
	var pageInfo request.PageInfoUser
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetClusterEvents(pageInfo); err != nil {
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
