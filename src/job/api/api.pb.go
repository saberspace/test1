// Code generated by protoc-gen-go. DO NOT EDIT.
// source: job/api.proto

package api

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	models "github.com/nirajgeorgian/gateway/src/job/models"
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

//  --------------- Job api --------------------
type CreateJobReq struct {
	Job                  *models.Job `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateJobReq) Reset()         { *m = CreateJobReq{} }
func (m *CreateJobReq) String() string { return proto.CompactTextString(m) }
func (*CreateJobReq) ProtoMessage()    {}
func (*CreateJobReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_69f9368fd342fe7b, []int{0}
}

func (m *CreateJobReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateJobReq.Unmarshal(m, b)
}
func (m *CreateJobReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateJobReq.Marshal(b, m, deterministic)
}
func (m *CreateJobReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateJobReq.Merge(m, src)
}
func (m *CreateJobReq) XXX_Size() int {
	return xxx_messageInfo_CreateJobReq.Size(m)
}
func (m *CreateJobReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateJobReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateJobReq proto.InternalMessageInfo

func (m *CreateJobReq) GetJob() *models.Job {
	if m != nil {
		return m.Job
	}
	return nil
}

type CreateJobRes struct {
	Job                  *models.Job `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateJobRes) Reset()         { *m = CreateJobRes{} }
func (m *CreateJobRes) String() string { return proto.CompactTextString(m) }
func (*CreateJobRes) ProtoMessage()    {}
func (*CreateJobRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_69f9368fd342fe7b, []int{1}
}

func (m *CreateJobRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateJobRes.Unmarshal(m, b)
}
func (m *CreateJobRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateJobRes.Marshal(b, m, deterministic)
}
func (m *CreateJobRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateJobRes.Merge(m, src)
}
func (m *CreateJobRes) XXX_Size() int {
	return xxx_messageInfo_CreateJobRes.Size(m)
}
func (m *CreateJobRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateJobRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreateJobRes proto.InternalMessageInfo

func (m *CreateJobRes) GetJob() *models.Job {
	if m != nil {
		return m.Job
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateJobReq)(nil), "job.CreateJobReq")
	proto.RegisterType((*CreateJobRes)(nil), "job.CreateJobRes")
}

func init() { proto.RegisterFile("job/api.proto", fileDescriptor_69f9368fd342fe7b) }

var fileDescriptor_69f9368fd342fe7b = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0xca, 0x4f, 0xd2,
	0x4f, 0x2c, 0xc8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0xca, 0x4f, 0x92, 0xd2,
	0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7,
	0x07, 0xcb, 0x25, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0x23, 0xc5, 0x0f, 0x32,
	0x22, 0x37, 0x3f, 0x25, 0x35, 0x07, 0x22, 0xa0, 0xa4, 0xc5, 0xc5, 0xe3, 0x5c, 0x94, 0x9a, 0x58,
	0x92, 0xea, 0x95, 0x9f, 0x14, 0x94, 0x5a, 0x28, 0x24, 0xc5, 0x05, 0x32, 0x56, 0x82, 0x51, 0x81,
	0x51, 0x83, 0xdb, 0x88, 0x43, 0x2f, 0x2b, 0x3f, 0x49, 0x0f, 0x24, 0x03, 0x12, 0x44, 0x53, 0x5b,
	0x8c, 0x4f, 0xad, 0x91, 0x23, 0x17, 0x97, 0x57, 0x7e, 0x52, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72,
	0xaa, 0x90, 0x31, 0x17, 0x27, 0x5c, 0xa7, 0x90, 0x20, 0x58, 0x25, 0xb2, 0xad, 0x52, 0x18, 0x42,
	0xc5, 0x4a, 0x0c, 0x4e, 0x66, 0x07, 0x1e, 0xcb, 0x31, 0x46, 0x19, 0x20, 0x79, 0x30, 0x2f, 0xb3,
	0x28, 0x31, 0x2b, 0x3d, 0x35, 0xbf, 0x28, 0x3d, 0x33, 0x31, 0x4f, 0x3f, 0x3d, 0xb1, 0x24, 0xb5,
	0x3c, 0xb1, 0x52, 0xbf, 0xb8, 0x28, 0x59, 0x1f, 0x1a, 0x32, 0xd6, 0x89, 0x05, 0x99, 0x49, 0x6c,
	0x60, 0x9f, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x92, 0xf8, 0x88, 0x85, 0x2f, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// JobServiceClient is the client API for JobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JobServiceClient interface {
	CreateJob(ctx context.Context, in *CreateJobReq, opts ...grpc.CallOption) (*CreateJobRes, error)
}

type jobServiceClient struct {
	cc *grpc.ClientConn
}

func NewJobServiceClient(cc *grpc.ClientConn) JobServiceClient {
	return &jobServiceClient{cc}
}

func (c *jobServiceClient) CreateJob(ctx context.Context, in *CreateJobReq, opts ...grpc.CallOption) (*CreateJobRes, error) {
	out := new(CreateJobRes)
	err := c.cc.Invoke(ctx, "/job.JobService/CreateJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServiceServer is the server API for JobService service.
type JobServiceServer interface {
	CreateJob(context.Context, *CreateJobReq) (*CreateJobRes, error)
}

// UnimplementedJobServiceServer can be embedded to have forward compatible implementations.
type UnimplementedJobServiceServer struct {
}

func (*UnimplementedJobServiceServer) CreateJob(ctx context.Context, req *CreateJobReq) (*CreateJobRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJob not implemented")
}

func RegisterJobServiceServer(s *grpc.Server, srv JobServiceServer) {
	s.RegisterService(&_JobService_serviceDesc, srv)
}

func _JobService_CreateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateJobReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).CreateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/CreateJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).CreateJob(ctx, req.(*CreateJobReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _JobService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "job.JobService",
	HandlerType: (*JobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJob",
			Handler:    _JobService_CreateJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job/api.proto",
}