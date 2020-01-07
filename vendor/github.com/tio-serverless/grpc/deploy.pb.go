// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deploy.proto

package tio_control_v1

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

type DeployRequest struct {
	Image                string   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Config               string   `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	Sid                  int32    `protobuf:"varint,3,opt,name=sid,proto3" json:"sid,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Stype                int32    `protobuf:"varint,5,opt,name=stype,proto3" json:"stype,omitempty"`
	InstanceNum          int32    `protobuf:"varint,6,opt,name=instanceNum,proto3" json:"instanceNum,omitempty"`
	InstanceMultiple     float64  `protobuf:"fixed64,7,opt,name=instanceMultiple,proto3" json:"instanceMultiple,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeployRequest) Reset()         { *m = DeployRequest{} }
func (m *DeployRequest) String() string { return proto.CompactTextString(m) }
func (*DeployRequest) ProtoMessage()    {}
func (*DeployRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_05f09e103004e384, []int{0}
}

func (m *DeployRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeployRequest.Unmarshal(m, b)
}
func (m *DeployRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeployRequest.Marshal(b, m, deterministic)
}
func (m *DeployRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployRequest.Merge(m, src)
}
func (m *DeployRequest) XXX_Size() int {
	return xxx_messageInfo_DeployRequest.Size(m)
}
func (m *DeployRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeployRequest proto.InternalMessageInfo

func (m *DeployRequest) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *DeployRequest) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

func (m *DeployRequest) GetSid() int32 {
	if m != nil {
		return m.Sid
	}
	return 0
}

func (m *DeployRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeployRequest) GetStype() int32 {
	if m != nil {
		return m.Stype
	}
	return 0
}

func (m *DeployRequest) GetInstanceNum() int32 {
	if m != nil {
		return m.InstanceNum
	}
	return 0
}

func (m *DeployRequest) GetInstanceMultiple() float64 {
	if m != nil {
		return m.InstanceMultiple
	}
	return 0
}

func init() {
	proto.RegisterType((*DeployRequest)(nil), "DeployRequest")
}

func init() { proto.RegisterFile("deploy.proto", fileDescriptor_05f09e103004e384) }

var fileDescriptor_05f09e103004e384 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0xff, 0xf9, 0xdb, 0x44, 0x72, 0xa3, 0x25, 0x5c, 0x44, 0x86, 0xae, 0x42, 0x57, 0xb1,
	0x8b, 0x80, 0xfa, 0x06, 0xe2, 0xc6, 0x85, 0x5d, 0xa4, 0x5d, 0xb9, 0x8b, 0xd3, 0xdb, 0x32, 0x30,
	0x99, 0x1b, 0x93, 0x49, 0x25, 0xcf, 0xe1, 0x43, 0xf9, 0x5a, 0xd2, 0x19, 0x05, 0x45, 0x10, 0x77,
	0xf7, 0x3b, 0xe7, 0x63, 0xe0, 0x0c, 0x9c, 0x6e, 0xa9, 0x35, 0x3c, 0x96, 0x6d, 0xc7, 0x8e, 0xe7,
	0xa9, 0x62, 0xdb, 0xbb, 0x00, 0x8b, 0x37, 0x01, 0x67, 0x77, 0xbe, 0xad, 0xe8, 0x79, 0xa0, 0xde,
	0xe1, 0x39, 0x44, 0xba, 0xa9, 0xf7, 0x24, 0x45, 0x2e, 0x8a, 0xa4, 0x0a, 0x80, 0x17, 0x10, 0x2b,
	0xb6, 0x3b, 0xbd, 0x97, 0xff, 0x7d, 0xfc, 0x41, 0x98, 0xc1, 0xa4, 0xd7, 0x5b, 0x39, 0xc9, 0x45,
	0x11, 0x55, 0xc7, 0x13, 0x11, 0xa6, 0xb6, 0x6e, 0x48, 0x4e, 0xbd, 0xe7, 0xef, 0xe3, 0x9b, 0xbd,
	0x1b, 0x5b, 0x92, 0x91, 0xf7, 0x02, 0x60, 0x0e, 0xa9, 0xb6, 0xbd, 0xab, 0xad, 0xa2, 0xd5, 0xd0,
	0xc8, 0xd8, 0x77, 0x5f, 0x23, 0x5c, 0x42, 0xf6, 0x89, 0x0f, 0x83, 0x71, 0xba, 0x35, 0x24, 0x4f,
	0x72, 0x51, 0x88, 0xea, 0x47, 0x7e, 0xfd, 0x2a, 0x20, 0xdb, 0x68, 0x0e, 0x63, 0xd6, 0xd4, 0x1d,
	0xb4, 0x22, 0x2c, 0x20, 0x59, 0xd1, 0x4b, 0xc8, 0x70, 0x56, 0x7e, 0x5b, 0x3a, 0x4f, 0xca, 0x8d,
	0xe6, 0x8a, 0x5a, 0x33, 0x2e, 0xfe, 0xe1, 0x12, 0xd2, 0xb5, 0xaa, 0x4d, 0xfd, 0x17, 0xf7, 0x12,
	0x20, 0xb4, 0xf7, 0x76, 0xc7, 0xbf, 0xaa, 0xb7, 0xd9, 0xe3, 0xcc, 0x69, 0x2e, 0x15, 0x5b, 0xd7,
	0xb1, 0x29, 0x0f, 0x57, 0x4f, 0xb1, 0xff, 0xf8, 0x9b, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x38,
	0x35, 0x0a, 0x5c, 0x95, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TioDeployServiceClient is the client API for TioDeployService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TioDeployServiceClient interface {
	NewDeploy(ctx context.Context, in *DeployRequest, opts ...grpc.CallOption) (*TioReply, error)
	ScalaDeploy(ctx context.Context, in *DeployRequest, opts ...grpc.CallOption) (*TioReply, error)
	DeployInfo(ctx context.Context, in *DeployRequest, opts ...grpc.CallOption) (*TioReply, error)
}

type tioDeployServiceClient struct {
	cc *grpc.ClientConn
}

func NewTioDeployServiceClient(cc *grpc.ClientConn) TioDeployServiceClient {
	return &tioDeployServiceClient{cc}
}

func (c *tioDeployServiceClient) NewDeploy(ctx context.Context, in *DeployRequest, opts ...grpc.CallOption) (*TioReply, error) {
	out := new(TioReply)
	err := c.cc.Invoke(ctx, "/TioDeployService/NewDeploy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tioDeployServiceClient) ScalaDeploy(ctx context.Context, in *DeployRequest, opts ...grpc.CallOption) (*TioReply, error) {
	out := new(TioReply)
	err := c.cc.Invoke(ctx, "/TioDeployService/ScalaDeploy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tioDeployServiceClient) DeployInfo(ctx context.Context, in *DeployRequest, opts ...grpc.CallOption) (*TioReply, error) {
	out := new(TioReply)
	err := c.cc.Invoke(ctx, "/TioDeployService/DeployInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TioDeployServiceServer is the server API for TioDeployService service.
type TioDeployServiceServer interface {
	NewDeploy(context.Context, *DeployRequest) (*TioReply, error)
	ScalaDeploy(context.Context, *DeployRequest) (*TioReply, error)
	DeployInfo(context.Context, *DeployRequest) (*TioReply, error)
}

// UnimplementedTioDeployServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTioDeployServiceServer struct {
}

func (*UnimplementedTioDeployServiceServer) NewDeploy(ctx context.Context, req *DeployRequest) (*TioReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewDeploy not implemented")
}
func (*UnimplementedTioDeployServiceServer) ScalaDeploy(ctx context.Context, req *DeployRequest) (*TioReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScalaDeploy not implemented")
}
func (*UnimplementedTioDeployServiceServer) DeployInfo(ctx context.Context, req *DeployRequest) (*TioReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeployInfo not implemented")
}

func RegisterTioDeployServiceServer(s *grpc.Server, srv TioDeployServiceServer) {
	s.RegisterService(&_TioDeployService_serviceDesc, srv)
}

func _TioDeployService_NewDeploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TioDeployServiceServer).NewDeploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TioDeployService/NewDeploy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TioDeployServiceServer).NewDeploy(ctx, req.(*DeployRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TioDeployService_ScalaDeploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TioDeployServiceServer).ScalaDeploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TioDeployService/ScalaDeploy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TioDeployServiceServer).ScalaDeploy(ctx, req.(*DeployRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TioDeployService_DeployInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TioDeployServiceServer).DeployInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TioDeployService/DeployInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TioDeployServiceServer).DeployInfo(ctx, req.(*DeployRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TioDeployService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TioDeployService",
	HandlerType: (*TioDeployServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewDeploy",
			Handler:    _TioDeployService_NewDeploy_Handler,
		},
		{
			MethodName: "ScalaDeploy",
			Handler:    _TioDeployService_ScalaDeploy_Handler,
		},
		{
			MethodName: "DeployInfo",
			Handler:    _TioDeployService_DeployInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deploy.proto",
}