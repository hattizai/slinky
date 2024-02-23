// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: slinky/marketmap/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

// GetMarketMapRequest is the query request for the GetMarketMap query.
// It takes no arguments.
type GetMarketMapRequest struct {
}

func (m *GetMarketMapRequest) Reset()         { *m = GetMarketMapRequest{} }
func (m *GetMarketMapRequest) String() string { return proto.CompactTextString(m) }
func (*GetMarketMapRequest) ProtoMessage()    {}
func (*GetMarketMapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5d6ff68f3c474a0, []int{0}
}
func (m *GetMarketMapRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetMarketMapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetMarketMapRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetMarketMapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMarketMapRequest.Merge(m, src)
}
func (m *GetMarketMapRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetMarketMapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMarketMapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMarketMapRequest proto.InternalMessageInfo

// GetMarketMapResponse is the query response for the GetMarketMap query.
type GetMarketMapResponse struct {
	// MarketMap defines the global set of market configurations for all providers
	// and markets.
	MarketMap MarketMap `protobuf:"bytes,1,opt,name=market_map,json=marketMap,proto3" json:"market_map"`
	// LastUpdated is the last block height that the market map was updated.
	// This field can be used as an optimization for clients checking if there
	// is a new update to the map.
	LastUpdated int64 `protobuf:"varint,2,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	// Version is the schema version for the MarketMap data structure and query
	// response.
	Version uint64 `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *GetMarketMapResponse) Reset()         { *m = GetMarketMapResponse{} }
func (m *GetMarketMapResponse) String() string { return proto.CompactTextString(m) }
func (*GetMarketMapResponse) ProtoMessage()    {}
func (*GetMarketMapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5d6ff68f3c474a0, []int{1}
}
func (m *GetMarketMapResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetMarketMapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetMarketMapResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetMarketMapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMarketMapResponse.Merge(m, src)
}
func (m *GetMarketMapResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetMarketMapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMarketMapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMarketMapResponse proto.InternalMessageInfo

func (m *GetMarketMapResponse) GetMarketMap() MarketMap {
	if m != nil {
		return m.MarketMap
	}
	return MarketMap{}
}

func (m *GetMarketMapResponse) GetLastUpdated() int64 {
	if m != nil {
		return m.LastUpdated
	}
	return 0
}

func (m *GetMarketMapResponse) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

// ParamsRequest is the request type for the Query/Params RPC method.
type ParamsRequest struct {
}

func (m *ParamsRequest) Reset()         { *m = ParamsRequest{} }
func (m *ParamsRequest) String() string { return proto.CompactTextString(m) }
func (*ParamsRequest) ProtoMessage()    {}
func (*ParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5d6ff68f3c474a0, []int{2}
}
func (m *ParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParamsRequest.Merge(m, src)
}
func (m *ParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *ParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ParamsRequest proto.InternalMessageInfo

// ParamsResponse is the response type for the Query/Params RPC method.
type ParamsResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *ParamsResponse) Reset()         { *m = ParamsResponse{} }
func (m *ParamsResponse) String() string { return proto.CompactTextString(m) }
func (*ParamsResponse) ProtoMessage()    {}
func (*ParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5d6ff68f3c474a0, []int{3}
}
func (m *ParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParamsResponse.Merge(m, src)
}
func (m *ParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *ParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ParamsResponse proto.InternalMessageInfo

func (m *ParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// GetLastUpdatedRequest is the request type for the Query/LastUpdated RPC
// method.
type GetLastUpdatedRequest struct {
}

func (m *GetLastUpdatedRequest) Reset()         { *m = GetLastUpdatedRequest{} }
func (m *GetLastUpdatedRequest) String() string { return proto.CompactTextString(m) }
func (*GetLastUpdatedRequest) ProtoMessage()    {}
func (*GetLastUpdatedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5d6ff68f3c474a0, []int{4}
}
func (m *GetLastUpdatedRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetLastUpdatedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetLastUpdatedRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetLastUpdatedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLastUpdatedRequest.Merge(m, src)
}
func (m *GetLastUpdatedRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetLastUpdatedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLastUpdatedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLastUpdatedRequest proto.InternalMessageInfo

// GetLastUpdatedResponse is the response type for the Query/LastUpdated RPC
// method.
type GetLastUpdatedResponse struct {
	LastUpdated int64 `protobuf:"varint,1,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
}

func (m *GetLastUpdatedResponse) Reset()         { *m = GetLastUpdatedResponse{} }
func (m *GetLastUpdatedResponse) String() string { return proto.CompactTextString(m) }
func (*GetLastUpdatedResponse) ProtoMessage()    {}
func (*GetLastUpdatedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5d6ff68f3c474a0, []int{5}
}
func (m *GetLastUpdatedResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetLastUpdatedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetLastUpdatedResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetLastUpdatedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLastUpdatedResponse.Merge(m, src)
}
func (m *GetLastUpdatedResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetLastUpdatedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLastUpdatedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLastUpdatedResponse proto.InternalMessageInfo

func (m *GetLastUpdatedResponse) GetLastUpdated() int64 {
	if m != nil {
		return m.LastUpdated
	}
	return 0
}

func init() {
	proto.RegisterType((*GetMarketMapRequest)(nil), "slinky.marketmap.v1.GetMarketMapRequest")
	proto.RegisterType((*GetMarketMapResponse)(nil), "slinky.marketmap.v1.GetMarketMapResponse")
	proto.RegisterType((*ParamsRequest)(nil), "slinky.marketmap.v1.ParamsRequest")
	proto.RegisterType((*ParamsResponse)(nil), "slinky.marketmap.v1.ParamsResponse")
	proto.RegisterType((*GetLastUpdatedRequest)(nil), "slinky.marketmap.v1.GetLastUpdatedRequest")
	proto.RegisterType((*GetLastUpdatedResponse)(nil), "slinky.marketmap.v1.GetLastUpdatedResponse")
}

func init() { proto.RegisterFile("slinky/marketmap/v1/query.proto", fileDescriptor_b5d6ff68f3c474a0) }

var fileDescriptor_b5d6ff68f3c474a0 = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcd, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xeb, 0x75, 0x14, 0xed, 0x2d, 0x1f, 0x92, 0xb7, 0x41, 0xd5, 0x81, 0x97, 0xa5, 0x12,
	0xca, 0xf8, 0x88, 0xb5, 0x71, 0x42, 0xdc, 0x86, 0xc4, 0x0e, 0x30, 0x09, 0x22, 0x71, 0xe1, 0x32,
	0x79, 0xcc, 0x0a, 0x51, 0xeb, 0xd8, 0x8b, 0x9d, 0x88, 0x5e, 0x39, 0x71, 0x04, 0x71, 0xe1, 0xc4,
	0xdf, 0xb3, 0xe3, 0x24, 0x2e, 0x9c, 0x10, 0x6a, 0xf9, 0x43, 0x50, 0x63, 0xa7, 0x94, 0x2d, 0xad,
	0x76, 0x4b, 0xde, 0xf7, 0xf1, 0xfb, 0xfc, 0x1e, 0x7f, 0xc0, 0xa6, 0x1e, 0x24, 0x69, 0x7f, 0x48,
	0x05, 0xcb, 0xfa, 0xdc, 0x08, 0xa6, 0x68, 0xb1, 0x43, 0x4f, 0x72, 0x9e, 0x0d, 0x43, 0x95, 0x49,
	0x23, 0xf1, 0xaa, 0x15, 0x84, 0x53, 0x41, 0x58, 0xec, 0x74, 0xd7, 0x62, 0x19, 0xcb, 0xb2, 0x4f,
	0x27, 0x5f, 0x56, 0xda, 0xbd, 0x13, 0x4b, 0x19, 0x0f, 0x38, 0x65, 0x2a, 0xa1, 0x2c, 0x4d, 0xa5,
	0x61, 0x26, 0x91, 0xa9, 0x76, 0x5d, 0xaf, 0xce, 0xc9, 0xfe, 0x2c, 0x52, 0x28, 0x96, 0x31, 0xe1,
	0x66, 0xf8, 0xeb, 0xb0, 0xba, 0xcf, 0xcd, 0x41, 0xd9, 0x3f, 0x60, 0x2a, 0xe2, 0x27, 0x39, 0xd7,
	0xc6, 0xff, 0x86, 0x60, 0xed, 0xff, 0xba, 0x56, 0x32, 0xd5, 0x1c, 0x3f, 0x03, 0xb0, 0xc3, 0x0e,
	0x05, 0x53, 0x1d, 0xe4, 0xa1, 0xa0, 0xbd, 0x4b, 0xc2, 0x9a, 0x44, 0xe1, 0x74, 0xed, 0xde, 0xf2,
	0xe9, 0xaf, 0xcd, 0x46, 0xb4, 0x22, 0xaa, 0x02, 0xde, 0x82, 0x6b, 0x03, 0xa6, 0xcd, 0x61, 0xae,
	0x8e, 0x99, 0xe1, 0xc7, 0x9d, 0x25, 0x0f, 0x05, 0xcd, 0xa8, 0x3d, 0xa9, 0xbd, 0xb1, 0x25, 0xdc,
	0x81, 0xab, 0x05, 0xcf, 0x74, 0x22, 0xd3, 0x4e, 0xd3, 0x43, 0xc1, 0x72, 0x54, 0xfd, 0xfa, 0x37,
	0xe1, 0xfa, 0xab, 0x32, 0x41, 0xc5, 0xfa, 0x02, 0x6e, 0x54, 0x05, 0x07, 0xf9, 0x04, 0x5a, 0x36,
	0xa4, 0x03, 0xdc, 0xa8, 0x05, 0xb4, 0x8b, 0x1c, 0x9d, 0x5b, 0xe0, 0xdf, 0x86, 0xf5, 0x7d, 0x6e,
	0x5e, 0xfe, 0x23, 0xa9, 0x5c, 0x9e, 0xc2, 0xad, 0xf3, 0x0d, 0xe7, 0x76, 0x3e, 0x0d, 0xba, 0x90,
	0x66, 0xf7, 0x7b, 0x13, 0xae, 0xbc, 0x9e, 0x5c, 0x01, 0xfc, 0x09, 0xc1, 0xca, 0x74, 0x67, 0x70,
	0x50, 0x0b, 0x56, 0x73, 0x20, 0xdd, 0xed, 0x4b, 0x28, 0x2d, 0x8f, 0x7f, 0xef, 0xe3, 0x8f, 0x3f,
	0x5f, 0x97, 0x3c, 0x4c, 0xe8, 0xfc, 0xfb, 0x21, 0x98, 0xc2, 0x5f, 0x10, 0xb4, 0x67, 0xf2, 0xe0,
	0xfb, 0xf3, 0x2c, 0x2e, 0xee, 0x46, 0xf7, 0xc1, 0xa5, 0xb4, 0x0e, 0x68, 0xbb, 0x04, 0xea, 0xe1,
	0xad, 0x5a, 0xa0, 0xd9, 0xbd, 0xc3, 0x05, 0xb4, 0xec, 0xb1, 0x60, 0x7f, 0xc1, 0x99, 0x55, 0x14,
	0xbd, 0x85, 0x1a, 0xe7, 0xde, 0x2b, 0xdd, 0xef, 0xe2, 0x0d, 0x3a, 0xff, 0x31, 0xec, 0x3d, 0x3f,
	0x1d, 0x11, 0x74, 0x36, 0x22, 0xe8, 0xf7, 0x88, 0xa0, 0xcf, 0x63, 0xd2, 0x38, 0x1b, 0x93, 0xc6,
	0xcf, 0x31, 0x69, 0xbc, 0x7d, 0x18, 0x27, 0xe6, 0x7d, 0x7e, 0x14, 0xbe, 0x93, 0x82, 0xea, 0x7e,
	0xa2, 0x1e, 0x09, 0x5e, 0x54, 0x93, 0x3e, 0xcc, 0xcc, 0x32, 0x43, 0xc5, 0xf5, 0x51, 0xab, 0x7c,
	0x55, 0x8f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x86, 0xe4, 0x0c, 0x8b, 0x05, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// MarketMap returns the full market map stored in the x/marketmap
	// module.
	MarketMap(ctx context.Context, in *GetMarketMapRequest, opts ...grpc.CallOption) (*GetMarketMapResponse, error)
	// LastUpdated returns the last height the market map was updated at.
	LastUpdated(ctx context.Context, in *GetLastUpdatedRequest, opts ...grpc.CallOption) (*GetLastUpdatedResponse, error)
	// Params returns the current x/marketmap module parameters.
	Params(ctx context.Context, in *ParamsRequest, opts ...grpc.CallOption) (*ParamsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) MarketMap(ctx context.Context, in *GetMarketMapRequest, opts ...grpc.CallOption) (*GetMarketMapResponse, error) {
	out := new(GetMarketMapResponse)
	err := c.cc.Invoke(ctx, "/slinky.marketmap.v1.Query/MarketMap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) LastUpdated(ctx context.Context, in *GetLastUpdatedRequest, opts ...grpc.CallOption) (*GetLastUpdatedResponse, error) {
	out := new(GetLastUpdatedResponse)
	err := c.cc.Invoke(ctx, "/slinky.marketmap.v1.Query/LastUpdated", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *ParamsRequest, opts ...grpc.CallOption) (*ParamsResponse, error) {
	out := new(ParamsResponse)
	err := c.cc.Invoke(ctx, "/slinky.marketmap.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// MarketMap returns the full market map stored in the x/marketmap
	// module.
	MarketMap(context.Context, *GetMarketMapRequest) (*GetMarketMapResponse, error)
	// LastUpdated returns the last height the market map was updated at.
	LastUpdated(context.Context, *GetLastUpdatedRequest) (*GetLastUpdatedResponse, error)
	// Params returns the current x/marketmap module parameters.
	Params(context.Context, *ParamsRequest) (*ParamsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) MarketMap(ctx context.Context, req *GetMarketMapRequest) (*GetMarketMapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarketMap not implemented")
}
func (*UnimplementedQueryServer) LastUpdated(ctx context.Context, req *GetLastUpdatedRequest) (*GetLastUpdatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LastUpdated not implemented")
}
func (*UnimplementedQueryServer) Params(ctx context.Context, req *ParamsRequest) (*ParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_MarketMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMarketMapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).MarketMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slinky.marketmap.v1.Query/MarketMap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).MarketMap(ctx, req.(*GetMarketMapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_LastUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLastUpdatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).LastUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slinky.marketmap.v1.Query/LastUpdated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).LastUpdated(ctx, req.(*GetLastUpdatedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slinky.marketmap.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*ParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "slinky.marketmap.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MarketMap",
			Handler:    _Query_MarketMap_Handler,
		},
		{
			MethodName: "LastUpdated",
			Handler:    _Query_LastUpdated_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "slinky/marketmap/v1/query.proto",
}

func (m *GetMarketMapRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetMarketMapRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetMarketMapRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetMarketMapResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetMarketMapResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetMarketMapResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Version != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x18
	}
	if m.LastUpdated != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.LastUpdated))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.MarketMap.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *GetLastUpdatedRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetLastUpdatedRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetLastUpdatedRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetLastUpdatedResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetLastUpdatedResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetLastUpdatedResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastUpdated != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.LastUpdated))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetMarketMapRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetMarketMapResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MarketMap.Size()
	n += 1 + l + sovQuery(uint64(l))
	if m.LastUpdated != 0 {
		n += 1 + sovQuery(uint64(m.LastUpdated))
	}
	if m.Version != 0 {
		n += 1 + sovQuery(uint64(m.Version))
	}
	return n
}

func (m *ParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *GetLastUpdatedRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetLastUpdatedResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LastUpdated != 0 {
		n += 1 + sovQuery(uint64(m.LastUpdated))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetMarketMapRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetMarketMapRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetMarketMapRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetMarketMapResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetMarketMapResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetMarketMapResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketMap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MarketMap.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUpdated", wireType)
			}
			m.LastUpdated = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastUpdated |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetLastUpdatedRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetLastUpdatedRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetLastUpdatedRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetLastUpdatedResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetLastUpdatedResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetLastUpdatedResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUpdated", wireType)
			}
			m.LastUpdated = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastUpdated |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)