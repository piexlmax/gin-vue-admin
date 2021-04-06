package core

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
	"gin-vue-admin/service"
	"go.uber.org/zap"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint { // 初始化redis服务
		initialize.Redis()
	}
	engine := initialize.Routers()
	engine.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, engine)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 Gin-Vue-Admin
	当前版本:V2.4.0
    加群方式:微信号：shouzi_1994 QQ群：622360840
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
	如果项目让您获得了收益，希望您能请团队喝杯可乐:https://www.gin-vue-admin.com/docs/coffee
`, address)

	if global.GVA_CONFIG.System.AutoRegisterRouter { // 自动把路由注册到sys_apis表
		go func() {
			_ = service.AutoRegisterRouter(engine)
		}()
	}
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
