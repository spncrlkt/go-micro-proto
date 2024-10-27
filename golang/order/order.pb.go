// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order/order.proto

package order

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type CreateOrderRequest struct {
	UserId               int64        `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderItems           []*OrderItem `protobuf:"bytes,2,rep,name=order_items,json=orderItems,proto3" json:"order_items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateOrderRequest) Reset()         { *m = CreateOrderRequest{} }
func (m *CreateOrderRequest) String() string { return proto.CompactTextString(m) }
func (*CreateOrderRequest) ProtoMessage()    {}
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa47a2077d8980ed, []int{0}
}

func (m *CreateOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderRequest.Unmarshal(m, b)
}
func (m *CreateOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderRequest.Marshal(b, m, deterministic)
}
func (m *CreateOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderRequest.Merge(m, src)
}
func (m *CreateOrderRequest) XXX_Size() int {
	return xxx_messageInfo_CreateOrderRequest.Size(m)
}
func (m *CreateOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderRequest proto.InternalMessageInfo

func (m *CreateOrderRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CreateOrderRequest) GetOrderItems() []*OrderItem {
	if m != nil {
		return m.OrderItems
	}
	return nil
}

type OrderItem struct {
	ProductCode          string   `protobuf:"bytes,1,opt,name=product_code,json=productCode,proto3" json:"product_code,omitempty"`
	UnitPrice            float32  `protobuf:"fixed32,2,opt,name=unit_price,json=unitPrice,proto3" json:"unit_price,omitempty"`
	Quantity             int32    `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderItem) Reset()         { *m = OrderItem{} }
func (m *OrderItem) String() string { return proto.CompactTextString(m) }
func (*OrderItem) ProtoMessage()    {}
func (*OrderItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa47a2077d8980ed, []int{1}
}

func (m *OrderItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderItem.Unmarshal(m, b)
}
func (m *OrderItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderItem.Marshal(b, m, deterministic)
}
func (m *OrderItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderItem.Merge(m, src)
}
func (m *OrderItem) XXX_Size() int {
	return xxx_messageInfo_OrderItem.Size(m)
}
func (m *OrderItem) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderItem.DiscardUnknown(m)
}

var xxx_messageInfo_OrderItem proto.InternalMessageInfo

func (m *OrderItem) GetProductCode() string {
	if m != nil {
		return m.ProductCode
	}
	return ""
}

func (m *OrderItem) GetUnitPrice() float32 {
	if m != nil {
		return m.UnitPrice
	}
	return 0
}

func (m *OrderItem) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type CreateOrderResponse struct {
	OrderId              int64    `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOrderResponse) Reset()         { *m = CreateOrderResponse{} }
func (m *CreateOrderResponse) String() string { return proto.CompactTextString(m) }
func (*CreateOrderResponse) ProtoMessage()    {}
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa47a2077d8980ed, []int{2}
}

func (m *CreateOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderResponse.Unmarshal(m, b)
}
func (m *CreateOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderResponse.Marshal(b, m, deterministic)
}
func (m *CreateOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderResponse.Merge(m, src)
}
func (m *CreateOrderResponse) XXX_Size() int {
	return xxx_messageInfo_CreateOrderResponse.Size(m)
}
func (m *CreateOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderResponse proto.InternalMessageInfo

func (m *CreateOrderResponse) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

type GetOrderRequest struct {
	OrderId              int64    `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOrderRequest) Reset()         { *m = GetOrderRequest{} }
func (m *GetOrderRequest) String() string { return proto.CompactTextString(m) }
func (*GetOrderRequest) ProtoMessage()    {}
func (*GetOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa47a2077d8980ed, []int{3}
}

func (m *GetOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderRequest.Unmarshal(m, b)
}
func (m *GetOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderRequest.Marshal(b, m, deterministic)
}
func (m *GetOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderRequest.Merge(m, src)
}
func (m *GetOrderRequest) XXX_Size() int {
	return xxx_messageInfo_GetOrderRequest.Size(m)
}
func (m *GetOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderRequest proto.InternalMessageInfo

func (m *GetOrderRequest) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

type GetOrderResponse struct {
	UserId               int64        `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderItems           []*OrderItem `protobuf:"bytes,2,rep,name=order_items,json=orderItems,proto3" json:"order_items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetOrderResponse) Reset()         { *m = GetOrderResponse{} }
func (m *GetOrderResponse) String() string { return proto.CompactTextString(m) }
func (*GetOrderResponse) ProtoMessage()    {}
func (*GetOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa47a2077d8980ed, []int{4}
}

func (m *GetOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderResponse.Unmarshal(m, b)
}
func (m *GetOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderResponse.Marshal(b, m, deterministic)
}
func (m *GetOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderResponse.Merge(m, src)
}
func (m *GetOrderResponse) XXX_Size() int {
	return xxx_messageInfo_GetOrderResponse.Size(m)
}
func (m *GetOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderResponse proto.InternalMessageInfo

func (m *GetOrderResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *GetOrderResponse) GetOrderItems() []*OrderItem {
	if m != nil {
		return m.OrderItems
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateOrderRequest)(nil), "CreateOrderRequest")
	proto.RegisterType((*OrderItem)(nil), "OrderItem")
	proto.RegisterType((*CreateOrderResponse)(nil), "CreateOrderResponse")
	proto.RegisterType((*GetOrderRequest)(nil), "GetOrderRequest")
	proto.RegisterType((*GetOrderResponse)(nil), "GetOrderResponse")
}

func init() {
	proto.RegisterFile("order/order.proto", fileDescriptor_fa47a2077d8980ed)
}

var fileDescriptor_fa47a2077d8980ed = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4d, 0x4f, 0xf2, 0x40,
	0x10, 0xc7, 0x29, 0x0d, 0x6f, 0xc3, 0x93, 0x3c, 0xb0, 0x98, 0x58, 0x49, 0x4c, 0x6a, 0x4f, 0x8d,
	0xc2, 0x62, 0x30, 0x7e, 0x01, 0x39, 0x10, 0x4e, 0x9a, 0x9e, 0x0c, 0x97, 0x06, 0xba, 0x13, 0xdc,
	0x48, 0xbb, 0x65, 0x77, 0x7a, 0xf0, 0xdb, 0x9b, 0x2e, 0x05, 0x03, 0x1a, 0x4f, 0x5e, 0x36, 0x3b,
	0xff, 0x79, 0xc9, 0xff, 0xb7, 0x3b, 0xd0, 0x57, 0x5a, 0xa0, 0x9e, 0xd8, 0x93, 0xe7, 0x5a, 0x91,
	0x0a, 0x96, 0xc0, 0x66, 0x1a, 0x57, 0x84, 0xcf, 0xa5, 0x18, 0xe1, 0xae, 0x40, 0x43, 0xec, 0x12,
	0x5a, 0x85, 0x41, 0x1d, 0x4b, 0xe1, 0x39, 0xbe, 0x13, 0xba, 0x51, 0xb3, 0x0c, 0x17, 0x82, 0xdd,
	0x41, 0xd7, 0x76, 0xc7, 0x92, 0x30, 0x35, 0x5e, 0xdd, 0x77, 0xc3, 0xee, 0x14, 0xb8, 0x6d, 0x5e,
	0x10, 0xa6, 0x11, 0xa8, 0xc3, 0xd5, 0x04, 0x12, 0x3a, 0xc7, 0x04, 0xbb, 0x81, 0x7f, 0xb9, 0x56,
	0xa2, 0x48, 0x28, 0x4e, 0x94, 0x40, 0x3b, 0xb7, 0x13, 0x75, 0x2b, 0x6d, 0xa6, 0x04, 0xb2, 0x6b,
	0x80, 0x22, 0x93, 0x14, 0xe7, 0x5a, 0x26, 0xe8, 0xd5, 0x7d, 0x27, 0xac, 0x47, 0x9d, 0x52, 0x79,
	0x29, 0x05, 0x36, 0x84, 0xf6, 0xae, 0x58, 0x65, 0x24, 0xe9, 0xc3, 0x73, 0x7d, 0x27, 0x6c, 0x44,
	0xc7, 0x38, 0xb8, 0x87, 0xc1, 0x09, 0x86, 0xc9, 0x55, 0x66, 0x90, 0x5d, 0x41, 0xbb, 0xb2, 0x7b,
	0x00, 0x69, 0xed, 0xfd, 0x89, 0x60, 0x04, 0xff, 0xe7, 0x48, 0x27, 0xd4, 0xbf, 0x54, 0xbf, 0x42,
	0xef, 0xab, 0xba, 0x1a, 0xfe, 0x27, 0x8f, 0x34, 0xdd, 0x42, 0xc3, 0x26, 0xd8, 0x23, 0x34, 0xf7,
	0x08, 0x6c, 0xc0, 0xbf, 0x7f, 0xc9, 0xf0, 0x82, 0xff, 0x00, 0x18, 0xd4, 0xd8, 0x08, 0xdc, 0x39,
	0x12, 0xeb, 0xf1, 0x33, 0x9a, 0x61, 0x9f, 0x9f, 0x3b, 0x0e, 0x6a, 0x4f, 0xb7, 0xcb, 0x70, 0x23,
	0xe9, 0xad, 0x58, 0xf3, 0x44, 0xa5, 0x13, 0x93, 0x67, 0x89, 0xde, 0xbe, 0xd3, 0x64, 0xa3, 0xc6,
	0xa9, 0x4c, 0xb4, 0x1a, 0xdb, 0xa5, 0xd8, 0x2f, 0xc8, 0xba, 0x69, 0x83, 0x87, 0xcf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x72, 0xc3, 0xfa, 0x39, 0x36, 0x02, 0x00, 0x00,
}
