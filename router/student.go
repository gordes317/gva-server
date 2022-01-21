package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitStudentRouter(Router *gin.RouterGroup) {
	StudentRouter := Router.Group("student").Use(middleware.OperationRecord())
	{
		StudentRouter.POST("createStudent", v1.CreateStudent)   // 新建Student
		StudentRouter.DELETE("deleteStudent", v1.DeleteStudent) // 删除Student
		StudentRouter.DELETE("deleteStudentByIds", v1.DeleteStudentByIds) // 批量删除Student
		StudentRouter.PUT("updateStudent", v1.UpdateStudent)    // 更新Student
		StudentRouter.GET("findStudent", v1.FindStudent)        // 根据ID获取Student
		StudentRouter.GET("getStudentList", v1.GetStudentList)  // 获取Student列表
	}
}
