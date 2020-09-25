// Copyright (c) 2020 TypeFox GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ws-manager-node.proto

package api

import (
	context "context"
	fmt "fmt"
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

type UidmapCanaryResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	ErrorCode            uint32   `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UidmapCanaryResponse) Reset()         { *m = UidmapCanaryResponse{} }
func (m *UidmapCanaryResponse) String() string { return proto.CompactTextString(m) }
func (*UidmapCanaryResponse) ProtoMessage()    {}
func (*UidmapCanaryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_022226a2c4a7acd9, []int{0}
}

func (m *UidmapCanaryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UidmapCanaryResponse.Unmarshal(m, b)
}
func (m *UidmapCanaryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UidmapCanaryResponse.Marshal(b, m, deterministic)
}
func (m *UidmapCanaryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UidmapCanaryResponse.Merge(m, src)
}
func (m *UidmapCanaryResponse) XXX_Size() int {
	return xxx_messageInfo_UidmapCanaryResponse.Size(m)
}
func (m *UidmapCanaryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UidmapCanaryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UidmapCanaryResponse proto.InternalMessageInfo

func (m *UidmapCanaryResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UidmapCanaryResponse) GetErrorCode() uint32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

type UidmapCanaryRequest struct {
	Pid                  int64                          `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Gid                  bool                           `protobuf:"varint,2,opt,name=gid,proto3" json:"gid,omitempty"`
	Mapping              []*UidmapCanaryRequest_Mapping `protobuf:"bytes,3,rep,name=mapping,proto3" json:"mapping,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *UidmapCanaryRequest) Reset()         { *m = UidmapCanaryRequest{} }
func (m *UidmapCanaryRequest) String() string { return proto.CompactTextString(m) }
func (*UidmapCanaryRequest) ProtoMessage()    {}
func (*UidmapCanaryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_022226a2c4a7acd9, []int{1}
}

func (m *UidmapCanaryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UidmapCanaryRequest.Unmarshal(m, b)
}
func (m *UidmapCanaryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UidmapCanaryRequest.Marshal(b, m, deterministic)
}
func (m *UidmapCanaryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UidmapCanaryRequest.Merge(m, src)
}
func (m *UidmapCanaryRequest) XXX_Size() int {
	return xxx_messageInfo_UidmapCanaryRequest.Size(m)
}
func (m *UidmapCanaryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UidmapCanaryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UidmapCanaryRequest proto.InternalMessageInfo

func (m *UidmapCanaryRequest) GetPid() int64 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *UidmapCanaryRequest) GetGid() bool {
	if m != nil {
		return m.Gid
	}
	return false
}

func (m *UidmapCanaryRequest) GetMapping() []*UidmapCanaryRequest_Mapping {
	if m != nil {
		return m.Mapping
	}
	return nil
}

type UidmapCanaryRequest_Mapping struct {
	ContainerId          uint32   `protobuf:"varint,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	HostId               uint32   `protobuf:"varint,2,opt,name=host_id,json=hostId,proto3" json:"host_id,omitempty"`
	Size                 uint32   `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UidmapCanaryRequest_Mapping) Reset()         { *m = UidmapCanaryRequest_Mapping{} }
func (m *UidmapCanaryRequest_Mapping) String() string { return proto.CompactTextString(m) }
func (*UidmapCanaryRequest_Mapping) ProtoMessage()    {}
func (*UidmapCanaryRequest_Mapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_022226a2c4a7acd9, []int{1, 0}
}

func (m *UidmapCanaryRequest_Mapping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UidmapCanaryRequest_Mapping.Unmarshal(m, b)
}
func (m *UidmapCanaryRequest_Mapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UidmapCanaryRequest_Mapping.Marshal(b, m, deterministic)
}
func (m *UidmapCanaryRequest_Mapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UidmapCanaryRequest_Mapping.Merge(m, src)
}
func (m *UidmapCanaryRequest_Mapping) XXX_Size() int {
	return xxx_messageInfo_UidmapCanaryRequest_Mapping.Size(m)
}
func (m *UidmapCanaryRequest_Mapping) XXX_DiscardUnknown() {
	xxx_messageInfo_UidmapCanaryRequest_Mapping.DiscardUnknown(m)
}

var xxx_messageInfo_UidmapCanaryRequest_Mapping proto.InternalMessageInfo

func (m *UidmapCanaryRequest_Mapping) GetContainerId() uint32 {
	if m != nil {
		return m.ContainerId
	}
	return 0
}

func (m *UidmapCanaryRequest_Mapping) GetHostId() uint32 {
	if m != nil {
		return m.HostId
	}
	return 0
}

func (m *UidmapCanaryRequest_Mapping) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func init() {
	proto.RegisterType((*UidmapCanaryResponse)(nil), "wsmannode.UidmapCanaryResponse")
	proto.RegisterType((*UidmapCanaryRequest)(nil), "wsmannode.UidmapCanaryRequest")
	proto.RegisterType((*UidmapCanaryRequest_Mapping)(nil), "wsmannode.UidmapCanaryRequest.Mapping")
}

func init() {
	proto.RegisterFile("ws-manager-node.proto", fileDescriptor_022226a2c4a7acd9)
}

var fileDescriptor_022226a2c4a7acd9 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x86, 0xad, 0x95, 0xd5, 0x9d, 0x6d, 0xa0, 0x51, 0xb1, 0x0c, 0xd4, 0xb9, 0x0b, 0xe9, 0xcd,
	0x3a, 0x9d, 0x2f, 0x20, 0xee, 0xc6, 0x5d, 0x88, 0x10, 0x18, 0xa2, 0x37, 0x23, 0x6b, 0x0e, 0x5d,
	0xd4, 0x26, 0x31, 0xe9, 0x18, 0xfa, 0xa0, 0x3e, 0x8f, 0x24, 0xeb, 0x86, 0x8a, 0x7a, 0xf7, 0x9f,
	0x8f, 0xf6, 0xcb, 0x9f, 0x1c, 0x38, 0x58, 0xd8, 0x5e, 0xc1, 0x24, 0xcb, 0xd1, 0xf4, 0xa4, 0xe2,
	0x98, 0x6a, 0xa3, 0x4a, 0x45, 0xea, 0x0b, 0x5b, 0x30, 0xe9, 0x40, 0xf7, 0x0e, 0xf6, 0xc7, 0x82,
	0x17, 0x4c, 0x0f, 0x99, 0x64, 0xe6, 0x8d, 0xa2, 0xd5, 0x4a, 0x5a, 0x24, 0x31, 0x44, 0x05, 0x5a,
	0xcb, 0x72, 0x8c, 0x83, 0x4e, 0x90, 0xd4, 0xe9, 0x6a, 0x24, 0x47, 0x00, 0x68, 0x8c, 0x32, 0x93,
	0x4c, 0x71, 0x8c, 0x37, 0x3b, 0x41, 0xd2, 0xa2, 0x75, 0x4f, 0x86, 0x4e, 0xf8, 0x11, 0xc0, 0xde,
	0x77, 0xe3, 0xeb, 0x1c, 0x6d, 0x49, 0x76, 0x20, 0xd4, 0x82, 0x7b, 0x59, 0x48, 0x5d, 0x74, 0x24,
	0x17, 0xdc, 0x1b, 0xb6, 0xa9, 0x8b, 0xe4, 0x0a, 0xa2, 0x82, 0x69, 0x2d, 0x64, 0x1e, 0x87, 0x9d,
	0x30, 0x69, 0x0c, 0xce, 0xd2, 0x75, 0xd3, 0xf4, 0x17, 0x69, 0x7a, 0xbb, 0xfc, 0x9a, 0xae, 0x7e,
	0x6b, 0x3f, 0x40, 0x54, 0x31, 0x72, 0x0a, 0xcd, 0x4c, 0xc9, 0x92, 0x09, 0x89, 0x66, 0x52, 0x9d,
	0xdc, 0xa2, 0x8d, 0x35, 0x1b, 0x71, 0x72, 0x08, 0xd1, 0x4c, 0xd9, 0x72, 0x52, 0xb5, 0x68, 0xd1,
	0x9a, 0x1b, 0x47, 0x9c, 0x10, 0xd8, 0xb2, 0xe2, 0x1d, 0xe3, 0xd0, 0x53, 0x9f, 0x07, 0x4f, 0xb0,
	0x3b, 0x92, 0xf7, 0xca, 0x3c, 0x5b, 0xcd, 0x32, 0xbc, 0xc1, 0x17, 0x8d, 0x86, 0x8c, 0xa1, 0xf9,
	0xb5, 0x17, 0x39, 0xf9, 0xb3, 0xf0, 0xf2, 0x5d, 0xdb, 0xc7, 0xff, 0xdf, 0xa8, 0xbb, 0x91, 0x04,
	0xe7, 0xc1, 0xf5, 0xc5, 0x63, 0x3f, 0x17, 0xe5, 0x6c, 0x3e, 0x4d, 0x33, 0x55, 0xb8, 0xa8, 0x15,
	0xef, 0x09, 0x55, 0xa5, 0xfe, 0x8f, 0xad, 0xf6, 0x99, 0x16, 0xd3, 0x9a, 0x5f, 0xed, 0xe5, 0x67,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x46, 0x38, 0x5f, 0xf3, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// InWorkspaceHelperClient is the client API for InWorkspaceHelper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InWorkspaceHelperClient interface {
	// UidmapCanary can establish a uid mapping of a new user namespace spawned within the workspace.
	UidmapCanary(ctx context.Context, opts ...grpc.CallOption) (InWorkspaceHelper_UidmapCanaryClient, error)
}

type inWorkspaceHelperClient struct {
	cc grpc.ClientConnInterface
}

func NewInWorkspaceHelperClient(cc grpc.ClientConnInterface) InWorkspaceHelperClient {
	return &inWorkspaceHelperClient{cc}
}

func (c *inWorkspaceHelperClient) UidmapCanary(ctx context.Context, opts ...grpc.CallOption) (InWorkspaceHelper_UidmapCanaryClient, error) {
	stream, err := c.cc.NewStream(ctx, &_InWorkspaceHelper_serviceDesc.Streams[0], "/wsmannode.InWorkspaceHelper/UidmapCanary", opts...)
	if err != nil {
		return nil, err
	}
	x := &inWorkspaceHelperUidmapCanaryClient{stream}
	return x, nil
}

type InWorkspaceHelper_UidmapCanaryClient interface {
	Send(*UidmapCanaryResponse) error
	Recv() (*UidmapCanaryRequest, error)
	grpc.ClientStream
}

type inWorkspaceHelperUidmapCanaryClient struct {
	grpc.ClientStream
}

func (x *inWorkspaceHelperUidmapCanaryClient) Send(m *UidmapCanaryResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *inWorkspaceHelperUidmapCanaryClient) Recv() (*UidmapCanaryRequest, error) {
	m := new(UidmapCanaryRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// InWorkspaceHelperServer is the server API for InWorkspaceHelper service.
type InWorkspaceHelperServer interface {
	// UidmapCanary can establish a uid mapping of a new user namespace spawned within the workspace.
	UidmapCanary(InWorkspaceHelper_UidmapCanaryServer) error
}

// UnimplementedInWorkspaceHelperServer can be embedded to have forward compatible implementations.
type UnimplementedInWorkspaceHelperServer struct {
}

func (*UnimplementedInWorkspaceHelperServer) UidmapCanary(srv InWorkspaceHelper_UidmapCanaryServer) error {
	return status.Errorf(codes.Unimplemented, "method UidmapCanary not implemented")
}

func RegisterInWorkspaceHelperServer(s *grpc.Server, srv InWorkspaceHelperServer) {
	s.RegisterService(&_InWorkspaceHelper_serviceDesc, srv)
}

func _InWorkspaceHelper_UidmapCanary_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(InWorkspaceHelperServer).UidmapCanary(&inWorkspaceHelperUidmapCanaryServer{stream})
}

type InWorkspaceHelper_UidmapCanaryServer interface {
	Send(*UidmapCanaryRequest) error
	Recv() (*UidmapCanaryResponse, error)
	grpc.ServerStream
}

type inWorkspaceHelperUidmapCanaryServer struct {
	grpc.ServerStream
}

func (x *inWorkspaceHelperUidmapCanaryServer) Send(m *UidmapCanaryRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *inWorkspaceHelperUidmapCanaryServer) Recv() (*UidmapCanaryResponse, error) {
	m := new(UidmapCanaryResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _InWorkspaceHelper_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wsmannode.InWorkspaceHelper",
	HandlerType: (*InWorkspaceHelperServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UidmapCanary",
			Handler:       _InWorkspaceHelper_UidmapCanary_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ws-manager-node.proto",
}
