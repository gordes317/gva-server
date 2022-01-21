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

// CreatePod
// @Tags Pod
// @Summary 创建Pod
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Pod true "创建Pod"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Pod/createPod [post]
func CreatePod(c *gin.Context) {
	var Pod model.PodUser
	_ = c.ShouldBindJSON(&Pod)
	if err := service.CreatePod(Pod.Name, Pod.UserName); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePod
// @Tags Pod
// @Summary 删除Pod
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Pod true "删除Pod"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Pod/deletePod [delete]
func DeletePod(c *gin.Context) {
	var Pod model.PodUser
	_ = c.ShouldBindJSON(&Pod)
	if err := service.DeletePod(Pod.Name, Pod.Namespace, Pod.UserName); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePodByIds
// @Tags Pod
// @Summary 批量删除Pod
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Pod"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Pod/deletePodByIds [delete]
func DeletePodByIds(c *gin.Context) {
	var Names request.NamesReqUser //需要修改
	_ = c.ShouldBindJSON(&Names)
	if err := service.DeletePodByIds(Names); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePod
// @Tags Pod
// @Summary 更新Pod
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Pod true "更新Pod"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Pod/updatePod [put]
func UpdatePod(c *gin.Context) {
	var Pod model.PodUser
	_ = c.ShouldBindJSON(&Pod)
	if err := service.UpdatePod(Pod, Pod.UserName); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPod
// @Tags Pod
// @Summary 用id查询Pod
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Pod true "用id查询Pod"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Pod/findPod [get]
func FindPod(c *gin.Context) {

	var namespace model.PodUser
	_ = c.ShouldBindQuery(&namespace)

	if err, rePod := service.GetPod(namespace.Name, namespace.UserName); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rePod": rePod}, c)
	}
}

// GetPodList
// @Tags Pod
// @Summary 分页获取Pod列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PodSearch true "分页获取Pod列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /namespace/getPodList [get]
func GetPodList(c *gin.Context) {
	var pageInfo request.PodSearchUser

	_ = c.ShouldBindQuery(&pageInfo)

	if err, list, total := service.GetPodInfoList(pageInfo); err != nil {
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

// GetPodLog
// @Tags Pod
// @Summary 获取pod日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Pod true "用id查询Pod"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pod/getPodLog [get]
func GetPodLog(c *gin.Context) {

	var pod model.PodUser
	_ = c.ShouldBindQuery(&pod)

	if err, rePod := service.GetPodLog(pod.Name, pod.Namespace, pod.ContainerName, pod.UserName); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rePod": rePod}, c)
	}
}

// ApplyYamlPod
// @Tags Pod
// @Summary 更新Pod Yaml
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Pod true "更新Pod Yaml"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pod/applyYamlPod [put]
func ApplyYamlPod(c *gin.Context) {
	var pod model.PodUser
	_ = c.ShouldBindJSON(&pod)
	if err := service.ApplyYamlPod(pod, pod.UserName); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// ReadYamlPod
// @Tags Pod
// @Summary 获取Pod对应yaml
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Pod true "获取Pod对应yaml"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pod/readYamlPod [get]
func ReadYamlPod(c *gin.Context) {

	var pod model.PodUser
	_ = c.ShouldBindQuery(&pod)

	if err, rePodYaml := service.ReadYamlPod(pod.Name, pod.Namespace, pod.UserName); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rePodYaml": rePodYaml}, c)
	}
}
