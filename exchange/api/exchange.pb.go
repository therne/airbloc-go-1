// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/exchange.proto

package api // import "github.com/airbloc/airbloc-go/exchange/api"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import api "github.com/airbloc/airbloc-go/common/api"
import api1 "github.com/airbloc/airbloc-go/data/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Contract_Type int32

const (
	Contract_RICHARDIAN Contract_Type = 0
	Contract_SMART      Contract_Type = 1
)

var Contract_Type_name = map[int32]string{
	0: "RICHARDIAN",
	1: "SMART",
}
var Contract_Type_value = map[string]int32{
	"RICHARDIAN": 0,
	"SMART":      1,
}

func (x Contract_Type) String() string {
	return proto.EnumName(Contract_Type_name, int32(x))
}
func (Contract_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_exchange_c82204f81f2966f6, []int{0, 0}
}

type Contract struct {
	Type                 Contract_Type `protobuf:"varint,1,opt,name=type,proto3,enum=airbloc.exchange.Contract_Type" json:"type,omitempty"`
	SmartContractAddress *api.Address  `protobuf:"bytes,2,opt,name=smartContractAddress,proto3" json:"smartContractAddress,omitempty"`
	RichardianHash       []byte        `protobuf:"bytes,3,opt,name=richardianHash,proto3" json:"richardianHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Contract) Reset()         { *m = Contract{} }
func (m *Contract) String() string { return proto.CompactTextString(m) }
func (*Contract) ProtoMessage()    {}
func (*Contract) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_c82204f81f2966f6, []int{0}
}
func (m *Contract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Contract.Unmarshal(m, b)
}
func (m *Contract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Contract.Marshal(b, m, deterministic)
}
func (dst *Contract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contract.Merge(dst, src)
}
func (m *Contract) XXX_Size() int {
	return xxx_messageInfo_Contract.Size(m)
}
func (m *Contract) XXX_DiscardUnknown() {
	xxx_messageInfo_Contract.DiscardUnknown(m)
}

var xxx_messageInfo_Contract proto.InternalMessageInfo

func (m *Contract) GetType() Contract_Type {
	if m != nil {
		return m.Type
	}
	return Contract_RICHARDIAN
}

func (m *Contract) GetSmartContractAddress() *api.Address {
	if m != nil {
		return m.SmartContractAddress
	}
	return nil
}

func (m *Contract) GetRichardianHash() []byte {
	if m != nil {
		return m.RichardianHash
	}
	return nil
}

type OrderRequest struct {
	Data                 *api1.BatchRequest `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Contract             *Contract          `protobuf:"bytes,2,opt,name=contract,proto3" json:"contract,omitempty"`
	Offeror              string             `protobuf:"bytes,3,opt,name=offeror,proto3" json:"offeror,omitempty"`
	Offeree              string             `protobuf:"bytes,4,opt,name=offeree,proto3" json:"offeree,omitempty"`
	Option               []string           `protobuf:"bytes,5,rep,name=option,proto3" json:"option,omitempty"`
	Amount               float64            `protobuf:"fixed64,6,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *OrderRequest) Reset()         { *m = OrderRequest{} }
func (m *OrderRequest) String() string { return proto.CompactTextString(m) }
func (*OrderRequest) ProtoMessage()    {}
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_c82204f81f2966f6, []int{1}
}
func (m *OrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderRequest.Unmarshal(m, b)
}
func (m *OrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderRequest.Marshal(b, m, deterministic)
}
func (dst *OrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderRequest.Merge(dst, src)
}
func (m *OrderRequest) XXX_Size() int {
	return xxx_messageInfo_OrderRequest.Size(m)
}
func (m *OrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OrderRequest proto.InternalMessageInfo

func (m *OrderRequest) GetData() *api1.BatchRequest {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *OrderRequest) GetContract() *Contract {
	if m != nil {
		return m.Contract
	}
	return nil
}

func (m *OrderRequest) GetOfferor() string {
	if m != nil {
		return m.Offeror
	}
	return ""
}

func (m *OrderRequest) GetOfferee() string {
	if m != nil {
		return m.Offeree
	}
	return ""
}

func (m *OrderRequest) GetOption() []string {
	if m != nil {
		return m.Option
	}
	return nil
}

func (m *OrderRequest) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type OrderIdRequest struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderIdRequest) Reset()         { *m = OrderIdRequest{} }
func (m *OrderIdRequest) String() string { return proto.CompactTextString(m) }
func (*OrderIdRequest) ProtoMessage()    {}
func (*OrderIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_c82204f81f2966f6, []int{2}
}
func (m *OrderIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderIdRequest.Unmarshal(m, b)
}
func (m *OrderIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderIdRequest.Marshal(b, m, deterministic)
}
func (dst *OrderIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderIdRequest.Merge(dst, src)
}
func (m *OrderIdRequest) XXX_Size() int {
	return xxx_messageInfo_OrderIdRequest.Size(m)
}
func (m *OrderIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OrderIdRequest proto.InternalMessageInfo

func (m *OrderIdRequest) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

type ContractId struct {
	ContractId           string   `protobuf:"bytes,1,opt,name=ContractId,proto3" json:"ContractId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContractId) Reset()         { *m = ContractId{} }
func (m *ContractId) String() string { return proto.CompactTextString(m) }
func (*ContractId) ProtoMessage()    {}
func (*ContractId) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_c82204f81f2966f6, []int{3}
}
func (m *ContractId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractId.Unmarshal(m, b)
}
func (m *ContractId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractId.Marshal(b, m, deterministic)
}
func (dst *ContractId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractId.Merge(dst, src)
}
func (m *ContractId) XXX_Size() int {
	return xxx_messageInfo_ContractId.Size(m)
}
func (m *ContractId) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractId.DiscardUnknown(m)
}

var xxx_messageInfo_ContractId proto.InternalMessageInfo

func (m *ContractId) GetContractId() string {
	if m != nil {
		return m.ContractId
	}
	return ""
}

type SettleResult struct {
	ContractId           *ContractId `protobuf:"bytes,1,opt,name=contractId,proto3" json:"contractId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SettleResult) Reset()         { *m = SettleResult{} }
func (m *SettleResult) String() string { return proto.CompactTextString(m) }
func (*SettleResult) ProtoMessage()    {}
func (*SettleResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_c82204f81f2966f6, []int{4}
}
func (m *SettleResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettleResult.Unmarshal(m, b)
}
func (m *SettleResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettleResult.Marshal(b, m, deterministic)
}
func (dst *SettleResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettleResult.Merge(dst, src)
}
func (m *SettleResult) XXX_Size() int {
	return xxx_messageInfo_SettleResult.Size(m)
}
func (m *SettleResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SettleResult.DiscardUnknown(m)
}

var xxx_messageInfo_SettleResult proto.InternalMessageInfo

func (m *SettleResult) GetContractId() *ContractId {
	if m != nil {
		return m.ContractId
	}
	return nil
}

func init() {
	proto.RegisterType((*Contract)(nil), "airbloc.exchange.Contract")
	proto.RegisterType((*OrderRequest)(nil), "airbloc.exchange.OrderRequest")
	proto.RegisterType((*OrderIdRequest)(nil), "airbloc.exchange.OrderIdRequest")
	proto.RegisterType((*ContractId)(nil), "airbloc.exchange.ContractId")
	proto.RegisterType((*SettleResult)(nil), "airbloc.exchange.SettleResult")
	proto.RegisterEnum("airbloc.exchange.Contract_Type", Contract_Type_name, Contract_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExchangeClient is the client API for Exchange service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExchangeClient interface {
	Order(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*api.Hash, error)
	Settle(ctx context.Context, in *api.Hash, opts ...grpc.CallOption) (*api.Result, error)
	Reject(ctx context.Context, in *OrderIdRequest, opts ...grpc.CallOption) (*api.Result, error)
	// after Open()
	CloseOrder(ctx context.Context, in *OrderIdRequest, opts ...grpc.CallOption) (*api.Result, error)
}

type exchangeClient struct {
	cc *grpc.ClientConn
}

func NewExchangeClient(cc *grpc.ClientConn) ExchangeClient {
	return &exchangeClient{cc}
}

func (c *exchangeClient) Order(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*api.Hash, error) {
	out := new(api.Hash)
	err := c.cc.Invoke(ctx, "/airbloc.exchange.Exchange/Order", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) Settle(ctx context.Context, in *api.Hash, opts ...grpc.CallOption) (*api.Result, error) {
	out := new(api.Result)
	err := c.cc.Invoke(ctx, "/airbloc.exchange.Exchange/Settle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) Reject(ctx context.Context, in *OrderIdRequest, opts ...grpc.CallOption) (*api.Result, error) {
	out := new(api.Result)
	err := c.cc.Invoke(ctx, "/airbloc.exchange.Exchange/Reject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) CloseOrder(ctx context.Context, in *OrderIdRequest, opts ...grpc.CallOption) (*api.Result, error) {
	out := new(api.Result)
	err := c.cc.Invoke(ctx, "/airbloc.exchange.Exchange/CloseOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExchangeServer is the server API for Exchange service.
type ExchangeServer interface {
	Order(context.Context, *OrderRequest) (*api.Hash, error)
	Settle(context.Context, *api.Hash) (*api.Result, error)
	Reject(context.Context, *OrderIdRequest) (*api.Result, error)
	// after Open()
	CloseOrder(context.Context, *OrderIdRequest) (*api.Result, error)
}

func RegisterExchangeServer(s *grpc.Server, srv ExchangeServer) {
	s.RegisterService(&_Exchange_serviceDesc, srv)
}

func _Exchange_Order_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).Order(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.exchange.Exchange/Order",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).Order(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_Settle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.Hash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).Settle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.exchange.Exchange/Settle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).Settle(ctx, req.(*api.Hash))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_Reject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).Reject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.exchange.Exchange/Reject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).Reject(ctx, req.(*OrderIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_CloseOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).CloseOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.exchange.Exchange/CloseOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).CloseOrder(ctx, req.(*OrderIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Exchange_serviceDesc = grpc.ServiceDesc{
	ServiceName: "airbloc.exchange.Exchange",
	HandlerType: (*ExchangeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Order",
			Handler:    _Exchange_Order_Handler,
		},
		{
			MethodName: "Settle",
			Handler:    _Exchange_Settle_Handler,
		},
		{
			MethodName: "Reject",
			Handler:    _Exchange_Reject_Handler,
		},
		{
			MethodName: "CloseOrder",
			Handler:    _Exchange_CloseOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/exchange.proto",
}

func init() { proto.RegisterFile("proto/exchange.proto", fileDescriptor_exchange_c82204f81f2966f6) }

var fileDescriptor_exchange_c82204f81f2966f6 = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x71, 0x9b, 0xa4, 0xc9, 0x24, 0x8a, 0xc2, 0x2a, 0x2a, 0x96, 0x85, 0x8a, 0xf1, 0x01,
	0x59, 0x55, 0x71, 0x24, 0x57, 0xea, 0x09, 0x0e, 0x49, 0x00, 0xd5, 0xe2, 0x9f, 0xb4, 0xed, 0x89,
	0xdb, 0x66, 0x3d, 0x8d, 0x8d, 0x12, 0xaf, 0x59, 0x6f, 0x24, 0xfa, 0x10, 0x3c, 0x1f, 0x12, 0x4f,
	0x83, 0xbc, 0x5e, 0x1b, 0x13, 0x9a, 0x0b, 0xa7, 0x68, 0xbe, 0xf9, 0xed, 0xe4, 0xfb, 0x66, 0xbd,
	0x30, 0xcd, 0xa5, 0x50, 0x62, 0x86, 0xdf, 0x79, 0xc2, 0xb2, 0x35, 0x06, 0xba, 0x24, 0x13, 0x96,
	0xca, 0xd5, 0x46, 0xf0, 0xa0, 0xd6, 0x9d, 0x49, 0xc5, 0xc5, 0x4c, 0xb1, 0x8a, 0x71, 0x1e, 0x57,
	0x8a, 0xba, 0xcf, 0xb1, 0xa8, 0x24, 0xef, 0xa7, 0x05, 0xfd, 0xa5, 0xc8, 0x94, 0x64, 0x5c, 0x91,
	0x4b, 0xe8, 0x94, 0x3d, 0xdb, 0x72, 0x2d, 0x7f, 0x1c, 0x3e, 0x0b, 0xf6, 0x47, 0x06, 0x35, 0x19,
	0xdc, 0xde, 0xe7, 0x48, 0x35, 0x4c, 0xde, 0xc3, 0xb4, 0xd8, 0x32, 0xa9, 0xea, 0xde, 0x3c, 0x8e,
	0x25, 0x16, 0x85, 0x7d, 0xe4, 0x5a, 0xfe, 0x30, 0x7c, 0xd2, 0x0c, 0xe1, 0x62, 0xbb, 0x15, 0x59,
	0x60, 0xda, 0xf4, 0xc1, 0x43, 0xe4, 0x05, 0x8c, 0x65, 0xca, 0x13, 0x26, 0xe3, 0x94, 0x65, 0xd7,
	0xac, 0x48, 0xec, 0x63, 0xd7, 0xf2, 0x47, 0x74, 0x4f, 0xf5, 0x9e, 0x43, 0xa7, 0xb4, 0x40, 0xc6,
	0x00, 0x34, 0x5a, 0x5e, 0xcf, 0xe9, 0x9b, 0x68, 0xfe, 0x69, 0xf2, 0x88, 0x0c, 0xa0, 0x7b, 0xf3,
	0x71, 0x4e, 0x6f, 0x27, 0x96, 0xf7, 0xcb, 0x82, 0xd1, 0x67, 0x19, 0xa3, 0xa4, 0xf8, 0x6d, 0x87,
	0x85, 0x22, 0x01, 0x74, 0xca, 0x5d, 0xe8, 0x74, 0xc3, 0xd0, 0x69, 0x8c, 0xe9, 0x05, 0x2d, 0x98,
	0xe2, 0x89, 0x21, 0xa9, 0xe6, 0xc8, 0x15, 0xf4, 0xb9, 0xb1, 0x67, 0xc2, 0x38, 0x87, 0x37, 0x42,
	0x1b, 0x96, 0xd8, 0x70, 0x22, 0xee, 0xee, 0x50, 0x0a, 0xa9, 0xcd, 0x0f, 0x68, 0x5d, 0x36, 0x1d,
	0x44, 0xbb, 0xd3, 0xea, 0x20, 0x92, 0x53, 0xe8, 0x89, 0x5c, 0xa5, 0x22, 0xb3, 0xbb, 0xee, 0xb1,
	0x3f, 0xa0, 0xa6, 0x2a, 0x75, 0xb6, 0x15, 0xbb, 0x4c, 0xd9, 0x3d, 0xd7, 0xf2, 0x2d, 0x6a, 0x2a,
	0xef, 0x1c, 0xc6, 0x3a, 0x5b, 0x14, 0xd7, 0xe9, 0x6c, 0x38, 0x31, 0x8a, 0x0e, 0x38, 0xa0, 0x75,
	0xe9, 0x5d, 0x00, 0xd4, 0x2e, 0xa3, 0x98, 0x9c, 0xb5, 0x2b, 0x83, 0xb6, 0x14, 0xef, 0x03, 0x8c,
	0x6e, 0x50, 0xa9, 0x0d, 0x52, 0x2c, 0x76, 0x1b, 0x45, 0x5e, 0x01, 0xf0, 0xbf, 0xf9, 0x61, 0xf8,
	0xf4, 0xf0, 0x1e, 0xa2, 0x98, 0xb6, 0xf8, 0xf0, 0xc7, 0x11, 0xf4, 0xdf, 0x1a, 0x86, 0xbc, 0x86,
	0xae, 0xf6, 0x44, 0xce, 0xfe, 0x3d, 0xdf, 0xbe, 0x29, 0x67, 0xba, 0xff, 0xd1, 0x94, 0x77, 0x4e,
	0xae, 0xa0, 0x57, 0x39, 0x23, 0x0f, 0xf6, 0x9d, 0xd3, 0x7d, 0xd5, 0x24, 0x58, 0x40, 0x8f, 0xe2,
	0x57, 0xe4, 0x8a, 0xb8, 0x07, 0xfe, 0xb7, 0xd9, 0xe2, 0xc1, 0x19, 0xef, 0x00, 0x96, 0x1b, 0x51,
	0x60, 0xe5, 0xff, 0xbf, 0xe7, 0x2c, 0x2e, 0xbe, 0x9c, 0xaf, 0x53, 0x95, 0xec, 0x56, 0xa5, 0x3e,
	0x33, 0x4c, 0xfd, 0xfb, 0x72, 0xfd, 0xe7, 0x55, 0xcf, 0x58, 0x9e, 0xae, 0x7a, 0xfa, 0x8d, 0x5e,
	0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x7f, 0xd4, 0xe7, 0xea, 0xf2, 0x03, 0x00, 0x00,
}
