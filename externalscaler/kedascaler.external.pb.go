// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kedascaler.external.proto

package externalscaler

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type ScaledObjectRef struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScaledObjectRef) Reset()         { *m = ScaledObjectRef{} }
func (m *ScaledObjectRef) String() string { return proto.CompactTextString(m) }
func (*ScaledObjectRef) ProtoMessage()    {}
func (*ScaledObjectRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e5752ad5f3bf705, []int{0}
}

func (m *ScaledObjectRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScaledObjectRef.Unmarshal(m, b)
}
func (m *ScaledObjectRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScaledObjectRef.Marshal(b, m, deterministic)
}
func (m *ScaledObjectRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScaledObjectRef.Merge(m, src)
}
func (m *ScaledObjectRef) XXX_Size() int {
	return xxx_messageInfo_ScaledObjectRef.Size(m)
}
func (m *ScaledObjectRef) XXX_DiscardUnknown() {
	xxx_messageInfo_ScaledObjectRef.DiscardUnknown(m)
}

var xxx_messageInfo_ScaledObjectRef proto.InternalMessageInfo

func (m *ScaledObjectRef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ScaledObjectRef) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type NewRequest struct {
	ScaledObjectRef      *ScaledObjectRef  `protobuf:"bytes,1,opt,name=scaledObjectRef,proto3" json:"scaledObjectRef,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *NewRequest) Reset()         { *m = NewRequest{} }
func (m *NewRequest) String() string { return proto.CompactTextString(m) }
func (*NewRequest) ProtoMessage()    {}
func (*NewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e5752ad5f3bf705, []int{1}
}

func (m *NewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewRequest.Unmarshal(m, b)
}
func (m *NewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewRequest.Marshal(b, m, deterministic)
}
func (m *NewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewRequest.Merge(m, src)
}
func (m *NewRequest) XXX_Size() int {
	return xxx_messageInfo_NewRequest.Size(m)
}
func (m *NewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewRequest proto.InternalMessageInfo

func (m *NewRequest) GetScaledObjectRef() *ScaledObjectRef {
	if m != nil {
		return m.ScaledObjectRef
	}
	return nil
}

func (m *NewRequest) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type IsActiveResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsActiveResponse) Reset()         { *m = IsActiveResponse{} }
func (m *IsActiveResponse) String() string { return proto.CompactTextString(m) }
func (*IsActiveResponse) ProtoMessage()    {}
func (*IsActiveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e5752ad5f3bf705, []int{2}
}

func (m *IsActiveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsActiveResponse.Unmarshal(m, b)
}
func (m *IsActiveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsActiveResponse.Marshal(b, m, deterministic)
}
func (m *IsActiveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsActiveResponse.Merge(m, src)
}
func (m *IsActiveResponse) XXX_Size() int {
	return xxx_messageInfo_IsActiveResponse.Size(m)
}
func (m *IsActiveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IsActiveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IsActiveResponse proto.InternalMessageInfo

func (m *IsActiveResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type GetMetricSpecResponse struct {
	MetricSpecs          []*MetricSpec `protobuf:"bytes,1,rep,name=metricSpecs,proto3" json:"metricSpecs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetMetricSpecResponse) Reset()         { *m = GetMetricSpecResponse{} }
func (m *GetMetricSpecResponse) String() string { return proto.CompactTextString(m) }
func (*GetMetricSpecResponse) ProtoMessage()    {}
func (*GetMetricSpecResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e5752ad5f3bf705, []int{3}
}

func (m *GetMetricSpecResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMetricSpecResponse.Unmarshal(m, b)
}
func (m *GetMetricSpecResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMetricSpecResponse.Marshal(b, m, deterministic)
}
func (m *GetMetricSpecResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMetricSpecResponse.Merge(m, src)
}
func (m *GetMetricSpecResponse) XXX_Size() int {
	return xxx_messageInfo_GetMetricSpecResponse.Size(m)
}
func (m *GetMetricSpecResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMetricSpecResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMetricSpecResponse proto.InternalMessageInfo

func (m *GetMetricSpecResponse) GetMetricSpecs() []*MetricSpec {
	if m != nil {
		return m.MetricSpecs
	}
	return nil
}

type MetricSpec struct {
	MetricName           string   `protobuf:"bytes,1,opt,name=metricName,proto3" json:"metricName,omitempty"`
	TargetSize           int64    `protobuf:"varint,2,opt,name=targetSize,proto3" json:"targetSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricSpec) Reset()         { *m = MetricSpec{} }
func (m *MetricSpec) String() string { return proto.CompactTextString(m) }
func (*MetricSpec) ProtoMessage()    {}
func (*MetricSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e5752ad5f3bf705, []int{4}
}

func (m *MetricSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricSpec.Unmarshal(m, b)
}
func (m *MetricSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricSpec.Marshal(b, m, deterministic)
}
func (m *MetricSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricSpec.Merge(m, src)
}
func (m *MetricSpec) XXX_Size() int {
	return xxx_messageInfo_MetricSpec.Size(m)
}
func (m *MetricSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricSpec.DiscardUnknown(m)
}

var xxx_messageInfo_MetricSpec proto.InternalMessageInfo

func (m *MetricSpec) GetMetricName() string {
	if m != nil {
		return m.MetricName
	}
	return ""
}

func (m *MetricSpec) GetTargetSize() int64 {
	if m != nil {
		return m.TargetSize
	}
	return 0
}

type GetMetricsRequest struct {
	ScaledObjectRef      *ScaledObjectRef `protobuf:"bytes,1,opt,name=scaledObjectRef,proto3" json:"scaledObjectRef,omitempty"`
	MetricName           string           `protobuf:"bytes,2,opt,name=metricName,proto3" json:"metricName,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetMetricsRequest) Reset()         { *m = GetMetricsRequest{} }
func (m *GetMetricsRequest) String() string { return proto.CompactTextString(m) }
func (*GetMetricsRequest) ProtoMessage()    {}
func (*GetMetricsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e5752ad5f3bf705, []int{5}
}

func (m *GetMetricsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMetricsRequest.Unmarshal(m, b)
}
func (m *GetMetricsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMetricsRequest.Marshal(b, m, deterministic)
}
func (m *GetMetricsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMetricsRequest.Merge(m, src)
}
func (m *GetMetricsRequest) XXX_Size() int {
	return xxx_messageInfo_GetMetricsRequest.Size(m)
}
func (m *GetMetricsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMetricsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMetricsRequest proto.InternalMessageInfo

func (m *GetMetricsRequest) GetScaledObjectRef() *ScaledObjectRef {
	if m != nil {
		return m.ScaledObjectRef
	}
	return nil
}

func (m *GetMetricsRequest) GetMetricName() string {
	if m != nil {
		return m.MetricName
	}
	return ""
}

type GetMetricsResponse struct {
	MetricValues         []*MetricValue `protobuf:"bytes,1,rep,name=metricValues,proto3" json:"metricValues,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetMetricsResponse) Reset()         { *m = GetMetricsResponse{} }
func (m *GetMetricsResponse) String() string { return proto.CompactTextString(m) }
func (*GetMetricsResponse) ProtoMessage()    {}
func (*GetMetricsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e5752ad5f3bf705, []int{6}
}

func (m *GetMetricsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMetricsResponse.Unmarshal(m, b)
}
func (m *GetMetricsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMetricsResponse.Marshal(b, m, deterministic)
}
func (m *GetMetricsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMetricsResponse.Merge(m, src)
}
func (m *GetMetricsResponse) XXX_Size() int {
	return xxx_messageInfo_GetMetricsResponse.Size(m)
}
func (m *GetMetricsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMetricsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMetricsResponse proto.InternalMessageInfo

func (m *GetMetricsResponse) GetMetricValues() []*MetricValue {
	if m != nil {
		return m.MetricValues
	}
	return nil
}

type MetricValue struct {
	MetricName           string   `protobuf:"bytes,1,opt,name=metricName,proto3" json:"metricName,omitempty"`
	MetricValue          int64    `protobuf:"varint,2,opt,name=metricValue,proto3" json:"metricValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricValue) Reset()         { *m = MetricValue{} }
func (m *MetricValue) String() string { return proto.CompactTextString(m) }
func (*MetricValue) ProtoMessage()    {}
func (*MetricValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e5752ad5f3bf705, []int{7}
}

func (m *MetricValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricValue.Unmarshal(m, b)
}
func (m *MetricValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricValue.Marshal(b, m, deterministic)
}
func (m *MetricValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricValue.Merge(m, src)
}
func (m *MetricValue) XXX_Size() int {
	return xxx_messageInfo_MetricValue.Size(m)
}
func (m *MetricValue) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricValue.DiscardUnknown(m)
}

var xxx_messageInfo_MetricValue proto.InternalMessageInfo

func (m *MetricValue) GetMetricName() string {
	if m != nil {
		return m.MetricName
	}
	return ""
}

func (m *MetricValue) GetMetricValue() int64 {
	if m != nil {
		return m.MetricValue
	}
	return 0
}

func init() {
	proto.RegisterType((*ScaledObjectRef)(nil), "externalscaler.ScaledObjectRef")
	proto.RegisterType((*NewRequest)(nil), "externalscaler.NewRequest")
	proto.RegisterMapType((map[string]string)(nil), "externalscaler.NewRequest.MetadataEntry")
	proto.RegisterType((*IsActiveResponse)(nil), "externalscaler.IsActiveResponse")
	proto.RegisterType((*GetMetricSpecResponse)(nil), "externalscaler.GetMetricSpecResponse")
	proto.RegisterType((*MetricSpec)(nil), "externalscaler.MetricSpec")
	proto.RegisterType((*GetMetricsRequest)(nil), "externalscaler.GetMetricsRequest")
	proto.RegisterType((*GetMetricsResponse)(nil), "externalscaler.GetMetricsResponse")
	proto.RegisterType((*MetricValue)(nil), "externalscaler.MetricValue")
}

func init() { proto.RegisterFile("kedascaler.external.proto", fileDescriptor_6e5752ad5f3bf705) }

var fileDescriptor_6e5752ad5f3bf705 = []byte{
	// 490 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x8d, 0x63, 0x5a, 0xa5, 0x13, 0xfa, 0xc1, 0x08, 0xaa, 0xe0, 0x22, 0x08, 0x2b, 0x21, 0x45,
	0x1c, 0x5c, 0x29, 0x5c, 0x10, 0x45, 0x42, 0x50, 0x22, 0x54, 0x89, 0x26, 0xd2, 0x46, 0xa9, 0xc4,
	0x71, 0xe3, 0x4c, 0xa3, 0x50, 0x27, 0x36, 0xde, 0x4d, 0x4b, 0x38, 0xf0, 0x2b, 0xf8, 0x75, 0xfc,
	0x1a, 0xe4, 0xf5, 0xf7, 0x2a, 0x21, 0x5c, 0x38, 0xd9, 0x3b, 0xef, 0xcd, 0xec, 0xbc, 0x37, 0xb3,
	0xf0, 0xf8, 0x86, 0x26, 0x42, 0x7a, 0xc2, 0xa7, 0xc8, 0xa5, 0xef, 0x8a, 0xa2, 0x85, 0xf0, 0xdd,
	0x30, 0x0a, 0x54, 0x80, 0x07, 0xd9, 0x39, 0x81, 0x9d, 0x93, 0x69, 0x10, 0x4c, 0x7d, 0x3a, 0xd5,
	0xe8, 0x78, 0x79, 0x7d, 0x4a, 0xf3, 0x50, 0xad, 0x12, 0x32, 0x3b, 0x87, 0xc3, 0x61, 0x4c, 0x9b,
	0x0c, 0xc6, 0x5f, 0xc9, 0x53, 0x9c, 0xae, 0x11, 0xe1, 0xde, 0x42, 0xcc, 0xa9, 0x65, 0xb5, 0xad,
	0xce, 0x1e, 0xd7, 0xff, 0xf8, 0x04, 0xf6, 0xe2, 0xaf, 0x0c, 0x85, 0x47, 0xad, 0xba, 0x06, 0x8a,
	0x00, 0xfb, 0x6d, 0x01, 0xf4, 0xe9, 0x8e, 0xd3, 0xb7, 0x25, 0x49, 0x85, 0x17, 0x70, 0x28, 0xab,
	0x35, 0x75, 0xad, 0x66, 0xf7, 0x99, 0x5b, 0x6d, 0xcd, 0x35, 0xae, 0xe6, 0x66, 0x1e, 0x7e, 0x84,
	0xc6, 0x9c, 0x94, 0x98, 0x08, 0x25, 0x5a, 0xf5, 0xb6, 0xdd, 0x69, 0x76, 0x3b, 0x66, 0x8d, 0xe2,
	0x62, 0xf7, 0x32, 0xa5, 0xf6, 0x16, 0x2a, 0x5a, 0xf1, 0x3c, 0xd3, 0x39, 0x83, 0xfd, 0x0a, 0x84,
	0x47, 0x60, 0xdf, 0xd0, 0x2a, 0x55, 0x18, 0xff, 0xe2, 0x43, 0xd8, 0xb9, 0x15, 0xfe, 0x32, 0x13,
	0x97, 0x1c, 0xde, 0xd4, 0x5f, 0x5b, 0xec, 0x25, 0x1c, 0x5d, 0xc8, 0xf7, 0x9e, 0x9a, 0xdd, 0x12,
	0x27, 0x19, 0x06, 0x0b, 0x49, 0x78, 0x0c, 0xbb, 0x11, 0xc9, 0xa5, 0xaf, 0x74, 0x89, 0x06, 0x4f,
	0x4f, 0x6c, 0x04, 0x8f, 0x3e, 0x91, 0xba, 0x24, 0x15, 0xcd, 0xbc, 0x61, 0x48, 0x5e, 0x9e, 0xf0,
	0x16, 0x9a, 0xf3, 0x3c, 0x2a, 0x5b, 0x96, 0x96, 0xe2, 0x98, 0x52, 0x4a, 0x89, 0x65, 0x3a, 0xfb,
	0x0c, 0x50, 0x40, 0xf8, 0x14, 0x20, 0x01, 0xfb, 0xc5, 0x94, 0x4a, 0x91, 0x18, 0x57, 0x22, 0x9a,
	0x92, 0x1a, 0xce, 0x7e, 0x24, 0x7a, 0x6c, 0x5e, 0x8a, 0xb0, 0x9f, 0xf0, 0x20, 0x6f, 0x52, 0xfe,
	0x87, 0x99, 0x55, 0xfb, 0xab, 0x9b, 0xfd, 0xb1, 0x11, 0x60, 0xf9, 0xfe, 0xd4, 0xa1, 0x77, 0x70,
	0x3f, 0xe1, 0x5c, 0xc5, 0xce, 0x67, 0x16, 0x9d, 0xac, 0xb7, 0x48, 0x73, 0x78, 0x25, 0x81, 0x0d,
	0xa0, 0x59, 0x02, 0xb7, 0xba, 0xd4, 0xce, 0x26, 0x72, 0x95, 0x8f, 0xdd, 0xe6, 0xe5, 0x50, 0xf7,
	0x97, 0x0d, 0x07, 0xbd, 0xf4, 0x76, 0x2d, 0x3a, 0xc2, 0x33, 0xb0, 0xfb, 0x74, 0x87, 0xce, 0xe6,
	0x1d, 0x74, 0x8e, 0xdd, 0xe4, 0xb9, 0xb9, 0xd9, 0x73, 0x73, 0x7b, 0xf1, 0x73, 0x63, 0x35, 0x1c,
	0x40, 0x23, 0x5b, 0x24, 0xdc, 0xe6, 0xaa, 0xd3, 0x36, 0x09, 0xe6, 0x0e, 0xb2, 0x1a, 0x7e, 0x81,
	0xfd, 0xca, 0xb6, 0x6d, 0xaf, 0xfa, 0xc2, 0x24, 0xac, 0xdd, 0x56, 0x56, 0xc3, 0x11, 0x40, 0x31,
	0x23, 0x7c, 0xbe, 0x31, 0x2d, 0xdb, 0x1f, 0x87, 0xfd, 0x8d, 0x92, 0x97, 0xfd, 0x00, 0x3b, 0xe7,
	0x7e, 0x20, 0xff, 0x41, 0xff, 0x46, 0x1b, 0xc7, 0xbb, 0x3a, 0xf2, 0xea, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x36, 0xaf, 0x11, 0xff, 0x02, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExternalScalerClient is the client API for ExternalScaler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExternalScalerClient interface {
	New(ctx context.Context, in *NewRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	IsActive(ctx context.Context, in *ScaledObjectRef, opts ...grpc.CallOption) (*IsActiveResponse, error)
	GetMetricSpec(ctx context.Context, in *ScaledObjectRef, opts ...grpc.CallOption) (*GetMetricSpecResponse, error)
	GetMetrics(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error)
	Close(ctx context.Context, in *ScaledObjectRef, opts ...grpc.CallOption) (*empty.Empty, error)
}

type externalScalerClient struct {
	cc *grpc.ClientConn
}

func NewExternalScalerClient(cc *grpc.ClientConn) ExternalScalerClient {
	return &externalScalerClient{cc}
}

func (c *externalScalerClient) New(ctx context.Context, in *NewRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/externalscaler.ExternalScaler/New", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalScalerClient) IsActive(ctx context.Context, in *ScaledObjectRef, opts ...grpc.CallOption) (*IsActiveResponse, error) {
	out := new(IsActiveResponse)
	err := c.cc.Invoke(ctx, "/externalscaler.ExternalScaler/IsActive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalScalerClient) GetMetricSpec(ctx context.Context, in *ScaledObjectRef, opts ...grpc.CallOption) (*GetMetricSpecResponse, error) {
	out := new(GetMetricSpecResponse)
	err := c.cc.Invoke(ctx, "/externalscaler.ExternalScaler/GetMetricSpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalScalerClient) GetMetrics(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error) {
	out := new(GetMetricsResponse)
	err := c.cc.Invoke(ctx, "/externalscaler.ExternalScaler/GetMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalScalerClient) Close(ctx context.Context, in *ScaledObjectRef, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/externalscaler.ExternalScaler/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExternalScalerServer is the server API for ExternalScaler service.
type ExternalScalerServer interface {
	New(context.Context, *NewRequest) (*empty.Empty, error)
	IsActive(context.Context, *ScaledObjectRef) (*IsActiveResponse, error)
	GetMetricSpec(context.Context, *ScaledObjectRef) (*GetMetricSpecResponse, error)
	GetMetrics(context.Context, *GetMetricsRequest) (*GetMetricsResponse, error)
	Close(context.Context, *ScaledObjectRef) (*empty.Empty, error)
}

// UnimplementedExternalScalerServer can be embedded to have forward compatible implementations.
type UnimplementedExternalScalerServer struct {
}

func (*UnimplementedExternalScalerServer) New(ctx context.Context, req *NewRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method New not implemented")
}
func (*UnimplementedExternalScalerServer) IsActive(ctx context.Context, req *ScaledObjectRef) (*IsActiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsActive not implemented")
}
func (*UnimplementedExternalScalerServer) GetMetricSpec(ctx context.Context, req *ScaledObjectRef) (*GetMetricSpecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetricSpec not implemented")
}
func (*UnimplementedExternalScalerServer) GetMetrics(ctx context.Context, req *GetMetricsRequest) (*GetMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetrics not implemented")
}
func (*UnimplementedExternalScalerServer) Close(ctx context.Context, req *ScaledObjectRef) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}

func RegisterExternalScalerServer(s *grpc.Server, srv ExternalScalerServer) {
	s.RegisterService(&_ExternalScaler_serviceDesc, srv)
}

func _ExternalScaler_New_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalScalerServer).New(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/externalscaler.ExternalScaler/New",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalScalerServer).New(ctx, req.(*NewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalScaler_IsActive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScaledObjectRef)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalScalerServer).IsActive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/externalscaler.ExternalScaler/IsActive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalScalerServer).IsActive(ctx, req.(*ScaledObjectRef))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalScaler_GetMetricSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScaledObjectRef)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalScalerServer).GetMetricSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/externalscaler.ExternalScaler/GetMetricSpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalScalerServer).GetMetricSpec(ctx, req.(*ScaledObjectRef))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalScaler_GetMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalScalerServer).GetMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/externalscaler.ExternalScaler/GetMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalScalerServer).GetMetrics(ctx, req.(*GetMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalScaler_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScaledObjectRef)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalScalerServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/externalscaler.ExternalScaler/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalScalerServer).Close(ctx, req.(*ScaledObjectRef))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExternalScaler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "externalscaler.ExternalScaler",
	HandlerType: (*ExternalScalerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "New",
			Handler:    _ExternalScaler_New_Handler,
		},
		{
			MethodName: "IsActive",
			Handler:    _ExternalScaler_IsActive_Handler,
		},
		{
			MethodName: "GetMetricSpec",
			Handler:    _ExternalScaler_GetMetricSpec_Handler,
		},
		{
			MethodName: "GetMetrics",
			Handler:    _ExternalScaler_GetMetrics_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _ExternalScaler_Close_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kedascaler.external.proto",
}
