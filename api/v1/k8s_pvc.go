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

// CreatePVC
// @Tags PVC
// @Summary 创建PVC
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PVC true "创建PVC"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PVC/createPVC [post]
func CreatePVC(c *gin.Context) {
	var PVC model.PVCUser
	_ = c.ShouldBindJSON(&PVC)
	if err := service.CreatePVC(PVC.Name, PVC.UserName); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePVC
// @Tags PVC
// @Summary 删除PVC
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PVC true "删除PVC"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PVC/deletePVC [delete]
func DeletePVC(c *gin.Context) {
	var PVC model.PVCUser
	_ = c.ShouldBindJSON(&PVC)
	if err := service.DeletePVC(PVC.Name, PVC.Namespace, PVC.UserName); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePVCByIds
// @Tags PVC
// @Summary 批量删除PVC
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PVC"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /PVC/deletePVCByIds [delete]
func DeletePVCByIds(c *gin.Context) {
	var Names request.NamesReqUser //需要修改
	_ = c.ShouldBindJSON(&Names)
	if err := service.DeletePVCByIds(Names); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePVC
// @Tags PVC
// @Summary 更新PVC
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PVC true "更新PVC"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PVC/updatePVC [put]
func UpdatePVC(c *gin.Context) {
	var PVC model.PVCUser
	_ = c.ShouldBindJSON(&PVC)
	if err := service.UpdatePVC(PVC); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPVC
// @Tags PVC
// @Summary 用id查询PVC
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PVC true "用id查询PVC"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PVC/findPVC [get]
func FindPVC(c *gin.Context) {

	var pvc model.PVCUser
	_ = c.ShouldBindQuery(&pvc)

	if err, rePVC := service.GetPVC(pvc.Name,pvc.Namespace, pvc.UserName); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rePVC": rePVC}, c)
	}
}

// GetPVCList
// @Tags PVC
// @Summary 分页获取PVC列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PVCSearch true "分页获取PVC列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pvc/getPVCList [get]
func GetPVCList(c *gin.Context) {
	var pageInfo request.PVCSearchUser

	_ = c.ShouldBindQuery(&pageInfo)

	if err, list, total := service.GetPVCInfoList(pageInfo); err != nil {
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
