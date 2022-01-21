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

// CreateStorageClass
// @Tags StorageClass
// @Summary 创建StorageClass
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StorageClass true "创建StorageClass"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /StorageClass/createStorageClass [post]
func CreateStorageClass(c *gin.Context) {
	var StorageClass model.StorageClassUser
	_ = c.ShouldBindJSON(&StorageClass)
	if err := service.CreateStorageClass(StorageClass.Name, StorageClass.UserName); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteStorageClass
// @Tags StorageClass
// @Summary 删除StorageClass
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StorageClass true "删除StorageClass"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /StorageClass/deleteStorageClass [delete]
func DeleteStorageClass(c *gin.Context) {
	var StorageClass model.StorageClassUser
	_ = c.ShouldBindJSON(&StorageClass)
	if err := service.DeleteStorageClass(StorageClass.Name, StorageClass.UserName); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteStorageClassByIds
// @Tags StorageClass
// @Summary 批量删除StorageClass
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除StorageClass"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /StorageClass/deleteStorageClassByIds [delete]
func DeleteStorageClassByIds(c *gin.Context) {
	var Names request.NamesReqUser //需要修改
	_ = c.ShouldBindJSON(&Names)
	if err := service.DeleteStorageClassByIds(Names); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateStorageClass
// @Tags StorageClass
// @Summary 更新StorageClass
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StorageClass true "更新StorageClass"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /StorageClass/updateStorageClass [put]
func UpdateStorageClass(c *gin.Context) {
	var StorageClass model.StorageClassUser
	_ = c.ShouldBindJSON(&StorageClass)
	if err := service.UpdateStorageClass(StorageClass); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindStorageClass
// @Tags StorageClass
// @Summary 用id查询StorageClass
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StorageClass true "用id查询StorageClass"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /StorageClass/findStorageClass [get]
func FindStorageClass(c *gin.Context) {

	var storageclass model.StorageClassUser
	_ = c.ShouldBindQuery(&storageclass)

	if err, reStorageClass := service.GetStorageClass(storageclass.Name, storageclass.UserName); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reStorageClass": reStorageClass}, c)
	}
}

// GetStorageClassList
// @Tags StorageClass
// @Summary 分页获取StorageClass列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.StorageClassSearch true "分页获取StorageClass列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /storageclass/getStorageClassList [get]
func GetStorageClassList(c *gin.Context) {
	var pageInfo request.StorageClassSearchUser

	_ = c.ShouldBindQuery(&pageInfo)

	if err, list, total := service.GetStorageClassInfoList(pageInfo); err != nil {
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
