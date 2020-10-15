// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipstate.proto

package ipstate

import (
	context "context"
	fmt "fmt"
	ipprovider "github.com/cisco-app-networking/nsm-nse/api/ipam/ipprovider"
	ipreceiver "github.com/cisco-app-networking/nsm-nse/api/ipam/ipreceiver"
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

type SubnetsState struct {
	Subnet               []*ipprovider.Subnet `protobuf:"bytes,1,rep,name=subnet,proto3" json:"subnet,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SubnetsState) Reset()         { *m = SubnetsState{} }
func (m *SubnetsState) String() string { return proto.CompactTextString(m) }
func (*SubnetsState) ProtoMessage()    {}
func (*SubnetsState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9c757518148bec6, []int{0}
}

func (m *SubnetsState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubnetsState.Unmarshal(m, b)
}
func (m *SubnetsState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubnetsState.Marshal(b, m, deterministic)
}
func (m *SubnetsState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubnetsState.Merge(m, src)
}
func (m *SubnetsState) XXX_Size() int {
	return xxx_messageInfo_SubnetsState.Size(m)
}
func (m *SubnetsState) XXX_DiscardUnknown() {
	xxx_messageInfo_SubnetsState.DiscardUnknown(m)
}

var xxx_messageInfo_SubnetsState proto.InternalMessageInfo

func (m *SubnetsState) GetSubnet() []*ipprovider.Subnet {
	if m != nil {
		return m.Subnet
	}
	return nil
}

type IpRangesState struct {
	Range                []*ipreceiver.IpRange `protobuf:"bytes,1,rep,name=range,proto3" json:"range,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *IpRangesState) Reset()         { *m = IpRangesState{} }
func (m *IpRangesState) String() string { return proto.CompactTextString(m) }
func (*IpRangesState) ProtoMessage()    {}
func (*IpRangesState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9c757518148bec6, []int{1}
}

func (m *IpRangesState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpRangesState.Unmarshal(m, b)
}
func (m *IpRangesState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpRangesState.Marshal(b, m, deterministic)
}
func (m *IpRangesState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpRangesState.Merge(m, src)
}
func (m *IpRangesState) XXX_Size() int {
	return xxx_messageInfo_IpRangesState.Size(m)
}
func (m *IpRangesState) XXX_DiscardUnknown() {
	xxx_messageInfo_IpRangesState.DiscardUnknown(m)
}

var xxx_messageInfo_IpRangesState proto.InternalMessageInfo

func (m *IpRangesState) GetRange() []*ipreceiver.IpRange {
	if m != nil {
		return m.Range
	}
	return nil
}

type StatusIdentifier struct {
	Fqdn                 string   `protobuf:"bytes,1,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	ConnectivityDomain   string   `protobuf:"bytes,2,opt,name=connectivity_domain,json=connectivityDomain,proto3" json:"connectivity_domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusIdentifier) Reset()         { *m = StatusIdentifier{} }
func (m *StatusIdentifier) String() string { return proto.CompactTextString(m) }
func (*StatusIdentifier) ProtoMessage()    {}
func (*StatusIdentifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9c757518148bec6, []int{2}
}

func (m *StatusIdentifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusIdentifier.Unmarshal(m, b)
}
func (m *StatusIdentifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusIdentifier.Marshal(b, m, deterministic)
}
func (m *StatusIdentifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusIdentifier.Merge(m, src)
}
func (m *StatusIdentifier) XXX_Size() int {
	return xxx_messageInfo_StatusIdentifier.Size(m)
}
func (m *StatusIdentifier) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusIdentifier.DiscardUnknown(m)
}

var xxx_messageInfo_StatusIdentifier proto.InternalMessageInfo

func (m *StatusIdentifier) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *StatusIdentifier) GetConnectivityDomain() string {
	if m != nil {
		return m.ConnectivityDomain
	}
	return ""
}

func init() {
	proto.RegisterType((*SubnetsState)(nil), "ippool.SubnetsState")
	proto.RegisterType((*IpRangesState)(nil), "ippool.IpRangesState")
	proto.RegisterType((*StatusIdentifier)(nil), "ippool.StatusIdentifier")
}

func init() { proto.RegisterFile("ipstate.proto", fileDescriptor_b9c757518148bec6) }

var fileDescriptor_b9c757518148bec6 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x41, 0x4b, 0x3b, 0x31,
	0x10, 0xc5, 0xff, 0xfb, 0x57, 0x2b, 0x8d, 0xd6, 0x4a, 0x6a, 0x61, 0xe9, 0xa9, 0x2c, 0x28, 0xbd,
	0x74, 0x17, 0x2a, 0xf4, 0xae, 0x14, 0xca, 0x82, 0xa7, 0xed, 0x41, 0xf0, 0x22, 0xe9, 0xee, 0x74,
	0x1d, 0xec, 0x26, 0x31, 0x99, 0x56, 0xfc, 0x42, 0x7e, 0x4e, 0xd9, 0x64, 0x2b, 0xb5, 0x07, 0x0f,
	0xde, 0x66, 0x78, 0x6f, 0x7e, 0x99, 0x97, 0x84, 0x75, 0x50, 0x5b, 0x12, 0x04, 0xb1, 0x36, 0x8a,
	0x14, 0x6f, 0xa1, 0xd6, 0x4a, 0xad, 0x07, 0x0f, 0x25, 0xd2, 0xcb, 0x66, 0x19, 0xe7, 0xaa, 0x4a,
	0x72, 0xb4, 0xb9, 0x1a, 0x0b, 0xad, 0xc7, 0x12, 0xe8, 0x5d, 0x99, 0x57, 0x94, 0x65, 0x22, 0x6d,
	0x35, 0x96, 0x16, 0x12, 0xa1, 0x31, 0x41, 0x2d, 0xaa, 0x04, 0xb5, 0x36, 0x6a, 0x8b, 0x05, 0x98,
	0xbd, 0xd2, 0x53, 0xff, 0x44, 0x33, 0x90, 0x03, 0x6e, 0x1d, 0x6d, 0x57, 0x7a, 0x5a, 0x34, 0x65,
	0xe7, 0x8b, 0xcd, 0x52, 0x02, 0xd9, 0x45, 0xbd, 0x39, 0xbf, 0x61, 0x2d, 0xeb, 0xfa, 0x30, 0x18,
	0x1e, 0x8d, 0xce, 0x26, 0x17, 0xb1, 0x0f, 0x11, 0x7b, 0x57, 0xd6, 0xa8, 0xd1, 0x94, 0x75, 0x52,
	0x9d, 0x09, 0x59, 0x42, 0x33, 0x78, 0xcd, 0x4e, 0x4c, 0xdd, 0x36, 0x73, 0xdd, 0xdd, 0x5c, 0xe3,
	0xca, 0xbc, 0x1a, 0x3d, 0xb2, 0xcb, 0xda, 0xbf, 0xb1, 0x69, 0x01, 0x92, 0x70, 0x85, 0x60, 0x38,
	0x67, 0xc7, 0xab, 0xb7, 0x42, 0x86, 0xc1, 0x30, 0x18, 0xb5, 0x33, 0x57, 0xf3, 0x84, 0xf5, 0x72,
	0x25, 0x25, 0xe4, 0x84, 0x5b, 0xa4, 0x8f, 0xe7, 0x42, 0x55, 0x02, 0x65, 0xf8, 0xdf, 0x59, 0xf8,
	0xbe, 0x34, 0x73, 0xca, 0xe4, 0x33, 0x60, 0xfd, 0x54, 0x8b, 0xca, 0x6d, 0x93, 0x4a, 0x02, 0x63,
	0x54, 0x29, 0x48, 0x19, 0x3e, 0x67, 0xbd, 0x39, 0xd0, 0xdd, 0x7a, 0xad, 0x72, 0x41, 0x50, 0x34,
	0x71, 0x79, 0xf8, 0x9d, 0xec, 0x60, 0x9f, 0xc1, 0xd5, 0xcf, 0xcc, 0x3e, 0x60, 0xf4, 0x8f, 0xcf,
	0x58, 0xb7, 0x06, 0xd5, 0xe7, 0x82, 0x8f, 0xfe, 0x0b, 0xa4, 0x7f, 0x70, 0x01, 0x3b, 0xca, 0x7d,
	0xfb, 0xe9, 0xb4, 0xf9, 0x26, 0xcb, 0x96, 0x7b, 0x83, 0xdb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x7b, 0xd5, 0xb2, 0xb2, 0x38, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IpamStateInterrogatorClient is the client API for IpamStateInterrogator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IpamStateInterrogatorClient interface {
	GetAllocatedSubnets(ctx context.Context, in *StatusIdentifier, opts ...grpc.CallOption) (*SubnetsState, error)
	GetActiveRanges(ctx context.Context, in *StatusIdentifier, opts ...grpc.CallOption) (*IpRangesState, error)
}

type ipamStateInterrogatorClient struct {
	cc *grpc.ClientConn
}

func NewIpamStateInterrogatorClient(cc *grpc.ClientConn) IpamStateInterrogatorClient {
	return &ipamStateInterrogatorClient{cc}
}

func (c *ipamStateInterrogatorClient) GetAllocatedSubnets(ctx context.Context, in *StatusIdentifier, opts ...grpc.CallOption) (*SubnetsState, error) {
	out := new(SubnetsState)
	err := c.cc.Invoke(ctx, "/ippool.IpamStateInterrogator/GetAllocatedSubnets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ipamStateInterrogatorClient) GetActiveRanges(ctx context.Context, in *StatusIdentifier, opts ...grpc.CallOption) (*IpRangesState, error) {
	out := new(IpRangesState)
	err := c.cc.Invoke(ctx, "/ippool.IpamStateInterrogator/GetActiveRanges", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IpamStateInterrogatorServer is the server API for IpamStateInterrogator service.
type IpamStateInterrogatorServer interface {
	GetAllocatedSubnets(context.Context, *StatusIdentifier) (*SubnetsState, error)
	GetActiveRanges(context.Context, *StatusIdentifier) (*IpRangesState, error)
}

// UnimplementedIpamStateInterrogatorServer can be embedded to have forward compatible implementations.
type UnimplementedIpamStateInterrogatorServer struct {
}

func (*UnimplementedIpamStateInterrogatorServer) GetAllocatedSubnets(ctx context.Context, req *StatusIdentifier) (*SubnetsState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllocatedSubnets not implemented")
}
func (*UnimplementedIpamStateInterrogatorServer) GetActiveRanges(ctx context.Context, req *StatusIdentifier) (*IpRangesState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveRanges not implemented")
}

func RegisterIpamStateInterrogatorServer(s *grpc.Server, srv IpamStateInterrogatorServer) {
	s.RegisterService(&_IpamStateInterrogator_serviceDesc, srv)
}

func _IpamStateInterrogator_GetAllocatedSubnets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IpamStateInterrogatorServer).GetAllocatedSubnets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ippool.IpamStateInterrogator/GetAllocatedSubnets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IpamStateInterrogatorServer).GetAllocatedSubnets(ctx, req.(*StatusIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _IpamStateInterrogator_GetActiveRanges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IpamStateInterrogatorServer).GetActiveRanges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ippool.IpamStateInterrogator/GetActiveRanges",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IpamStateInterrogatorServer).GetActiveRanges(ctx, req.(*StatusIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

var _IpamStateInterrogator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ippool.IpamStateInterrogator",
	HandlerType: (*IpamStateInterrogatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllocatedSubnets",
			Handler:    _IpamStateInterrogator_GetAllocatedSubnets_Handler,
		},
		{
			MethodName: "GetActiveRanges",
			Handler:    _IpamStateInterrogator_GetActiveRanges_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ipstate.proto",
}
