package v1beta1

import (
	fmt "fmt"

	io "io"

	proto "github.com/gogo/protobuf/proto"
	v11 "k8s.io/api/core/v1"

	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

func (m *HTTPIngressPath) Reset()      { *m = HTTPIngressPath{} }
func (*HTTPIngressPath) ProtoMessage() {}
func (*HTTPIngressPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bea11de0ceb8f53, []int{0}
}
func (m *HTTPIngressPath) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HTTPIngressPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *HTTPIngressPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPIngressPath.Merge(m, src)
}
func (m *HTTPIngressPath) XXX_Size() int {
	return m.Size()
}
func (m *HTTPIngressPath) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTPIngressPath.DiscardUnknown(m)
}

var xxx_messageInfo_HTTPIngressPath proto.InternalMessageInfo

func (m *HTTPIngressRuleValue) Reset()      { *m = HTTPIngressRuleValue{} }
func (*HTTPIngressRuleValue) ProtoMessage() {}
func (*HTTPIngressRuleValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bea11de0ceb8f53, []int{1}
}
func (m *HTTPIngressRuleValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HTTPIngressRuleValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *HTTPIngressRuleValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPIngressRuleValue.Merge(m, src)
}
func (m *HTTPIngressRuleValue) XXX_Size() int {
	return m.Size()
}
func (m *HTTPIngressRuleValue) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTPIngressRuleValue.DiscardUnknown(m)
}

var xxx_messageInfo_HTTPIngressRuleValue proto.InternalMessageInfo

func (m *Ingress) Reset()      { *m = Ingress{} }
func (*Ingress) ProtoMessage() {}
func (*Ingress) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bea11de0ceb8f53, []int{2}
}
func (m *Ingress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Ingress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Ingress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ingress.Merge(m, src)
}
func (m *Ingress) XXX_Size() int {
	return m.Size()
}
func (m *Ingress) XXX_DiscardUnknown() {
	xxx_messageInfo_Ingress.DiscardUnknown(m)
}

var xxx_messageInfo_Ingress proto.InternalMessageInfo

func (m *IngressBackend) Reset()      { *m = IngressBackend{} }
func (*IngressBackend) ProtoMessage() {}
func (*IngressBackend) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bea11de0ceb8f53, []int{3}
}
func (m *IngressBackend) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IngressBackend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *IngressBackend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IngressBackend.Merge(m, src)
}
func (m *IngressBackend) XXX_Size() int {
	return m.Size()
}
func (m *IngressBackend) XXX_DiscardUnknown() {
	xxx_messageInfo_IngressBackend.DiscardUnknown(m)
}

var xxx_messageInfo_IngressBackend proto.InternalMessageInfo

func (m *IngressClass) Reset()      { *m = IngressClass{} }
func (*IngressClass) ProtoMessage() {}
func (*IngressClass) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bea11de0ceb8f53, []int{4}
}
func (m *IngressClass) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IngressClass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *IngressClass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IngressClass.Merge(m, src)
}
func (m *IngressClass) XXX_Size() int {
	return m.Size()
}
func (m *IngressClass) XXX_DiscardUnknown() {
	xxx_messageInfo_IngressClass.DiscardUnknown(m)
}

var xxx_messageInfo_IngressClass proto.InternalMessageInfo

func (m *IngressClassList) Reset()      { *m = IngressClassList{} }
func (*IngressClassList) ProtoMessage() {}
func (*IngressClassList) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bea11de0ceb8f53, []int{5}
}
func (m *IngressClassList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IngressClassList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *IngressClassList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IngressClassList.Merge(m, src)
}
func (m *IngressClassList) XXX_Size() int {
	return m.Size()
}
func (m *IngressClassList) XXX_DiscardUnknown() {
	xxx_messageInfo_IngressClassList.DiscardUnknown(m)
}

var xxx_messageInfo_IngressClassList proto.InternalMessageInfo

func (m *IngressClassParametersReference) Reset()      { *m = IngressClassParametersReference{} }
func (*IngressClassParametersReference) ProtoMessage() {}
func (*IngressClassParametersReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bea11de0ceb8f53, []int{6}
}
func (m *IngressClassParametersReference) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IngressClassParametersReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *IngressClassParametersReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IngressClassParametersReference.Merge(m, src)
}
func (m *IngressClassParametersReference) XXX_Size() int {
	return m.Size()
}
func (m *IngressClassParametersReference) XXX_DiscardUnknown() {
	xxx_messageInfo_IngressClassParametersReference.DiscardUnknown(m)
}

var xxx_messageInfo_IngressClassParametersReference proto.InternalMessageInfo

func (m *IngressClassSpec) Reset()      { *m = IngressClassSpec{} }
func (*IngressClassSpec) ProtoMessage() {}
func (*IngressClassSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bea11de0ceb8f53, []int{7}
}
func (m *IngressClassSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IngressClassSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *IngressClassSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IngressClassSpec.Merge(m, src)
}
func (m *IngressClassSpec) XXX_Size() int {
	return m.Size()
}
func (m *IngressClassSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_IngressClassSpec.DiscardUnknown(m)
}

var xxx_messageInfo_IngressClassSpec proto.InternalMessageInfo

func (m *IngressList) Reset()      { *m = IngressList{} }
func (*IngressList) ProtoMessage() {}
func (*IngressList) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bea11de0ceb8f53, []int{8}
}
func (m *IngressList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IngressList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *IngressList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IngressList.Merge(m, src)
}
func (m *IngressList) XXX_Size() int {
	return m.Size()
}
func (m *IngressList) XXX_DiscardUnknown() {
	xxx_messageInfo_IngressList.DiscardUnknown(m)
}

var xxx_messageInfo_IngressList proto.InternalMessageInfo

func (m *IngressRule) Reset()      { *m = IngressRule{} }
func (*IngressRule) ProtoMessage() {}
func (*IngressRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bea11de0ceb8f53, []int{9}
}
func (m *IngressRule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IngressRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *IngressRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IngressRule.Merge(m, src)
}
func (m *IngressRule) XXX_Size() int {
	return m.Size()
}
func (m *IngressRule) XXX_DiscardUnknown() {
	xxx_messageInfo_IngressRule.DiscardUnknown(m)
}

var xxx_messageInfo_IngressRule proto.InternalMessageInfo

func (m *IngressRuleValue) Reset()      { *m = IngressRuleValue{} }
func (*IngressRuleValue) ProtoMessage() {}
func (*IngressRuleValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bea11de0ceb8f53, []int{10}
}
func (m *IngressRuleValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IngressRuleValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *IngressRuleValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IngressRuleValue.Merge(m, src)
}
func (m *IngressRuleValue) XXX_Size() int {
	return m.Size()
}
func (m *IngressRuleValue) XXX_DiscardUnknown() {
	xxx_messageInfo_IngressRuleValue.DiscardUnknown(m)
}

var xxx_messageInfo_IngressRuleValue proto.InternalMessageInfo

func (m *IngressSpec) Reset()      { *m = IngressSpec{} }
func (*IngressSpec) ProtoMessage() {}
func (*IngressSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bea11de0ceb8f53, []int{11}
}
func (m *IngressSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IngressSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *IngressSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IngressSpec.Merge(m, src)
}
func (m *IngressSpec) XXX_Size() int {
	return m.Size()
}
func (m *IngressSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_IngressSpec.DiscardUnknown(m)
}

var xxx_messageInfo_IngressSpec proto.InternalMessageInfo

func (m *IngressStatus) Reset()      { *m = IngressStatus{} }
func (*IngressStatus) ProtoMessage() {}
func (*IngressStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bea11de0ceb8f53, []int{12}
}
func (m *IngressStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IngressStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *IngressStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IngressStatus.Merge(m, src)
}
func (m *IngressStatus) XXX_Size() int {
	return m.Size()
}
func (m *IngressStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_IngressStatus.DiscardUnknown(m)
}

var xxx_messageInfo_IngressStatus proto.InternalMessageInfo

func (m *IngressTLS) Reset()      { *m = IngressTLS{} }
func (*IngressTLS) ProtoMessage() {}
func (*IngressTLS) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bea11de0ceb8f53, []int{13}
}
func (m *IngressTLS) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IngressTLS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *IngressTLS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IngressTLS.Merge(m, src)
}
func (m *IngressTLS) XXX_Size() int {
	return m.Size()
}
func (m *IngressTLS) XXX_DiscardUnknown() {
	xxx_messageInfo_IngressTLS.DiscardUnknown(m)
}

var xxx_messageInfo_IngressTLS proto.InternalMessageInfo

func init() {
	proto.RegisterType((*HTTPIngressPath)(nil), "k8s.io.api.networking.v1beta1.HTTPIngressPath")
	proto.RegisterType((*HTTPIngressRuleValue)(nil), "k8s.io.api.networking.v1beta1.HTTPIngressRuleValue")
	proto.RegisterType((*Ingress)(nil), "k8s.io.api.networking.v1beta1.Ingress")
	proto.RegisterType((*IngressBackend)(nil), "k8s.io.api.networking.v1beta1.IngressBackend")
	proto.RegisterType((*IngressClass)(nil), "k8s.io.api.networking.v1beta1.IngressClass")
	proto.RegisterType((*IngressClassList)(nil), "k8s.io.api.networking.v1beta1.IngressClassList")
	proto.RegisterType((*IngressClassParametersReference)(nil), "k8s.io.api.networking.v1beta1.IngressClassParametersReference")
	proto.RegisterType((*IngressClassSpec)(nil), "k8s.io.api.networking.v1beta1.IngressClassSpec")
	proto.RegisterType((*IngressList)(nil), "k8s.io.api.networking.v1beta1.IngressList")
	proto.RegisterType((*IngressRule)(nil), "k8s.io.api.networking.v1beta1.IngressRule")
	proto.RegisterType((*IngressRuleValue)(nil), "k8s.io.api.networking.v1beta1.IngressRuleValue")
	proto.RegisterType((*IngressSpec)(nil), "k8s.io.api.networking.v1beta1.IngressSpec")
	proto.RegisterType((*IngressStatus)(nil), "k8s.io.api.networking.v1beta1.IngressStatus")
	proto.RegisterType((*IngressTLS)(nil), "k8s.io.api.networking.v1beta1.IngressTLS")
}

func init() {
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/api/networking/v1beta1/generated.proto", fileDescriptor_5bea11de0ceb8f53)
}

var fileDescriptor_5bea11de0ceb8f53 = []byte{
	// 1085 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x41, 0x6f, 0xe3, 0x44,
	0x14, 0xae, 0x93, 0x66, 0x9b, 0x4e, 0xd2, 0x6e, 0x35, 0xf4, 0x10, 0x55, 0x22, 0xa9, 0x7c, 0x40,
	0x85, 0xa5, 0x36, 0xcd, 0x2e, 0x88, 0x13, 0x02, 0xaf, 0x04, 0xad, 0x36, 0x6c, 0xc3, 0x24, 0x02,
	0x84, 0x38, 0xec, 0xc4, 0x79, 0xeb, 0x98, 0x38, 0xb6, 0x99, 0x19, 0x07, 0xed, 0x8d, 0x2b, 0x27,
	0xf8, 0x15, 0xfc, 0x04, 0xc4, 0x11, 0xc1, 0xa5, 0xc7, 0x3d, 0xee, 0x85, 0x8a, 0x86, 0x7f, 0xc1,
	0x09, 0xcd, 0x78, 0x62, 0x3b, 0x49, 0xcb, 0xa6, 0x1c, 0xf6, 0x94, 0xcc, 0x7b, 0xdf, 0x7b, 0x6f,
	0xde, 0x7b, 0xdf, 0xbc, 0x67, 0xf4, 0xf1, 0xf8, 0x7d, 0x6e, 0xf9, 0x91, 0x3d, 0x4e, 0x06, 0xc0,
	0x42, 0x10, 0xc0, 0xed, 0x29, 0x84, 0xc3, 0x88, 0xd9, 0x5a, 0x41, 0x63, 0xdf, 0x0e, 0x41, 0x7c,
	0x17, 0xb1, 0xb1, 0x1f, 0x7a, 0xf6, 0xf4, 0x64, 0x00, 0x82, 0x9e, 0xd8, 0x1e, 0x84, 0xc0, 0xa8,
	0x80, 0xa1, 0x15, 0xb3, 0x48, 0x44, 0xf8, 0xf5, 0x14, 0x6e, 0xd1, 0xd8, 0xb7, 0x72, 0xb8, 0xa5,
	0xe1, 0x07, 0xc7, 0x9e, 0x2f, 0x46, 0xc9, 0xc0, 0x72, 0xa3, 0x89, 0xed, 0x45, 0x5e, 0x64, 0x2b,
	0xab, 0x41, 0xf2, 0x54, 0x9d, 0xd4, 0x41, 0xfd, 0x4b, 0xbd, 0x1d, 0x98, 0x85, 0xe0, 0x6e, 0xc4,
	0xc0, 0x9e, 0xae, 0x44, 0x3c, 0x78, 0x90, 0x63, 0x26, 0xd4, 0x1d, 0xf9, 0x21, 0xb0, 0x67, 0x76,
	0x3c, 0xf6, 0xa4, 0x80, 0xdb, 0x13, 0x10, 0xf4, 0x3a, 0x2b, 0xfb, 0x26, 0x2b, 0x96, 0x84, 0xc2,
	0x9f, 0xc0, 0x8a, 0xc1, 0x7b, 0x2f, 0x33, 0xe0, 0xee, 0x08, 0x26, 0x74, 0xc5, 0xee, 0xfe, 0x4d,
	0x76, 0x89, 0xf0, 0x03, 0xdb, 0x0f, 0x05, 0x17, 0x6c, 0xd9, 0xc8, 0xfc, 0xc3, 0x40, 0x77, 0x4f,
	0xfb, 0xfd, 0xee, 0x59, 0xe8, 0x31, 0xe0, 0xbc, 0x4b, 0xc5, 0x08, 0x1f, 0xa2, 0xcd, 0x98, 0x8a,
	0x51, 0xc3, 0x38, 0x34, 0x8e, 0xb6, 0x9d, 0xfa, 0xc5, 0x65, 0x6b, 0x63, 0x76, 0xd9, 0xda, 0x94,
	0x3a, 0xa2, 0x34, 0xf8, 0x01, 0xaa, 0xca, 0xdf, 0xfe, 0xb3, 0x18, 0x1a, 0x65, 0x85, 0x6a, 0xcc,
	0x2e, 0x5b, 0xd5, 0xae, 0x96, 0xfd, 0x53, 0xf8, 0x4f, 0x32, 0x24, 0xfe, 0x12, 0x6d, 0x0d, 0xa8,
	0x3b, 0x86, 0x70, 0xd8, 0x28, 0x1d, 0x1a, 0x47, 0xb5, 0xf6, 0xb1, 0xf5, 0x9f, 0x3d, 0xb4, 0xf4,
	0xa5, 0x9c, 0xd4, 0xc8, 0xb9, 0xab, 0x6f, 0xb2, 0xa5, 0x05, 0x64, 0xee, 0xce, 0x1c, 0xa3, 0xfd,
	0x42, 0x12, 0x24, 0x09, 0xe0, 0x73, 0x1a, 0x24, 0x80, 0x7b, 0xa8, 0x22, 0xa3, 0xf3, 0x86, 0x71,
	0x58, 0x3e, 0xaa, 0xb5, 0xad, 0x97, 0xc4, 0x5b, 0x2a, 0x84, 0xb3, 0xa3, 0x03, 0x56, 0xe4, 0x89,
	0x93, 0xd4, 0x97, 0xf9, 0x63, 0x09, 0x6d, 0x69, 0x14, 0x7e, 0x82, 0xaa, 0xb2, 0xef, 0x43, 0x2a,
	0xa8, 0x2a, 0x57, 0xad, 0xfd, 0x4e, 0x21, 0x46, 0xd6, 0x06, 0x2b, 0x1e, 0x7b, 0x52, 0xc0, 0x2d,
	0x89, 0xb6, 0xa6, 0x27, 0xd6, 0xf9, 0xe0, 0x1b, 0x70, 0xc5, 0xa7, 0x20, 0xa8, 0x83, 0x75, 0x14,
	0x94, 0xcb, 0x48, 0xe6, 0x15, 0x77, 0xd0, 0x26, 0x8f, 0xc1, 0xd5, 0x15, 0x7b, 0x6b, 0xbd, 0x8a,
	0xf5, 0x62, 0x70, 0xf3, 0xc6, 0xc9, 0x13, 0x51, 0x5e, 0x70, 0x1f, 0xdd, 0xe1, 0x82, 0x8a, 0x84,
	0xab, 0xb6, 0xd5, 0xda, 0x6f, 0xaf, 0xe9, 0x4f, 0xd9, 0x38, 0xbb, 0xda, 0xe3, 0x9d, 0xf4, 0x4c,
	0xb4, 0x2f, 0xf3, 0x87, 0x12, 0xda, 0x5d, 0xec, 0x15, 0x7e, 0x17, 0xd5, 0x38, 0xb0, 0xa9, 0xef,
	0xc2, 0x63, 0x3a, 0x01, 0x4d, 0xa5, 0xd7, 0xb4, 0x7d, 0xad, 0x97, 0xab, 0x48, 0x11, 0x87, 0xbd,
	0xcc, 0xac, 0x1b, 0x31, 0xa1, 0x93, 0xbe, 0xb9, 0xa4, 0x92, 0xd9, 0x56, 0xca, 0x6c, 0xeb, 0x2c,
	0x14, 0xe7, 0xac, 0x27, 0x98, 0x1f, 0x7a, 0x2b, 0x81, 0xa4, 0x33, 0x52, 0xf4, 0x8c, 0xbf, 0x40,
	0x55, 0x06, 0x3c, 0x4a, 0x98, 0x0b, 0xba, 0x14, 0x0b, 0x64, 0x94, 0x23, 0x40, 0xb6, 0x49, 0xf2,
	0x76, 0xd8, 0x89, 0x5c, 0x1a, 0xa4, 0xcd, 0x21, 0xf0, 0x14, 0x18, 0x84, 0x2e, 0x38, 0x75, 0x49,
	0x78, 0xa2, 0x5d, 0x90, 0xcc, 0x99, 0x7c, 0x50, 0x75, 0x5d, 0x8b, 0x87, 0x01, 0x7d, 0x25, 0x14,
	0xf9, 0x6c, 0x81, 0x22, 0xf6, 0x7a, 0x2d, 0x55, 0x97, 0xbb, 0x89, 0x27, 0xe6, 0xef, 0x06, 0xda,
	0x2b, 0x02, 0x3b, 0x3e, 0x17, 0xf8, 0xeb, 0x95, 0x4c, 0xac, 0xf5, 0x32, 0x91, 0xd6, 0x2a, 0x8f,
	0x3d, 0x1d, 0xaa, 0x3a, 0x97, 0x14, 0xb2, 0xe8, 0xa2, 0x8a, 0x2f, 0x60, 0xc2, 0x1b, 0x25, 0xf5,
	0x56, 0xef, 0xdd, 0x22, 0x8d, 0xfc, 0xa1, 0x9e, 0x49, 0x0f, 0x24, 0x75, 0x64, 0xfe, 0x69, 0xa0,
	0x56, 0x11, 0xd6, 0xa5, 0x8c, 0x4e, 0x40, 0x00, 0xe3, 0x59, 0x1b, 0xf1, 0x11, 0xaa, 0xd2, 0xee,
	0xd9, 0x27, 0x2c, 0x4a, 0xe2, 0xf9, 0xbc, 0x93, 0xf7, 0xfb, 0x48, 0xcb, 0x48, 0xa6, 0x95, 0x53,
	0x71, 0xec, 0xeb, 0xd1, 0x55, 0x98, 0x8a, 0x8f, 0xfc, 0x70, 0x48, 0x94, 0x46, 0x22, 0x42, 0x49,
	0xf6, 0xf2, 0x22, 0x42, 0xb1, 0x5c, 0x69, 0x70, 0x0b, 0x55, 0xb8, 0x1b, 0xc5, 0xd0, 0xd8, 0x54,
	0x90, 0x6d, 0x79, 0xe5, 0x9e, 0x14, 0x90, 0x54, 0x8e, 0xef, 0xa1, 0x6d, 0x09, 0xe4, 0x31, 0x75,
	0xa1, 0x51, 0x51, 0xa0, 0x9d, 0xd9, 0x65, 0x6b, 0xfb, 0xf1, 0x5c, 0x48, 0x72, 0xbd, 0xf9, 0xcb,
	0x52, 0x93, 0x64, 0xff, 0x70, 0x1b, 0x21, 0x37, 0x0a, 0x05, 0x8b, 0x82, 0x00, 0x98, 0x4e, 0x29,
	0xa3, 0xcf, 0xc3, 0x4c, 0x43, 0x0a, 0x28, 0x1c, 0x22, 0x14, 0x67, 0xb5, 0xd1, 0x34, 0xfa, 0xe0,
	0x16, 0xf5, 0xbf, 0xa6, 0xb0, 0xce, 0xae, 0x8c, 0x57, 0x50, 0x14, 0x22, 0x98, 0xbf, 0x1a, 0xa8,
	0xa6, 0xed, 0x5f, 0x01, 0xb1, 0x1e, 0x2d, 0x12, 0xeb, 0x8d, 0x35, 0x97, 0xce, 0xf5, 0x9c, 0xfa,
	0x39, 0xbf, 0xba, 0x5c, 0x33, 0xb2, 0xe7, 0xa3, 0x88, 0x8b, 0xe5, 0x5d, 0x79, 0x1a, 0x71, 0x41,
	0x94, 0x06, 0x27, 0x68, 0xcf, 0x5f, 0xda, 0x4b, 0xb7, 0x7b, 0xa9, 0x99, 0x99, 0xd3, 0xd0, 0xee,
	0xf7, 0x96, 0x35, 0x64, 0x25, 0x84, 0x09, 0x68, 0x05, 0x25, 0x07, 0xc5, 0x48, 0x88, 0x58, 0xd7,
	0xf8, 0xfe, 0xfa, 0xdb, 0x30, 0xbf, 0x42, 0x55, 0x65, 0xd7, 0xef, 0x77, 0x89, 0x72, 0x65, 0xfe,
	0x56, 0xca, 0xea, 0xa1, 0xe8, 0xf7, 0x61, 0x96, 0xad, 0x62, 0x86, 0x1a, 0xfe, 0x29, 0xd9, 0xf7,
	0x0b, 0x17, 0xcf, 0x74, 0x64, 0x05, 0x8d, 0xfb, 0xf9, 0x57, 0x82, 0xf1, 0x7f, 0xbe, 0x12, 0x6a,
	0xd7, 0x7d, 0x21, 0xe0, 0x53, 0x54, 0x16, 0xc1, 0x9c, 0x02, 0x6f, 0xae, 0xe7, 0xb1, 0xdf, 0xe9,
	0x39, 0x35, 0x5d, 0xf2, 0x72, 0xbf, 0xd3, 0x23, 0xd2, 0x05, 0x3e, 0x47, 0x15, 0x96, 0x04, 0x20,
	0x37, 0x68, 0x79, 0xfd, 0x8d, 0x2c, 0x2b, 0x98, 0x53, 0x4a, 0x9e, 0x38, 0x49, 0xfd, 0x98, 0xdf,
	0xa2, 0x9d, 0x85, 0x35, 0x8b, 0x9f, 0xa0, 0x7a, 0x10, 0xd1, 0xa1, 0x43, 0x03, 0x1a, 0xba, 0xfa,
	0x11, 0x2f, 0xf1, 0x76, 0xbe, 0x9f, 0x3a, 0x05, 0x9c, 0x5e, 0xd2, 0xfb, 0x3a, 0x48, 0xbd, 0xa8,
	0x23, 0x0b, 0x1e, 0x4d, 0x8a, 0x50, 0x9e, 0xa3, 0x9c, 0x4a, 0x92, 0xa9, 0xe9, 0x57, 0x92, 0x9e,
	0x4a, 0x92, 0xc0, 0x9c, 0xa4, 0x72, 0x39, 0x53, 0x38, 0xb8, 0x0c, 0x84, 0x6a, 0x67, 0x69, 0x71,
	0xa6, 0xf4, 0x32, 0x0d, 0x29, 0xa0, 0x9c, 0xe3, 0x8b, 0xab, 0xe6, 0xc6, 0xf3, 0xab, 0xe6, 0xc6,
	0x8b, 0xab, 0xe6, 0xc6, 0xf7, 0xb3, 0xa6, 0x71, 0x31, 0x6b, 0x1a, 0xcf, 0x67, 0x4d, 0xe3, 0xc5,
	0xac, 0x69, 0xfc, 0x35, 0x6b, 0x1a, 0x3f, 0xfd, 0xdd, 0xdc, 0xf8, 0x6a, 0x4b, 0x97, 0xe9, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x8b, 0x3b, 0x2e, 0x16, 0x0c, 0x00, 0x00,
}

func (m *HTTPIngressPath) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HTTPIngressPath) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HTTPIngressPath) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PathType != nil {
		i -= len(*m.PathType)
		copy(dAtA[i:], *m.PathType)
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.PathType)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.Backend.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	i -= len(m.Path)
	copy(dAtA[i:], m.Path)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Path)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *HTTPIngressRuleValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HTTPIngressRuleValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HTTPIngressRuleValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Paths) > 0 {
		for iNdEx := len(m.Paths) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Paths[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Ingress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ingress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Ingress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Status.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Spec.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ObjectMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *IngressBackend) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IngressBackend) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IngressBackend) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Resource != nil {
		{
			size, err := m.Resource.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.ServicePort.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	i -= len(m.ServiceName)
	copy(dAtA[i:], m.ServiceName)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.ServiceName)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *IngressClass) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IngressClass) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IngressClass) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Spec.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ObjectMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *IngressClassList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IngressClassList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IngressClassList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.ListMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *IngressClassParametersReference) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IngressClassParametersReference) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IngressClassParametersReference) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Namespace != nil {
		i -= len(*m.Namespace)
		copy(dAtA[i:], *m.Namespace)
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.Namespace)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Scope != nil {
		i -= len(*m.Scope)
		copy(dAtA[i:], *m.Scope)
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.Scope)))
		i--
		dAtA[i] = 0x22
	}
	i -= len(m.Name)
	copy(dAtA[i:], m.Name)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Name)))
	i--
	dAtA[i] = 0x1a
	i -= len(m.Kind)
	copy(dAtA[i:], m.Kind)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Kind)))
	i--
	dAtA[i] = 0x12
	if m.APIGroup != nil {
		i -= len(*m.APIGroup)
		copy(dAtA[i:], *m.APIGroup)
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.APIGroup)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IngressClassSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IngressClassSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IngressClassSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Parameters != nil {
		{
			size, err := m.Parameters.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	i -= len(m.Controller)
	copy(dAtA[i:], m.Controller)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Controller)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *IngressList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IngressList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IngressList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.ListMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *IngressRule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IngressRule) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IngressRule) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.IngressRuleValue.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	i -= len(m.Host)
	copy(dAtA[i:], m.Host)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Host)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *IngressRuleValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IngressRuleValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IngressRuleValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.HTTP != nil {
		{
			size, err := m.HTTP.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IngressSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IngressSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IngressSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IngressClassName != nil {
		i -= len(*m.IngressClassName)
		copy(dAtA[i:], *m.IngressClassName)
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.IngressClassName)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Rules) > 0 {
		for iNdEx := len(m.Rules) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rules[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.TLS) > 0 {
		for iNdEx := len(m.TLS) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TLS[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Backend != nil {
		{
			size, err := m.Backend.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IngressStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IngressStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IngressStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.LoadBalancer.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *IngressTLS) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IngressTLS) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IngressTLS) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.SecretName)
	copy(dAtA[i:], m.SecretName)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.SecretName)))
	i--
	dAtA[i] = 0x12
	if len(m.Hosts) > 0 {
		for iNdEx := len(m.Hosts) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Hosts[iNdEx])
			copy(dAtA[i:], m.Hosts[iNdEx])
			i = encodeVarintGenerated(dAtA, i, uint64(len(m.Hosts[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenerated(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HTTPIngressPath) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Backend.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if m.PathType != nil {
		l = len(*m.PathType)
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *HTTPIngressRuleValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Paths) > 0 {
		for _, e := range m.Paths {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *Ingress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *IngressBackend) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ServiceName)
	n += 1 + l + sovGenerated(uint64(l))
	l = m.ServicePort.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if m.Resource != nil {
		l = m.Resource.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *IngressClass) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *IngressClassList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *IngressClassParametersReference) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.APIGroup != nil {
		l = len(*m.APIGroup)
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.Kind)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Name)
	n += 1 + l + sovGenerated(uint64(l))
	if m.Scope != nil {
		l = len(*m.Scope)
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Namespace != nil {
		l = len(*m.Namespace)
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *IngressClassSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Controller)
	n += 1 + l + sovGenerated(uint64(l))
	if m.Parameters != nil {
		l = m.Parameters.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *IngressList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *IngressRule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Host)
	n += 1 + l + sovGenerated(uint64(l))
	l = m.IngressRuleValue.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *IngressRuleValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HTTP != nil {
		l = m.HTTP.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *IngressSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Backend != nil {
		l = m.Backend.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if len(m.TLS) > 0 {
		for _, e := range m.TLS {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if len(m.Rules) > 0 {
		for _, e := range m.Rules {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if m.IngressClassName != nil {
		l = len(*m.IngressClassName)
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *IngressStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.LoadBalancer.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *IngressTLS) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Hosts) > 0 {
		for _, s := range m.Hosts {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	l = len(m.SecretName)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *HTTPIngressPath) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HTTPIngressPath{`,
		`Path:` + fmt.Sprintf("%v", this.Path) + `,`,
		`Backend:` + strings.Replace(strings.Replace(this.Backend.String(), "IngressBackend", "IngressBackend", 1), `&`, ``, 1) + `,`,
		`PathType:` + valueToStringGenerated(this.PathType) + `,`,
		`}`,
	}, "")
	return s
}
func (this *HTTPIngressRuleValue) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForPaths := "[]HTTPIngressPath{"
	for _, f := range this.Paths {
		repeatedStringForPaths += strings.Replace(strings.Replace(f.String(), "HTTPIngressPath", "HTTPIngressPath", 1), `&`, ``, 1) + ","
	}
	repeatedStringForPaths += "}"
	s := strings.Join([]string{`&HTTPIngressRuleValue{`,
		`Paths:` + repeatedStringForPaths + `,`,
		`}`,
	}, "")
	return s
}
func (this *Ingress) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Ingress{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ObjectMeta), "ObjectMeta", "v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "IngressSpec", "IngressSpec", 1), `&`, ``, 1) + `,`,
		`Status:` + strings.Replace(strings.Replace(this.Status.String(), "IngressStatus", "IngressStatus", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *IngressBackend) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&IngressBackend{`,
		`ServiceName:` + fmt.Sprintf("%v", this.ServiceName) + `,`,
		`ServicePort:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ServicePort), "IntOrString", "intstr.IntOrString", 1), `&`, ``, 1) + `,`,
		`Resource:` + strings.Replace(fmt.Sprintf("%v", this.Resource), "TypedLocalObjectReference", "v11.TypedLocalObjectReference", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *IngressClass) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&IngressClass{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ObjectMeta), "ObjectMeta", "v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "IngressClassSpec", "IngressClassSpec", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *IngressClassList) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForItems := "[]IngressClass{"
	for _, f := range this.Items {
		repeatedStringForItems += strings.Replace(strings.Replace(f.String(), "IngressClass", "IngressClass", 1), `&`, ``, 1) + ","
	}
	repeatedStringForItems += "}"
	s := strings.Join([]string{`&IngressClassList{`,
		`ListMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ListMeta), "ListMeta", "v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + repeatedStringForItems + `,`,
		`}`,
	}, "")
	return s
}
func (this *IngressClassParametersReference) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&IngressClassParametersReference{`,
		`APIGroup:` + valueToStringGenerated(this.APIGroup) + `,`,
		`Kind:` + fmt.Sprintf("%v", this.Kind) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Scope:` + valueToStringGenerated(this.Scope) + `,`,
		`Namespace:` + valueToStringGenerated(this.Namespace) + `,`,
		`}`,
	}, "")
	return s
}
func (this *IngressClassSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&IngressClassSpec{`,
		`Controller:` + fmt.Sprintf("%v", this.Controller) + `,`,
		`Parameters:` + strings.Replace(this.Parameters.String(), "IngressClassParametersReference", "IngressClassParametersReference", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *IngressList) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForItems := "[]Ingress{"
	for _, f := range this.Items {
		repeatedStringForItems += strings.Replace(strings.Replace(f.String(), "Ingress", "Ingress", 1), `&`, ``, 1) + ","
	}
	repeatedStringForItems += "}"
	s := strings.Join([]string{`&IngressList{`,
		`ListMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ListMeta), "ListMeta", "v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + repeatedStringForItems + `,`,
		`}`,
	}, "")
	return s
}
func (this *IngressRule) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&IngressRule{`,
		`Host:` + fmt.Sprintf("%v", this.Host) + `,`,
		`IngressRuleValue:` + strings.Replace(strings.Replace(this.IngressRuleValue.String(), "IngressRuleValue", "IngressRuleValue", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *IngressRuleValue) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&IngressRuleValue{`,
		`HTTP:` + strings.Replace(this.HTTP.String(), "HTTPIngressRuleValue", "HTTPIngressRuleValue", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *IngressSpec) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForTLS := "[]IngressTLS{"
	for _, f := range this.TLS {
		repeatedStringForTLS += strings.Replace(strings.Replace(f.String(), "IngressTLS", "IngressTLS", 1), `&`, ``, 1) + ","
	}
	repeatedStringForTLS += "}"
	repeatedStringForRules := "[]IngressRule{"
	for _, f := range this.Rules {
		repeatedStringForRules += strings.Replace(strings.Replace(f.String(), "IngressRule", "IngressRule", 1), `&`, ``, 1) + ","
	}
	repeatedStringForRules += "}"
	s := strings.Join([]string{`&IngressSpec{`,
		`Backend:` + strings.Replace(this.Backend.String(), "IngressBackend", "IngressBackend", 1) + `,`,
		`TLS:` + repeatedStringForTLS + `,`,
		`Rules:` + repeatedStringForRules + `,`,
		`IngressClassName:` + valueToStringGenerated(this.IngressClassName) + `,`,
		`}`,
	}, "")
	return s
}
func (this *IngressStatus) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&IngressStatus{`,
		`LoadBalancer:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.LoadBalancer), "LoadBalancerStatus", "v11.LoadBalancerStatus", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *IngressTLS) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&IngressTLS{`,
		`Hosts:` + fmt.Sprintf("%v", this.Hosts) + `,`,
		`SecretName:` + fmt.Sprintf("%v", this.SecretName) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *HTTPIngressPath) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: HTTPIngressPath: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HTTPIngressPath: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Backend", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Backend.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PathType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := PathType(dAtA[iNdEx:postIndex])
			m.PathType = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *HTTPIngressRuleValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: HTTPIngressRuleValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HTTPIngressRuleValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Paths", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Paths = append(m.Paths, HTTPIngressPath{})
			if err := m.Paths[len(m.Paths)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *Ingress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: Ingress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ingress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *IngressBackend) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: IngressBackend: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IngressBackend: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServicePort", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ServicePort.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resource", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Resource == nil {
				m.Resource = &v11.TypedLocalObjectReference{}
			}
			if err := m.Resource.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *IngressClass) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: IngressClass: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IngressClass: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *IngressClassList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: IngressClassList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IngressClassList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, IngressClass{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *IngressClassParametersReference) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: IngressClassParametersReference: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IngressClassParametersReference: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field APIGroup", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.APIGroup = &s
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Kind = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scope", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Scope = &s
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Namespace = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *IngressClassSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: IngressClassSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IngressClassSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Controller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Controller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parameters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Parameters == nil {
				m.Parameters = &IngressClassParametersReference{}
			}
			if err := m.Parameters.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *IngressList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: IngressList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IngressList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, Ingress{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *IngressRule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: IngressRule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IngressRule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Host", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Host = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IngressRuleValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.IngressRuleValue.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *IngressRuleValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: IngressRuleValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IngressRuleValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HTTP", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HTTP == nil {
				m.HTTP = &HTTPIngressRuleValue{}
			}
			if err := m.HTTP.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *IngressSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: IngressSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IngressSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Backend", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Backend == nil {
				m.Backend = &IngressBackend{}
			}
			if err := m.Backend.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TLS", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TLS = append(m.TLS, IngressTLS{})
			if err := m.TLS[len(m.TLS)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rules = append(m.Rules, IngressRule{})
			if err := m.Rules[len(m.Rules)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IngressClassName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.IngressClassName = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *IngressStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: IngressStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IngressStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LoadBalancer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LoadBalancer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *IngressTLS) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: IngressTLS: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IngressTLS: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hosts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hosts = append(m.Hosts, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecretName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SecretName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenerated
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenerated
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenerated        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenerated = fmt.Errorf("proto: unexpected end of group")
)
