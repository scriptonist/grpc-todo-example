// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TodoAPIClient is the client API for TodoAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoAPIClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Void, error)
	Read(ctx context.Context, in *Void, opts ...grpc.CallOption) (TodoAPI_ReadClient, error)
}

type todoAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoAPIClient(cc grpc.ClientConnInterface) TodoAPIClient {
	return &todoAPIClient{cc}
}

func (c *todoAPIClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/todo.TodoAPI/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoAPIClient) Read(ctx context.Context, in *Void, opts ...grpc.CallOption) (TodoAPI_ReadClient, error) {
	stream, err := c.cc.NewStream(ctx, &TodoAPI_ServiceDesc.Streams[0], "/todo.TodoAPI/Read", opts...)
	if err != nil {
		return nil, err
	}
	x := &todoAPIReadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TodoAPI_ReadClient interface {
	Recv() (*Todo, error)
	grpc.ClientStream
}

type todoAPIReadClient struct {
	grpc.ClientStream
}

func (x *todoAPIReadClient) Recv() (*Todo, error) {
	m := new(Todo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TodoAPIServer is the server API for TodoAPI service.
// All implementations must embed UnimplementedTodoAPIServer
// for forward compatibility
type TodoAPIServer interface {
	Create(context.Context, *CreateRequest) (*Void, error)
	Read(*Void, TodoAPI_ReadServer) error
	mustEmbedUnimplementedTodoAPIServer()
}

// UnimplementedTodoAPIServer must be embedded to have forward compatible implementations.
type UnimplementedTodoAPIServer struct {
}

func (UnimplementedTodoAPIServer) Create(context.Context, *CreateRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTodoAPIServer) Read(*Void, TodoAPI_ReadServer) error {
	return status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedTodoAPIServer) mustEmbedUnimplementedTodoAPIServer() {}

// UnsafeTodoAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoAPIServer will
// result in compilation errors.
type UnsafeTodoAPIServer interface {
	mustEmbedUnimplementedTodoAPIServer()
}

func RegisterTodoAPIServer(s grpc.ServiceRegistrar, srv TodoAPIServer) {
	s.RegisterService(&TodoAPI_ServiceDesc, srv)
}

func _TodoAPI_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoAPIServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoAPI/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoAPIServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoAPI_Read_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Void)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TodoAPIServer).Read(m, &todoAPIReadServer{stream})
}

type TodoAPI_ReadServer interface {
	Send(*Todo) error
	grpc.ServerStream
}

type todoAPIReadServer struct {
	grpc.ServerStream
}

func (x *todoAPIReadServer) Send(m *Todo) error {
	return x.ServerStream.SendMsg(m)
}

// TodoAPI_ServiceDesc is the grpc.ServiceDesc for TodoAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todo.TodoAPI",
	HandlerType: (*TodoAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TodoAPI_Create_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Read",
			Handler:       _TodoAPI_Read_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "internal/api/api.proto",
}