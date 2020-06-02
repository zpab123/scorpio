// /////////////////////////////////////////////////////////////////////////////
// grpc 消息服务

package rpc

import (
	"context"

	"github.com/zpab123/sco/protocol"
)

// /////////////////////////////////////////////////////////////////////////////
// GrpcService

// grpc 消息服务
type GrpcService struct {
	remoteService IRemoteService // remote 服务
}

// 新建1个 GrpcService
func NewGrpcService(svc IRemoteService) protocol.GrpcServer {
	g := GrpcService{
		remoteService: svc,
	}

	return &g
}

// hander 调用
func (this GrpcService) Call(ctx context.Context, req *protocol.GrpcRequest) (*protocol.GrpcResponse, error) {
	if nil != this.remoteService {
		return nil, nil
	}

	res := this.remoteService.OnData(req.Data)
	req := protocol.GrpcResponse{
		Data: res,
	}

	return req, nil
}

// remote 调用
