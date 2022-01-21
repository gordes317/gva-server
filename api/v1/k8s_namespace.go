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

// @Tags Namespace
// @Summary 创建Namespace
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Namespace true "创建Namespace"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Namespace/createNamespace [post]
func CreateNamespace(c *gin.Context) {
	var Namespace model.NamespaceUser
	_ = c.ShouldBindJSON(&Namespace)
	if err := service.CreateNamespace(Namespace.Name, Namespace.UserName); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Namespace
// @Summary 删除Namespace
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Namespace true "删除Namespace"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Namespace/deleteNamespace [delete]
func DeleteNamespace(c *gin.Context) {
	var Namespace model.NamespaceUser
	_ = c.ShouldBindJSON(&Namespace)
	if err := service.DeleteNamespace(Namespace.Name, Namespace.UserName); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Namespace
// @Summary 批量删除Namespace
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Namespace"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Namespace/deleteNamespaceByIds [delete]
func DeleteNamespaceByIds(c *gin.Context) {
	var Names request.NamesReq         //需要修改
	var NamesUser request.NamesReqUser //需要修改
	_ = c.ShouldBindJSON(&Names)
	if err := service.DeleteNamespaceByIds(Names, NamesUser.UserName); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Namespace
// @Summary 更新Namespace
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Namespace true "更新Namespace"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Namespace/updateNamespace [put]
func UpdateNamespace(c *gin.Context) {
	//var Namespace model.Namespace
	var NamespaceUser model.NamespaceUser
	//_ = c.ShouldBindJSON(&Namespace)
	_ = c.ShouldBindJSON(&NamespaceUser)
	// 先删除
	if err := service.DeleteNamespace(NamespaceUser.OldName, NamespaceUser.UserName); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	}
	// 后新增
	if err := service.CreateNamespace(NamespaceUser.Name, NamespaceUser.UserName); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
	//if err := service.UpdateNamespace(NamespaceUser); err != nil {
	//	global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
	//	response.FailWithMessage("更新失败", c)
	//} else {
	//	response.OkWithMessage("更新成功", c)
	//}
}

// @Tags Namespace
// @Summary 用id查询Namespace
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Namespace true "用id查询Namespace"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Namespace/findNamespace [get]
func FindNamespace(c *gin.Context) {

	var Namespace model.NamespaceUser
	_ = c.ShouldBindQuery(&Namespace)

	if err, reNamespace := service.GetNamespace(Namespace.Name, Namespace.UserName); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reNamespace": reNamespace}, c)
	}
}

// @Tags Namespace
// @Summary 分页获取Namespace列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.NamespaceSearch true "分页获取Namespace列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /namespace/getNamespaceList [get]
func GetNamespaceList(c *gin.Context) {
	var pageInfo request.NamespaceList

	_ = c.ShouldBindQuery(&pageInfo)
	fmt.Printf("namespace type:%T,value:%v\n", pageInfo, pageInfo)
	if err, list, total := service.GetNamespaceInfoList(pageInfo); err != nil {
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
