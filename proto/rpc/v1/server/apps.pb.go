// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/rpc/v1/server/apps.proto

package server

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
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

type NewOwnerRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewOwnerRequest) Reset()         { *m = NewOwnerRequest{} }
func (m *NewOwnerRequest) String() string { return proto.CompactTextString(m) }
func (*NewOwnerRequest) ProtoMessage()    {}
func (*NewOwnerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d1e9fb88ade88f6, []int{0}
}

func (m *NewOwnerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewOwnerRequest.Unmarshal(m, b)
}
func (m *NewOwnerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewOwnerRequest.Marshal(b, m, deterministic)
}
func (m *NewOwnerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewOwnerRequest.Merge(m, src)
}
func (m *NewOwnerRequest) XXX_Size() int {
	return xxx_messageInfo_NewOwnerRequest.Size(m)
}
func (m *NewOwnerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewOwnerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewOwnerRequest proto.InternalMessageInfo

func (m *NewOwnerRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *NewOwnerRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type CheckOwnerRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckOwnerRequest) Reset()         { *m = CheckOwnerRequest{} }
func (m *CheckOwnerRequest) String() string { return proto.CompactTextString(m) }
func (*CheckOwnerRequest) ProtoMessage()    {}
func (*CheckOwnerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d1e9fb88ade88f6, []int{1}
}

func (m *CheckOwnerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckOwnerRequest.Unmarshal(m, b)
}
func (m *CheckOwnerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckOwnerRequest.Marshal(b, m, deterministic)
}
func (m *CheckOwnerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckOwnerRequest.Merge(m, src)
}
func (m *CheckOwnerRequest) XXX_Size() int {
	return xxx_messageInfo_CheckOwnerRequest.Size(m)
}
func (m *CheckOwnerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckOwnerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckOwnerRequest proto.InternalMessageInfo

func (m *CheckOwnerRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *CheckOwnerRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type CheckOwnerResult struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckOwnerResult) Reset()         { *m = CheckOwnerResult{} }
func (m *CheckOwnerResult) String() string { return proto.CompactTextString(m) }
func (*CheckOwnerResult) ProtoMessage()    {}
func (*CheckOwnerResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d1e9fb88ade88f6, []int{2}
}

func (m *CheckOwnerResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckOwnerResult.Unmarshal(m, b)
}
func (m *CheckOwnerResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckOwnerResult.Marshal(b, m, deterministic)
}
func (m *CheckOwnerResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckOwnerResult.Merge(m, src)
}
func (m *CheckOwnerResult) XXX_Size() int {
	return xxx_messageInfo_CheckOwnerResult.Size(m)
}
func (m *CheckOwnerResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckOwnerResult.DiscardUnknown(m)
}

var xxx_messageInfo_CheckOwnerResult proto.InternalMessageInfo

func (m *CheckOwnerResult) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type RegisterRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d1e9fb88ade88f6, []int{3}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RegisterResult struct {
	AppId                string   `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResult) Reset()         { *m = RegisterResult{} }
func (m *RegisterResult) String() string { return proto.CompactTextString(m) }
func (*RegisterResult) ProtoMessage()    {}
func (*RegisterResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d1e9fb88ade88f6, []int{4}
}

func (m *RegisterResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResult.Unmarshal(m, b)
}
func (m *RegisterResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResult.Marshal(b, m, deterministic)
}
func (m *RegisterResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResult.Merge(m, src)
}
func (m *RegisterResult) XXX_Size() int {
	return xxx_messageInfo_RegisterResult.Size(m)
}
func (m *RegisterResult) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResult.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResult proto.InternalMessageInfo

func (m *RegisterResult) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

type UnregisterRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnregisterRequest) Reset()         { *m = UnregisterRequest{} }
func (m *UnregisterRequest) String() string { return proto.CompactTextString(m) }
func (*UnregisterRequest) ProtoMessage()    {}
func (*UnregisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d1e9fb88ade88f6, []int{5}
}

func (m *UnregisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnregisterRequest.Unmarshal(m, b)
}
func (m *UnregisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnregisterRequest.Marshal(b, m, deterministic)
}
func (m *UnregisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnregisterRequest.Merge(m, src)
}
func (m *UnregisterRequest) XXX_Size() int {
	return xxx_messageInfo_UnregisterRequest.Size(m)
}
func (m *UnregisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnregisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnregisterRequest proto.InternalMessageInfo

func (m *UnregisterRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func init() {
	proto.RegisterType((*NewOwnerRequest)(nil), "airbloc.rpc.v1.NewOwnerRequest")
	proto.RegisterType((*CheckOwnerRequest)(nil), "airbloc.rpc.v1.CheckOwnerRequest")
	proto.RegisterType((*CheckOwnerResult)(nil), "airbloc.rpc.v1.CheckOwnerResult")
	proto.RegisterType((*RegisterRequest)(nil), "airbloc.rpc.v1.RegisterRequest")
	proto.RegisterType((*RegisterResult)(nil), "airbloc.rpc.v1.RegisterResult")
	proto.RegisterType((*UnregisterRequest)(nil), "airbloc.rpc.v1.UnregisterRequest")
}

func init() { proto.RegisterFile("proto/rpc/v1/server/apps.proto", fileDescriptor_1d1e9fb88ade88f6) }

var fileDescriptor_1d1e9fb88ade88f6 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0x59, 0x99, 0x32, 0xef, 0xc3, 0xe6, 0x82, 0xc8, 0xa8, 0x30, 0x67, 0x40, 0xd1, 0x07,
	0x13, 0xe6, 0x9e, 0x45, 0x74, 0x88, 0x88, 0xa0, 0x58, 0xf0, 0xc5, 0xb7, 0xb6, 0xc6, 0x6e, 0xf4,
	0x4f, 0x62, 0x92, 0x76, 0xf8, 0xbd, 0xfc, 0x80, 0xd2, 0x74, 0xa5, 0xda, 0x76, 0x3e, 0xf8, 0xd4,
	0xde, 0x7b, 0x7f, 0x3d, 0xf7, 0x34, 0x27, 0x30, 0x16, 0x92, 0x6b, 0x4e, 0xa5, 0xf0, 0x69, 0x36,
	0xa5, 0x8a, 0xc9, 0x8c, 0x49, 0xea, 0x0a, 0xa1, 0x88, 0x19, 0xa0, 0xbe, 0xbb, 0x94, 0x5e, 0xc4,
	0x7d, 0x22, 0x85, 0x4f, 0xb2, 0xa9, 0x7d, 0x10, 0x70, 0x1e, 0x44, 0x8c, 0x9a, 0xa9, 0x97, 0xbe,
	0x53, 0x16, 0x0b, 0xfd, 0x59, 0xc0, 0xf8, 0x12, 0x06, 0x8f, 0x6c, 0xf5, 0xb4, 0x4a, 0x98, 0x74,
	0xd8, 0x47, 0xca, 0x94, 0x46, 0x7b, 0xb0, 0xe5, 0x0a, 0x71, 0xff, 0x36, 0xea, 0x4c, 0x3a, 0xa7,
	0x3b, 0x4e, 0x51, 0xe4, 0x5d, 0x9e, 0x53, 0x23, 0xab, 0xe8, 0x9a, 0x02, 0x5f, 0xc1, 0x70, 0xbe,
	0x60, 0x7e, 0xf8, 0x6f, 0x01, 0x0c, 0xbb, 0x3f, 0x05, 0x54, 0x1a, 0x69, 0xd4, 0x07, 0x8b, 0x87,
	0xe6, 0xe3, 0x9e, 0x63, 0xf1, 0x10, 0x1f, 0xc3, 0xc0, 0x61, 0xc1, 0x52, 0xe9, 0x6a, 0x05, 0x82,
	0x6e, 0xe2, 0xc6, 0x6c, 0xbd, 0xc1, 0xbc, 0xe3, 0x13, 0xe8, 0x57, 0x98, 0x11, 0x6a, 0x35, 0x82,
	0xcf, 0x60, 0xf8, 0x92, 0xc8, 0x9a, 0x60, 0x2b, 0x7a, 0xf1, 0x65, 0x41, 0xf7, 0x5a, 0x08, 0x85,
	0x1e, 0xa0, 0x57, 0x6a, 0xa3, 0x43, 0xf2, 0xfb, 0x80, 0x49, 0xcd, 0x9c, 0x3d, 0xde, 0x0c, 0x18,
	0x5b, 0x77, 0x00, 0x95, 0x01, 0x74, 0x54, 0xa7, 0x1b, 0xe6, 0xec, 0x7d, 0x52, 0x44, 0x48, 0xca,
	0x08, 0xc9, 0x6d, 0x1e, 0x21, 0x9a, 0x43, 0xaf, 0x0c, 0xaf, 0xe9, 0xaa, 0x16, 0xeb, 0x46, 0x91,
	0x67, 0x80, 0x2a, 0x81, 0xa6, 0x9b, 0x46, 0xbc, 0xf6, 0xe4, 0x2f, 0x24, 0xff, 0xc1, 0x9b, 0xd9,
	0xeb, 0x34, 0x58, 0xea, 0x45, 0xea, 0x11, 0x9f, 0xc7, 0x74, 0x4d, 0x97, 0xcf, 0xf3, 0x80, 0xd3,
	0x96, 0x1b, 0xec, 0x6d, 0x9b, 0xe6, 0xec, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x03, 0x01, 0xea,
	0xdf, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AppsClient is the client API for Apps service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AppsClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResult, error)
	Unregister(ctx context.Context, in *UnregisterRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	NewOwner(ctx context.Context, in *NewOwnerRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	CheckOwner(ctx context.Context, in *CheckOwnerRequest, opts ...grpc.CallOption) (*CheckOwnerResult, error)
}

type appsClient struct {
	cc *grpc.ClientConn
}

func NewAppsClient(cc *grpc.ClientConn) AppsClient {
	return &appsClient{cc}
}

func (c *appsClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResult, error) {
	out := new(RegisterResult)
	err := c.cc.Invoke(ctx, "/airbloc.rpc.v1.Apps/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appsClient) Unregister(ctx context.Context, in *UnregisterRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/airbloc.rpc.v1.Apps/Unregister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appsClient) NewOwner(ctx context.Context, in *NewOwnerRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/airbloc.rpc.v1.Apps/NewOwner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appsClient) CheckOwner(ctx context.Context, in *CheckOwnerRequest, opts ...grpc.CallOption) (*CheckOwnerResult, error) {
	out := new(CheckOwnerResult)
	err := c.cc.Invoke(ctx, "/airbloc.rpc.v1.Apps/CheckOwner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppsServer is the server API for Apps service.
type AppsServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResult, error)
	Unregister(context.Context, *UnregisterRequest) (*empty.Empty, error)
	NewOwner(context.Context, *NewOwnerRequest) (*empty.Empty, error)
	CheckOwner(context.Context, *CheckOwnerRequest) (*CheckOwnerResult, error)
}

func RegisterAppsServer(s *grpc.Server, srv AppsServer) {
	s.RegisterService(&_Apps_serviceDesc, srv)
}

func _Apps_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppsServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.rpc.v1.Apps/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppsServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apps_Unregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnregisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppsServer).Unregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.rpc.v1.Apps/Unregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppsServer).Unregister(ctx, req.(*UnregisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apps_NewOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppsServer).NewOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.rpc.v1.Apps/NewOwner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppsServer).NewOwner(ctx, req.(*NewOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apps_CheckOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppsServer).CheckOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.rpc.v1.Apps/CheckOwner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppsServer).CheckOwner(ctx, req.(*CheckOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Apps_serviceDesc = grpc.ServiceDesc{
	ServiceName: "airbloc.rpc.v1.Apps",
	HandlerType: (*AppsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Apps_Register_Handler,
		},
		{
			MethodName: "Unregister",
			Handler:    _Apps_Unregister_Handler,
		},
		{
			MethodName: "NewOwner",
			Handler:    _Apps_NewOwner_Handler,
		},
		{
			MethodName: "CheckOwner",
			Handler:    _Apps_CheckOwner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rpc/v1/server/apps.proto",
}
