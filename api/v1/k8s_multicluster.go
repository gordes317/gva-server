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

// @Tags MultiCluster
// @Summary 创建MultiCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MultiCluster true "创建MultiCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /multicluster/createMultiCluster [post]
func CreateMultiCluster(c *gin.Context) {
	var multicluster model.MultiCluster
	_ = c.ShouldBindJSON(&multicluster)
	if err := service.CreateMultiCluster(multicluster); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags MultiCluster
// @Summary 删除MultiCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MultiCluster true "删除MultiCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /multicluster/deleteMultiCluster [delete]
func DeleteMultiCluster(c *gin.Context) {
	var multicluster model.MultiCluster
	_ = c.ShouldBindJSON(&multicluster)
	if err := service.DeleteMultiCluster(multicluster); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags MultiCluster
// @Summary 批量删除MultiCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MultiCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /multicluster/deleteMultiClusterByIds [delete]
func DeleteMultiClusterByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteMultiClusterByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags MultiCluster
// @Summary 更新MultiCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MultiCluster true "更新MultiCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /multicluster/updateMultiCluster [put]
func UpdateMultiCluster(c *gin.Context) {
	var multicluster model.MultiCluster
	_ = c.ShouldBindJSON(&multicluster)
	if err := service.UpdateMultiCluster(multicluster); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags MultiCluster
// @Summary 用id查询MultiCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MultiCluster true "用id查询MultiCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /multicluster/findMultiCluster [get]
func FindMultiCluster(c *gin.Context) {
	var multicluster model.MultiCluster
	_ = c.ShouldBindQuery(&multicluster)
	if err, remulticluster := service.GetMultiCluster(multicluster.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remulticluster": remulticluster}, c)
	}
}

// @Tags MultiCluster
// @Summary 分页获取MultiCluster列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.MultiClusterSearch true "分页获取MultiCluster列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /multicluster/getMultiClusterList [get]
func GetMultiClusterList(c *gin.Context) {
	var pageInfo request.MultiClusterSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetMultiClusterInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		fmt.Printf("类型为%T\n", list)
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
