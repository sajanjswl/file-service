// Code generated by protoc-gen-go. DO NOT EDIT.
// source: file-service.proto

package v1

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

type UploadStatusCode int32

const (
	UploadStatusCode_Unknown UploadStatusCode = 0
	UploadStatusCode_Ok      UploadStatusCode = 1
	UploadStatusCode_Failed  UploadStatusCode = 2
)

var UploadStatusCode_name = map[int32]string{
	0: "Unknown",
	1: "Ok",
	2: "Failed",
}

var UploadStatusCode_value = map[string]int32{
	"Unknown": 0,
	"Ok":      1,
	"Failed":  2,
}

func (x UploadStatusCode) String() string {
	return proto.EnumName(UploadStatusCode_name, int32(x))
}

func (UploadStatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b50cba0e4ffba287, []int{0}
}

type CreateUserRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b50cba0e4ffba287, []int{0}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CreateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CreateUserResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b50cba0e4ffba287, []int{1}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Chunk struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chunk) Reset()         { *m = Chunk{} }
func (m *Chunk) String() string { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()    {}
func (*Chunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_b50cba0e4ffba287, []int{2}
}

func (m *Chunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chunk.Unmarshal(m, b)
}
func (m *Chunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chunk.Marshal(b, m, deterministic)
}
func (m *Chunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chunk.Merge(m, src)
}
func (m *Chunk) XXX_Size() int {
	return xxx_messageInfo_Chunk.Size(m)
}
func (m *Chunk) XXX_DiscardUnknown() {
	xxx_messageInfo_Chunk.DiscardUnknown(m)
}

var xxx_messageInfo_Chunk proto.InternalMessageInfo

func (m *Chunk) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type UploadStatus struct {
	Message              string           `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code                 UploadStatusCode `protobuf:"varint,2,opt,name=code,proto3,enum=v1.UploadStatusCode" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UploadStatus) Reset()         { *m = UploadStatus{} }
func (m *UploadStatus) String() string { return proto.CompactTextString(m) }
func (*UploadStatus) ProtoMessage()    {}
func (*UploadStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_b50cba0e4ffba287, []int{3}
}

func (m *UploadStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadStatus.Unmarshal(m, b)
}
func (m *UploadStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadStatus.Marshal(b, m, deterministic)
}
func (m *UploadStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadStatus.Merge(m, src)
}
func (m *UploadStatus) XXX_Size() int {
	return xxx_messageInfo_UploadStatus.Size(m)
}
func (m *UploadStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadStatus.DiscardUnknown(m)
}

var xxx_messageInfo_UploadStatus proto.InternalMessageInfo

func (m *UploadStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UploadStatus) GetCode() UploadStatusCode {
	if m != nil {
		return m.Code
	}
	return UploadStatusCode_Unknown
}

func init() {
	proto.RegisterEnum("v1.UploadStatusCode", UploadStatusCode_name, UploadStatusCode_value)
	proto.RegisterType((*CreateUserRequest)(nil), "v1.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "v1.CreateUserResponse")
	proto.RegisterType((*Chunk)(nil), "v1.Chunk")
	proto.RegisterType((*UploadStatus)(nil), "v1.UploadStatus")
}

func init() { proto.RegisterFile("file-service.proto", fileDescriptor_b50cba0e4ffba287) }

var fileDescriptor_b50cba0e4ffba287 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x6c, 0x52, 0x68, 0xe9, 0x6b, 0x85, 0xcc, 0x13, 0xa0, 0xa8, 0x13, 0x78, 0xaa, 0x90, 0x1a,
	0xa9, 0xed, 0xcc, 0x14, 0xa9, 0x2b, 0x52, 0xaa, 0x8c, 0x0c, 0xa6, 0x79, 0x40, 0x94, 0xd4, 0x0e,
	0xb6, 0x93, 0x0e, 0xfc, 0x3c, 0x72, 0x42, 0xa0, 0x0a, 0x62, 0xf3, 0xbd, 0x3b, 0xdf, 0x3b, 0x9f,
	0x01, 0x5f, 0xb3, 0x82, 0x96, 0x86, 0x74, 0x9d, 0xed, 0x29, 0x2c, 0xb5, 0xb2, 0x0a, 0xfd, 0x7a,
	0xc5, 0x9f, 0xe1, 0x2a, 0xd2, 0x24, 0x2c, 0x25, 0x86, 0x74, 0x4c, 0x1f, 0x15, 0x19, 0x8b, 0x0c,
	0x86, 0xa2, 0xcc, 0x02, 0xef, 0xce, 0x5b, 0x4c, 0x62, 0x77, 0xc4, 0x39, 0x5c, 0x54, 0x86, 0xb4,
	0x14, 0x07, 0x0a, 0xfc, 0x66, 0xfc, 0x83, 0x1d, 0x57, 0x0a, 0x63, 0x8e, 0x4a, 0xa7, 0xc1, 0xb0,
	0xe5, 0x3a, 0xcc, 0x43, 0xc0, 0x53, 0x7b, 0x53, 0x2a, 0x69, 0x08, 0x03, 0x18, 0x1f, 0xc8, 0x18,
	0xf1, 0x46, 0xdf, 0x3b, 0x3a, 0xc8, 0xef, 0xe1, 0x3c, 0x7a, 0xaf, 0x64, 0xee, 0x24, 0x91, 0x92,
	0x96, 0xa4, 0x6d, 0x24, 0xb3, 0xb8, 0x83, 0x3c, 0x86, 0x59, 0x52, 0x16, 0x4a, 0xa4, 0x3b, 0x2b,
	0x6c, 0x65, 0xfe, 0x37, 0xc3, 0x05, 0x9c, 0xed, 0x55, 0xda, 0x06, 0xbe, 0x5c, 0x5f, 0x87, 0xf5,
	0x2a, 0x3c, 0xbd, 0x19, 0xa9, 0x94, 0xe2, 0x46, 0xf1, 0xb0, 0x01, 0xd6, 0x67, 0x70, 0x0a, 0xe3,
	0x44, 0xe6, 0x52, 0x1d, 0x25, 0x1b, 0xe0, 0x08, 0xfc, 0xa7, 0x9c, 0x79, 0x08, 0x30, 0xda, 0x8a,
	0xac, 0xa0, 0x94, 0xf9, 0xeb, 0x4f, 0x98, 0x6e, 0xb3, 0x82, 0x76, 0x6d, 0xa7, 0xf8, 0x08, 0xf0,
	0xfb, 0x54, 0xbc, 0x71, 0xdb, 0xfe, 0x34, 0x3b, 0xbf, 0xed, 0x8f, 0xdb, 0x46, 0xf8, 0x00, 0x97,
	0x00, 0x6d, 0x04, 0xe7, 0x89, 0x93, 0x46, 0xe7, 0x9a, 0x98, 0xb3, 0x7e, 0x6e, 0x3e, 0x58, 0x78,
	0x2f, 0xa3, 0xe6, 0x0b, 0x37, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x60, 0x35, 0x38, 0x70, 0xd8,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileServiceClient interface {
	//simple RPC
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	//client side straming RPC
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (FileService_UploadFileClient, error)
}

type fileServiceClient struct {
	cc *grpc.ClientConn
}

func NewFileServiceClient(cc *grpc.ClientConn) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/v1.FileService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (FileService_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileService_serviceDesc.Streams[0], "/v1.FileService/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileServiceUploadFileClient{stream}
	return x, nil
}

type FileService_UploadFileClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*UploadStatus, error)
	grpc.ClientStream
}

type fileServiceUploadFileClient struct {
	grpc.ClientStream
}

func (x *fileServiceUploadFileClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileServiceUploadFileClient) CloseAndRecv() (*UploadStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileServiceServer is the server API for FileService service.
type FileServiceServer interface {
	//simple RPC
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	//client side straming RPC
	UploadFile(FileService_UploadFileServer) error
}

// UnimplementedFileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (*UnimplementedFileServiceServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedFileServiceServer) UploadFile(srv FileService_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}

func RegisterFileServiceServer(s *grpc.Server, srv FileServiceServer) {
	s.RegisterService(&_FileService_serviceDesc, srv)
}

func _FileService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.FileService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileServiceServer).UploadFile(&fileServiceUploadFileServer{stream})
}

type FileService_UploadFileServer interface {
	SendAndClose(*UploadStatus) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type fileServiceUploadFileServer struct {
	grpc.ServerStream
}

func (x *fileServiceUploadFileServer) SendAndClose(m *UploadStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileServiceUploadFileServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _FileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _FileService_CreateUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _FileService_UploadFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "file-service.proto",
}
