// /////////////////////////////////////////////////////////////////////////////
// 常量-接扣-types

package rpc

import (
	// goContext "context"

	"github.com/zpab123/sco/discovery" // 服务发现
	"github.com/zpab123/sco/model"     // 全局模型
	"github.com/zpab123/sco/protocol"  // 消息协议
	"golang.org/x/net/context"         // golang 上下文
	"google.golang.org/grpc"           // grpc
)

// /////////////////////////////////////////////////////////////////////////////
// 常量
const (
	C_SVC_NAME    = "sco.rpcService" // sco rpc服务名称
	C_METHOD_CALL = "Call"           // sco Call 方法
	C_RPC_GC      = "rpc.GrpcClient" // GrpcClient 组件名字
)

// /////////////////////////////////////////////////////////////////////////////
// 接口

// rpc server 服务
type IServer interface {
	SetRpcService(ss protocol.ScoGrpcServer) // 设置引擎服务
	model.IComponent                         // 接口继承：组件
}

// rpc client 服务
type IClient interface {
	model.IComponent    // 接口继承：组件
	discovery.IListener // 接口继承：服务发现侦听
}

// sco 引擎服务
type IScoService interface {
	Call(ctx context.Context, req *protocol.RpcRequest) (*protocol.RpcResponse, error) // 方法调用
}

// 连接对象接口
type IConn interface {
	Call(ctx context.Context, req *protocol.RpcRequest, opts ...grpc.CallOption) (*protocol.RpcResponse, error) // 远程调用
}
