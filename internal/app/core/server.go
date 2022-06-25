package core

import (
	"fmt"
	"go-admin/internal/app/global"
	"go-admin/internal/app/initialize"
	"time"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.SYS_CONFIG.System.UseMultipoint || global.SYS_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	Router := initialize.Routers()
	//Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.SYS_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.SYS_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 go-admin
	当前版本:V0.0.1
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
`, address)
	global.SYS_LOG.Error(s.ListenAndServe().Error())
}
