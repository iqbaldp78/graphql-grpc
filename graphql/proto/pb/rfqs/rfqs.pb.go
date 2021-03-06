// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/mbizmarket/dmp/graphql/proto/pb/rfqs/rfqs.proto

package rfqs

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Rfqs struct {
	ID                   int64                `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TransactionId        int64                `protobuf:"varint,2,opt,name=TransactionId,proto3" json:"TransactionId,omitempty"`
	RequestedBy          int64                `protobuf:"varint,3,opt,name=RequestedBy,proto3" json:"RequestedBy,omitempty"`
	CreatedBy            int64                `protobuf:"varint,4,opt,name=CreatedBy,proto3" json:"CreatedBy,omitempty"`
	PaymentType          int32                `protobuf:"varint,6,opt,name=PaymentType,proto3" json:"PaymentType,omitempty"`
	PaymentDuration      int64                `protobuf:"varint,7,opt,name=PaymentDuration,proto3" json:"PaymentDuration,omitempty"`
	CompanyBuyerId       int64                `protobuf:"varint,8,opt,name=CompanyBuyerId,proto3" json:"CompanyBuyerId,omitempty"`
	CompanySellerId      int64                `protobuf:"varint,9,opt,name=CompanySellerId,proto3" json:"CompanySellerId,omitempty"`
	RfqNo                string               `protobuf:"bytes,10,opt,name=RfqNo,proto3" json:"RfqNo,omitempty"`
	ReferenceNo          string               `protobuf:"bytes,11,opt,name=ReferenceNo,proto3" json:"ReferenceNo,omitempty"`
	Notes                string               `protobuf:"bytes,12,opt,name=Notes,proto3" json:"Notes,omitempty"`
	NoteForSeller        string               `protobuf:"bytes,13,opt,name=NoteForSeller,proto3" json:"NoteForSeller,omitempty"`
	SubTotal             float32              `protobuf:"fixed32,14,opt,name=SubTotal,proto3" json:"SubTotal,omitempty"`
	TaxBasis             float32              `protobuf:"fixed32,15,opt,name=TaxBasis,proto3" json:"TaxBasis,omitempty"`
	Ppn                  float32              `protobuf:"fixed32,16,opt,name=Ppn,proto3" json:"Ppn,omitempty"`
	Pph                  float32              `protobuf:"fixed32,17,opt,name=Pph,proto3" json:"Pph,omitempty"`
	Rounding             float32              `protobuf:"fixed32,18,opt,name=Rounding,proto3" json:"Rounding,omitempty"`
	GrandTotal           float32              `protobuf:"fixed32,19,opt,name=GrandTotal,proto3" json:"GrandTotal,omitempty"`
	Status               int32                `protobuf:"varint,20,opt,name=Status,proto3" json:"Status,omitempty"`
	StatusReason         string               `protobuf:"bytes,21,opt,name=StatusReason,proto3" json:"StatusReason,omitempty"`
	ExpiredAt            *timestamp.Timestamp `protobuf:"bytes,22,opt,name=ExpiredAt,proto3" json:"ExpiredAt,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,23,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,24,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	DeletedAt            *timestamp.Timestamp `protobuf:"bytes,25,opt,name=DeletedAt,proto3" json:"DeletedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Rfqs) Reset()         { *m = Rfqs{} }
func (m *Rfqs) String() string { return proto.CompactTextString(m) }
func (*Rfqs) ProtoMessage()    {}
func (*Rfqs) Descriptor() ([]byte, []int) {
	return fileDescriptor_3932bcfd2247a7ec, []int{0}
}

func (m *Rfqs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rfqs.Unmarshal(m, b)
}
func (m *Rfqs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rfqs.Marshal(b, m, deterministic)
}
func (m *Rfqs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rfqs.Merge(m, src)
}
func (m *Rfqs) XXX_Size() int {
	return xxx_messageInfo_Rfqs.Size(m)
}
func (m *Rfqs) XXX_DiscardUnknown() {
	xxx_messageInfo_Rfqs.DiscardUnknown(m)
}

var xxx_messageInfo_Rfqs proto.InternalMessageInfo

func (m *Rfqs) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Rfqs) GetTransactionId() int64 {
	if m != nil {
		return m.TransactionId
	}
	return 0
}

func (m *Rfqs) GetRequestedBy() int64 {
	if m != nil {
		return m.RequestedBy
	}
	return 0
}

func (m *Rfqs) GetCreatedBy() int64 {
	if m != nil {
		return m.CreatedBy
	}
	return 0
}

func (m *Rfqs) GetPaymentType() int32 {
	if m != nil {
		return m.PaymentType
	}
	return 0
}

func (m *Rfqs) GetPaymentDuration() int64 {
	if m != nil {
		return m.PaymentDuration
	}
	return 0
}

func (m *Rfqs) GetCompanyBuyerId() int64 {
	if m != nil {
		return m.CompanyBuyerId
	}
	return 0
}

func (m *Rfqs) GetCompanySellerId() int64 {
	if m != nil {
		return m.CompanySellerId
	}
	return 0
}

func (m *Rfqs) GetRfqNo() string {
	if m != nil {
		return m.RfqNo
	}
	return ""
}

func (m *Rfqs) GetReferenceNo() string {
	if m != nil {
		return m.ReferenceNo
	}
	return ""
}

func (m *Rfqs) GetNotes() string {
	if m != nil {
		return m.Notes
	}
	return ""
}

func (m *Rfqs) GetNoteForSeller() string {
	if m != nil {
		return m.NoteForSeller
	}
	return ""
}

func (m *Rfqs) GetSubTotal() float32 {
	if m != nil {
		return m.SubTotal
	}
	return 0
}

func (m *Rfqs) GetTaxBasis() float32 {
	if m != nil {
		return m.TaxBasis
	}
	return 0
}

func (m *Rfqs) GetPpn() float32 {
	if m != nil {
		return m.Ppn
	}
	return 0
}

func (m *Rfqs) GetPph() float32 {
	if m != nil {
		return m.Pph
	}
	return 0
}

func (m *Rfqs) GetRounding() float32 {
	if m != nil {
		return m.Rounding
	}
	return 0
}

func (m *Rfqs) GetGrandTotal() float32 {
	if m != nil {
		return m.GrandTotal
	}
	return 0
}

func (m *Rfqs) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Rfqs) GetStatusReason() string {
	if m != nil {
		return m.StatusReason
	}
	return ""
}

func (m *Rfqs) GetExpiredAt() *timestamp.Timestamp {
	if m != nil {
		return m.ExpiredAt
	}
	return nil
}

func (m *Rfqs) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Rfqs) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Rfqs) GetDeletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

type RespData struct {
	Data                 []*Rfqs  `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespData) Reset()         { *m = RespData{} }
func (m *RespData) String() string { return proto.CompactTextString(m) }
func (*RespData) ProtoMessage()    {}
func (*RespData) Descriptor() ([]byte, []int) {
	return fileDescriptor_3932bcfd2247a7ec, []int{1}
}

func (m *RespData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespData.Unmarshal(m, b)
}
func (m *RespData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespData.Marshal(b, m, deterministic)
}
func (m *RespData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespData.Merge(m, src)
}
func (m *RespData) XXX_Size() int {
	return xxx_messageInfo_RespData.Size(m)
}
func (m *RespData) XXX_DiscardUnknown() {
	xxx_messageInfo_RespData.DiscardUnknown(m)
}

var xxx_messageInfo_RespData proto.InternalMessageInfo

func (m *RespData) GetData() []*Rfqs {
	if m != nil {
		return m.Data
	}
	return nil
}

type Req struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Req) Reset()         { *m = Req{} }
func (m *Req) String() string { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()    {}
func (*Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_3932bcfd2247a7ec, []int{2}
}

func (m *Req) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Req.Unmarshal(m, b)
}
func (m *Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Req.Marshal(b, m, deterministic)
}
func (m *Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Req.Merge(m, src)
}
func (m *Req) XXX_Size() int {
	return xxx_messageInfo_Req.Size(m)
}
func (m *Req) XXX_DiscardUnknown() {
	xxx_messageInfo_Req.DiscardUnknown(m)
}

var xxx_messageInfo_Req proto.InternalMessageInfo

func (m *Req) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func init() {
	proto.RegisterType((*Rfqs)(nil), "rfqs.Rfqs")
	proto.RegisterType((*RespData)(nil), "rfqs.RespData")
	proto.RegisterType((*Req)(nil), "rfqs.Req")
}

func init() { proto.RegisterFile("github.com/mbizmarket/dmp/graphql/proto/pb/rfqs/rfqs.proto", fileDescriptor_3932bcfd2247a7ec) }

var fileDescriptor_3932bcfd2247a7ec = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0x41, 0x6f, 0xda, 0x3e,
	0x18, 0xc6, 0xff, 0x01, 0xca, 0xbf, 0xbc, 0x14, 0xe8, 0xbc, 0xb6, 0xf3, 0xd0, 0xd4, 0x45, 0x68,
	0x9a, 0xa2, 0x69, 0x02, 0xa9, 0xbb, 0xec, 0x36, 0x95, 0xb2, 0x55, 0x5c, 0x50, 0x65, 0xe8, 0x65,
	0x37, 0x43, 0x5e, 0x20, 0x12, 0xc4, 0x89, 0xed, 0x4c, 0xcd, 0x57, 0xdc, 0xa7, 0x9a, 0x6c, 0x27,
	0x14, 0xb8, 0x70, 0x81, 0xf7, 0xf9, 0x3d, 0xcf, 0x63, 0x2b, 0x60, 0x07, 0x68, 0x22, 0x85, 0x16,
	0x83, 0x64, 0x3e, 0x90, 0xcb, 0x54, 0xd9, 0x8f, 0xbe, 0x45, 0xa4, 0x66, 0xe6, 0xee, 0xc7, 0x95,
	0x10, 0xab, 0x0d, 0x0e, 0x2c, 0x9b, 0x67, 0xcb, 0x81, 0x8e, 0xb6, 0xa8, 0x34, 0xdf, 0x26, 0x2e,
	0xd6, 0xfb, 0x5b, 0x87, 0x1a, 0x5b, 0xa6, 0x8a, 0xb4, 0xa1, 0x32, 0x1e, 0x51, 0xcf, 0xf7, 0x82,
	0x2a, 0xab, 0x8c, 0x47, 0xe4, 0x13, 0xb4, 0x66, 0x92, 0xc7, 0x8a, 0x2f, 0x74, 0x24, 0xe2, 0x71,
	0x48, 0x2b, 0xd6, 0x3a, 0x84, 0xc4, 0x87, 0x26, 0xc3, 0x34, 0x43, 0xa5, 0x31, 0x1c, 0xe6, 0xb4,
	0x6a, 0x33, 0xfb, 0x88, 0x7c, 0x80, 0xc6, 0x83, 0x44, 0xee, 0xfc, 0x9a, 0xf5, 0x5f, 0x81, 0xe9,
	0x3f, 0xf1, 0x7c, 0x8b, 0xb1, 0x9e, 0xe5, 0x09, 0xd2, 0xba, 0xef, 0x05, 0x67, 0x6c, 0x1f, 0x91,
	0x00, 0x3a, 0x85, 0x1c, 0x65, 0x92, 0x9b, 0x6d, 0xe9, 0xff, 0x76, 0x95, 0x63, 0x4c, 0x3e, 0x43,
	0xfb, 0x41, 0x6c, 0x13, 0x1e, 0xe7, 0xc3, 0x2c, 0x47, 0x39, 0x0e, 0xe9, 0xb9, 0x0d, 0x1e, 0x51,
	0xb3, 0x62, 0x41, 0xa6, 0xb8, 0xd9, 0xd8, 0x60, 0xc3, 0xad, 0x78, 0x84, 0xc9, 0x15, 0x9c, 0xb1,
	0x65, 0x3a, 0x11, 0x14, 0x7c, 0x2f, 0x68, 0x30, 0x27, 0xdc, 0x33, 0x2f, 0x51, 0x62, 0xbc, 0xc0,
	0x89, 0xa0, 0x4d, 0xeb, 0xed, 0x23, 0xd3, 0x9b, 0x08, 0x8d, 0x8a, 0x5e, 0xb8, 0x9e, 0x15, 0xe6,
	0x17, 0x35, 0xc3, 0x2f, 0x21, 0xdd, 0x06, 0xb4, 0x65, 0xdd, 0x43, 0x48, 0xba, 0x70, 0x3e, 0xcd,
	0xe6, 0x33, 0xa1, 0xf9, 0x86, 0xb6, 0x7d, 0x2f, 0xa8, 0xb0, 0x9d, 0x36, 0xde, 0x8c, 0xbf, 0x0c,
	0xb9, 0x8a, 0x14, 0xed, 0x38, 0xaf, 0xd4, 0xe4, 0x12, 0xaa, 0x4f, 0x49, 0x4c, 0x2f, 0x2d, 0x36,
	0xa3, 0x23, 0x6b, 0xfa, 0xa6, 0x24, 0x6b, 0xd3, 0x67, 0x22, 0x8b, 0xc3, 0x28, 0x5e, 0x51, 0xe2,
	0xfa, 0xa5, 0x26, 0xb7, 0x00, 0x8f, 0x92, 0xc7, 0xa1, 0xdb, 0xf9, 0xad, 0x75, 0xf7, 0x08, 0xb9,
	0x81, 0xfa, 0x54, 0x73, 0x9d, 0x29, 0x7a, 0x65, 0xff, 0xa4, 0x42, 0x91, 0x1e, 0x5c, 0xb8, 0x89,
	0x21, 0x57, 0x22, 0xa6, 0xd7, 0xf6, 0xa1, 0x0e, 0x18, 0xf9, 0x0e, 0x8d, 0x9f, 0x2f, 0x49, 0x24,
	0x31, 0xbc, 0xd7, 0xf4, 0xc6, 0xf7, 0x82, 0xe6, 0x5d, 0xb7, 0xef, 0x4e, 0x66, 0xbf, 0x3c, 0x99,
	0xfd, 0x59, 0x79, 0x32, 0xd9, 0x6b, 0xd8, 0x34, 0x8b, 0xc3, 0x72, 0xaf, 0xe9, 0xbb, 0xd3, 0xcd,
	0x5d, 0xd8, 0x34, 0x9f, 0x93, 0xb0, 0x68, 0xd2, 0xd3, 0xcd, 0x5d, 0xd8, 0x34, 0x47, 0xb8, 0x41,
	0xd7, 0x7c, 0x7f, 0xba, 0xb9, 0x0b, 0xf7, 0xbe, 0xc0, 0x39, 0x43, 0x95, 0x8c, 0xb8, 0xe6, 0xe4,
	0x16, 0x6a, 0xe6, 0x9b, 0x7a, 0x7e, 0x35, 0x68, 0xde, 0x41, 0xdf, 0x5e, 0x4d, 0x73, 0xd3, 0x98,
	0xe5, 0xbd, 0x6b, 0xa8, 0x32, 0x4c, 0x8f, 0xaf, 0xdd, 0xdd, 0x0f, 0xe8, 0xdc, 0xaf, 0xd6, 0x3c,
	0x8e, 0xb6, 0x53, 0x94, 0x7f, 0xa2, 0x05, 0x2a, 0xf2, 0x15, 0x5a, 0x8f, 0xa8, 0x4d, 0x75, 0x98,
	0x3f, 0x2b, 0x94, 0xa4, 0x51, 0x2c, 0x86, 0x69, 0xb7, 0x5d, 0x8e, 0x6e, 0xd7, 0xde, 0x7f, 0xc3,
	0xce, 0xef, 0xd6, 0xc1, 0x3b, 0x61, 0x5e, 0xb7, 0xf2, 0xdb, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x05, 0x62, 0x5f, 0x2b, 0x2b, 0x04, 0x00, 0x00,
}
