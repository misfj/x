package api

import (
	_ "coredx/core/api/docs"
	"coredx/core/api/middleware"
	"coredx/core/api/service"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) initRoute() {
	//绑定全局跨域
	s.eng.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//业务应用登录
	s.eng.POST("/v1/app/login", service.AppLogin)
	s.eng.Use(middleware.Cors(), middleware.Logger())
	appG := s.eng.Group("/v1/app")
	{
		//业务应用注册
		appG.POST("/register", service.AppRegister)
		//后续接口可能需要JWT进行支持,届时分组(平台数据目录注册)
		appG.POST("/directory/register", service.DirectRegister)
		//测试接口,验证API服务是否启动
		appG.POST("/test", service.AppTest)
	}

	authG := appG.Use(middleware.Auth())
	{
		//业务应用代理用户进行注册
		authG.POST("/user/register", service.AppUserRegister)
		//业务应用代理用户进行用户账号进行绑定,绑定多个平台的用户账号
		authG.POST("/account/bind", service.AppAccountBind)
		//获取业务应用信息
		authG.POST("/info", service.AppInfo)
		//业务应用程序代理用户获取用户证书
		authG.POST("/user/ca", service.AppUserCA)

		//数据上传
		/*todo 如果对方的回调地址是核心服务器API地址的话记得跳过某地址,
		todo 因为会走路由组,造成请求中间件Panic,后续优化。配置一下即可
		*/
		//业务应用代理用户进行数据存储,支持Form-data格式格式和Body RawJson格式
		authG.POST("/service/data/store", service.Upload)
		//业务引用进行数据目录订阅
		authG.POST("/service/directory/subscribe", service.Subscribe)
		//查询业务应用订阅的服务
		authG.GET("/service/subscribe", service.SubscribeService)
		//业务应用购买服务
		authG.POST("/service/buy", service.BuyService)
		//业务应用获取服务列表
		authG.POST("/service/list", service.ServiceList)
		//获取数据列表
		authG.POST("/service/data/list", service.DataList)
		//添加水印
		authG.POST("/service/water", service.Water)
		//请求授权确权
		authG.POST("/service/cg", service.CG)
		//请求智能服务
		authG.POST("/service/ai", service.AI)
		//请求支付服务
		authG.POST("/service/pay", service.Pay)
		//请求加密解密(暂时指支持了form-data接口)
		authG.POST("/service/crypto", service.CryptoXX)
		//业务应用请求数据存储
		authG.POST("/service/data/store2", service.Store)
		//获取用户空间信息
		authG.POST("/user/space/info", service.SpaceInfo)
		//用户升级个人空间
		authG.POST("/user/space/expand", service.Expand)
		//请求更新数据文件
		authG.POST("/data/update", service.Update)
		//请求删除数据文件
		authG.POST("/data/delete", service.Delete)
		//请求共享数据
		authG.POST("/data/share", service.Share)
		//设置数据共享策略
		authG.POST("/data/share/rules", service.Rules)

	}
}
