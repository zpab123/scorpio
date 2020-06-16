// /////////////////////////////////////////////////////////////////////////////
// sco

package sco

import (
	"github.com/zpab123/sco/app"
)

// /////////////////////////////////////////////////////////////////////////////
// 变量

// app 单利
var scoApp *app.Application = app.NewApplication()

// /////////////////////////////////////////////////////////////////////////////
// public-api

// 启动服务器
func Run() {
	scoApp.Run()
}

// 获取app
func GetApp() *app.Application {
	return scoApp
}

// 设置 handler
func SetHandler(handler app.IHandler) {
	scoApp.SetHandler(handler)
}

// 以 rpc 的方式获取远程数据
func Call(mid uint16, data []byte) []byte {
	return scoApp.Call(mid, data)
}
