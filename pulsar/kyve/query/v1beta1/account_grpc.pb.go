// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: kyve/query/v1beta1/account.proto

package queryv1beta1

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

const (
	QueryAccount_AccountAssets_FullMethodName               = "/kyve.query.v1beta1.QueryAccount/AccountAssets"
	QueryAccount_AccountDelegationUnbondings_FullMethodName = "/kyve.query.v1beta1.QueryAccount/AccountDelegationUnbondings"
	QueryAccount_AccountFundedList_FullMethodName           = "/kyve.query.v1beta1.QueryAccount/AccountFundedList"
	QueryAccount_AccountRedelegation_FullMethodName         = "/kyve.query.v1beta1.QueryAccount/AccountRedelegation"
)

// QueryAccountClient is the client API for QueryAccount service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryAccountClient interface {
	// AccountAssets returns an overview of the sum of all balances for a given user. e.g. balance, staking, funding, etc.
	AccountAssets(ctx context.Context, in *QueryAccountAssetsRequest, opts ...grpc.CallOption) (*QueryAccountAssetsResponse, error)
	// AccountDelegationUnbondings ...
	AccountDelegationUnbondings(ctx context.Context, in *QueryAccountDelegationUnbondingsRequest, opts ...grpc.CallOption) (*QueryAccountDelegationUnbondingsResponse, error)
	// AccountFundedList returns all pools the given user has funded into.
	AccountFundedList(ctx context.Context, in *QueryAccountFundedListRequest, opts ...grpc.CallOption) (*QueryAccountFundedListResponse, error)
	// AccountRedelegation ...
	AccountRedelegation(ctx context.Context, in *QueryAccountRedelegationRequest, opts ...grpc.CallOption) (*QueryAccountRedelegationResponse, error)
}

type queryAccountClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryAccountClient(cc grpc.ClientConnInterface) QueryAccountClient {
	return &queryAccountClient{cc}
}

func (c *queryAccountClient) AccountAssets(ctx context.Context, in *QueryAccountAssetsRequest, opts ...grpc.CallOption) (*QueryAccountAssetsResponse, error) {
	out := new(QueryAccountAssetsResponse)
	err := c.cc.Invoke(ctx, QueryAccount_AccountAssets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryAccountClient) AccountDelegationUnbondings(ctx context.Context, in *QueryAccountDelegationUnbondingsRequest, opts ...grpc.CallOption) (*QueryAccountDelegationUnbondingsResponse, error) {
	out := new(QueryAccountDelegationUnbondingsResponse)
	err := c.cc.Invoke(ctx, QueryAccount_AccountDelegationUnbondings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryAccountClient) AccountFundedList(ctx context.Context, in *QueryAccountFundedListRequest, opts ...grpc.CallOption) (*QueryAccountFundedListResponse, error) {
	out := new(QueryAccountFundedListResponse)
	err := c.cc.Invoke(ctx, QueryAccount_AccountFundedList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryAccountClient) AccountRedelegation(ctx context.Context, in *QueryAccountRedelegationRequest, opts ...grpc.CallOption) (*QueryAccountRedelegationResponse, error) {
	out := new(QueryAccountRedelegationResponse)
	err := c.cc.Invoke(ctx, QueryAccount_AccountRedelegation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryAccountServer is the server API for QueryAccount service.
// All implementations must embed UnimplementedQueryAccountServer
// for forward compatibility
type QueryAccountServer interface {
	// AccountAssets returns an overview of the sum of all balances for a given user. e.g. balance, staking, funding, etc.
	AccountAssets(context.Context, *QueryAccountAssetsRequest) (*QueryAccountAssetsResponse, error)
	// AccountDelegationUnbondings ...
	AccountDelegationUnbondings(context.Context, *QueryAccountDelegationUnbondingsRequest) (*QueryAccountDelegationUnbondingsResponse, error)
	// AccountFundedList returns all pools the given user has funded into.
	AccountFundedList(context.Context, *QueryAccountFundedListRequest) (*QueryAccountFundedListResponse, error)
	// AccountRedelegation ...
	AccountRedelegation(context.Context, *QueryAccountRedelegationRequest) (*QueryAccountRedelegationResponse, error)
	mustEmbedUnimplementedQueryAccountServer()
}

// UnimplementedQueryAccountServer must be embedded to have forward compatible implementations.
type UnimplementedQueryAccountServer struct {
}

func (UnimplementedQueryAccountServer) AccountAssets(context.Context, *QueryAccountAssetsRequest) (*QueryAccountAssetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountAssets not implemented")
}
func (UnimplementedQueryAccountServer) AccountDelegationUnbondings(context.Context, *QueryAccountDelegationUnbondingsRequest) (*QueryAccountDelegationUnbondingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountDelegationUnbondings not implemented")
}
func (UnimplementedQueryAccountServer) AccountFundedList(context.Context, *QueryAccountFundedListRequest) (*QueryAccountFundedListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountFundedList not implemented")
}
func (UnimplementedQueryAccountServer) AccountRedelegation(context.Context, *QueryAccountRedelegationRequest) (*QueryAccountRedelegationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountRedelegation not implemented")
}
func (UnimplementedQueryAccountServer) mustEmbedUnimplementedQueryAccountServer() {}

// UnsafeQueryAccountServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryAccountServer will
// result in compilation errors.
type UnsafeQueryAccountServer interface {
	mustEmbedUnimplementedQueryAccountServer()
}

func RegisterQueryAccountServer(s grpc.ServiceRegistrar, srv QueryAccountServer) {
	s.RegisterService(&QueryAccount_ServiceDesc, srv)
}

func _QueryAccount_AccountAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAccountAssetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryAccountServer).AccountAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueryAccount_AccountAssets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryAccountServer).AccountAssets(ctx, req.(*QueryAccountAssetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryAccount_AccountDelegationUnbondings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAccountDelegationUnbondingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryAccountServer).AccountDelegationUnbondings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueryAccount_AccountDelegationUnbondings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryAccountServer).AccountDelegationUnbondings(ctx, req.(*QueryAccountDelegationUnbondingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryAccount_AccountFundedList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAccountFundedListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryAccountServer).AccountFundedList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueryAccount_AccountFundedList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryAccountServer).AccountFundedList(ctx, req.(*QueryAccountFundedListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryAccount_AccountRedelegation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAccountRedelegationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryAccountServer).AccountRedelegation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueryAccount_AccountRedelegation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryAccountServer).AccountRedelegation(ctx, req.(*QueryAccountRedelegationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QueryAccount_ServiceDesc is the grpc.ServiceDesc for QueryAccount service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QueryAccount_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kyve.query.v1beta1.QueryAccount",
	HandlerType: (*QueryAccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AccountAssets",
			Handler:    _QueryAccount_AccountAssets_Handler,
		},
		{
			MethodName: "AccountDelegationUnbondings",
			Handler:    _QueryAccount_AccountDelegationUnbondings_Handler,
		},
		{
			MethodName: "AccountFundedList",
			Handler:    _QueryAccount_AccountFundedList_Handler,
		},
		{
			MethodName: "AccountRedelegation",
			Handler:    _QueryAccount_AccountRedelegation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kyve/query/v1beta1/account.proto",
}