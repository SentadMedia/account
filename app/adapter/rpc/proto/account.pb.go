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

type SignInResponse struct {
	UserId               uint64   `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInResponse) Reset()         { *m = SignInResponse{} }
func (m *SignInResponse) String() string { return proto.CompactTextString(m) }
func (*SignInResponse) ProtoMessage()    {}
func (*SignInResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_66d0219933e10f98, []int{3}
}

func (m *SignInResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInResponse.Unmarshal(m, b)
}
func (m *SignInResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInResponse.Marshal(b, m, deterministic)
}
func (m *SignInResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInResponse.Merge(m, src)
}
func (m *SignInResponse) XXX_Size() int {
	return xxx_messageInfo_SignInResponse.Size(m)
}
func (m *SignInResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignInResponse proto.InternalMessageInfo

func (m *SignInResponse) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func init() {
	proto.RegisterType((*Void)(nil), "proto.Void")
	proto.RegisterType((*RegisterAccountRequest)(nil), "proto.RegisterAccountRequest")
	proto.RegisterType((*SignInRequest)(nil), "proto.SignInRequest")
	proto.RegisterType((*SignInResponse)(nil), "proto.SignInResponse")
}

func init() {
	proto.RegisterFile("app/adapter/rpc/proto/account.proto", fileDescriptor_66d0219933e10f98)
}

var fileDescriptor_66d0219933e10f98 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x4b, 0xfb, 0x30,
	0x18, 0xc7, 0x97, 0xdf, 0x6f, 0xed, 0xb6, 0x67, 0xa8, 0x10, 0xe6, 0x28, 0x45, 0xa1, 0xd4, 0x4b,
	0x4f, 0x2b, 0xe8, 0xc1, 0xa3, 0x78, 0x92, 0x5d, 0x44, 0x32, 0xf4, 0x1e, 0xdb, 0xc7, 0x11, 0xe8,
	0x9a, 0x98, 0xa4, 0xf8, 0x02, 0x7c, 0x97, 0xbe, 0x1a, 0x69, 0x92, 0x4e, 0x1d, 0x1e, 0x3c, 0x95,
	0xcf, 0xf7, 0xdb, 0x6f, 0x9e, 0x7f, 0x70, 0xc1, 0x95, 0x2a, 0x79, 0xcd, 0x95, 0x45, 0x5d, 0x6a,
	0x55, 0x95, 0x4a, 0x4b, 0x2b, 0x4b, 0x5e, 0x55, 0xb2, 0x6b, 0xed, 0xca, 0x11, 0x8d, 0xdc, 0x27,
	0x8f, 0x61, 0xfc, 0x24, 0x45, 0x9d, 0x7f, 0x10, 0x58, 0x32, 0xdc, 0x0a, 0x63, 0x51, 0xdf, 0xfa,
	0x1f, 0x19, 0xbe, 0x76, 0x68, 0x2c, 0x4d, 0x61, 0xda, 0x19, 0xd4, 0xf7, 0x7c, 0x87, 0x09, 0xc9,
	0x48, 0x31, 0x63, 0x7b, 0xa6, 0x67, 0x30, 0x7b, 0x11, 0xda, 0x58, 0x67, 0xfe, 0x73, 0xe6, 0x97,
	0xd0, 0x27, 0x1b, 0x1e, 0xcc, 0xff, 0x3e, 0x39, 0x30, 0x5d, 0x40, 0x84, 0x3b, 0x2e, 0x9a, 0x64,
	0xec, 0x0c, 0x0f, 0x7d, 0x42, 0x71, 0x63, 0xde, 0xa4, 0xae, 0x93, 0xc8, 0x27, 0x06, 0xa6, 0x19,
	0xcc, 0x85, 0xd9, 0x74, 0x0a, 0xf5, 0xa3, 0x41, 0x9d, 0xc4, 0x19, 0x29, 0xa6, 0xec, 0xbb, 0x44,
	0x97, 0x10, 0x6b, 0xd9, 0xe0, 0xba, 0x4e, 0x26, 0x19, 0x29, 0x22, 0x16, 0x28, 0xbf, 0x83, 0xa3,
	0x8d, 0xd8, 0xb6, 0xeb, 0xf6, 0x2f, 0x23, 0xa5, 0x30, 0x7d, 0x18, 0x5a, 0xf0, 0x13, 0xed, 0x39,
	0x2f, 0xe0, 0x78, 0x78, 0xc8, 0x28, 0xd9, 0x1a, 0xec, 0x4b, 0xf6, 0xc9, 0x75, 0xed, 0xde, 0x19,
	0xb3, 0x40, 0x97, 0xef, 0x04, 0x26, 0x61, 0x8f, 0xf4, 0x06, 0x4e, 0x0e, 0x56, 0x4b, 0xcf, 0xfd,
	0x15, 0x56, 0xbf, 0xaf, 0x3c, 0x9d, 0x07, 0xdb, 0x9d, 0x66, 0x44, 0xaf, 0x21, 0xf6, 0x65, 0xe9,
	0x22, 0x18, 0x3f, 0xc6, 0x49, 0x4f, 0x0f, 0x54, 0xdf, 0x5b, 0x3e, 0x7a, 0x8e, 0x9d, 0x7e, 0xf5,
	0x19, 0x00, 0x00, 0xff, 0xff, 0x46, 0x8c, 0x41, 0xa9, 0x12, 0x02, 0x00, 0x00,
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
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
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

func (c *accountClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	out := new(SignInResponse)
	err := c.cc.Invoke(ctx, "/proto.Account/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
type AccountServer interface {
	RegisterAccount(context.Context, *RegisterAccountRequest) (*Void, error)
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
}

// UnimplementedAccountServer can be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (*UnimplementedAccountServer) RegisterAccount(ctx context.Context, req *RegisterAccountRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAccount not implemented")
}
func (*UnimplementedAccountServer) SignIn(ctx context.Context, req *SignInRequest) (*SignInResponse, error) {
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
