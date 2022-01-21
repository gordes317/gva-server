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

// @Tags Template
// @Summary 创建Template
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Template true "创建Template"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /template/createTemplate [post]
func CreateTemplate(c *gin.Context) {
	var template model.Template
	_ = c.ShouldBindJSON(&template)
	if err := service.CreateTemplate(template); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Template
// @Summary 删除Template
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Template true "删除Template"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /template/deleteTemplate [delete]
func DeleteTemplate(c *gin.Context) {
	var template model.Template
	_ = c.ShouldBindJSON(&template)
	if err := service.DeleteTemplate(template); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Template
// @Summary 批量删除Template
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Template"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /template/deleteTemplateByIds [delete]
func DeleteTemplateByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteTemplateByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Template
// @Summary 更新Template
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Template true "更新Template"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /template/updateTemplate [put]
func UpdateTemplate(c *gin.Context) {
	var template model.Template
	_ = c.ShouldBindJSON(&template)
	if err := service.UpdateTemplate(template); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Template
// @Summary 用id查询Template
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Template true "用id查询Template"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /template/findTemplate [get]
func FindTemplate(c *gin.Context) {
	var template model.Template
	_ = c.ShouldBindQuery(&template)
	if err, retemplate := service.GetTemplate(template.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retemplate": retemplate}, c)
	}
}

// @Tags Template
// @Summary 分页获取Template列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TemplateSearch true "分页获取Template列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /template/getTemplateList [get]
func GetTemplateList(c *gin.Context) {
	var pageInfo request.TemplateSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetTemplateInfoList(pageInfo); err != nil {
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
