package app

import (
	"context"

	"github.com/zpab123/sco/protocol"
)

type Remote struct {
}

func (this *Remote) Call(ctx context.Context, rq *protocol.GrpcRequest) (*protocol.GrpcResponse, error) {
	return nil, nil
}
