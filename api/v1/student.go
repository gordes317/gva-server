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

// @Tags Student
// @Summary 创建Student
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Student true "创建Student"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /student/createStudent [post]
func CreateStudent(c *gin.Context) {
	var student model.Student
	_ = c.ShouldBindJSON(&student)
	if err := service.CreateStudent(student); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Student
// @Summary 删除Student
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Student true "删除Student"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /student/deleteStudent [delete]
func DeleteStudent(c *gin.Context) {
	var student model.Student
	_ = c.ShouldBindJSON(&student)
	if err := service.DeleteStudent(student); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Student
// @Summary 批量删除Student
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Student"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /student/deleteStudentByIds [delete]
func DeleteStudentByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteStudentByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Student
// @Summary 更新Student
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Student true "更新Student"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /student/updateStudent [put]
func UpdateStudent(c *gin.Context) {
	var student model.Student
	_ = c.ShouldBindJSON(&student)
	if err := service.UpdateStudent(student); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Student
// @Summary 用id查询Student
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Student true "用id查询Student"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /student/findStudent [get]
func FindStudent(c *gin.Context) {
	var student model.Student
	_ = c.ShouldBindQuery(&student)
	if err, restudent := service.GetStudent(student.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"restudent": restudent}, c)
	}
}

// @Tags Student
// @Summary 分页获取Student列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.StudentSearch true "分页获取Student列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /student/getStudentList [get]
func GetStudentList(c *gin.Context) {
	var pageInfo request.StudentSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetStudentInfoList(pageInfo); err != nil {
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
