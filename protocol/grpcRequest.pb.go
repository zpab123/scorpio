// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpcRequest.proto

package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 请求数据
type GrpcRequest struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpcRequest) Reset()         { *m = GrpcRequest{} }
func (m *GrpcRequest) String() string { return proto.CompactTextString(m) }
func (*GrpcRequest) ProtoMessage()    {}
func (*GrpcRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcRequest_53672b4b8a3ea362, []int{0}
}
func (m *GrpcRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcRequest.Unmarshal(m, b)
}
func (m *GrpcRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcRequest.Marshal(b, m, deterministic)
}
func (dst *GrpcRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcRequest.Merge(dst, src)
}
func (m *GrpcRequest) XXX_Size() int {
	return xxx_messageInfo_GrpcRequest.Size(m)
}
func (m *GrpcRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcRequest proto.InternalMessageInfo

func (m *GrpcRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GrpcRequest)(nil), "protocol.GrpcRequest")
}

func init() { proto.RegisterFile("grpcRequest.proto", fileDescriptor_grpcRequest_53672b4b8a3ea362) }

var fileDescriptor_grpcRequest_53672b4b8a3ea362 = []byte{
	// 78 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x2f, 0x2a, 0x48,
	0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00,
	0x53, 0xc9, 0xf9, 0x39, 0x4a, 0x8a, 0x5c, 0xdc, 0xee, 0x08, 0x69, 0x21, 0x21, 0x2e, 0x96, 0x94,
	0xc4, 0x92, 0x44, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x9e, 0x20, 0x30, 0x3b, 0x89, 0x0d, 0xac, 0xd8,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x13, 0xd5, 0x4e, 0xfc, 0x48, 0x00, 0x00, 0x00,
}
