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

type RegisterAccountRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	FirstName            string   `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string   `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email                string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	IsSuperUser          bool     `protobuf:"varint,7,opt,name=isSuperUser,proto3" json:"isSuperUser,omitempty"`
	RoleId               int32    `protobuf:"varint,8,opt,name=roleId,proto3" json:"roleId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterAccountRequest) Reset()         { *m = RegisterAccountRequest{} }
func (m *RegisterAccountRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterAccountRequest) ProtoMessage()    {}
func (*RegisterAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_66d0219933e10f98, []int{0}
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

func (m *RegisterAccountRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

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

type RegisterAccountResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterAccountResponse) Reset()         { *m = RegisterAccountResponse{} }
func (m *RegisterAccountResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterAccountResponse) ProtoMessage()    {}
func (*RegisterAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_66d0219933e10f98, []int{1}
}

func (m *RegisterAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterAccountResponse.Unmarshal(m, b)
}
func (m *RegisterAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterAccountResponse.Marshal(b, m, deterministic)
}
func (m *RegisterAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterAccountResponse.Merge(m, src)
}
func (m *RegisterAccountResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterAccountResponse.Size(m)
}
func (m *RegisterAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterAccountResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RegisterAccountRequest)(nil), "proto.RegisterAccountRequest")
	proto.RegisterType((*RegisterAccountResponse)(nil), "proto.RegisterAccountResponse")
}

func init() {
	proto.RegisterFile("app/adapter/rpc/proto/account.proto", fileDescriptor_66d0219933e10f98)
}

var fileDescriptor_66d0219933e10f98 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4f, 0x4b, 0xf4, 0x30,
	0x10, 0xc6, 0xdf, 0xf6, 0xb5, 0xdd, 0xee, 0x08, 0x0a, 0x41, 0xd6, 0xb8, 0xa8, 0x94, 0x7a, 0xe9,
	0x69, 0x0b, 0xfa, 0x09, 0x3c, 0x7a, 0xf1, 0x10, 0xf1, 0xe8, 0x21, 0x36, 0xa3, 0x04, 0xba, 0x9b,
	0x38, 0x93, 0xe2, 0x77, 0xf6, 0x53, 0xc8, 0x26, 0x5d, 0x15, 0xff, 0x9c, 0xc2, 0x6f, 0x7e, 0xcf,
	0x13, 0x92, 0x81, 0x0b, 0xed, 0x7d, 0xa7, 0x8d, 0xf6, 0x01, 0xa9, 0x23, 0xdf, 0x77, 0x9e, 0x5c,
	0x70, 0x9d, 0xee, 0x7b, 0x37, 0x6e, 0xc2, 0x2a, 0x92, 0x28, 0xe2, 0xd1, 0xbc, 0x65, 0xb0, 0x50,
	0xf8, 0x6c, 0x39, 0x20, 0x5d, 0xa7, 0x80, 0xc2, 0x97, 0x11, 0x39, 0x88, 0x03, 0xc8, 0xad, 0x91,
	0x59, 0x9d, 0xb5, 0x85, 0xca, 0xad, 0x11, 0x4b, 0xa8, 0x46, 0x46, 0xba, 0xd5, 0x6b, 0x94, 0x79,
	0x9d, 0xb5, 0x73, 0xf5, 0xc1, 0xe2, 0x14, 0xe6, 0x4f, 0x96, 0x38, 0x44, 0xf9, 0x3f, 0xca, 0xcf,
	0xc1, 0xb6, 0x39, 0xe8, 0x49, 0xee, 0xa5, 0xe6, 0x8e, 0xc5, 0x11, 0x14, 0xb8, 0xd6, 0x76, 0x90,
	0x45, 0x14, 0x09, 0xb6, 0x0d, 0xaf, 0x99, 0x5f, 0x1d, 0x19, 0x59, 0xa6, 0xc6, 0x8e, 0x45, 0x0d,
	0xfb, 0x96, 0xef, 0x46, 0x8f, 0x74, 0xcf, 0x48, 0x72, 0x56, 0x67, 0x6d, 0xa5, 0xbe, 0x8e, 0xc4,
	0x02, 0x4a, 0x72, 0x03, 0xde, 0x18, 0x59, 0xc5, 0xd7, 0x4f, 0xd4, 0x9c, 0xc0, 0xf1, 0x8f, 0xbf,
	0xb2, 0x77, 0x1b, 0xc6, 0xcb, 0x07, 0x98, 0x4d, 0x23, 0xa1, 0xe0, 0xf0, 0x5b, 0x4a, 0x9c, 0xa5,
	0xa5, 0xad, 0x7e, 0xdf, 0xd4, 0xf2, 0xfc, 0x2f, 0x9d, 0x2e, 0x6f, 0xfe, 0x3d, 0x96, 0x31, 0x70,
	0xf5, 0x1e, 0x00, 0x00, 0xff, 0xff, 0xd0, 0xc1, 0x76, 0xf9, 0x9b, 0x01, 0x00, 0x00,
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
	RegisterAccount(ctx context.Context, in *RegisterAccountRequest, opts ...grpc.CallOption) (*RegisterAccountResponse, error)
}

type accountClient struct {
	cc *grpc.ClientConn
}

func NewAccountClient(cc *grpc.ClientConn) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) RegisterAccount(ctx context.Context, in *RegisterAccountRequest, opts ...grpc.CallOption) (*RegisterAccountResponse, error) {
	out := new(RegisterAccountResponse)
	err := c.cc.Invoke(ctx, "/proto.Account/RegisterAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
type AccountServer interface {
	RegisterAccount(context.Context, *RegisterAccountRequest) (*RegisterAccountResponse, error)
}

// UnimplementedAccountServer can be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (*UnimplementedAccountServer) RegisterAccount(ctx context.Context, req *RegisterAccountRequest) (*RegisterAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAccount not implemented")
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

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterAccount",
			Handler:    _Account_RegisterAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/adapter/rpc/proto/account.proto",
}
