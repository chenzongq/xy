package core

import (
	"github.com/xyxy/admin-service/global"
	"github.com/xyxy/admin-service/initialize"
	"github.com/xyxy/admin-service/service/system"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.XY_CONFIG.System.UseMultipoint || global.XY_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	// 从db加载jwt数据
	if global.XY_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	//address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	//s := initServer(address, Router)
	//// 保证文本顺序输出
	//// In order to ensure that the text order output can be deleted
	//time.Sleep(10 * time.Microsecond)
	//global.GVA_LOG.Info("server run success on ", zap.String("address", address))
	//global.GVA_LOG.Error(s.ListenAndServe().Error())
}
