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

// CreateSecret
// @Tags Secret
// @Summary 创建Secret
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Secret true "创建Secret"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Secret/createSecret [post]
func CreateSecret(c *gin.Context) {
	var Secret model.SecretUser
	_ = c.ShouldBindJSON(&Secret)
	if err := service.CreateSecret(Secret.Name, Secret.UserName); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSecret
// @Tags Secret
// @Summary 删除Secret
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Secret true "删除Secret"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Secret/deleteSecret [delete]
func DeleteSecret(c *gin.Context) {
	var Secret model.SecretUser
	_ = c.ShouldBindJSON(&Secret)
	if err := service.DeleteSecret(Secret.Name, Secret.UserName); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSecretByIds
// @Tags Secret
// @Summary 批量删除Secret
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Secret"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Secret/deleteSecretByIds [delete]
func DeleteSecretByIds(c *gin.Context) {
	var Names request.NamesReqUser //需要修改
	_ = c.ShouldBindJSON(&Names)
	if err := service.DeleteSecretByIds(Names); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSecret
// @Tags Secret
// @Summary 更新Secret
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Secret true "更新Secret"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Secret/updateSecret [put]
func UpdateSecret(c *gin.Context) {
	var Secret model.Secret
	var SecretUser model.SecretUser
	_ = c.ShouldBindJSON(&Secret)
	_ = c.ShouldBindJSON(&SecretUser)
	if err := service.UpdateSecret(Secret, SecretUser.UserName); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSecret
// @Tags Secret
// @Summary 用id查询Secret
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Secret true "用id查询Secret"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Secret/findSecret [get]
func FindSecret(c *gin.Context) {

	var secret model.SecretUser
	_ = c.ShouldBindQuery(&secret)

	if err, reSecret := service.GetSecret(secret.Name,secret.Namespace, secret.UserName); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reSecret": reSecret}, c)
	}
}

// GetSecretList
// @Tags Secret
// @Summary 分页获取Secret列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SecretSearch true "分页获取Secret列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /secret/getSecretList [get]
func GetSecretList(c *gin.Context) {
	var pageInfo request.SecretSearchUser

	_ = c.ShouldBindQuery(&pageInfo)

	if err, list, total := service.GetSecretInfoList(pageInfo); err != nil {
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

// ApplyYamlSecret
// @Tags Secret
// @Summary 更新Secret Yaml
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Secret true "更新Secret Yaml"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /secret/applyYamlSecret [put]
func ApplyYamlSecret(c *gin.Context) {
	var secret model.SecretUser
	_ = c.ShouldBindJSON(&secret)
	if err := service.ApplyYamlSecret(secret); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// ReadYamlSecret
// @Tags Secret
// @Summary 获取Secret对应yaml
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Secret true "获取Secret对应yaml"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /secret/readYamlSecret [get]
func ReadYamlSecret(c *gin.Context) {

	var secret model.SecretUser
	_ = c.ShouldBindQuery(&secret)

	if err, reSecretYaml := service.ReadYamlSecret(secret.Name, secret.Namespace, secret.UserName); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reSecretYaml": reSecretYaml}, c)
	}
}
