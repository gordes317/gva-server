package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateIngress
// @Tags Ingress
// @Summary 创建Ingress
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Ingress true "创建Ingress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Ingress/createIngress [post]
func CreateIngress(c *gin.Context) {
	var Ingress model.IngressUser
	_ = c.ShouldBindJSON(&Ingress)
	kubeVersion := utils.GetKubeVersion(Ingress.UserName)
	if kubeVersion >= 119 {
		if err := service.CreateIngress(Ingress.Name, Ingress.UserName); err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
			response.FailWithMessage("创建失败", c)
		} else {
			response.OkWithMessage("创建成功", c)
		}
	} else {
		if err := service.CreateIngressUnder119(Ingress.Name, Ingress.UserName); err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
			response.FailWithMessage("创建失败", c)
		} else {
			response.OkWithMessage("创建成功", c)
		}
	}

}

// DeleteIngress
// @Tags Ingress
// @Summary 删除Ingress
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Ingress true "删除Ingress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Ingress/deleteIngress [delete]
func DeleteIngress(c *gin.Context) {
	var Ingress model.IngressUser
	_ = c.ShouldBindJSON(&Ingress)
	kubeVersion := utils.GetKubeVersion(Ingress.UserName)
	if kubeVersion >= 119 {
		if err := service.DeleteIngress(Ingress.Name, Ingress.Namespace, Ingress.UserName); err != nil {
			global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
			response.FailWithMessage("删除失败", c)
		} else {
			response.OkWithMessage("删除成功", c)
		}
	} else {
		if err := service.DeleteIngressUnder119(Ingress.Name, Ingress.Namespace, Ingress.UserName); err != nil {
			global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
			response.FailWithMessage("删除失败", c)
		} else {
			response.OkWithMessage("删除成功", c)
		}
	}

}

// DeleteIngressByIds
// @Tags Ingress
// @Summary 批量删除Ingress
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Ingress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Ingress/deleteIngressByIds [delete]
func DeleteIngressByIds(c *gin.Context) {
	var Names request.NamesReqUser //需要修改
	_ = c.ShouldBindJSON(&Names)
	kubeVersion := utils.GetKubeVersion(Names.UserName)
	if kubeVersion >= 119 {
		if err := service.DeleteIngressByIds(Names, Names.UserName); err != nil {
			global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
			response.FailWithMessage("批量删除失败", c)
		} else {
			response.OkWithMessage("批量删除成功", c)
		}
	} else {
		if err := service.DeleteIngressByIdsUnder119(Names, Names.UserName); err != nil {
			global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
			response.FailWithMessage("批量删除失败", c)
		} else {
			response.OkWithMessage("批量删除成功", c)
		}
	}

}

// UpdateIngress
// @Tags Ingress
// @Summary 更新Ingress
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Ingress true "更新Ingress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Ingress/updateIngress [put]
func UpdateIngress(c *gin.Context) {
	var Ingress model.IngressUser
	_ = c.ShouldBindJSON(&Ingress)
	kubeVersion := utils.GetKubeVersion(Ingress.UserName)
	if kubeVersion >= 119 {
		if err := service.UpdateIngress(Ingress); err != nil {
			global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
			response.FailWithMessage("更新失败", c)
		} else {
			response.OkWithMessage("更新成功", c)
		}
	} else {
		if err := service.UpdateIngressUnder119(Ingress); err != nil {
			global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
			response.FailWithMessage("更新失败", c)
		} else {
			response.OkWithMessage("更新成功", c)
		}
	}

}

// FindIngress
// @Tags Ingress
// @Summary 用id查询Ingress
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Ingress true "用id查询Ingress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Ingress/findIngress [get]
func FindIngress(c *gin.Context) {

	var configmap model.IngressUser
	_ = c.ShouldBindQuery(&configmap)
	kubeVersion := utils.GetKubeVersion(configmap.UserName)
	if kubeVersion >= 119 {
		if err, reIngress := service.GetIngress(configmap.Name, configmap.Namespace, configmap.UserName); err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
			response.FailWithMessage("查询失败", c)
		} else {
			response.OkWithData(gin.H{"reIngress": reIngress}, c)
		}
	} else {
		if err, reIngress := service.GetIngressUnder119(configmap.Name, configmap.Namespace, configmap.UserName); err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
			response.FailWithMessage("查询失败", c)
		} else {
			response.OkWithData(gin.H{"reIngress": reIngress}, c)
		}
	}

}

// GetIngressList
// @Tags Ingress
// @Summary 分页获取Ingress列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IngressSearch true "分页获取Ingress列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /configmap/getIngressList [get]
func GetIngressList(c *gin.Context) {
	var pageInfo request.IngressSearchUser

	_ = c.ShouldBindQuery(&pageInfo)
	kubeVersion := utils.GetKubeVersion(pageInfo.UserName)
	if kubeVersion >= 119 {
		if err, list, total := service.GetIngressInfoList(pageInfo); err != nil {
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
	} else {
		if err, list, total := service.GetIngressInfoListUnder119(pageInfo); err != nil {
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

}

// ApplyYamlIngress
// @Tags Ingress
// @Summary 更新Ingress Yaml
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Ingress true "更新Ingress Yaml"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /configmap/applyYamlIngress [put]
func ApplyYamlIngress(c *gin.Context) {
	var configmap model.IngressUser
	_ = c.ShouldBindJSON(&configmap)
	if err := service.ApplyYamlIngress(configmap); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// ReadYamlIngress
// @Tags Ingress
// @Summary 获取Ingress对应yaml
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Ingress true "获取Ingress对应yaml"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /configmap/readYamlIngress [get]
func ReadYamlIngress(c *gin.Context) {

	var configmap model.IngressUser
	_ = c.ShouldBindQuery(&configmap)

	kubeVersion := utils.GetKubeVersion(configmap.UserName)
	if kubeVersion >= 119 {
		if err, reIngressYaml := service.ReadYamlIngress(configmap.Name, configmap.Namespace, configmap.UserName); err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
			response.FailWithMessage("查询失败", c)
		} else {
			response.OkWithData(gin.H{"reIngressYaml": reIngressYaml}, c)
		}
	} else {
		if err, reIngressYaml := service.ReadYamlIngressUnder119(configmap.Name, configmap.Namespace, configmap.UserName); err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
			response.FailWithMessage("查询失败", c)
		} else {
			response.OkWithData(gin.H{"reIngressYaml": reIngressYaml}, c)
		}
	}

}
