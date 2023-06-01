// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: sync/sync.proto

package sync

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SyncableDB_GetMerkleRoot_FullMethodName     = "/sync.SyncableDB/GetMerkleRoot"
	SyncableDB_GetProof_FullMethodName          = "/sync.SyncableDB/GetProof"
	SyncableDB_GetChangeProof_FullMethodName    = "/sync.SyncableDB/GetChangeProof"
	SyncableDB_VerifyChangeProof_FullMethodName = "/sync.SyncableDB/VerifyChangeProof"
	SyncableDB_CommitChangeProof_FullMethodName = "/sync.SyncableDB/CommitChangeProof"
	SyncableDB_GetRangeProof_FullMethodName     = "/sync.SyncableDB/GetRangeProof"
	SyncableDB_CommitRangeProof_FullMethodName  = "/sync.SyncableDB/CommitRangeProof"
)

// SyncableDBClient is the client API for SyncableDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SyncableDBClient interface {
	GetMerkleRoot(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetMerkleRootResponse, error)
	GetProof(ctx context.Context, in *GetProofRequest, opts ...grpc.CallOption) (*GetProofResponse, error)
	GetChangeProof(ctx context.Context, in *GetChangeProofRequest, opts ...grpc.CallOption) (*GetChangeProofResponse, error)
	VerifyChangeProof(ctx context.Context, in *VerifyChangeProofRequest, opts ...grpc.CallOption) (*VerifyChangeProofResponse, error)
	CommitChangeProof(ctx context.Context, in *CommitChangeProofRequest, opts ...grpc.CallOption) (*CommitChangeProofResponse, error)
	GetRangeProof(ctx context.Context, in *GetRangeProofRequest, opts ...grpc.CallOption) (*GetRangeProofResponse, error)
	CommitRangeProof(ctx context.Context, in *CommitRangeProofRequest, opts ...grpc.CallOption) (*CommitRangeProofResponse, error)
}

type syncableDBClient struct {
	cc grpc.ClientConnInterface
}

func NewSyncableDBClient(cc grpc.ClientConnInterface) SyncableDBClient {
	return &syncableDBClient{cc}
}

func (c *syncableDBClient) GetMerkleRoot(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetMerkleRootResponse, error) {
	out := new(GetMerkleRootResponse)
	err := c.cc.Invoke(ctx, SyncableDB_GetMerkleRoot_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncableDBClient) GetProof(ctx context.Context, in *GetProofRequest, opts ...grpc.CallOption) (*GetProofResponse, error) {
	out := new(GetProofResponse)
	err := c.cc.Invoke(ctx, SyncableDB_GetProof_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncableDBClient) GetChangeProof(ctx context.Context, in *GetChangeProofRequest, opts ...grpc.CallOption) (*GetChangeProofResponse, error) {
	out := new(GetChangeProofResponse)
	err := c.cc.Invoke(ctx, SyncableDB_GetChangeProof_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncableDBClient) VerifyChangeProof(ctx context.Context, in *VerifyChangeProofRequest, opts ...grpc.CallOption) (*VerifyChangeProofResponse, error) {
	out := new(VerifyChangeProofResponse)
	err := c.cc.Invoke(ctx, SyncableDB_VerifyChangeProof_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncableDBClient) CommitChangeProof(ctx context.Context, in *CommitChangeProofRequest, opts ...grpc.CallOption) (*CommitChangeProofResponse, error) {
	out := new(CommitChangeProofResponse)
	err := c.cc.Invoke(ctx, SyncableDB_CommitChangeProof_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncableDBClient) GetRangeProof(ctx context.Context, in *GetRangeProofRequest, opts ...grpc.CallOption) (*GetRangeProofResponse, error) {
	out := new(GetRangeProofResponse)
	err := c.cc.Invoke(ctx, SyncableDB_GetRangeProof_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncableDBClient) CommitRangeProof(ctx context.Context, in *CommitRangeProofRequest, opts ...grpc.CallOption) (*CommitRangeProofResponse, error) {
	out := new(CommitRangeProofResponse)
	err := c.cc.Invoke(ctx, SyncableDB_CommitRangeProof_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SyncableDBServer is the server API for SyncableDB service.
// All implementations must embed UnimplementedSyncableDBServer
// for forward compatibility
type SyncableDBServer interface {
	GetMerkleRoot(context.Context, *emptypb.Empty) (*GetMerkleRootResponse, error)
	GetProof(context.Context, *GetProofRequest) (*GetProofResponse, error)
	GetChangeProof(context.Context, *GetChangeProofRequest) (*GetChangeProofResponse, error)
	VerifyChangeProof(context.Context, *VerifyChangeProofRequest) (*VerifyChangeProofResponse, error)
	CommitChangeProof(context.Context, *CommitChangeProofRequest) (*CommitChangeProofResponse, error)
	GetRangeProof(context.Context, *GetRangeProofRequest) (*GetRangeProofResponse, error)
	CommitRangeProof(context.Context, *CommitRangeProofRequest) (*CommitRangeProofResponse, error)
	mustEmbedUnimplementedSyncableDBServer()
}

// UnimplementedSyncableDBServer must be embedded to have forward compatible implementations.
type UnimplementedSyncableDBServer struct {
}

func (UnimplementedSyncableDBServer) GetMerkleRoot(context.Context, *emptypb.Empty) (*GetMerkleRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMerkleRoot not implemented")
}
func (UnimplementedSyncableDBServer) GetProof(context.Context, *GetProofRequest) (*GetProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProof not implemented")
}
func (UnimplementedSyncableDBServer) GetChangeProof(context.Context, *GetChangeProofRequest) (*GetChangeProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChangeProof not implemented")
}
func (UnimplementedSyncableDBServer) VerifyChangeProof(context.Context, *VerifyChangeProofRequest) (*VerifyChangeProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyChangeProof not implemented")
}
func (UnimplementedSyncableDBServer) CommitChangeProof(context.Context, *CommitChangeProofRequest) (*CommitChangeProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommitChangeProof not implemented")
}
func (UnimplementedSyncableDBServer) GetRangeProof(context.Context, *GetRangeProofRequest) (*GetRangeProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRangeProof not implemented")
}
func (UnimplementedSyncableDBServer) CommitRangeProof(context.Context, *CommitRangeProofRequest) (*CommitRangeProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommitRangeProof not implemented")
}
func (UnimplementedSyncableDBServer) mustEmbedUnimplementedSyncableDBServer() {}

// UnsafeSyncableDBServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SyncableDBServer will
// result in compilation errors.
type UnsafeSyncableDBServer interface {
	mustEmbedUnimplementedSyncableDBServer()
}

func RegisterSyncableDBServer(s grpc.ServiceRegistrar, srv SyncableDBServer) {
	s.RegisterService(&SyncableDB_ServiceDesc, srv)
}

func _SyncableDB_GetMerkleRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncableDBServer).GetMerkleRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncableDB_GetMerkleRoot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncableDBServer).GetMerkleRoot(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncableDB_GetProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncableDBServer).GetProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncableDB_GetProof_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncableDBServer).GetProof(ctx, req.(*GetProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncableDB_GetChangeProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChangeProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncableDBServer).GetChangeProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncableDB_GetChangeProof_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncableDBServer).GetChangeProof(ctx, req.(*GetChangeProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncableDB_VerifyChangeProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyChangeProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncableDBServer).VerifyChangeProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncableDB_VerifyChangeProof_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncableDBServer).VerifyChangeProof(ctx, req.(*VerifyChangeProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncableDB_CommitChangeProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitChangeProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncableDBServer).CommitChangeProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncableDB_CommitChangeProof_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncableDBServer).CommitChangeProof(ctx, req.(*CommitChangeProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncableDB_GetRangeProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRangeProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncableDBServer).GetRangeProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncableDB_GetRangeProof_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncableDBServer).GetRangeProof(ctx, req.(*GetRangeProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncableDB_CommitRangeProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitRangeProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncableDBServer).CommitRangeProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncableDB_CommitRangeProof_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncableDBServer).CommitRangeProof(ctx, req.(*CommitRangeProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SyncableDB_ServiceDesc is the grpc.ServiceDesc for SyncableDB service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SyncableDB_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sync.SyncableDB",
	HandlerType: (*SyncableDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMerkleRoot",
			Handler:    _SyncableDB_GetMerkleRoot_Handler,
		},
		{
			MethodName: "GetProof",
			Handler:    _SyncableDB_GetProof_Handler,
		},
		{
			MethodName: "GetChangeProof",
			Handler:    _SyncableDB_GetChangeProof_Handler,
		},
		{
			MethodName: "VerifyChangeProof",
			Handler:    _SyncableDB_VerifyChangeProof_Handler,
		},
		{
			MethodName: "CommitChangeProof",
			Handler:    _SyncableDB_CommitChangeProof_Handler,
		},
		{
			MethodName: "GetRangeProof",
			Handler:    _SyncableDB_GetRangeProof_Handler,
		},
		{
			MethodName: "CommitRangeProof",
			Handler:    _SyncableDB_CommitRangeProof_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sync/sync.proto",
}
