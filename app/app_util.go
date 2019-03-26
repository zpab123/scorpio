// /////////////////////////////////////////////////////////////////////////////
// Application 一些辅助函数

package app

import (
	"flag"
	"os"

	"github.com/zpab123/sco/config" // 配置管理
	"github.com/zpab123/zaplog"     // log 库
)

// 完成 app 的默认设置
func defaultConfig(app *Application) {
	// 解析启动参数
	parseArgs(app)

	// 获取服务器信息

	// 设置 log 信息

	// 默认组件参数
}

// 解析 命令行参数
func parseArgs(app *Application) {
	// 参数定义
	//serverType := flag.String("type", "serverType", "server type") // 服务器类型，例如 gate connect area ...等
	//gid := flag.Uint("gid", 0, "gid")                                       // 服务器进程id
	name := flag.String("name", "gate_1", "server name") // 服务器名字
	//frontend := flag.Bool("frontend", false, "is frontend server")          // 是否是前端服务器
	//host := flag.String("host", "127.0.0.1", "server host")                 // 服务器IP地址
	//port := flag.Uint("port", 0, "server port")                             // 服务器端口
	//clientHost := flag.String("clientHost", "127.0.0.1", "for client host") // 面向客户端的 IP地址
	//cTcpPort := flag.Uint("cTcpPort", 0, "tcp port")                        // tcp 连接端口
	//cWsPort := flag.Uint("cWsPort", 0, "websocket port")                    // websocket 连接端口

	// 解析参数
	flag.Parse()

	// 赋值
	//cmdParam := &cmd.CmdParam{
	//ServerType: *serverType,
	//Gid:        *gid,
	//Name: *name,
	//Frontend:   *frontend,
	//Host:       *host,
	//Port:       *port,
	//ClientHost: *clientHost,
	//CTcpPort:   *cTcpPort,
	//CWsPort:    *cWsPort,
	//}

	// 设置 app 名字
	app.baseInfo.Name = *name
}

// 获取 server.json 信息
func getServerJson(app *Application) {
	// 根据 AppType 和 Name 获取 服务器配置参数
	appType := app.baseInfo.AppType
	name := app.baseInfo.Name
	list, ok := config.GetServerMap()[appType]

	if nil == list || len(list) <= 0 || !ok {
		zaplog.Fatal("app 获取 appType 信息失败。 appType=%s", appType)

		os.Exit(1)
	}

	// 获取服务器信息
	for _, info := range list {
		if info.Name == name {
			app.serverInfo = info

			break
		}
	}

	if app.serverInfo == nil {
		zaplog.Fatal("app 获取 server.json 信息失败。 appName=%s", app.baseInfo.Name)

		os.Exit(1)
	}
}
