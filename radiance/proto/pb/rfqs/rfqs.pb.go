// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/pb/rfqs/rfqs.proto

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

func init() { proto.RegisterFile("proto/pb/rfqs/rfqs.proto", fileDescriptor_3932bcfd2247a7ec) }

var fileDescriptor_3932bcfd2247a7ec = []byte{
	// 564 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0x41, 0x6f, 0xda, 0x30,
	0x1c, 0xc5, 0x17, 0xa0, 0xac, 0x98, 0x96, 0x76, 0x5e, 0xdb, 0x79, 0x68, 0xea, 0x22, 0x34, 0x4d,
	0xd1, 0x34, 0x91, 0x89, 0x5d, 0x76, 0x9b, 0x4a, 0xd9, 0x2a, 0x2e, 0x55, 0x65, 0xe8, 0x65, 0x37,
	0x87, 0xfc, 0x81, 0x68, 0xc4, 0x4e, 0x6c, 0x67, 0x6a, 0xf6, 0x11, 0xf7, 0xa9, 0x26, 0xdb, 0x09,
	0xa5, 0x5c, 0x7a, 0x81, 0xff, 0xfb, 0xbd, 0xf7, 0x6c, 0x05, 0xec, 0x20, 0x92, 0x49, 0xa1, 0x45,
	0x98, 0x45, 0xa1, 0x5c, 0xe6, 0xca, 0x7e, 0x0c, 0x2d, 0xc2, 0x2d, 0x33, 0xf7, 0xdf, 0xaf, 0x84,
	0x58, 0x6d, 0x20, 0xb4, 0x2c, 0x2a, 0x96, 0xa1, 0x4e, 0x52, 0x50, 0x9a, 0xa5, 0x99, 0x8b, 0x0d,
	0xfe, 0xb5, 0x51, 0x8b, 0x2e, 0x73, 0x85, 0x7b, 0xa8, 0x31, 0x9d, 0x10, 0xcf, 0xf7, 0x82, 0x26,
	0x6d, 0x4c, 0x27, 0xf8, 0x03, 0x3a, 0x9e, 0x4b, 0xc6, 0x15, 0x5b, 0xe8, 0x44, 0xf0, 0x69, 0x4c,
	0x1a, 0xd6, 0x7a, 0x0a, 0xb1, 0x8f, 0xba, 0x14, 0xf2, 0x02, 0x94, 0x86, 0x78, 0x5c, 0x92, 0xa6,
	0xcd, 0xec, 0x22, 0xfc, 0x0e, 0x75, 0xae, 0x25, 0x30, 0xe7, 0xb7, 0xac, 0xff, 0x08, 0x4c, 0xff,
	0x8e, 0x95, 0x29, 0x70, 0x3d, 0x2f, 0x33, 0x20, 0x6d, 0xdf, 0x0b, 0x0e, 0xe8, 0x2e, 0xc2, 0x01,
	0x3a, 0xa9, 0xe4, 0xa4, 0x90, 0xcc, 0x6c, 0x4b, 0x5e, 0xda, 0x55, 0xf6, 0x31, 0xfe, 0x88, 0x7a,
	0xd7, 0x22, 0xcd, 0x18, 0x2f, 0xc7, 0x45, 0x09, 0x72, 0x1a, 0x93, 0x43, 0x1b, 0xdc, 0xa3, 0x66,
	0xc5, 0x8a, 0xcc, 0x60, 0xb3, 0xb1, 0xc1, 0x8e, 0x5b, 0x71, 0x0f, 0xe3, 0x33, 0x74, 0x40, 0x97,
	0xf9, 0xad, 0x20, 0xc8, 0xf7, 0x82, 0x0e, 0x75, 0xc2, 0x3d, 0xf3, 0x12, 0x24, 0xf0, 0x05, 0xdc,
	0x0a, 0xd2, 0xb5, 0xde, 0x2e, 0x32, 0xbd, 0x5b, 0xa1, 0x41, 0x91, 0x23, 0xd7, 0xb3, 0xc2, 0xfc,
	0xa2, 0x66, 0xf8, 0x29, 0xa4, 0xdb, 0x80, 0x1c, 0x5b, 0xf7, 0x29, 0xc4, 0x7d, 0x74, 0x38, 0x2b,
	0xa2, 0xb9, 0xd0, 0x6c, 0x43, 0x7a, 0xbe, 0x17, 0x34, 0xe8, 0x56, 0x1b, 0x6f, 0xce, 0x1e, 0xc6,
	0x4c, 0x25, 0x8a, 0x9c, 0x38, 0xaf, 0xd6, 0xf8, 0x14, 0x35, 0xef, 0x32, 0x4e, 0x4e, 0x2d, 0x36,
	0xa3, 0x23, 0x6b, 0xf2, 0xaa, 0x26, 0x6b, 0xd3, 0xa7, 0xa2, 0xe0, 0x71, 0xc2, 0x57, 0x04, 0xbb,
	0x7e, 0xad, 0xf1, 0x25, 0x42, 0x37, 0x92, 0xf1, 0xd8, 0xed, 0xfc, 0xda, 0xba, 0x3b, 0x04, 0x5f,
	0xa0, 0xf6, 0x4c, 0x33, 0x5d, 0x28, 0x72, 0x66, 0xff, 0xa4, 0x4a, 0xe1, 0x01, 0x3a, 0x72, 0x13,
	0x05, 0xa6, 0x04, 0x27, 0xe7, 0xf6, 0xa1, 0x9e, 0x30, 0xfc, 0x0d, 0x75, 0x7e, 0x3c, 0x64, 0x89,
	0x84, 0xf8, 0x4a, 0x93, 0x0b, 0xdf, 0x0b, 0xba, 0xa3, 0xfe, 0xd0, 0x9d, 0xcc, 0x61, 0x7d, 0x32,
	0x87, 0xf3, 0xfa, 0x64, 0xd2, 0xc7, 0xb0, 0x69, 0x56, 0x87, 0xe5, 0x4a, 0x93, 0x37, 0xcf, 0x37,
	0xb7, 0x61, 0xd3, 0xbc, 0xcf, 0xe2, 0xaa, 0x49, 0x9e, 0x6f, 0x6e, 0xc3, 0xa6, 0x39, 0x81, 0x0d,
	0xb8, 0xe6, 0xdb, 0xe7, 0x9b, 0xdb, 0xf0, 0xe0, 0x13, 0x3a, 0xa4, 0xa0, 0xb2, 0x09, 0xd3, 0x0c,
	0x5f, 0xa2, 0x96, 0xf9, 0x26, 0x9e, 0xdf, 0x0c, 0xba, 0x23, 0x34, 0xb4, 0x57, 0xd3, 0xdc, 0x34,
	0x6a, 0xf9, 0xe0, 0x1c, 0x35, 0x29, 0xe4, 0xfb, 0xd7, 0x6e, 0xf4, 0x1d, 0x9d, 0x5c, 0xad, 0xd6,
	0x8c, 0x27, 0xe9, 0x0c, 0xe4, 0x9f, 0x64, 0x01, 0x0a, 0x7f, 0x46, 0xc7, 0x37, 0xa0, 0x4d, 0x75,
	0x5c, 0xde, 0x2b, 0x90, 0xb8, 0x53, 0x2d, 0x06, 0x79, 0xbf, 0x57, 0x8f, 0x6e, 0xd7, 0xc1, 0x8b,
	0xf1, 0xe8, 0xd7, 0x97, 0x55, 0xa2, 0xd7, 0x45, 0x34, 0x5c, 0x88, 0x34, 0x4c, 0xa3, 0xe4, 0x6f,
	0xca, 0xe4, 0x6f, 0xd0, 0x61, 0x9c, 0x66, 0xa1, 0x64, 0x71, 0xc2, 0xf8, 0xa2, 0x7a, 0x1f, 0xd4,
	0xaf, 0x8d, 0xa8, 0x6d, 0xe5, 0xd7, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x63, 0xde, 0x4b, 0x1c,
	0x4e, 0x04, 0x00, 0x00,
}
