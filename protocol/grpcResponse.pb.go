// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpcResponse.proto

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

// 回复数据
type GrpcResponse struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpcResponse) Reset()         { *m = GrpcResponse{} }
func (m *GrpcResponse) String() string { return proto.CompactTextString(m) }
func (*GrpcResponse) ProtoMessage()    {}
func (*GrpcResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcResponse_e76eb89149b00f32, []int{0}
}
func (m *GrpcResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcResponse.Unmarshal(m, b)
}
func (m *GrpcResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcResponse.Marshal(b, m, deterministic)
}
func (dst *GrpcResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcResponse.Merge(dst, src)
}
func (m *GrpcResponse) XXX_Size() int {
	return xxx_messageInfo_GrpcResponse.Size(m)
}
func (m *GrpcResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcResponse proto.InternalMessageInfo

func (m *GrpcResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GrpcResponse)(nil), "protocol.GrpcResponse")
}

func init() { proto.RegisterFile("grpcResponse.proto", fileDescriptor_grpcResponse_e76eb89149b00f32) }

var fileDescriptor_grpcResponse_e76eb89149b00f32 = []byte{
	// 79 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0x2f, 0x2a, 0x48,
	0x0e, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x00, 0x53, 0xc9, 0xf9, 0x39, 0x4a, 0x4a, 0x5c, 0x3c, 0xee, 0x48, 0xf2, 0x42, 0x42, 0x5c, 0x2c,
	0x29, 0x89, 0x25, 0x89, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x60, 0x76, 0x12, 0x1b, 0x58,
	0xb5, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xd4, 0xf7, 0xdc, 0x8f, 0x4a, 0x00, 0x00, 0x00,
}