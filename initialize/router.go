package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/tcxcn/finance-a/docs"
	"github.com/tcxcn/finance-a/middleware"
	"github.com/tcxcn/finance-a/router"
	"github.com/xyxy/admin-service/global"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()
	InstallPlugin(Router) // 安装插件
	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example
	// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build。在打开下面4行注释
	// Router.LoadHTMLGlob("./dist/*.html") // npm打包成dist的路径
	// Router.Static("/favicon.ico", "./dist/favicon.ico")
	// Router.Static("/static", "./dist/assets")   // dist里面的静态资源
	// Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面

	Router.StaticFS(global.GVA_CONFIG.Local.StorePath, http.Dir(global.GVA_CONFIG.Local.StorePath)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 如果需要使用https 请打开此中间件 然后前往 core/server.go 将启动模式 更变为 Router.RunTLS("端口","你的cre/pem文件","你的key文件")
	// 跨域，如需跨域可以打开下面的注释
	Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	global.GVA_LOG.Info("use middleware cors")
	docs.SwaggerInfo.BasePath = global.GVA_CONFIG.System.RouterPrefix
	Router.GET(global.GVA_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		//systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		//systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}
	//PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	//PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	//{
	//	systemRouter.InitApiRouter(PrivateGroup)                 // 注册功能api路由
	//systemRouter.InitJwtRouter(PrivateGroup) // jwt相关路由
	//	systemRouter.InitUserRouter(PrivateGroup)                // 注册用户路由
	//	systemRouter.InitMenuRouter(PrivateGroup)                // 注册menu路由
	//	systemRouter.InitSystemRouter(PrivateGroup)              // system相关路由
	//	systemRouter.InitCasbinRouter(PrivateGroup)              // 权限相关路由
	//	systemRouter.InitAutoCodeRouter(PrivateGroup)            // 创建自动化代码
	//	systemRouter.InitAuthorityRouter(PrivateGroup)           // 注册角色路由
	//	systemRouter.InitSysDictionaryRouter(PrivateGroup)       // 字典管理
	//	systemRouter.InitAutoCodeHistoryRouter(PrivateGroup)     // 自动化代码历史
	//	systemRouter.InitSysOperationRecordRouter(PrivateGroup)  // 操作记录
	//	systemRouter.InitSysDictionaryDetailRouter(PrivateGroup) // 字典详情管理
	//	systemRouter.InitAuthorityBtnRouterRouter(PrivateGroup)  // 字典详情管理
	//
	//	exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由
	//
	//}
	{
		//tenantRouter := router.RouterGroupApp.Tenant //区域管理
		//tenantRouter.InitTenantsRouter(PrivateGroup)
	}
	{
		//financeRouter := router.RouterGroupApp.Finance                       //财务报表
		//financeRouter.InitFinanceReconciliationRouter(PrivateGroup)          //财务报表
		//financeRouter.InitFinanceOverdueSettingsRouter(PrivateGroup)         //超期设置
		//financeRouter.InitFinanceLogRouter(PrivateGroup)                     //操作日志
		//financeRouter.InitAccountingSettingsRouter(PrivateGroup)             //核账设置
		//financeRouter.InitFinanceZpRecordsRouter(PrivateGroup)               //直赔记录
		//financeRouter.InitFinanceDownPaymentRecordRouter(PrivateGroup)       //向下付款
		//financeRouter.InitFinanceInvoiceInformationRouter(PrivateGroup)      //发票信息
		//financeRouter.InitFinanceInvoiceRouter(PrivateGroup)                 //发票列表
		//financeRouter.InitFinanceDashboardRouter(PrivateGroup)               //仪表盘
		//financeRouter.InitOutputValueRouterRouter(PrivateGroup, PublicGroup) //产值列表
		//financeRouter.InitPayableConfigRouter(PrivateGroup)                  //合作伙伴向下支付配置
		//financeRouter.InitPayableManageRouter(PrivateGroup)                  //合作伙伴向下支付管理
		//financeRouter.InitHistoryDataRouterRouter(PrivateGroup)
		////  financeRouter.InitDriverRecordRouter(PrivateGroup)                   //驾驶员列表
		////	financeRouter.InitSalaryRouter(PrivateGroup)                         //驾驶员工资列表
		//financeRouter.InitPrepaidRouter(PrivateGroup)              //预付
		//financeRouter.InitDeductionRouterRouter(PrivateGroup)      //预付对应的抵减
		//financeRouter.InitPaymentRecordsRouter(PrivateGroup)       //预付对应的付款
		//financeRouter.InitDownPaymentDashboardRouter(PrivateGroup) // 向下付款仪表盘统计
		//financeRouter.InitBadDebtRouter(PrivateGroup)              // 坏账
	}
	{
		//messageRouter := router.RouterGroupApp.Message
		//messageRouter.InitMessageRouter(PrivateGroup)
	}
	{
		//initRouter := router.RouterGroupApp.Init
		//initRouter.InitApiRouter(PublicGroup)
	}

	{
		//v2Router := router.RouterGroupApp.V2
		//v2Router.InitV2Router(PublicGroup)
	}
	// 测试路由
	{
		//PublicGroup.GET("/fix/outDispatcher", fix.OutDispatcher)
		//PublicGroup.GET("/fix/league", fix.League)
		//PublicGroup.GET("/fix/genNotPaidByExcels", fix.GenNotPaidByExcels)
		//PublicGroup.GET("/fix/charge", fix.Charge)
		//PublicGroup.GET("/fix/badDebtInit", fix.BadDebtInitData)
		//PublicGroup.GET("/fix/processNotPaid", fix.ProcessNotPaid)
	}

	global.XY_LOG.Info("router register success")
	return Router
}
