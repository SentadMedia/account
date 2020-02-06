// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/adapter/rpc/proto/account.proto

package proto

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

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_66d0219933e10f98, []int{0}
}

func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

type RegisterAccountRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	IsSuperUser          bool     `protobuf:"varint,6,opt,name=isSuperUser,proto3" json:"isSuperUser,omitempty"`
	RoleId               int32    `protobuf:"varint,7,opt,name=roleId,proto3" json:"roleId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterAccountRequest) Reset()         { *m = RegisterAccountRequest{} }
func (m *RegisterAccountRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterAccountRequest) ProtoMessage()    {}
func (*RegisterAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_66d0219933e10f98, []int{1}
}

func (m *RegisterAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterAccountRequest.Unmarshal(m, b)
}
func (m *RegisterAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterAccountRequest.Marshal(b, m, deterministic)
}
func (m *RegisterAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterAccountRequest.Merge(m, src)
}
func (m *RegisterAccountRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterAccountRequest.Size(m)
}
func (m *RegisterAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterAccountRequest proto.InternalMessageInfo

func (m *RegisterAccountRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *RegisterAccountRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *RegisterAccountRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *RegisterAccountRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterAccountRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterAccountRequest) GetIsSuperUser() bool {
	if m != nil {
		return m.IsSuperUser
	}
	return false
}

func (m *RegisterAccountRequest) GetRoleId() int32 {
	if m != nil {
		return m.RoleId
	}
	return 0
}

type SignInRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInRequest) Reset()         { *m = SignInRequest{} }
func (m *SignInRequest) String() string { return proto.CompactTextString(m) }
func (*SignInRequest) ProtoMessage()    {}
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_66d0219933e10f98, []int{2}
}

func (m *SignInRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInRequest.Unmarshal(m, b)
}
func (m *SignInRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInRequest.Marshal(b, m, deterministic)
}
func (m *SignInRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInRequest.Merge(m, src)
}
func (m *SignInRequest) XXX_Size() int {
	return xxx_messageInfo_SignInRequest.Size(m)
}
func (m *SignInRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignInRequest proto.InternalMessageInfo

func (m *SignInRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *SignInRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*Void)(nil), "proto.Void")
	proto.RegisterType((*RegisterAccountRequest)(nil), "proto.RegisterAccountRequest")
	proto.RegisterType((*SignInRequest)(nil), "proto.SignInRequest")
}

func init() {
	proto.RegisterFile("app/adapter/rpc/proto/account.proto", fileDescriptor_66d0219933e10f98)
}

var fileDescriptor_66d0219933e10f98 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x8d, 0x36, 0x69, 0x3b, 0x45, 0x84, 0xa1, 0x94, 0x10, 0x14, 0x42, 0xbc, 0xe4, 0x62,
	0x03, 0xfa, 0x00, 0xe2, 0x49, 0x7a, 0x11, 0x49, 0xd1, 0xfb, 0x9a, 0x8c, 0x65, 0x21, 0xcd, 0xae,
	0xb3, 0x1b, 0xc4, 0x57, 0xf5, 0x69, 0xa4, 0xbb, 0x49, 0xd5, 0xe2, 0xc1, 0xd3, 0xf2, 0xed, 0x37,
	0x3f, 0xcc, 0xfc, 0x70, 0x29, 0xb4, 0x2e, 0x44, 0x2d, 0xb4, 0x25, 0x2e, 0x58, 0x57, 0x85, 0x66,
	0x65, 0x55, 0x21, 0xaa, 0x4a, 0x75, 0xad, 0x5d, 0x3a, 0xc2, 0xd0, 0x3d, 0x59, 0x04, 0xa3, 0x67,
	0x25, 0xeb, 0xec, 0x33, 0x80, 0x45, 0x49, 0x1b, 0x69, 0x2c, 0xf1, 0x9d, 0x1f, 0x2c, 0xe9, 0xad,
	0x23, 0x63, 0x31, 0x81, 0x49, 0x67, 0x88, 0x1f, 0xc4, 0x96, 0xe2, 0x20, 0x0d, 0xf2, 0x69, 0xb9,
	0x67, 0x3c, 0x87, 0xe9, 0xab, 0x64, 0x63, 0x9d, 0x3c, 0x76, 0xf2, 0xfb, 0x63, 0x97, 0x6c, 0x44,
	0x2f, 0x4f, 0x7c, 0x72, 0x60, 0x9c, 0x43, 0x48, 0x5b, 0x21, 0x9b, 0x78, 0xe4, 0x84, 0x87, 0x5d,
	0x42, 0x0b, 0x63, 0xde, 0x15, 0xd7, 0x71, 0xe8, 0x13, 0x03, 0x63, 0x0a, 0x33, 0x69, 0xd6, 0x9d,
	0x26, 0x7e, 0x32, 0xc4, 0x71, 0x94, 0x06, 0xf9, 0xa4, 0xfc, 0xf9, 0x85, 0x0b, 0x88, 0x58, 0x35,
	0xb4, 0xaa, 0xe3, 0x71, 0x1a, 0xe4, 0x61, 0xd9, 0x53, 0x76, 0x0f, 0xa7, 0x6b, 0xb9, 0x69, 0x57,
	0xed, 0x7f, 0x4e, 0x4a, 0x60, 0xf2, 0x38, 0xac, 0xe0, 0x2f, 0xda, 0xf3, 0xf5, 0x07, 0x8c, 0xfb,
	0x72, 0xf0, 0x16, 0xce, 0x0e, 0xfa, 0xc2, 0x0b, 0x5f, 0xed, 0xf2, 0xef, 0x1e, 0x93, 0x59, 0xaf,
	0x5d, 0xdf, 0x47, 0x78, 0x05, 0x91, 0x5f, 0x0a, 0xe7, 0xbd, 0xf8, 0xb5, 0xe3, 0xc1, 0xf8, 0x4b,
	0xe4, 0xe8, 0xe6, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x22, 0x7e, 0x87, 0x24, 0xdd, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountClient interface {
	RegisterAccount(ctx context.Context, in *RegisterAccountRequest, opts ...grpc.CallOption) (*Void, error)
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*Void, error)
}

type accountClient struct {
	cc *grpc.ClientConn
}

func NewAccountClient(cc *grpc.ClientConn) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) RegisterAccount(ctx context.Context, in *RegisterAccountRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/proto.Account/RegisterAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/proto.Account/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
type AccountServer interface {
	RegisterAccount(context.Context, *RegisterAccountRequest) (*Void, error)
	SignIn(context.Context, *SignInRequest) (*Void, error)
}

// UnimplementedAccountServer can be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (*UnimplementedAccountServer) RegisterAccount(ctx context.Context, req *RegisterAccountRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAccount not implemented")
}
func (*UnimplementedAccountServer) SignIn(ctx context.Context, req *SignInRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}

func RegisterAccountServer(s *grpc.Server, srv AccountServer) {
	s.RegisterService(&_Account_serviceDesc, srv)
}

func _Account_RegisterAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).RegisterAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Account/RegisterAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).RegisterAccount(ctx, req.(*RegisterAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Account/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterAccount",
			Handler:    _Account_RegisterAccount_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _Account_SignIn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/adapter/rpc/proto/account.proto",
}
