// /////////////////////////////////////////////////////////////////////////////
// 常量-接口-types

package app

import (
	"time"
)

// /////////////////////////////////////////////////////////////////////////////
// 常量

const (
	C_STOP_OUT_TIME          = 30 * time.Second // 关闭app的时候，超过此时间，就会强制关闭
	C_APP_TYPE_FRONTEND byte = 1                // 前端 app
	C_APP_TYPE_BACKEND  byte = 2                // 后端 app
)

// /////////////////////////////////////////////////////////////////////////////
// 接口

// handler 服务
type IHandler interface {
	OnData(data []byte) (bool, []byte) // 收到数据
}

// reomte 服务
type IRemote interface {
	OnData(data []byte) []byte // 收到数据
}
