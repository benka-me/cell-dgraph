// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/benka-me/cell-dgraph/protobuf/dgraph.proto

package dgraph

import (
	context "context"
	fmt "fmt"
	hive "github.com/benka-me/hive/go-pkg/hive"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("github.com/benka-me/cell-dgraph/protobuf/dgraph.proto", fileDescriptor_a8528d43d4c2783c)
}

var fileDescriptor_a8528d43d4c2783c = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0x86, 0x5b, 0x3e, 0xbe, 0x94, 0x9e, 0xaa, 0xc8, 0x2c, 0xb3, 0x88, 0x3f, 0x14, 0x44, 0x30,
	0x0d, 0x28, 0xbd, 0x01, 0xad, 0x54, 0x41, 0x5c, 0xb4, 0x14, 0xc1, 0x5d, 0xd2, 0x1c, 0x27, 0x43,
	0x33, 0x99, 0x74, 0x32, 0x51, 0xbc, 0x09, 0xf1, 0xb2, 0x5c, 0x76, 0xe9, 0x52, 0xda, 0x1b, 0x91,
	0xcc, 0x24, 0x36, 0x8a, 0xd4, 0xe5, 0x73, 0xe6, 0x79, 0xdf, 0xc0, 0xc9, 0x81, 0x3e, 0x65, 0x2a,
	0xca, 0x83, 0xde, 0x54, 0x70, 0x2f, 0xc0, 0x64, 0xe6, 0xbb, 0x1c, 0xbd, 0x29, 0xc6, 0xb1, 0x1b,
	0x52, 0xe9, 0xa7, 0x91, 0x97, 0x4a, 0xa1, 0x44, 0x90, 0x3f, 0x78, 0x86, 0x7b, 0x9a, 0x89, 0x65,
	0xc8, 0x3e, 0xf9, 0x2d, 0x1e, 0xb1, 0x47, 0x5c, 0xe7, 0x0a, 0x32, 0xa9, 0xd3, 0x97, 0x7f, 0x00,
	0x26, 0x38, 0xc9, 0x50, 0x92, 0x43, 0xb0, 0xae, 0x13, 0xa6, 0x06, 0x01, 0xe9, 0xf4, 0xb4, 0x75,
	0xc9, 0x53, 0xf5, 0x6c, 0xd7, 0x81, 0x1c, 0xc3, 0xd6, 0x08, 0x29, 0xcb, 0x14, 0x4a, 0x9d, 0xd9,
	0x36, 0x8f, 0xb7, 0xf8, 0x54, 0xa0, 0x0d, 0x06, 0xf5, 0xd3, 0x11, 0xfc, 0xbf, 0x11, 0x94, 0x25,
	0x84, 0x98, 0xa1, 0x86, 0x11, 0xce, 0x73, 0xcc, 0xd4, 0x37, 0x71, 0x0f, 0xac, 0x09, 0x0b, 0x87,
	0xa8, 0x48, 0xdb, 0x4c, 0xc7, 0x4a, 0xfe, 0x10, 0xda, 0x13, 0x16, 0x0e, 0x30, 0x46, 0x85, 0x75,
	0xa7, 0x55, 0x7e, 0x9c, 0x74, 0xa1, 0x53, 0x88, 0x89, 0xcf, 0x71, 0x43, 0x4d, 0x17, 0x76, 0x2a,
	0x6b, 0x43, 0xd7, 0x3e, 0xb4, 0x0a, 0x6b, 0x8c, 0x8a, 0xd4, 0xc2, 0xf6, 0x5a, 0x25, 0x07, 0x00,
	0xc5, 0xa8, 0xec, 0xa8, 0x4b, 0x5f, 0x25, 0x7d, 0xd8, 0xad, 0xd6, 0x74, 0x91, 0x67, 0x4a, 0x70,
	0x94, 0xd5, 0x1a, 0xae, 0xfc, 0x24, 0xcc, 0x22, 0x7f, 0x86, 0x23, 0x9c, 0xdb, 0xe5, 0xfa, 0xee,
	0x30, 0x9e, 0x0a, 0x8e, 0xe7, 0xc3, 0xb7, 0xa5, 0xd3, 0x5c, 0x2c, 0x9d, 0xe6, 0xc7, 0xd2, 0x69,
	0xbe, 0xae, 0x9c, 0xc6, 0x62, 0xe5, 0x34, 0xde, 0x57, 0x4e, 0xe3, 0xde, 0xfd, 0xeb, 0x2e, 0xa8,
	0x70, 0xd3, 0x19, 0x2d, 0xaf, 0x22, 0xb0, 0xf4, 0x0f, 0x3e, 0xfb, 0x0c, 0x00, 0x00, 0xff, 0xff,
	0x0e, 0x53, 0xf7, 0x08, 0x4f, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DgraphUserClient is the client API for DgraphUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DgraphUserClient interface {
	InitDb(ctx context.Context, in *hive.Empty, opts ...grpc.CallOption) (*hive.Empty, error)
	RegisterUser(ctx context.Context, in *hive.NewUser, opts ...grpc.CallOption) (*hive.User, error)
	Login(ctx context.Context, in *hive.LoginRequest, opts ...grpc.CallOption) (*hive.User, error)
	UidGet(ctx context.Context, in *hive.Str, opts ...grpc.CallOption) (*hive.User, error)
	UidDelete(ctx context.Context, in *hive.Str, opts ...grpc.CallOption) (*hive.N, error)
	UsernameGet(ctx context.Context, in *hive.Str, opts ...grpc.CallOption) (*hive.User, error)
	UsernameDelete(ctx context.Context, in *hive.Str, opts ...grpc.CallOption) (*hive.N, error)
	UserSet(ctx context.Context, in *hive.User, opts ...grpc.CallOption) (*hive.Str, error)
	UserDelete(ctx context.Context, in *hive.User, opts ...grpc.CallOption) (*hive.N, error)
	RegisterCustomer(ctx context.Context, in *hive.HandshakeReq, opts ...grpc.CallOption) (*hive.Welcome, error)
}

type dgraphUserClient struct {
	cc *grpc.ClientConn
}

func NewDgraphUserClient(cc *grpc.ClientConn) DgraphUserClient {
	return &dgraphUserClient{cc}
}

func (c *dgraphUserClient) InitDb(ctx context.Context, in *hive.Empty, opts ...grpc.CallOption) (*hive.Empty, error) {
	out := new(hive.Empty)
	err := c.cc.Invoke(ctx, "/dgraph.dgraphUser/InitDb", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dgraphUserClient) RegisterUser(ctx context.Context, in *hive.NewUser, opts ...grpc.CallOption) (*hive.User, error) {
	out := new(hive.User)
	err := c.cc.Invoke(ctx, "/dgraph.dgraphUser/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dgraphUserClient) Login(ctx context.Context, in *hive.LoginRequest, opts ...grpc.CallOption) (*hive.User, error) {
	out := new(hive.User)
	err := c.cc.Invoke(ctx, "/dgraph.dgraphUser/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dgraphUserClient) UidGet(ctx context.Context, in *hive.Str, opts ...grpc.CallOption) (*hive.User, error) {
	out := new(hive.User)
	err := c.cc.Invoke(ctx, "/dgraph.dgraphUser/UidGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dgraphUserClient) UidDelete(ctx context.Context, in *hive.Str, opts ...grpc.CallOption) (*hive.N, error) {
	out := new(hive.N)
	err := c.cc.Invoke(ctx, "/dgraph.dgraphUser/UidDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dgraphUserClient) UsernameGet(ctx context.Context, in *hive.Str, opts ...grpc.CallOption) (*hive.User, error) {
	out := new(hive.User)
	err := c.cc.Invoke(ctx, "/dgraph.dgraphUser/UsernameGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dgraphUserClient) UsernameDelete(ctx context.Context, in *hive.Str, opts ...grpc.CallOption) (*hive.N, error) {
	out := new(hive.N)
	err := c.cc.Invoke(ctx, "/dgraph.dgraphUser/UsernameDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dgraphUserClient) UserSet(ctx context.Context, in *hive.User, opts ...grpc.CallOption) (*hive.Str, error) {
	out := new(hive.Str)
	err := c.cc.Invoke(ctx, "/dgraph.dgraphUser/UserSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dgraphUserClient) UserDelete(ctx context.Context, in *hive.User, opts ...grpc.CallOption) (*hive.N, error) {
	out := new(hive.N)
	err := c.cc.Invoke(ctx, "/dgraph.dgraphUser/UserDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dgraphUserClient) RegisterCustomer(ctx context.Context, in *hive.HandshakeReq, opts ...grpc.CallOption) (*hive.Welcome, error) {
	out := new(hive.Welcome)
	err := c.cc.Invoke(ctx, "/dgraph.dgraphUser/RegisterCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DgraphUserServer is the server API for DgraphUser service.
type DgraphUserServer interface {
	InitDb(context.Context, *hive.Empty) (*hive.Empty, error)
	RegisterUser(context.Context, *hive.NewUser) (*hive.User, error)
	Login(context.Context, *hive.LoginRequest) (*hive.User, error)
	UidGet(context.Context, *hive.Str) (*hive.User, error)
	UidDelete(context.Context, *hive.Str) (*hive.N, error)
	UsernameGet(context.Context, *hive.Str) (*hive.User, error)
	UsernameDelete(context.Context, *hive.Str) (*hive.N, error)
	UserSet(context.Context, *hive.User) (*hive.Str, error)
	UserDelete(context.Context, *hive.User) (*hive.N, error)
	RegisterCustomer(context.Context, *hive.HandshakeReq) (*hive.Welcome, error)
}

// UnimplementedDgraphUserServer can be embedded to have forward compatible implementations.
type UnimplementedDgraphUserServer struct {
}

func (*UnimplementedDgraphUserServer) InitDb(ctx context.Context, req *hive.Empty) (*hive.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitDb not implemented")
}
func (*UnimplementedDgraphUserServer) RegisterUser(ctx context.Context, req *hive.NewUser) (*hive.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (*UnimplementedDgraphUserServer) Login(ctx context.Context, req *hive.LoginRequest) (*hive.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedDgraphUserServer) UidGet(ctx context.Context, req *hive.Str) (*hive.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UidGet not implemented")
}
func (*UnimplementedDgraphUserServer) UidDelete(ctx context.Context, req *hive.Str) (*hive.N, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UidDelete not implemented")
}
func (*UnimplementedDgraphUserServer) UsernameGet(ctx context.Context, req *hive.Str) (*hive.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsernameGet not implemented")
}
func (*UnimplementedDgraphUserServer) UsernameDelete(ctx context.Context, req *hive.Str) (*hive.N, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsernameDelete not implemented")
}
func (*UnimplementedDgraphUserServer) UserSet(ctx context.Context, req *hive.User) (*hive.Str, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSet not implemented")
}
func (*UnimplementedDgraphUserServer) UserDelete(ctx context.Context, req *hive.User) (*hive.N, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDelete not implemented")
}
func (*UnimplementedDgraphUserServer) RegisterCustomer(ctx context.Context, req *hive.HandshakeReq) (*hive.Welcome, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterCustomer not implemented")
}

func RegisterDgraphUserServer(s *grpc.Server, srv DgraphUserServer) {
	s.RegisterService(&_DgraphUser_serviceDesc, srv)
}

func _DgraphUser_InitDb_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hive.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DgraphUserServer).InitDb(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dgraph.dgraphUser/InitDb",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DgraphUserServer).InitDb(ctx, req.(*hive.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DgraphUser_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hive.NewUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DgraphUserServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dgraph.dgraphUser/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DgraphUserServer).RegisterUser(ctx, req.(*hive.NewUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _DgraphUser_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hive.LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DgraphUserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dgraph.dgraphUser/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DgraphUserServer).Login(ctx, req.(*hive.LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DgraphUser_UidGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hive.Str)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DgraphUserServer).UidGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dgraph.dgraphUser/UidGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DgraphUserServer).UidGet(ctx, req.(*hive.Str))
	}
	return interceptor(ctx, in, info, handler)
}

func _DgraphUser_UidDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hive.Str)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DgraphUserServer).UidDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dgraph.dgraphUser/UidDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DgraphUserServer).UidDelete(ctx, req.(*hive.Str))
	}
	return interceptor(ctx, in, info, handler)
}

func _DgraphUser_UsernameGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hive.Str)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DgraphUserServer).UsernameGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dgraph.dgraphUser/UsernameGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DgraphUserServer).UsernameGet(ctx, req.(*hive.Str))
	}
	return interceptor(ctx, in, info, handler)
}

func _DgraphUser_UsernameDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hive.Str)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DgraphUserServer).UsernameDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dgraph.dgraphUser/UsernameDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DgraphUserServer).UsernameDelete(ctx, req.(*hive.Str))
	}
	return interceptor(ctx, in, info, handler)
}

func _DgraphUser_UserSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hive.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DgraphUserServer).UserSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dgraph.dgraphUser/UserSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DgraphUserServer).UserSet(ctx, req.(*hive.User))
	}
	return interceptor(ctx, in, info, handler)
}

func _DgraphUser_UserDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hive.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DgraphUserServer).UserDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dgraph.dgraphUser/UserDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DgraphUserServer).UserDelete(ctx, req.(*hive.User))
	}
	return interceptor(ctx, in, info, handler)
}

func _DgraphUser_RegisterCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hive.HandshakeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DgraphUserServer).RegisterCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dgraph.dgraphUser/RegisterCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DgraphUserServer).RegisterCustomer(ctx, req.(*hive.HandshakeReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _DgraphUser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dgraph.dgraphUser",
	HandlerType: (*DgraphUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitDb",
			Handler:    _DgraphUser_InitDb_Handler,
		},
		{
			MethodName: "RegisterUser",
			Handler:    _DgraphUser_RegisterUser_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _DgraphUser_Login_Handler,
		},
		{
			MethodName: "UidGet",
			Handler:    _DgraphUser_UidGet_Handler,
		},
		{
			MethodName: "UidDelete",
			Handler:    _DgraphUser_UidDelete_Handler,
		},
		{
			MethodName: "UsernameGet",
			Handler:    _DgraphUser_UsernameGet_Handler,
		},
		{
			MethodName: "UsernameDelete",
			Handler:    _DgraphUser_UsernameDelete_Handler,
		},
		{
			MethodName: "UserSet",
			Handler:    _DgraphUser_UserSet_Handler,
		},
		{
			MethodName: "UserDelete",
			Handler:    _DgraphUser_UserDelete_Handler,
		},
		{
			MethodName: "RegisterCustomer",
			Handler:    _DgraphUser_RegisterCustomer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/benka-me/cell-dgraph/protobuf/dgraph.proto",
}
