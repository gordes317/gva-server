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

// CreatePV
// @Tags PV
// @Summary 创建PV
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PV true "创建PV"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PV/createPV [post]
func CreatePV(c *gin.Context) {
	var PV model.PVUser
	_ = c.ShouldBindJSON(&PV)
	if err := service.CreatePV(PV.Name, PV.UserName); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePV
// @Tags PV
// @Summary 删除PV
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PV true "删除PV"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PV/deletePV [delete]
func DeletePV(c *gin.Context) {
	var PV model.PVUser
	_ = c.ShouldBindJSON(&PV)
	if err := service.DeletePV(PV.Name, PV.UserName); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePVByIds
// @Tags PV
// @Summary 批量删除PV
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PV"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /PV/deletePVByIds [delete]
func DeletePVByIds(c *gin.Context) {
	var Names request.NamesReqUser //需要修改
	_ = c.ShouldBindJSON(&Names)
	if err := service.DeletePVByIds(Names); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePV
// @Tags PV
// @Summary 更新PV
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PV true "更新PV"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PV/updatePV [put]
func UpdatePV(c *gin.Context) {
	var PV model.PVUser
	_ = c.ShouldBindJSON(&PV)
	if err := service.UpdatePV(PV); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPV
// @Tags PV
// @Summary 用id查询PV
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PV true "用id查询PV"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PV/findPV [get]
func FindPV(c *gin.Context) {

	var pv model.PVUser
	_ = c.ShouldBindQuery(&pv)

	if err, rePV := service.GetPV(pv.Name, pv.UserName); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rePV": rePV}, c)
	}
}

// GetPVList
// @Tags PV
// @Summary 分页获取PV列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PVSearch true "分页获取PV列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pv/getPVList [get]
func GetPVList(c *gin.Context) {
	var pageInfo request.PVSearchUser

	_ = c.ShouldBindQuery(&pageInfo)

	if err, list, total := service.GetPVInfoList(pageInfo); err != nil {
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
