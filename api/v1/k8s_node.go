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

// CreateNode
// @Tags Node
// @Summary 创建Node
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Node true "创建Node"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /node/createNode [post]
func CreateNode(c *gin.Context) {
	var Node model.NodeUser
	_ = c.ShouldBindJSON(&Node)
	if err := service.CreateNode(Node.Name, Node.UserName); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteNode
// @Tags Node
// @Summary 删除Node
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Node true "删除Node"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /node/deleteNode [delete]
func DeleteNode(c *gin.Context) {
	var Node model.NodeUser
	_ = c.ShouldBindJSON(&Node)
	if err := service.DeleteNode(Node.Name, Node.UserName); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteNodeByIds
// @Tags Node
// @Summary 批量删除Node
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Node"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /node/deleteNodeByIds [delete]
func DeleteNodeByIds(c *gin.Context) {
	var Names request.NamesReqUser //需要修改
	_ = c.ShouldBindJSON(&Names)
	if err := service.DeleteNodeByIds(Names); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateNode
// @Tags Node
// @Summary 更新Node
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Node true "更新Node"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /node/updateNode [put]
func UpdateNode(c *gin.Context) {
	var Node model.NodeUser
	_ = c.ShouldBindJSON(&Node)
	if err := service.UpdateNode(Node); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindNode
// @Tags Node
// @Summary 用id查询Node
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Node true "用id查询Node"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /node/findNode [get]
func FindNode(c *gin.Context) {

	//var node model.Node
	var node request.NodeSearchUser
	_ = c.ShouldBindQuery(&node)

	if err, reNode := service.GetNode(node.Name, node.UserName); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reNode": reNode}, c)
	}
}

// GetNodeList
// @Tags Node
// @Summary 分页获取Node列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.NodeSearch true "分页获取Node列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /node/getNodeList [get]
func GetNodeList(c *gin.Context) {
	var pageInfo request.NodeSearchUser

	_ = c.ShouldBindQuery(&pageInfo)

	if err, list, total := service.GetNodeInfoList(pageInfo); err != nil {
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
