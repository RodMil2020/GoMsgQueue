// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message_queue.proto

package messagequeue

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

type Response_Status int32

const (
	Response_SUCCESS Response_Status = 0
	Response_ERROR   Response_Status = 1
)

var Response_Status_name = map[int32]string{
	0: "SUCCESS",
	1: "ERROR",
}

var Response_Status_value = map[string]int32{
	"SUCCESS": 0,
	"ERROR":   1,
}

func (x Response_Status) String() string {
	return proto.EnumName(Response_Status_name, int32(x))
}

func (Response_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_47f0b7e138abaa83, []int{2, 0}
}

// QueueName, is the name of the queue.
type QueueName struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueueName) Reset()         { *m = QueueName{} }
func (m *QueueName) String() string { return proto.CompactTextString(m) }
func (*QueueName) ProtoMessage()    {}
func (*QueueName) Descriptor() ([]byte, []int) {
	return fileDescriptor_47f0b7e138abaa83, []int{0}
}

func (m *QueueName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueName.Unmarshal(m, b)
}
func (m *QueueName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueName.Marshal(b, m, deterministic)
}
func (m *QueueName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueName.Merge(m, src)
}
func (m *QueueName) XXX_Size() int {
	return xxx_messageInfo_QueueName.Size(m)
}
func (m *QueueName) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueName.DiscardUnknown(m)
}

var xxx_messageInfo_QueueName proto.InternalMessageInfo

func (m *QueueName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// QueueMessage is the message you want to pass to the other server / app.
type QueueMessage struct {
	Queue                *QueueName `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty"`
	MessageJson          string     `protobuf:"bytes,2,opt,name=message_json,json=messageJson,proto3" json:"message_json,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *QueueMessage) Reset()         { *m = QueueMessage{} }
func (m *QueueMessage) String() string { return proto.CompactTextString(m) }
func (*QueueMessage) ProtoMessage()    {}
func (*QueueMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_47f0b7e138abaa83, []int{1}
}

func (m *QueueMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueMessage.Unmarshal(m, b)
}
func (m *QueueMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueMessage.Marshal(b, m, deterministic)
}
func (m *QueueMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueMessage.Merge(m, src)
}
func (m *QueueMessage) XXX_Size() int {
	return xxx_messageInfo_QueueMessage.Size(m)
}
func (m *QueueMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueMessage.DiscardUnknown(m)
}

var xxx_messageInfo_QueueMessage proto.InternalMessageInfo

func (m *QueueMessage) GetQueue() *QueueName {
	if m != nil {
		return m.Queue
	}
	return nil
}

func (m *QueueMessage) GetMessageJson() string {
	if m != nil {
		return m.MessageJson
	}
	return ""
}

// this is to hold the different status of a Queue message
type Response struct {
	Status               Response_Status `protobuf:"varint,1,opt,name=status,proto3,enum=messagequeue.Response_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_47f0b7e138abaa83, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() Response_Status {
	if m != nil {
		return m.Status
	}
	return Response_SUCCESS
}

func init() {
	proto.RegisterEnum("messagequeue.Response_Status", Response_Status_name, Response_Status_value)
	proto.RegisterType((*QueueName)(nil), "messagequeue.QueueName")
	proto.RegisterType((*QueueMessage)(nil), "messagequeue.QueueMessage")
	proto.RegisterType((*Response)(nil), "messagequeue.Response")
}

func init() { proto.RegisterFile("message_queue.proto", fileDescriptor_47f0b7e138abaa83) }

var fileDescriptor_47f0b7e138abaa83 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4e, 0x83, 0x40,
	0x10, 0x86, 0xc5, 0x58, 0x94, 0x01, 0x4d, 0x1d, 0x13, 0x6d, 0x48, 0x8c, 0x95, 0x93, 0x17, 0xf7,
	0x80, 0xf1, 0xe0, 0xd1, 0x92, 0x5e, 0x4c, 0xd4, 0xba, 0xc4, 0x73, 0x45, 0x32, 0x21, 0x9a, 0xc0,
	0x22, 0xbb, 0xbc, 0x8a, 0xcf, 0x6b, 0x3a, 0x2c, 0x0d, 0x4d, 0x6a, 0x6f, 0x93, 0x9f, 0x6f, 0x3e,
	0x66, 0x76, 0xe0, 0xac, 0x24, 0xad, 0xb3, 0x82, 0x96, 0x3f, 0x2d, 0xb5, 0x24, 0xea, 0x46, 0x19,
	0x85, 0x81, 0x0d, 0x39, 0x8b, 0xae, 0xc0, 0x7b, 0x5b, 0x15, 0x2f, 0x59, 0x49, 0x88, 0x70, 0x50,
	0x65, 0x25, 0x4d, 0x9c, 0xa9, 0x73, 0xe3, 0x49, 0xae, 0xa3, 0x0f, 0x08, 0x18, 0x78, 0xee, 0xba,
	0xf0, 0x16, 0x46, 0xdc, 0xc9, 0x90, 0x1f, 0x5f, 0x88, 0xa1, 0x4e, 0xac, 0x5d, 0xb2, 0xa3, 0xf0,
	0x1a, 0xfa, 0xff, 0x2d, 0xbf, 0xb5, 0xaa, 0x26, 0xfb, 0xac, 0xf6, 0x6d, 0xf6, 0xa4, 0x55, 0x15,
	0xe5, 0x70, 0x24, 0x49, 0xd7, 0xaa, 0xd2, 0x84, 0xf7, 0xe0, 0x6a, 0x93, 0x99, 0x56, 0xb3, 0xfe,
	0x24, 0xbe, 0xdc, 0xd4, 0xf7, 0x9c, 0x48, 0x19, 0x92, 0x16, 0x8e, 0xa6, 0xe0, 0x76, 0x09, 0xfa,
	0x70, 0x98, 0xbe, 0x27, 0xc9, 0x3c, 0x4d, 0xc7, 0x7b, 0xe8, 0xc1, 0x68, 0x2e, 0xe5, 0xab, 0x1c,
	0x3b, 0xf1, 0xaf, 0x03, 0x81, 0x5d, 0x81, 0x67, 0xc4, 0x04, 0x8e, 0xf3, 0x86, 0x32, 0xb3, 0x5e,
	0x2c, 0xdc, 0xb2, 0x89, 0xfd, 0x16, 0x9e, 0x6f, 0x1f, 0x03, 0x1f, 0x01, 0x0a, 0x32, 0xbd, 0xe1,
	0xbf, 0xb7, 0x08, 0x77, 0xa8, 0x67, 0x0f, 0x10, 0x7d, 0x29, 0x51, 0x34, 0x75, 0x2e, 0x36, 0xaf,
	0x35, 0x6c, 0x99, 0x9d, 0x0e, 0x67, 0x5f, 0xac, 0xee, 0xb8, 0x70, 0x3e, 0x5d, 0x3e, 0xe8, 0xdd,
	0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x77, 0x67, 0xd6, 0xef, 0xe7, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MessageQueueClient is the client API for MessageQueue service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageQueueClient interface {
	// creates a new message inside the given queue.
	// Returns a status message signifying if the operation was successful.
	CreateMessage(ctx context.Context, in *QueueMessage, opts ...grpc.CallOption) (*Response, error)
	// gets the next message available in the given queue name.
	// Returns QueueData or nil if there exist no more messages in the queue.
	GetMessage(ctx context.Context, in *QueueName, opts ...grpc.CallOption) (*QueueMessage, error)
}

type messageQueueClient struct {
	cc *grpc.ClientConn
}

func NewMessageQueueClient(cc *grpc.ClientConn) MessageQueueClient {
	return &messageQueueClient{cc}
}

func (c *messageQueueClient) CreateMessage(ctx context.Context, in *QueueMessage, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/messagequeue.MessageQueue/createMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageQueueClient) GetMessage(ctx context.Context, in *QueueName, opts ...grpc.CallOption) (*QueueMessage, error) {
	out := new(QueueMessage)
	err := c.cc.Invoke(ctx, "/messagequeue.MessageQueue/getMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageQueueServer is the server API for MessageQueue service.
type MessageQueueServer interface {
	// creates a new message inside the given queue.
	// Returns a status message signifying if the operation was successful.
	CreateMessage(context.Context, *QueueMessage) (*Response, error)
	// gets the next message available in the given queue name.
	// Returns QueueData or nil if there exist no more messages in the queue.
	GetMessage(context.Context, *QueueName) (*QueueMessage, error)
}

// UnimplementedMessageQueueServer can be embedded to have forward compatible implementations.
type UnimplementedMessageQueueServer struct {
}

func (*UnimplementedMessageQueueServer) CreateMessage(ctx context.Context, req *QueueMessage) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
}
func (*UnimplementedMessageQueueServer) GetMessage(ctx context.Context, req *QueueName) (*QueueMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}

func RegisterMessageQueueServer(s *grpc.Server, srv MessageQueueServer) {
	s.RegisterService(&_MessageQueue_serviceDesc, srv)
}

func _MessageQueue_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueueMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageQueueServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messagequeue.MessageQueue/CreateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageQueueServer).CreateMessage(ctx, req.(*QueueMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageQueue_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueueName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageQueueServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messagequeue.MessageQueue/GetMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageQueueServer).GetMessage(ctx, req.(*QueueName))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessageQueue_serviceDesc = grpc.ServiceDesc{
	ServiceName: "messagequeue.MessageQueue",
	HandlerType: (*MessageQueueServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createMessage",
			Handler:    _MessageQueue_CreateMessage_Handler,
		},
		{
			MethodName: "getMessage",
			Handler:    _MessageQueue_GetMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message_queue.proto",
}
