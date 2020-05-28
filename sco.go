// /////////////////////////////////////////////////////////////////////////////
// sco

package sco

import (
	"github.com/zpab123/sco/app"
	"github.com/zpab123/sco/network"
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

// 注册 handler
func RegisterHandler(handler network.IHandler) {
	scoApp.RegisterHandler(handler)
}
