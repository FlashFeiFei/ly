// Code generated by protoc-gen-go. DO NOT EDIT.
// source: finance.proto

package finance

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

//用户
type User struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_c04e2e1c1ba79a81, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

//充值
type RechargeInfo struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Money                int64    `protobuf:"varint,2,opt,name=money,proto3" json:"money,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RechargeInfo) Reset()         { *m = RechargeInfo{} }
func (m *RechargeInfo) String() string { return proto.CompactTextString(m) }
func (*RechargeInfo) ProtoMessage()    {}
func (*RechargeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c04e2e1c1ba79a81, []int{1}
}

func (m *RechargeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RechargeInfo.Unmarshal(m, b)
}
func (m *RechargeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RechargeInfo.Marshal(b, m, deterministic)
}
func (m *RechargeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RechargeInfo.Merge(m, src)
}
func (m *RechargeInfo) XXX_Size() int {
	return xxx_messageInfo_RechargeInfo.Size(m)
}
func (m *RechargeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RechargeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RechargeInfo proto.InternalMessageInfo

func (m *RechargeInfo) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *RechargeInfo) GetMoney() int64 {
	if m != nil {
		return m.Money
	}
	return 0
}

//余额
type Money struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Money                int64    `protobuf:"varint,2,opt,name=money,proto3" json:"money,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Money) Reset()         { *m = Money{} }
func (m *Money) String() string { return proto.CompactTextString(m) }
func (*Money) ProtoMessage()    {}
func (*Money) Descriptor() ([]byte, []int) {
	return fileDescriptor_c04e2e1c1ba79a81, []int{2}
}

func (m *Money) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Money.Unmarshal(m, b)
}
func (m *Money) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Money.Marshal(b, m, deterministic)
}
func (m *Money) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Money.Merge(m, src)
}
func (m *Money) XXX_Size() int {
	return xxx_messageInfo_Money.Size(m)
}
func (m *Money) XXX_DiscardUnknown() {
	xxx_messageInfo_Money.DiscardUnknown(m)
}

var xxx_messageInfo_Money proto.InternalMessageInfo

func (m *Money) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Money) GetMoney() int64 {
	if m != nil {
		return m.Money
	}
	return 0
}

//结果
type Result struct {
	Result               string        `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	RechargeInfo         *RechargeInfo `protobuf:"bytes,2,opt,name=recharge_info,json=rechargeInfo,proto3" json:"recharge_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_c04e2e1c1ba79a81, []int{3}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *Result) GetRechargeInfo() *RechargeInfo {
	if m != nil {
		return m.RechargeInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "finance.User")
	proto.RegisterType((*RechargeInfo)(nil), "finance.RechargeInfo")
	proto.RegisterType((*Money)(nil), "finance.Money")
	proto.RegisterType((*Result)(nil), "finance.Result")
}

func init() { proto.RegisterFile("finance.proto", fileDescriptor_c04e2e1c1ba79a81) }

var fileDescriptor_c04e2e1c1ba79a81 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xcb, 0xcc, 0x4b,
	0xcc, 0x4b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0xa4, 0xb8,
	0x58, 0x42, 0x8b, 0x53, 0x8b, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25, 0x77, 0x2e, 0x9e, 0xa0, 0xd4, 0xe4, 0x8c, 0xc4, 0xa2,
	0xf4, 0x54, 0xcf, 0xbc, 0xb4, 0x7c, 0x21, 0x45, 0x2e, 0x96, 0xd2, 0xe2, 0xd4, 0x22, 0xb0, 0x1a,
	0x6e, 0x23, 0x5e, 0x3d, 0x98, 0x91, 0x20, 0x03, 0x82, 0xc0, 0x52, 0x42, 0x22, 0x5c, 0xac, 0xb9,
	0xf9, 0x79, 0xa9, 0x95, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x10, 0x8e, 0x92, 0x03, 0x17,
	0xab, 0x2f, 0x88, 0x41, 0xbe, 0x09, 0x31, 0x5c, 0x6c, 0x41, 0xa9, 0xc5, 0xa5, 0x39, 0x25, 0x42,
	0x62, 0x5c, 0x6c, 0x45, 0x60, 0x16, 0xd4, 0xa9, 0x50, 0x9e, 0x90, 0x15, 0x17, 0x6f, 0x11, 0xd4,
	0xb1, 0xf1, 0x99, 0x79, 0x69, 0xf9, 0x60, 0xfd, 0xdc, 0x46, 0xa2, 0x70, 0x3b, 0x90, 0xbd, 0x12,
	0xc4, 0x53, 0x84, 0xc4, 0x33, 0xca, 0xe2, 0x62, 0x77, 0x83, 0xa8, 0x12, 0x32, 0xe2, 0xe2, 0x80,
	0x29, 0x14, 0xc2, 0xae, 0x57, 0x8a, 0x1f, 0x49, 0x18, 0x6c, 0xb5, 0x36, 0x17, 0x57, 0x60, 0x69,
	0x6a, 0x51, 0x25, 0xc4, 0x8f, 0xa8, 0xbe, 0x92, 0xe2, 0x83, 0x73, 0xc1, 0xd2, 0x49, 0x6c, 0xe0,
	0x08, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x3b, 0x55, 0x32, 0x91, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FinanceClient is the client API for Finance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FinanceClient interface {
	//充值服务
	Recharge(ctx context.Context, in *RechargeInfo, opts ...grpc.CallOption) (*Result, error)
	//查询充值记录
	QueryMoney(ctx context.Context, in *User, opts ...grpc.CallOption) (*Money, error)
}

type financeClient struct {
	cc *grpc.ClientConn
}

func NewFinanceClient(cc *grpc.ClientConn) FinanceClient {
	return &financeClient{cc}
}

func (c *financeClient) Recharge(ctx context.Context, in *RechargeInfo, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/finance.Finance/Recharge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeClient) QueryMoney(ctx context.Context, in *User, opts ...grpc.CallOption) (*Money, error) {
	out := new(Money)
	err := c.cc.Invoke(ctx, "/finance.Finance/QueryMoney", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FinanceServer is the server API for Finance service.
type FinanceServer interface {
	//充值服务
	Recharge(context.Context, *RechargeInfo) (*Result, error)
	//查询充值记录
	QueryMoney(context.Context, *User) (*Money, error)
}

// UnimplementedFinanceServer can be embedded to have forward compatible implementations.
type UnimplementedFinanceServer struct {
}

func (*UnimplementedFinanceServer) Recharge(ctx context.Context, req *RechargeInfo) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recharge not implemented")
}
func (*UnimplementedFinanceServer) QueryMoney(ctx context.Context, req *User) (*Money, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryMoney not implemented")
}

func RegisterFinanceServer(s *grpc.Server, srv FinanceServer) {
	s.RegisterService(&_Finance_serviceDesc, srv)
}

func _Finance_Recharge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RechargeInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServer).Recharge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/finance.Finance/Recharge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServer).Recharge(ctx, req.(*RechargeInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Finance_QueryMoney_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServer).QueryMoney(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/finance.Finance/QueryMoney",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServer).QueryMoney(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _Finance_serviceDesc = grpc.ServiceDesc{
	ServiceName: "finance.Finance",
	HandlerType: (*FinanceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Recharge",
			Handler:    _Finance_Recharge_Handler,
		},
		{
			MethodName: "QueryMoney",
			Handler:    _Finance_QueryMoney_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "finance.proto",
}
