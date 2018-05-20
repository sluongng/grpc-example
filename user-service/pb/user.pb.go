// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// Requests
type GetUserRequest struct {
	UserId               int32    `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_d52ecf8f7b1f5980, []int{0}
}
func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (dst *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(dst, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

// Replies
type UserReply struct {
	User                 *User    `protobuf:"bytes,1,opt,name=pb" json:"pb,omitempty"`
	Roles                []*Role  `protobuf:"bytes,2,rep,name=roles" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserReply) Reset()         { *m = UserReply{} }
func (m *UserReply) String() string { return proto.CompactTextString(m) }
func (*UserReply) ProtoMessage()    {}
func (*UserReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_d52ecf8f7b1f5980, []int{1}
}
func (m *UserReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserReply.Unmarshal(m, b)
}
func (m *UserReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserReply.Marshal(b, m, deterministic)
}
func (dst *UserReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserReply.Merge(dst, src)
}
func (m *UserReply) XXX_Size() int {
	return xxx_messageInfo_UserReply.Size(m)
}
func (m *UserReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UserReply.DiscardUnknown(m)
}

var xxx_messageInfo_UserReply proto.InternalMessageInfo

func (m *UserReply) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UserReply) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

type User struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_d52ecf8f7b1f5980, []int{2}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type Role struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_d52ecf8f7b1f5980, []int{3}
}
func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (dst *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(dst, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*GetUserRequest)(nil), "pb.GetUserRequest")
	proto.RegisterType((*UserReply)(nil), "pb.UserReply")
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*Role)(nil), "pb.Role")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserReply, error)
}

type usersClient struct {
	cc *grpc.ClientConn
}

func NewUsersClient(cc *grpc.ClientConn) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/pb.Users/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
type UsersServer interface {
	GetUser(context.Context, *GetUserRequest) (*UserReply, error)
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Users/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _Users_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb.proto",
}

func init() { proto.RegisterFile("pb.proto", fileDescriptor_user_d52ecf8f7b1f5980) }

var fileDescriptor_user_d52ecf8f7b1f5980 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x3f, 0x4f, 0xc4, 0x30,
	0x0c, 0xc5, 0x69, 0x2e, 0xb9, 0xd3, 0x19, 0xa9, 0x48, 0x56, 0x25, 0x22, 0x06, 0x14, 0x65, 0x2a,
	0x0c, 0x1d, 0xca, 0xc8, 0xc2, 0x86, 0x18, 0x58, 0x22, 0x31, 0xa3, 0xa2, 0x7a, 0xa8, 0x94, 0x92,
	0x92, 0xb4, 0x03, 0xdf, 0x1e, 0xc5, 0x41, 0xfc, 0xd9, 0x6e, 0xb3, 0xfd, 0x5e, 0x9e, 0xf3, 0x33,
	0xc0, 0x96, 0x28, 0x76, 0x4b, 0x0c, 0x6b, 0x40, 0x99, 0x6b, 0x7b, 0x03, 0xf5, 0x23, 0xad, 0x2f,
	0x89, 0xa2, 0xa3, 0x8f, 0x8d, 0xd2, 0x8a, 0x97, 0x70, 0xc8, 0xca, 0xeb, 0x34, 0xea, 0xca, 0x54,
	0xad, 0x72, 0xfb, 0xdc, 0x3e, 0x8d, 0xf6, 0x19, 0x8e, 0xc5, 0xb7, 0xf8, 0x4f, 0xbc, 0x06, 0x7e,
	0xcf, 0x96, 0xf3, 0x1e, 0x3a, 0x0e, 0x66, 0x99, 0xe7, 0x68, 0x40, 0xc5, 0xe0, 0x29, 0x69, 0x61,
	0x76, 0xbf, 0x06, 0x17, 0x3c, 0xb9, 0x22, 0xd8, 0x07, 0x90, 0xd9, 0x8f, 0x35, 0x88, 0x9f, 0x55,
	0x62, 0x1a, 0x11, 0x41, 0xbe, 0x0f, 0x33, 0x69, 0x61, 0xaa, 0xf6, 0xe8, 0xb8, 0xc6, 0x06, 0x14,
	0xcd, 0xc3, 0xe4, 0xf5, 0x8e, 0x87, 0xa5, 0xb1, 0xb7, 0x20, 0x73, 0xe0, 0x29, 0x09, 0xfd, 0x3d,
	0xa8, 0xbc, 0x2d, 0x61, 0x0f, 0x87, 0x6f, 0x60, 0x6c, 0xca, 0xa7, 0xfe, 0xf3, 0x5f, 0x5d, 0xfc,
	0x61, 0xc9, 0xa8, 0xf6, 0xec, 0x6d, 0xcf, 0x17, 0xbb, 0xfb, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xf7,
	0x1a, 0xa3, 0xc7, 0x3f, 0x01, 0x00, 0x00,
}