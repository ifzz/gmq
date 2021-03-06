// Code generated by protoc-gen-go. DO NOT EDIT.
// source: comment.proto

package commentModule

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

type RequestPutComment struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	TargetId             int32    `protobuf:"varint,2,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	ParentId             int32    `protobuf:"varint,3,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	UserId               int32    `protobuf:"varint,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ClientId             string   `protobuf:"bytes,6,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestPutComment) Reset()         { *m = RequestPutComment{} }
func (m *RequestPutComment) String() string { return proto.CompactTextString(m) }
func (*RequestPutComment) ProtoMessage()    {}
func (*RequestPutComment) Descriptor() ([]byte, []int) {
	return fileDescriptor_749aee09ea917828, []int{0}
}

func (m *RequestPutComment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestPutComment.Unmarshal(m, b)
}
func (m *RequestPutComment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestPutComment.Marshal(b, m, deterministic)
}
func (m *RequestPutComment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestPutComment.Merge(m, src)
}
func (m *RequestPutComment) XXX_Size() int {
	return xxx_messageInfo_RequestPutComment.Size(m)
}
func (m *RequestPutComment) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestPutComment.DiscardUnknown(m)
}

var xxx_messageInfo_RequestPutComment proto.InternalMessageInfo

func (m *RequestPutComment) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *RequestPutComment) GetTargetId() int32 {
	if m != nil {
		return m.TargetId
	}
	return 0
}

func (m *RequestPutComment) GetParentId() int32 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

func (m *RequestPutComment) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *RequestPutComment) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *RequestPutComment) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type ResponsePutComment struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponsePutComment) Reset()         { *m = ResponsePutComment{} }
func (m *ResponsePutComment) String() string { return proto.CompactTextString(m) }
func (*ResponsePutComment) ProtoMessage()    {}
func (*ResponsePutComment) Descriptor() ([]byte, []int) {
	return fileDescriptor_749aee09ea917828, []int{1}
}

func (m *ResponsePutComment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponsePutComment.Unmarshal(m, b)
}
func (m *ResponsePutComment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponsePutComment.Marshal(b, m, deterministic)
}
func (m *ResponsePutComment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponsePutComment.Merge(m, src)
}
func (m *ResponsePutComment) XXX_Size() int {
	return xxx_messageInfo_ResponsePutComment.Size(m)
}
func (m *ResponsePutComment) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponsePutComment.DiscardUnknown(m)
}

var xxx_messageInfo_ResponsePutComment proto.InternalMessageInfo

func (m *ResponsePutComment) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ResponsePutComment) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Comment struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	TargetId             int32    `protobuf:"varint,3,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	UserId               int32    `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ParentId             int32    `protobuf:"varint,5,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Content              string   `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt            int32    `protobuf:"varint,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Comment) Reset()         { *m = Comment{} }
func (m *Comment) String() string { return proto.CompactTextString(m) }
func (*Comment) ProtoMessage()    {}
func (*Comment) Descriptor() ([]byte, []int) {
	return fileDescriptor_749aee09ea917828, []int{2}
}

func (m *Comment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Comment.Unmarshal(m, b)
}
func (m *Comment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Comment.Marshal(b, m, deterministic)
}
func (m *Comment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Comment.Merge(m, src)
}
func (m *Comment) XXX_Size() int {
	return xxx_messageInfo_Comment.Size(m)
}
func (m *Comment) XXX_DiscardUnknown() {
	xxx_messageInfo_Comment.DiscardUnknown(m)
}

var xxx_messageInfo_Comment proto.InternalMessageInfo

func (m *Comment) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Comment) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Comment) GetTargetId() int32 {
	if m != nil {
		return m.TargetId
	}
	return 0
}

func (m *Comment) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Comment) GetParentId() int32 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

func (m *Comment) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Comment) GetCreatedAt() int32 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

type RequestGetComments struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	TargetId             int32    `protobuf:"varint,2,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	PageSize             int32    `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	CurrentPage          int32    `protobuf:"varint,4,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty"`
	Total                int32    `protobuf:"varint,5,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestGetComments) Reset()         { *m = RequestGetComments{} }
func (m *RequestGetComments) String() string { return proto.CompactTextString(m) }
func (*RequestGetComments) ProtoMessage()    {}
func (*RequestGetComments) Descriptor() ([]byte, []int) {
	return fileDescriptor_749aee09ea917828, []int{3}
}

func (m *RequestGetComments) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestGetComments.Unmarshal(m, b)
}
func (m *RequestGetComments) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestGetComments.Marshal(b, m, deterministic)
}
func (m *RequestGetComments) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestGetComments.Merge(m, src)
}
func (m *RequestGetComments) XXX_Size() int {
	return xxx_messageInfo_RequestGetComments.Size(m)
}
func (m *RequestGetComments) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestGetComments.DiscardUnknown(m)
}

var xxx_messageInfo_RequestGetComments proto.InternalMessageInfo

func (m *RequestGetComments) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *RequestGetComments) GetTargetId() int32 {
	if m != nil {
		return m.TargetId
	}
	return 0
}

func (m *RequestGetComments) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *RequestGetComments) GetCurrentPage() int32 {
	if m != nil {
		return m.CurrentPage
	}
	return 0
}

func (m *RequestGetComments) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type ResponseGetComments struct {
	Comments             []*Comment `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
	PageSize             int32      `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	CurrentPage          int32      `protobuf:"varint,3,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty"`
	Total                int32      `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ResponseGetComments) Reset()         { *m = ResponseGetComments{} }
func (m *ResponseGetComments) String() string { return proto.CompactTextString(m) }
func (*ResponseGetComments) ProtoMessage()    {}
func (*ResponseGetComments) Descriptor() ([]byte, []int) {
	return fileDescriptor_749aee09ea917828, []int{4}
}

func (m *ResponseGetComments) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseGetComments.Unmarshal(m, b)
}
func (m *ResponseGetComments) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseGetComments.Marshal(b, m, deterministic)
}
func (m *ResponseGetComments) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseGetComments.Merge(m, src)
}
func (m *ResponseGetComments) XXX_Size() int {
	return xxx_messageInfo_ResponseGetComments.Size(m)
}
func (m *ResponseGetComments) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseGetComments.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseGetComments proto.InternalMessageInfo

func (m *ResponseGetComments) GetComments() []*Comment {
	if m != nil {
		return m.Comments
	}
	return nil
}

func (m *ResponseGetComments) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ResponseGetComments) GetCurrentPage() int32 {
	if m != nil {
		return m.CurrentPage
	}
	return 0
}

func (m *ResponseGetComments) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*RequestPutComment)(nil), "commentModule.RequestPutComment")
	proto.RegisterType((*ResponsePutComment)(nil), "commentModule.ResponsePutComment")
	proto.RegisterType((*Comment)(nil), "commentModule.Comment")
	proto.RegisterType((*RequestGetComments)(nil), "commentModule.RequestGetComments")
	proto.RegisterType((*ResponseGetComments)(nil), "commentModule.ResponseGetComments")
}

func init() { proto.RegisterFile("comment.proto", fileDescriptor_749aee09ea917828) }

var fileDescriptor_749aee09ea917828 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcd, 0xaa, 0xd3, 0x40,
	0x14, 0xc7, 0xef, 0xe4, 0xb3, 0x39, 0xd5, 0x8b, 0x8e, 0xa2, 0xe1, 0x8a, 0x90, 0xce, 0xaa, 0xab,
	0x2e, 0xea, 0xce, 0x9d, 0xb8, 0x90, 0x2e, 0x84, 0x4b, 0x0a, 0x6e, 0x4b, 0xcc, 0x1c, 0x42, 0xa0,
	0xcd, 0xc4, 0xcc, 0x44, 0xb0, 0xaf, 0xe2, 0xc2, 0x67, 0x70, 0xe9, 0xc2, 0x77, 0x93, 0xcc, 0x4c,
	0x9a, 0xa4, 0xb5, 0x0a, 0x77, 0x37, 0xe7, 0x33, 0xbf, 0xff, 0x39, 0x27, 0xf0, 0x38, 0x17, 0x87,
	0x03, 0x56, 0x6a, 0x55, 0x37, 0x42, 0x09, 0xda, 0x9b, 0x1f, 0x05, 0x6f, 0xf7, 0xc8, 0x7e, 0x12,
	0x78, 0x9a, 0xe2, 0x97, 0x16, 0xa5, 0xba, 0x6f, 0xd5, 0x7b, 0x13, 0xa3, 0x14, 0x3c, 0xf5, 0xad,
	0xc6, 0x98, 0x24, 0x64, 0x19, 0xa5, 0xfa, 0x4d, 0x5f, 0x41, 0xa4, 0xb2, 0xa6, 0x40, 0xb5, 0x2b,
	0x79, 0xec, 0x24, 0x64, 0xe9, 0xa7, 0x33, 0xe3, 0xd8, 0xf0, 0x2e, 0x58, 0x67, 0x0d, 0x56, 0x3a,
	0xe8, 0x9a, 0xa0, 0x71, 0x6c, 0x38, 0x8d, 0x21, 0xcc, 0x45, 0xa5, 0xb0, 0x52, 0xb1, 0xa7, 0x1b,
	0xf6, 0x26, 0x7d, 0x09, 0x61, 0x2b, 0xb1, 0xe9, 0x8a, 0x7c, 0x5d, 0x14, 0x74, 0xa6, 0xe9, 0x97,
	0xef, 0x4b, 0xdb, 0x2f, 0xd0, 0x45, 0x33, 0xe3, 0xd8, 0x70, 0xf6, 0x16, 0x68, 0x8a, 0xb2, 0x16,
	0x95, 0xc4, 0x29, 0x73, 0x2e, 0xb8, 0x61, 0xf6, 0x53, 0xfd, 0xa6, 0x4f, 0xc0, 0x3d, 0xc8, 0x42,
	0xd3, 0x46, 0x69, 0xf7, 0x64, 0xbf, 0x08, 0x84, 0x7d, 0xc5, 0x2d, 0x38, 0x25, 0xb7, 0xf9, 0x4e,
	0xc9, 0x4f, 0xaa, 0x9d, 0x6b, 0xaa, 0xdd, 0x33, 0xd5, 0x23, 0x7c, 0xef, 0x1c, 0x7f, 0x18, 0x87,
	0x7f, 0x7d, 0x1c, 0xc1, 0x74, 0x1c, 0xaf, 0x01, 0xf2, 0x06, 0x33, 0x85, 0x7c, 0x97, 0xa9, 0x38,
	0xd4, 0x75, 0x91, 0xf5, 0xbc, 0x53, 0xec, 0x3b, 0xe9, 0x84, 0xeb, 0x5d, 0x7d, 0xc0, 0x5e, 0xb7,
	0x7c, 0xe0, 0xb2, 0x0a, 0xdc, 0xc9, 0xf2, 0x88, 0xc3, 0xb2, 0x0a, 0xdc, 0x96, 0x47, 0xa4, 0x0b,
	0x78, 0x94, 0xb7, 0x8d, 0x66, 0xef, 0x7c, 0x56, 0xd8, 0xdc, 0xfa, 0xee, 0xb3, 0x02, 0xe9, 0x73,
	0xf0, 0x95, 0x50, 0xd9, 0xde, 0x2a, 0x33, 0x06, 0xfb, 0x41, 0xe0, 0x59, 0xbf, 0x96, 0x31, 0xde,
	0x1a, 0x66, 0xf6, 0xe4, 0x64, 0x4c, 0x12, 0x77, 0x39, 0x5f, 0xbf, 0x58, 0x4d, 0x6e, 0x70, 0x65,
	0x53, 0xd3, 0x53, 0xde, 0x94, 0xd0, 0xf9, 0x0f, 0xa1, 0xfb, 0x0f, 0x42, 0x6f, 0x44, 0xb8, 0xfe,
	0x4d, 0xe0, 0xd6, 0x7e, 0x6b, 0x8b, 0xcd, 0xd7, 0x32, 0x47, 0xfa, 0x09, 0xe6, 0xc5, 0x88, 0x75,
	0x71, 0x46, 0x76, 0x39, 0xed, 0x3b, 0x76, 0x91, 0x72, 0x21, 0x99, 0xdd, 0xd0, 0x2d, 0x40, 0x3d,
	0x9c, 0x66, 0xf2, 0xf7, 0xb6, 0xc3, 0xf1, 0xde, 0x2d, 0xae, 0x74, 0x1d, 0x52, 0xd8, 0xcd, 0xe7,
	0x40, 0xff, 0xc1, 0x6f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x97, 0xe6, 0x22, 0x53, 0xd2, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommentServiceClient is the client API for CommentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommentServiceClient interface {
	GetComments(ctx context.Context, in *RequestGetComments, opts ...grpc.CallOption) (*ResponseGetComments, error)
	PutComment(ctx context.Context, in *RequestPutComment, opts ...grpc.CallOption) (*ResponsePutComment, error)
}

type commentServiceClient struct {
	cc *grpc.ClientConn
}

func NewCommentServiceClient(cc *grpc.ClientConn) CommentServiceClient {
	return &commentServiceClient{cc}
}

func (c *commentServiceClient) GetComments(ctx context.Context, in *RequestGetComments, opts ...grpc.CallOption) (*ResponseGetComments, error) {
	out := new(ResponseGetComments)
	err := c.cc.Invoke(ctx, "/commentModule.CommentService/getComments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) PutComment(ctx context.Context, in *RequestPutComment, opts ...grpc.CallOption) (*ResponsePutComment, error) {
	out := new(ResponsePutComment)
	err := c.cc.Invoke(ctx, "/commentModule.CommentService/putComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServiceServer is the server API for CommentService service.
type CommentServiceServer interface {
	GetComments(context.Context, *RequestGetComments) (*ResponseGetComments, error)
	PutComment(context.Context, *RequestPutComment) (*ResponsePutComment, error)
}

// UnimplementedCommentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCommentServiceServer struct {
}

func (*UnimplementedCommentServiceServer) GetComments(ctx context.Context, req *RequestGetComments) (*ResponseGetComments, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComments not implemented")
}
func (*UnimplementedCommentServiceServer) PutComment(ctx context.Context, req *RequestPutComment) (*ResponsePutComment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutComment not implemented")
}

func RegisterCommentServiceServer(s *grpc.Server, srv CommentServiceServer) {
	s.RegisterService(&_CommentService_serviceDesc, srv)
}

func _CommentService_GetComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGetComments)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commentModule.CommentService/GetComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetComments(ctx, req.(*RequestGetComments))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_PutComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestPutComment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).PutComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commentModule.CommentService/PutComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).PutComment(ctx, req.(*RequestPutComment))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "commentModule.CommentService",
	HandlerType: (*CommentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getComments",
			Handler:    _CommentService_GetComments_Handler,
		},
		{
			MethodName: "putComment",
			Handler:    _CommentService_PutComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment.proto",
}
