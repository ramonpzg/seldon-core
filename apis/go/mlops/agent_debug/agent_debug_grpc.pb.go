/*
Copyright 2022 Seldon Technologies Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: mlops/agent_debug/agent_debug.proto

package agent_debug

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

// AgentDebugServiceClient is the client API for AgentDebugService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentDebugServiceClient interface {
	ReplicaStatus(ctx context.Context, in *ReplicaStatusRequest, opts ...grpc.CallOption) (*ReplicaStatusResponse, error)
}

type agentDebugServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentDebugServiceClient(cc grpc.ClientConnInterface) AgentDebugServiceClient {
	return &agentDebugServiceClient{cc}
}

func (c *agentDebugServiceClient) ReplicaStatus(ctx context.Context, in *ReplicaStatusRequest, opts ...grpc.CallOption) (*ReplicaStatusResponse, error) {
	out := new(ReplicaStatusResponse)
	err := c.cc.Invoke(ctx, "/seldon.mlops.agent_debug.AgentDebugService/ReplicaStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentDebugServiceServer is the server API for AgentDebugService service.
// All implementations must embed UnimplementedAgentDebugServiceServer
// for forward compatibility
type AgentDebugServiceServer interface {
	ReplicaStatus(context.Context, *ReplicaStatusRequest) (*ReplicaStatusResponse, error)
	mustEmbedUnimplementedAgentDebugServiceServer()
}

// UnimplementedAgentDebugServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAgentDebugServiceServer struct {
}

func (UnimplementedAgentDebugServiceServer) ReplicaStatus(context.Context, *ReplicaStatusRequest) (*ReplicaStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplicaStatus not implemented")
}
func (UnimplementedAgentDebugServiceServer) mustEmbedUnimplementedAgentDebugServiceServer() {}

// UnsafeAgentDebugServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentDebugServiceServer will
// result in compilation errors.
type UnsafeAgentDebugServiceServer interface {
	mustEmbedUnimplementedAgentDebugServiceServer()
}

func RegisterAgentDebugServiceServer(s grpc.ServiceRegistrar, srv AgentDebugServiceServer) {
	s.RegisterService(&AgentDebugService_ServiceDesc, srv)
}

func _AgentDebugService_ReplicaStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplicaStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentDebugServiceServer).ReplicaStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seldon.mlops.agent_debug.AgentDebugService/ReplicaStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentDebugServiceServer).ReplicaStatus(ctx, req.(*ReplicaStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AgentDebugService_ServiceDesc is the grpc.ServiceDesc for AgentDebugService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentDebugService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "seldon.mlops.agent_debug.AgentDebugService",
	HandlerType: (*AgentDebugServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReplicaStatus",
			Handler:    _AgentDebugService_ReplicaStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mlops/agent_debug/agent_debug.proto",
}
