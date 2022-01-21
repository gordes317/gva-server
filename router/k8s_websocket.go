
package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitWSocketRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base").Use(middleware.NeedInit())
	{
		//pod console websocket
		BaseRouter.GET("termialSocket", func(c *gin.Context) {v1.WsHandler(c.Writer, c.Request)})
		BaseRouter.GET("buildPlanSocket", func(c *gin.Context) {v1.WsGetBuildHistory(c.Writer, c.Request)})
		BaseRouter.GET("taskRunSocket", func(c *gin.Context) {v1.WsGetTaskRun(c.Writer, c.Request)})
	}
	return BaseRouter

}


