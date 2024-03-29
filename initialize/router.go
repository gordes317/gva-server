package initialize

import (
	_ "gin-vue-admin/docs"
	"gin-vue-admin/global"
	"gin-vue-admin/middleware"
	"gin-vue-admin/router"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.GVA_LOG.Info("use middleware logger")
	// 跨域
	//Router.Use(middleware.Cors()) // 如需跨域可以打开
	global.GVA_LOG.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	PublicGroup := Router.Group("")
	{
		router.InitBaseRouter(PublicGroup)    // 注册基础功能路由 不做鉴权
		router.InitInitRouter(PublicGroup)    // 自动初始化相关
		router.InitWSocketRouter(PublicGroup) //web socket router

	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		router.InitApiRouter(PrivateGroup)                   // 注册功能api路由
		router.InitJwtRouter(PrivateGroup)                   // jwt相关路由
		router.InitUserRouter(PrivateGroup)                  // 注册用户路由
		router.InitMenuRouter(PrivateGroup)                  // 注册menu路由
		router.InitEmailRouter(PrivateGroup)                 // 邮件相关路由
		router.InitSystemRouter(PrivateGroup)                // system相关路由
		router.InitCasbinRouter(PrivateGroup)                // 权限相关路由
		router.InitCustomerRouter(PrivateGroup)              // 客户路由
		router.InitAutoCodeRouter(PrivateGroup)              // 创建自动化代码
		router.InitAuthorityRouter(PrivateGroup)             // 注册角色路由
		router.InitSimpleUploaderRouter(PrivateGroup)        // 断点续传（插件版）
		router.InitSysDictionaryRouter(PrivateGroup)         // 字典管理
		router.InitSysOperationRecordRouter(PrivateGroup)    // 操作记录
		router.InitSysDictionaryDetailRouter(PrivateGroup)   // 字典详情管理
		router.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由
		router.InitExcelRouter(PrivateGroup)                 // 表格导入导出
		router.InitStudentRouter(PrivateGroup)               //测试学生CRUD
		router.InitNamespaceRouter(PrivateGroup)             // 注册NameSpace路由
		router.InitPodRouter(PrivateGroup)                   //注册Pod路由
		router.InitBaseAPIRouter(PrivateGroup)               //注册基础接口信息 getSelectNamespaceList
		router.InitDeploymentRouter(PrivateGroup)            //注册Deployment路由
		router.InitServiceRouter(PrivateGroup)               //注册Service 路由
		router.InitNodeRouter(PrivateGroup)                  //注册Node路由
		router.InitStorageClassRouter(PrivateGroup)          //注册StorageClass路由
		router.InitPVRouter(PrivateGroup)                    //注册PV路由
		router.InitPVCRouter(PrivateGroup)                   //注册PVC路由
		router.InitConfigMapRouter(PrivateGroup)             //注册ConfigMap 路由
		router.InitSecretRouter(PrivateGroup)                //注册Secret路由
		router.InitTemplateRouter(PrivateGroup)              // 注册模版路由
		router.InitMultiClusterRouter(PrivateGroup)          //注册多集群管理路由
		router.InitJobRouter(PrivateGroup)                   //注册Job路由
		router.InitCronJobRouter(PrivateGroup)               //注册CronJob路由
		router.InitDaemonSetRouter(PrivateGroup)             //注册DaemonSet路由
		router.InitStatefulSetRouter(PrivateGroup)           //注册StatefulSet路由
		router.InitIngressRouter(PrivateGroup)               //注册Ingress路由
		router.InitPipelinesRouter(PrivateGroup)               //注册Pipelines路由

		// Code generated by gin-vue-admin Begin; DO NOT EDIT.
		// Code generated by gin-vue-admin End; DO NOT EDIT.
	}
	global.GVA_LOG.Info("router register success")
	return Router
}
