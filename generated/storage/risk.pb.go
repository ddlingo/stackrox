// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/risk.proto

package storage

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// Next tag: 9
type RiskSubjectType int32

const (
	RiskSubjectType_UNKNOWN         RiskSubjectType = 0
	RiskSubjectType_DEPLOYMENT      RiskSubjectType = 1
	RiskSubjectType_NAMESPACE       RiskSubjectType = 2
	RiskSubjectType_CLUSTER         RiskSubjectType = 3
	RiskSubjectType_NODE            RiskSubjectType = 7
	RiskSubjectType_NODE_COMPONENT  RiskSubjectType = 8
	RiskSubjectType_IMAGE           RiskSubjectType = 4
	RiskSubjectType_IMAGE_COMPONENT RiskSubjectType = 6
	RiskSubjectType_SERVICEACCOUNT  RiskSubjectType = 5
)

var RiskSubjectType_name = map[int32]string{
	0: "UNKNOWN",
	1: "DEPLOYMENT",
	2: "NAMESPACE",
	3: "CLUSTER",
	7: "NODE",
	8: "NODE_COMPONENT",
	4: "IMAGE",
	6: "IMAGE_COMPONENT",
	5: "SERVICEACCOUNT",
}

var RiskSubjectType_value = map[string]int32{
	"UNKNOWN":         0,
	"DEPLOYMENT":      1,
	"NAMESPACE":       2,
	"CLUSTER":         3,
	"NODE":            7,
	"NODE_COMPONENT":  8,
	"IMAGE":           4,
	"IMAGE_COMPONENT": 6,
	"SERVICEACCOUNT":  5,
}

func (x RiskSubjectType) String() string {
	return proto.EnumName(RiskSubjectType_name, int32(x))
}

func (RiskSubjectType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_76ed3b6b4b68c2b1, []int{0}
}

type Risk struct {
	Id                   string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Subject              *RiskSubject   `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Score                float32        `protobuf:"fixed32,3,opt,name=score,proto3" json:"score,omitempty" search:"Risk Score,hidden"`
	Results              []*Risk_Result `protobuf:"bytes,4,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Risk) Reset()         { *m = Risk{} }
func (m *Risk) String() string { return proto.CompactTextString(m) }
func (*Risk) ProtoMessage()    {}
func (*Risk) Descriptor() ([]byte, []int) {
	return fileDescriptor_76ed3b6b4b68c2b1, []int{0}
}
func (m *Risk) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Risk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Risk.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Risk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Risk.Merge(m, src)
}
func (m *Risk) XXX_Size() int {
	return m.Size()
}
func (m *Risk) XXX_DiscardUnknown() {
	xxx_messageInfo_Risk.DiscardUnknown(m)
}

var xxx_messageInfo_Risk proto.InternalMessageInfo

func (m *Risk) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Risk) GetSubject() *RiskSubject {
	if m != nil {
		return m.Subject
	}
	return nil
}

func (m *Risk) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Risk) GetResults() []*Risk_Result {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *Risk) MessageClone() proto.Message {
	return m.Clone()
}
func (m *Risk) Clone() *Risk {
	if m == nil {
		return nil
	}
	cloned := new(Risk)
	*cloned = *m

	cloned.Subject = m.Subject.Clone()
	if m.Results != nil {
		cloned.Results = make([]*Risk_Result, len(m.Results))
		for idx, v := range m.Results {
			cloned.Results[idx] = v.Clone()
		}
	}
	return cloned
}

type Risk_Result struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Factors              []*Risk_Result_Factor `protobuf:"bytes,2,rep,name=factors,proto3" json:"factors,omitempty"`
	Score                float32               `protobuf:"fixed32,3,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Risk_Result) Reset()         { *m = Risk_Result{} }
func (m *Risk_Result) String() string { return proto.CompactTextString(m) }
func (*Risk_Result) ProtoMessage()    {}
func (*Risk_Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_76ed3b6b4b68c2b1, []int{0, 0}
}
func (m *Risk_Result) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Risk_Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Risk_Result.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Risk_Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Risk_Result.Merge(m, src)
}
func (m *Risk_Result) XXX_Size() int {
	return m.Size()
}
func (m *Risk_Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Risk_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Risk_Result proto.InternalMessageInfo

func (m *Risk_Result) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Risk_Result) GetFactors() []*Risk_Result_Factor {
	if m != nil {
		return m.Factors
	}
	return nil
}

func (m *Risk_Result) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Risk_Result) MessageClone() proto.Message {
	return m.Clone()
}
func (m *Risk_Result) Clone() *Risk_Result {
	if m == nil {
		return nil
	}
	cloned := new(Risk_Result)
	*cloned = *m

	if m.Factors != nil {
		cloned.Factors = make([]*Risk_Result_Factor, len(m.Factors))
		for idx, v := range m.Factors {
			cloned.Factors[idx] = v.Clone()
		}
	}
	return cloned
}

type Risk_Result_Factor struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Risk_Result_Factor) Reset()         { *m = Risk_Result_Factor{} }
func (m *Risk_Result_Factor) String() string { return proto.CompactTextString(m) }
func (*Risk_Result_Factor) ProtoMessage()    {}
func (*Risk_Result_Factor) Descriptor() ([]byte, []int) {
	return fileDescriptor_76ed3b6b4b68c2b1, []int{0, 0, 0}
}
func (m *Risk_Result_Factor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Risk_Result_Factor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Risk_Result_Factor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Risk_Result_Factor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Risk_Result_Factor.Merge(m, src)
}
func (m *Risk_Result_Factor) XXX_Size() int {
	return m.Size()
}
func (m *Risk_Result_Factor) XXX_DiscardUnknown() {
	xxx_messageInfo_Risk_Result_Factor.DiscardUnknown(m)
}

var xxx_messageInfo_Risk_Result_Factor proto.InternalMessageInfo

func (m *Risk_Result_Factor) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Risk_Result_Factor) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Risk_Result_Factor) MessageClone() proto.Message {
	return m.Clone()
}
func (m *Risk_Result_Factor) Clone() *Risk_Result_Factor {
	if m == nil {
		return nil
	}
	cloned := new(Risk_Result_Factor)
	*cloned = *m

	return cloned
}

type RiskSubject struct {
	Id                   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Namespace            string          `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty" search:"Namespace,store"`
	ClusterId            string          `protobuf:"bytes,3,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty" search:"Cluster ID,store,hidden"`
	Type                 RiskSubjectType `protobuf:"varint,4,opt,name=type,proto3,enum=storage.RiskSubjectType" json:"type,omitempty" search:"Risk Subject Type,hidden"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RiskSubject) Reset()         { *m = RiskSubject{} }
func (m *RiskSubject) String() string { return proto.CompactTextString(m) }
func (*RiskSubject) ProtoMessage()    {}
func (*RiskSubject) Descriptor() ([]byte, []int) {
	return fileDescriptor_76ed3b6b4b68c2b1, []int{1}
}
func (m *RiskSubject) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RiskSubject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RiskSubject.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RiskSubject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RiskSubject.Merge(m, src)
}
func (m *RiskSubject) XXX_Size() int {
	return m.Size()
}
func (m *RiskSubject) XXX_DiscardUnknown() {
	xxx_messageInfo_RiskSubject.DiscardUnknown(m)
}

var xxx_messageInfo_RiskSubject proto.InternalMessageInfo

func (m *RiskSubject) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RiskSubject) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *RiskSubject) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *RiskSubject) GetType() RiskSubjectType {
	if m != nil {
		return m.Type
	}
	return RiskSubjectType_UNKNOWN
}

func (m *RiskSubject) MessageClone() proto.Message {
	return m.Clone()
}
func (m *RiskSubject) Clone() *RiskSubject {
	if m == nil {
		return nil
	}
	cloned := new(RiskSubject)
	*cloned = *m

	return cloned
}

func init() {
	proto.RegisterEnum("storage.RiskSubjectType", RiskSubjectType_name, RiskSubjectType_value)
	proto.RegisterType((*Risk)(nil), "storage.Risk")
	proto.RegisterType((*Risk_Result)(nil), "storage.Risk.Result")
	proto.RegisterType((*Risk_Result_Factor)(nil), "storage.Risk.Result.Factor")
	proto.RegisterType((*RiskSubject)(nil), "storage.RiskSubject")
}

func init() { proto.RegisterFile("storage/risk.proto", fileDescriptor_76ed3b6b4b68c2b1) }

var fileDescriptor_76ed3b6b4b68c2b1 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xed, 0x38, 0x6e, 0x5c, 0xdf, 0xe8, 0x4b, 0xad, 0xfb, 0x65, 0x61, 0x02, 0x4a, 0x4c, 0x04,
	0x52, 0x84, 0x2a, 0x57, 0x2a, 0x61, 0xd3, 0x5d, 0xe2, 0x18, 0x14, 0xd1, 0xd8, 0xd1, 0x38, 0x01,
	0xc1, 0xa6, 0x72, 0xed, 0x21, 0x35, 0x49, 0xeb, 0xc8, 0xe3, 0x48, 0xf4, 0x4d, 0x58, 0x20, 0xf1,
	0x3a, 0xac, 0x10, 0x4f, 0x10, 0xa1, 0xf0, 0x06, 0xe1, 0x05, 0xd0, 0xd8, 0x71, 0x69, 0x69, 0x77,
	0xd7, 0xf7, 0xfc, 0xdc, 0xa3, 0x63, 0x0d, 0x20, 0x4f, 0xe3, 0xc4, 0x9f, 0xb2, 0xc3, 0x24, 0xe2,
	0x33, 0x73, 0x91, 0xc4, 0x69, 0x8c, 0xca, 0x76, 0x57, 0xaf, 0x4d, 0xe3, 0x69, 0x9c, 0xed, 0x0e,
	0xc5, 0x94, 0xc3, 0xad, 0xef, 0x12, 0xc8, 0x34, 0xe2, 0x33, 0xac, 0x82, 0x14, 0x85, 0x3a, 0x31,
	0x48, 0x5b, 0xa5, 0x52, 0x14, 0xa2, 0x09, 0x0a, 0x5f, 0x9e, 0x7d, 0x64, 0x41, 0xaa, 0x4b, 0x06,
	0x69, 0x57, 0x8e, 0x6a, 0xe6, 0xd6, 0xc9, 0x14, 0x7c, 0x2f, 0xc7, 0x68, 0x41, 0xc2, 0x0e, 0xec,
	0xf2, 0x20, 0x4e, 0x98, 0x5e, 0x32, 0x48, 0x5b, 0xea, 0x35, 0x36, 0xab, 0x66, 0x9d, 0x33, 0x3f,
	0x09, 0xce, 0x8f, 0x5b, 0x42, 0x60, 0x78, 0x02, 0x3d, 0x38, 0x8f, 0xc2, 0x90, 0x5d, 0xb6, 0x68,
	0x4e, 0x16, 0x57, 0x12, 0xc6, 0x97, 0xf3, 0x94, 0xeb, 0xb2, 0x51, 0xba, 0x73, 0xc5, 0xa4, 0x19,
	0x48, 0x0b, 0x52, 0xfd, 0x2b, 0x81, 0x72, 0xbe, 0x43, 0x04, 0xf9, 0xd2, 0xbf, 0x60, 0xdb, 0xc8,
	0xd9, 0x8c, 0x2f, 0x40, 0xf9, 0xe0, 0x07, 0x69, 0x9c, 0x70, 0x5d, 0xca, 0xec, 0x1e, 0xde, 0x67,
	0x67, 0xbe, 0xcc, 0x38, 0xb4, 0xe0, 0x62, 0xed, 0x56, 0xf6, 0x6d, 0xb6, 0x7a, 0x07, 0xca, 0x39,
	0x11, 0x75, 0x50, 0x2e, 0x18, 0xe7, 0xfe, 0xb4, 0xb8, 0x56, 0x7c, 0xa2, 0x06, 0xa5, 0x65, 0x32,
	0xcf, 0x1a, 0x52, 0xa9, 0x18, 0x5b, 0xbf, 0x09, 0x54, 0x6e, 0x14, 0x74, 0xa7, 0xd7, 0x63, 0x50,
	0x45, 0x54, 0xbe, 0xf0, 0x03, 0x96, 0xeb, 0x7a, 0x8f, 0x36, 0xab, 0xa6, 0x5e, 0x74, 0xe5, 0x14,
	0xe0, 0x81, 0x48, 0xce, 0x5a, 0xf4, 0x2f, 0x1d, 0x2d, 0x80, 0x60, 0xbe, 0xe4, 0x29, 0x4b, 0x4e,
	0xa3, 0x30, 0x0b, 0xab, 0xf6, 0x9e, 0x6c, 0x56, 0x4d, 0xa3, 0x10, 0x5b, 0x39, 0x6a, 0x0c, 0xfa,
	0xb9, 0xfa, 0xba, 0x6e, 0x75, 0xab, 0x1b, 0x84, 0xe8, 0x81, 0x9c, 0x5e, 0x2d, 0x98, 0x2e, 0x1b,
	0xa4, 0x5d, 0x3d, 0xd2, 0xef, 0xfb, 0xab, 0xe3, 0xab, 0x05, 0xeb, 0x3d, 0xdd, 0xac, 0x9a, 0x8f,
	0x6f, 0xff, 0xc1, 0x1c, 0x35, 0x04, 0x7c, 0xed, 0x9c, 0x99, 0x3d, 0xfb, 0x42, 0x60, 0xff, 0x1f,
	0x03, 0xac, 0x80, 0x32, 0x71, 0x5e, 0x3b, 0xee, 0x5b, 0x47, 0xdb, 0xc1, 0x2a, 0x40, 0xdf, 0x1e,
	0x9d, 0xb8, 0xef, 0x86, 0xb6, 0x33, 0xd6, 0x08, 0xfe, 0x07, 0xaa, 0xd3, 0x1d, 0xda, 0xde, 0xa8,
	0x6b, 0xd9, 0x9a, 0x24, 0xb8, 0xd6, 0xc9, 0xc4, 0x1b, 0xdb, 0x54, 0x2b, 0xe1, 0x1e, 0xc8, 0x8e,
	0xdb, 0xb7, 0x35, 0x05, 0x11, 0xaa, 0x62, 0x3a, 0xb5, 0xdc, 0xe1, 0xc8, 0x75, 0x84, 0x72, 0x0f,
	0x55, 0xd8, 0x1d, 0x0c, 0xbb, 0xaf, 0x6c, 0x4d, 0xc6, 0xff, 0x61, 0x3f, 0x1b, 0x6f, 0xe0, 0x65,
	0xa1, 0xf1, 0x6c, 0xfa, 0x66, 0x60, 0xd9, 0x5d, 0xcb, 0x72, 0x27, 0xce, 0x58, 0xdb, 0xed, 0x75,
	0xbe, 0xad, 0x1b, 0xe4, 0xc7, 0xba, 0x41, 0x7e, 0xae, 0x1b, 0xe4, 0xf3, 0xaf, 0xc6, 0x0e, 0x3c,
	0x88, 0x62, 0x93, 0xa7, 0x7e, 0x30, 0x4b, 0xe2, 0x4f, 0xf9, 0x53, 0x28, 0x8a, 0x78, 0x5f, 0xbc,
	0x98, 0xb3, 0x72, 0xb6, 0x7f, 0xfe, 0x27, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xfb, 0x85, 0xe6, 0x57,
	0x03, 0x00, 0x00,
}

func (m *Risk) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Risk) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Risk) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Results) > 0 {
		for iNdEx := len(m.Results) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Results[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRisk(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Score != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Score))))
		i--
		dAtA[i] = 0x1d
	}
	if m.Subject != nil {
		{
			size, err := m.Subject.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRisk(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintRisk(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Risk_Result) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Risk_Result) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Risk_Result) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Score != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Score))))
		i--
		dAtA[i] = 0x1d
	}
	if len(m.Factors) > 0 {
		for iNdEx := len(m.Factors) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Factors[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRisk(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintRisk(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Risk_Result_Factor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Risk_Result_Factor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Risk_Result_Factor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintRisk(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintRisk(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RiskSubject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RiskSubject) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RiskSubject) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Type != 0 {
		i = encodeVarintRisk(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ClusterId) > 0 {
		i -= len(m.ClusterId)
		copy(dAtA[i:], m.ClusterId)
		i = encodeVarintRisk(dAtA, i, uint64(len(m.ClusterId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintRisk(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintRisk(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRisk(dAtA []byte, offset int, v uint64) int {
	offset -= sovRisk(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Risk) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovRisk(uint64(l))
	}
	if m.Subject != nil {
		l = m.Subject.Size()
		n += 1 + l + sovRisk(uint64(l))
	}
	if m.Score != 0 {
		n += 5
	}
	if len(m.Results) > 0 {
		for _, e := range m.Results {
			l = e.Size()
			n += 1 + l + sovRisk(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Risk_Result) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRisk(uint64(l))
	}
	if len(m.Factors) > 0 {
		for _, e := range m.Factors {
			l = e.Size()
			n += 1 + l + sovRisk(uint64(l))
		}
	}
	if m.Score != 0 {
		n += 5
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Risk_Result_Factor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovRisk(uint64(l))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovRisk(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RiskSubject) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovRisk(uint64(l))
	}
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovRisk(uint64(l))
	}
	l = len(m.ClusterId)
	if l > 0 {
		n += 1 + l + sovRisk(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovRisk(uint64(m.Type))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRisk(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRisk(x uint64) (n int) {
	return sovRisk(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Risk) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRisk
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
			return fmt.Errorf("proto: Risk: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Risk: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRisk
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
				return ErrInvalidLengthRisk
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRisk
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRisk
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
				return ErrInvalidLengthRisk
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRisk
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Subject == nil {
				m.Subject = &RiskSubject{}
			}
			if err := m.Subject.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Score", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Score = float32(math.Float32frombits(v))
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Results", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRisk
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
				return ErrInvalidLengthRisk
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRisk
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Results = append(m.Results, &Risk_Result{})
			if err := m.Results[len(m.Results)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRisk(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRisk
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Risk_Result) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRisk
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
			return fmt.Errorf("proto: Result: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Result: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRisk
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
				return ErrInvalidLengthRisk
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRisk
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Factors", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRisk
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
				return ErrInvalidLengthRisk
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRisk
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Factors = append(m.Factors, &Risk_Result_Factor{})
			if err := m.Factors[len(m.Factors)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Score", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Score = float32(math.Float32frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipRisk(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRisk
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Risk_Result_Factor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRisk
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
			return fmt.Errorf("proto: Factor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Factor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRisk
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
				return ErrInvalidLengthRisk
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRisk
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRisk
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
				return ErrInvalidLengthRisk
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRisk
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRisk(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRisk
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RiskSubject) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRisk
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
			return fmt.Errorf("proto: RiskSubject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RiskSubject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRisk
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
				return ErrInvalidLengthRisk
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRisk
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRisk
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
				return ErrInvalidLengthRisk
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRisk
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRisk
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
				return ErrInvalidLengthRisk
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRisk
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRisk
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= RiskSubjectType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRisk(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRisk
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRisk(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRisk
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
					return 0, ErrIntOverflowRisk
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
					return 0, ErrIntOverflowRisk
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
				return 0, ErrInvalidLengthRisk
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRisk
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRisk
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRisk        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRisk          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRisk = fmt.Errorf("proto: unexpected end of group")
)
