// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpccontnet/pb/testpb.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MathArgRequest struct {
	A                    int64    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int64    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MathArgRequest) Reset()         { *m = MathArgRequest{} }
func (m *MathArgRequest) String() string { return proto.CompactTextString(m) }
func (*MathArgRequest) ProtoMessage()    {}
func (*MathArgRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5dadfce7ea647f2, []int{0}
}

func (m *MathArgRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MathArgRequest.Unmarshal(m, b)
}
func (m *MathArgRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MathArgRequest.Marshal(b, m, deterministic)
}
func (m *MathArgRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MathArgRequest.Merge(m, src)
}
func (m *MathArgRequest) XXX_Size() int {
	return xxx_messageInfo_MathArgRequest.Size(m)
}
func (m *MathArgRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MathArgRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MathArgRequest proto.InternalMessageInfo

func (m *MathArgRequest) GetA() int64 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *MathArgRequest) GetB() int64 {
	if m != nil {
		return m.B
	}
	return 0
}

type MathMutliResponse struct {
	Ab1                  int64    `protobuf:"varint,1,opt,name=ab1,proto3" json:"ab1,omitempty"`
	Ab2                  int64    `protobuf:"varint,2,opt,name=ab2,proto3" json:"ab2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MathMutliResponse) Reset()         { *m = MathMutliResponse{} }
func (m *MathMutliResponse) String() string { return proto.CompactTextString(m) }
func (*MathMutliResponse) ProtoMessage()    {}
func (*MathMutliResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5dadfce7ea647f2, []int{1}
}

func (m *MathMutliResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MathMutliResponse.Unmarshal(m, b)
}
func (m *MathMutliResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MathMutliResponse.Marshal(b, m, deterministic)
}
func (m *MathMutliResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MathMutliResponse.Merge(m, src)
}
func (m *MathMutliResponse) XXX_Size() int {
	return xxx_messageInfo_MathMutliResponse.Size(m)
}
func (m *MathMutliResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MathMutliResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MathMutliResponse proto.InternalMessageInfo

func (m *MathMutliResponse) GetAb1() int64 {
	if m != nil {
		return m.Ab1
	}
	return 0
}

func (m *MathMutliResponse) GetAb2() int64 {
	if m != nil {
		return m.Ab2
	}
	return 0
}

func init() {
	proto.RegisterType((*MathArgRequest)(nil), "pb.MathArgRequest")
	proto.RegisterType((*MathMutliResponse)(nil), "pb.MathMutliResponse")
}

func init() { proto.RegisterFile("grpccontnet/pb/testpb.proto", fileDescriptor_b5dadfce7ea647f2) }

var fileDescriptor_b5dadfce7ea647f2 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0x2f, 0x2a, 0x48,
	0x4e, 0xce, 0xcf, 0x2b, 0xc9, 0x4b, 0x2d, 0xd1, 0x2f, 0x48, 0xd2, 0x2f, 0x49, 0x2d, 0x2e, 0x29,
	0x48, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe1, 0xe2, 0xf3,
	0x4d, 0x2c, 0xc9, 0x70, 0x2c, 0x4a, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xe2, 0xe1,
	0x62, 0x4c, 0x94, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x62, 0x4c, 0x04, 0xf1, 0x92, 0x24, 0x98,
	0x20, 0xbc, 0x24, 0x25, 0x73, 0x2e, 0x41, 0x90, 0x6a, 0xdf, 0xd2, 0x92, 0x9c, 0xcc, 0xa0, 0xd4,
	0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x01, 0x2e, 0xe6, 0xc4, 0x24, 0x43, 0xa8, 0x16, 0x10,
	0x13, 0x22, 0x62, 0x04, 0xd5, 0x06, 0x62, 0x1a, 0xb9, 0x72, 0x71, 0x83, 0x34, 0x06, 0xa7, 0x16,
	0x95, 0x65, 0x26, 0xa7, 0x0a, 0x99, 0x71, 0x71, 0xc2, 0xcd, 0x11, 0x12, 0xd2, 0x2b, 0x48, 0xd2,
	0x43, 0x75, 0x84, 0x94, 0x28, 0x4c, 0x0c, 0xc5, 0xaa, 0x24, 0x36, 0xb0, 0xc3, 0x8d, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x90, 0xaa, 0x62, 0x67, 0xd7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MathServiceClient is the client API for MathService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MathServiceClient interface {
	MathMutli(ctx context.Context, in *MathArgRequest, opts ...grpc.CallOption) (*MathMutliResponse, error)
}

type mathServiceClient struct {
	cc *grpc.ClientConn
}

func NewMathServiceClient(cc *grpc.ClientConn) MathServiceClient {
	return &mathServiceClient{cc}
}

func (c *mathServiceClient) MathMutli(ctx context.Context, in *MathArgRequest, opts ...grpc.CallOption) (*MathMutliResponse, error) {
	out := new(MathMutliResponse)
	err := c.cc.Invoke(ctx, "/pb.MathService/MathMutli", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MathServiceServer is the server API for MathService service.
type MathServiceServer interface {
	MathMutli(context.Context, *MathArgRequest) (*MathMutliResponse, error)
}

// UnimplementedMathServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMathServiceServer struct {
}

func (*UnimplementedMathServiceServer) MathMutli(ctx context.Context, req *MathArgRequest) (*MathMutliResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MathMutli not implemented")
}

func RegisterMathServiceServer(s *grpc.Server, srv MathServiceServer) {
	s.RegisterService(&_MathService_serviceDesc, srv)
}

func _MathService_MathMutli_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MathArgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServiceServer).MathMutli(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MathService/MathMutli",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServiceServer).MathMutli(ctx, req.(*MathArgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MathService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MathService",
	HandlerType: (*MathServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MathMutli",
			Handler:    _MathService_MathMutli_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpccontnet/pb/testpb.proto",
}
