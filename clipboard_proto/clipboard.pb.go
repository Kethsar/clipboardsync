// Code generated by protoc-gen-go. DO NOT EDIT.
// source: clipboard.proto

package clipboardsync

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Clipboard struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Clipboard) Reset()         { *m = Clipboard{} }
func (m *Clipboard) String() string { return proto.CompactTextString(m) }
func (*Clipboard) ProtoMessage()    {}
func (*Clipboard) Descriptor() ([]byte, []int) {
	return fileDescriptor_72275e738ef73aac, []int{0}
}

func (m *Clipboard) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Clipboard.Unmarshal(m, b)
}
func (m *Clipboard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Clipboard.Marshal(b, m, deterministic)
}
func (m *Clipboard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Clipboard.Merge(m, src)
}
func (m *Clipboard) XXX_Size() int {
	return xxx_messageInfo_Clipboard.Size(m)
}
func (m *Clipboard) XXX_DiscardUnknown() {
	xxx_messageInfo_Clipboard.DiscardUnknown(m)
}

var xxx_messageInfo_Clipboard proto.InternalMessageInfo

func (m *Clipboard) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*Clipboard)(nil), "clipboardsync.Clipboard")
}

func init() { proto.RegisterFile("clipboard.proto", fileDescriptor_72275e738ef73aac) }

var fileDescriptor_72275e738ef73aac = []byte{
	// 110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xce, 0xc9, 0x2c,
	0x48, 0xca, 0x4f, 0x2c, 0x4a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x85, 0x0b, 0x14,
	0x57, 0xe6, 0x25, 0x2b, 0xc9, 0x73, 0x71, 0x3a, 0xc3, 0x04, 0x84, 0x84, 0xb8, 0x58, 0x52, 0x12,
	0x4b, 0x12, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0xa3, 0x40, 0x2e, 0x5e, 0xb8,
	0x82, 0xe0, 0xca, 0xbc, 0x64, 0x21, 0x07, 0x2e, 0x16, 0x30, 0x2d, 0xa1, 0x87, 0x62, 0x92, 0x1e,
	0x5c, 0x95, 0x14, 0x4e, 0x19, 0x25, 0x06, 0x0d, 0x46, 0x03, 0xc6, 0x24, 0x36, 0xb0, 0x4b, 0x8c,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x0b, 0x12, 0xc6, 0x9c, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClipboardSyncClient is the client API for ClipboardSync service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClipboardSyncClient interface {
	Sync(ctx context.Context, opts ...grpc.CallOption) (ClipboardSync_SyncClient, error)
}

type clipboardSyncClient struct {
	cc *grpc.ClientConn
}

func NewClipboardSyncClient(cc *grpc.ClientConn) ClipboardSyncClient {
	return &clipboardSyncClient{cc}
}

func (c *clipboardSyncClient) Sync(ctx context.Context, opts ...grpc.CallOption) (ClipboardSync_SyncClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ClipboardSync_serviceDesc.Streams[0], "/clipboardsync.ClipboardSync/Sync", opts...)
	if err != nil {
		return nil, err
	}
	x := &clipboardSyncSyncClient{stream}
	return x, nil
}

type ClipboardSync_SyncClient interface {
	Send(*Clipboard) error
	Recv() (*Clipboard, error)
	grpc.ClientStream
}

type clipboardSyncSyncClient struct {
	grpc.ClientStream
}

func (x *clipboardSyncSyncClient) Send(m *Clipboard) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clipboardSyncSyncClient) Recv() (*Clipboard, error) {
	m := new(Clipboard)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClipboardSyncServer is the server API for ClipboardSync service.
type ClipboardSyncServer interface {
	Sync(ClipboardSync_SyncServer) error
}

func RegisterClipboardSyncServer(s *grpc.Server, srv ClipboardSyncServer) {
	s.RegisterService(&_ClipboardSync_serviceDesc, srv)
}

func _ClipboardSync_Sync_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClipboardSyncServer).Sync(&clipboardSyncSyncServer{stream})
}

type ClipboardSync_SyncServer interface {
	Send(*Clipboard) error
	Recv() (*Clipboard, error)
	grpc.ServerStream
}

type clipboardSyncSyncServer struct {
	grpc.ServerStream
}

func (x *clipboardSyncSyncServer) Send(m *Clipboard) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clipboardSyncSyncServer) Recv() (*Clipboard, error) {
	m := new(Clipboard)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ClipboardSync_serviceDesc = grpc.ServiceDesc{
	ServiceName: "clipboardsync.ClipboardSync",
	HandlerType: (*ClipboardSyncServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Sync",
			Handler:       _ClipboardSync_Sync_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "clipboard.proto",
}
