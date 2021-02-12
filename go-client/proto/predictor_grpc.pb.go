// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// PredictorClient is the client API for Predictor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PredictorClient interface {
	Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictResponse, error)
	PredictProcessed(ctx context.Context, in *PredictProcessedRequest, opts ...grpc.CallOption) (*PredictResponse, error)
}

type predictorClient struct {
	cc grpc.ClientConnInterface
}

func NewPredictorClient(cc grpc.ClientConnInterface) PredictorClient {
	return &predictorClient{cc}
}

func (c *predictorClient) Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictResponse, error) {
	out := new(PredictResponse)
	err := c.cc.Invoke(ctx, "/predictor.Predictor/Predict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictorClient) PredictProcessed(ctx context.Context, in *PredictProcessedRequest, opts ...grpc.CallOption) (*PredictResponse, error) {
	out := new(PredictResponse)
	err := c.cc.Invoke(ctx, "/predictor.Predictor/PredictProcessed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PredictorServer is the server API for Predictor service.
// All implementations must embed UnimplementedPredictorServer
// for forward compatibility
type PredictorServer interface {
	Predict(context.Context, *PredictRequest) (*PredictResponse, error)
	PredictProcessed(context.Context, *PredictProcessedRequest) (*PredictResponse, error)
	mustEmbedUnimplementedPredictorServer()
}

// UnimplementedPredictorServer must be embedded to have forward compatible implementations.
type UnimplementedPredictorServer struct {
}

func (UnimplementedPredictorServer) Predict(context.Context, *PredictRequest) (*PredictResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Predict not implemented")
}
func (UnimplementedPredictorServer) PredictProcessed(context.Context, *PredictProcessedRequest) (*PredictResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PredictProcessed not implemented")
}
func (UnimplementedPredictorServer) mustEmbedUnimplementedPredictorServer() {}

// UnsafePredictorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PredictorServer will
// result in compilation errors.
type UnsafePredictorServer interface {
	mustEmbedUnimplementedPredictorServer()
}

func RegisterPredictorServer(s grpc.ServiceRegistrar, srv PredictorServer) {
	s.RegisterService(&Predictor_ServiceDesc, srv)
}

func _Predictor_Predict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictorServer).Predict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/predictor.Predictor/Predict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictorServer).Predict(ctx, req.(*PredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Predictor_PredictProcessed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictProcessedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictorServer).PredictProcessed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/predictor.Predictor/PredictProcessed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictorServer).PredictProcessed(ctx, req.(*PredictProcessedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Predictor_ServiceDesc is the grpc.ServiceDesc for Predictor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Predictor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "predictor.Predictor",
	HandlerType: (*PredictorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Predict",
			Handler:    _Predictor_Predict_Handler,
		},
		{
			MethodName: "PredictProcessed",
			Handler:    _Predictor_PredictProcessed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/predictor.proto",
}