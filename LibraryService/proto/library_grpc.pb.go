// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: proto/library.proto

package proto

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

// LibrarySearchClient is the client API for LibrarySearch service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LibrarySearchClient interface {
	GetAuthor(ctx context.Context, in *GetAuthorRequest, opts ...grpc.CallOption) (*GetAuthorResponse, error)
	GetBooks(ctx context.Context, in *GetBooksRequest, opts ...grpc.CallOption) (*GetBooksResponse, error)
}

type librarySearchClient struct {
	cc grpc.ClientConnInterface
}

func NewLibrarySearchClient(cc grpc.ClientConnInterface) LibrarySearchClient {
	return &librarySearchClient{cc}
}

func (c *librarySearchClient) GetAuthor(ctx context.Context, in *GetAuthorRequest, opts ...grpc.CallOption) (*GetAuthorResponse, error) {
	out := new(GetAuthorResponse)
	err := c.cc.Invoke(ctx, "/library.search.v1.LibrarySearch/GetAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *librarySearchClient) GetBooks(ctx context.Context, in *GetBooksRequest, opts ...grpc.CallOption) (*GetBooksResponse, error) {
	out := new(GetBooksResponse)
	err := c.cc.Invoke(ctx, "/library.search.v1.LibrarySearch/GetBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LibrarySearchServer is the server API for LibrarySearch service.
// All implementations must embed UnimplementedLibrarySearchServer
// for forward compatibility
type LibrarySearchServer interface {
	GetAuthor(context.Context, *GetAuthorRequest) (*GetAuthorResponse, error)
	GetBooks(context.Context, *GetBooksRequest) (*GetBooksResponse, error)
	mustEmbedUnimplementedLibrarySearchServer()
}

// UnimplementedLibrarySearchServer must be embedded to have forward compatible implementations.
type UnimplementedLibrarySearchServer struct {
}

func (UnimplementedLibrarySearchServer) GetAuthor(context.Context, *GetAuthorRequest) (*GetAuthorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthor not implemented")
}
func (UnimplementedLibrarySearchServer) GetBooks(context.Context, *GetBooksRequest) (*GetBooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBooks not implemented")
}
func (UnimplementedLibrarySearchServer) mustEmbedUnimplementedLibrarySearchServer() {}

// UnsafeLibrarySearchServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LibrarySearchServer will
// result in compilation errors.
type UnsafeLibrarySearchServer interface {
	mustEmbedUnimplementedLibrarySearchServer()
}

func RegisterLibrarySearchServer(s grpc.ServiceRegistrar, srv LibrarySearchServer) {
	s.RegisterService(&LibrarySearch_ServiceDesc, srv)
}

func _LibrarySearch_GetAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibrarySearchServer).GetAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.search.v1.LibrarySearch/GetAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibrarySearchServer).GetAuthor(ctx, req.(*GetAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibrarySearch_GetBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibrarySearchServer).GetBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.search.v1.LibrarySearch/GetBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibrarySearchServer).GetBooks(ctx, req.(*GetBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LibrarySearch_ServiceDesc is the grpc.ServiceDesc for LibrarySearch service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LibrarySearch_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "library.search.v1.LibrarySearch",
	HandlerType: (*LibrarySearchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAuthor",
			Handler:    _LibrarySearch_GetAuthor_Handler,
		},
		{
			MethodName: "GetBooks",
			Handler:    _LibrarySearch_GetBooks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/library.proto",
}
