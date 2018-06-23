// Code generated by protoc-gen-go. DO NOT EDIT.
// source: analyze.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Results struct {
	Results              []*Result `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Results) Reset()         { *m = Results{} }
func (m *Results) String() string { return proto.CompactTextString(m) }
func (*Results) ProtoMessage()    {}
func (*Results) Descriptor() ([]byte, []int) {
	return fileDescriptor_analyze_b614ff30b2243d3a, []int{0}
}
func (m *Results) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Results.Unmarshal(m, b)
}
func (m *Results) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Results.Marshal(b, m, deterministic)
}
func (dst *Results) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Results.Merge(dst, src)
}
func (m *Results) XXX_Size() int {
	return xxx_messageInfo_Results.Size(m)
}
func (m *Results) XXX_DiscardUnknown() {
	xxx_messageInfo_Results.DiscardUnknown(m)
}

var xxx_messageInfo_Results proto.InternalMessageInfo

func (m *Results) GetResults() []*Result {
	if m != nil {
		return m.Results
	}
	return nil
}

type AnalyzeApiRequest struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	AnalyzeTemplateId    string   `protobuf:"bytes,2,opt,name=analyzeTemplateId" json:"analyzeTemplateId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnalyzeApiRequest) Reset()         { *m = AnalyzeApiRequest{} }
func (m *AnalyzeApiRequest) String() string { return proto.CompactTextString(m) }
func (*AnalyzeApiRequest) ProtoMessage()    {}
func (*AnalyzeApiRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_analyze_b614ff30b2243d3a, []int{1}
}
func (m *AnalyzeApiRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnalyzeApiRequest.Unmarshal(m, b)
}
func (m *AnalyzeApiRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnalyzeApiRequest.Marshal(b, m, deterministic)
}
func (dst *AnalyzeApiRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnalyzeApiRequest.Merge(dst, src)
}
func (m *AnalyzeApiRequest) XXX_Size() int {
	return xxx_messageInfo_AnalyzeApiRequest.Size(m)
}
func (m *AnalyzeApiRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AnalyzeApiRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AnalyzeApiRequest proto.InternalMessageInfo

func (m *AnalyzeApiRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *AnalyzeApiRequest) GetAnalyzeTemplateId() string {
	if m != nil {
		return m.AnalyzeTemplateId
	}
	return ""
}

type AnalyzeRequest struct {
	Value                string   `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	Fields               []string `protobuf:"bytes,2,rep,name=fields" json:"fields,omitempty"`
	MinProbability       string   `protobuf:"bytes,3,opt,name=minProbability" json:"minProbability,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnalyzeRequest) Reset()         { *m = AnalyzeRequest{} }
func (m *AnalyzeRequest) String() string { return proto.CompactTextString(m) }
func (*AnalyzeRequest) ProtoMessage()    {}
func (*AnalyzeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_analyze_b614ff30b2243d3a, []int{2}
}
func (m *AnalyzeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnalyzeRequest.Unmarshal(m, b)
}
func (m *AnalyzeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnalyzeRequest.Marshal(b, m, deterministic)
}
func (dst *AnalyzeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnalyzeRequest.Merge(dst, src)
}
func (m *AnalyzeRequest) XXX_Size() int {
	return xxx_messageInfo_AnalyzeRequest.Size(m)
}
func (m *AnalyzeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AnalyzeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AnalyzeRequest proto.InternalMessageInfo

func (m *AnalyzeRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *AnalyzeRequest) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *AnalyzeRequest) GetMinProbability() string {
	if m != nil {
		return m.MinProbability
	}
	return ""
}

func init() {
	proto.RegisterType((*Results)(nil), "types.Results")
	proto.RegisterType((*AnalyzeApiRequest)(nil), "types.AnalyzeApiRequest")
	proto.RegisterType((*AnalyzeRequest)(nil), "types.AnalyzeRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AnalyzeService service

type AnalyzeServiceClient interface {
	Apply(ctx context.Context, in *AnalyzeRequest, opts ...grpc.CallOption) (*Results, error)
}

type analyzeServiceClient struct {
	cc *grpc.ClientConn
}

func NewAnalyzeServiceClient(cc *grpc.ClientConn) AnalyzeServiceClient {
	return &analyzeServiceClient{cc}
}

func (c *analyzeServiceClient) Apply(ctx context.Context, in *AnalyzeRequest, opts ...grpc.CallOption) (*Results, error) {
	out := new(Results)
	err := grpc.Invoke(ctx, "/types.AnalyzeService/Apply", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AnalyzeService service

type AnalyzeServiceServer interface {
	Apply(context.Context, *AnalyzeRequest) (*Results, error)
}

func RegisterAnalyzeServiceServer(s *grpc.Server, srv AnalyzeServiceServer) {
	s.RegisterService(&_AnalyzeService_serviceDesc, srv)
}

func _AnalyzeService_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnalyzeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzeServiceServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.AnalyzeService/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzeServiceServer).Apply(ctx, req.(*AnalyzeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AnalyzeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.AnalyzeService",
	HandlerType: (*AnalyzeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Apply",
			Handler:    _AnalyzeService_Apply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "analyze.proto",
}

func init() { proto.RegisterFile("analyze.proto", fileDescriptor_analyze_b614ff30b2243d3a) }

var fileDescriptor_analyze_b614ff30b2243d3a = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xdf, 0x4a, 0xc3, 0x30,
	0x14, 0xc6, 0xed, 0x6a, 0x37, 0x76, 0x74, 0x85, 0x1d, 0x54, 0xca, 0xae, 0x4a, 0x2f, 0xb4, 0x17,
	0x52, 0xa4, 0x3e, 0x41, 0xbd, 0xf3, 0x4e, 0xa2, 0x3e, 0x40, 0xba, 0x9d, 0x41, 0x20, 0x6d, 0x62,
	0x93, 0x0e, 0xeb, 0xd3, 0x0b, 0x49, 0x26, 0xcc, 0xdd, 0x9d, 0x7f, 0xdf, 0x2f, 0x5f, 0x3e, 0x58,
	0xf1, 0x9e, 0xcb, 0xe9, 0x87, 0x2a, 0x3d, 0x28, 0xab, 0x30, 0xb1, 0x93, 0x26, 0xb3, 0xb9, 0xde,
	0xaa, 0xae, 0x53, 0xbd, 0x1f, 0x16, 0x35, 0x2c, 0x18, 0x99, 0x51, 0x5a, 0x83, 0x0f, 0xb0, 0x18,
	0x7c, 0x99, 0x45, 0x79, 0x5c, 0x5e, 0xd5, 0xab, 0xca, 0x29, 0x2a, 0x7f, 0xc0, 0x8e, 0xdb, 0xe2,
	0x13, 0xd6, 0x8d, 0x27, 0x37, 0x5a, 0x30, 0xfa, 0x1a, 0xc9, 0x58, 0x44, 0xb8, 0xb4, 0xf4, 0x6d,
	0xb3, 0x28, 0x8f, 0xca, 0x25, 0x73, 0x35, 0x3e, 0xc2, 0x3a, 0x58, 0xf8, 0xa0, 0x4e, 0x4b, 0x6e,
	0xe9, 0x75, 0x97, 0xcd, 0xdc, 0xc1, 0xf9, 0xa2, 0xd8, 0x43, 0x1a, 0xb0, 0x47, 0xe6, 0x0d, 0x24,
	0x07, 0x2e, 0x47, 0x0a, 0x50, 0xdf, 0xe0, 0x1d, 0xcc, 0xf7, 0x82, 0xe4, 0xce, 0x64, 0xb3, 0x3c,
	0x2e, 0x97, 0x2c, 0x74, 0x78, 0x0f, 0x69, 0x27, 0xfa, 0xb7, 0x41, 0xb5, 0xbc, 0x15, 0x52, 0xd8,
	0x29, 0x8b, 0x9d, 0xec, 0xdf, 0xb4, 0x7e, 0xf9, 0x7b, 0xe7, 0x9d, 0x86, 0x83, 0xd8, 0x12, 0x3e,
	0x41, 0xd2, 0x68, 0x2d, 0x27, 0xbc, 0x0d, 0x3f, 0x3e, 0xf5, 0xb1, 0x49, 0x4f, 0x82, 0x30, 0xc5,
	0x45, 0x3b, 0x77, 0xe9, 0x3d, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x79, 0x17, 0x12, 0x63,
	0x01, 0x00, 0x00,
}
