// Code generated by protoc-gen-go. DO NOT EDIT.
// source: serverdata.proto

package chain

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

// Must be synced with `PlaceOrderRequest`.
type Order struct {
	OrderId          int64  `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Trader           string `protobuf:"bytes,2,opt,name=trader,proto3" json:"trader,omitempty"`
	UserFrom         string `protobuf:"bytes,3,opt,name=user_from,json=userFrom,proto3" json:"user_from,omitempty"`
	Action           int32  `protobuf:"varint,4,opt,name=action,proto3" json:"action,omitempty"`
	Type             int32  `protobuf:"varint,5,opt,name=type,proto3" json:"type,omitempty"`
	StockTokenId     int64  `protobuf:"varint,6,opt,name=stock_token_id,json=stockTokenId,proto3" json:"stock_token_id,omitempty"`
	StockTokenSymbol string `protobuf:"bytes,7,opt,name=stock_token_symbol,json=stockTokenSymbol,proto3" json:"stock_token_symbol,omitempty"`
	CashTokenId      int64  `protobuf:"varint,8,opt,name=cash_token_id,json=cashTokenId,proto3" json:"cash_token_id,omitempty"`
	CashTokenSymbol  string `protobuf:"bytes,9,opt,name=cash_token_symbol,json=cashTokenSymbol,proto3" json:"cash_token_symbol,omitempty"`
	// How many cash tokens that 1e8 stock tokens is worth at this price.
	PriceE8 uint64 `protobuf:"varint,10,opt,name=price_e8,json=priceE8,proto3" json:"price_e8,omitempty"`
	// The total amount (multiplied by 1e8) of stock token to buy/sell.
	AmountTotalE8     uint64 `protobuf:"varint,11,opt,name=amount_total_e8,json=amountTotalE8,proto3" json:"amount_total_e8,omitempty"`
	AmountLeftE8      uint64 `protobuf:"varint,12,opt,name=amount_left_e8,json=amountLeftE8,proto3" json:"amount_left_e8,omitempty"`
	AccumulatedFundE8 uint64 `protobuf:"varint,13,opt,name=accumulated_fund_e8,json=accumulatedFundE8,proto3" json:"accumulated_fund_e8,omitempty"`
	OrderSign         string `protobuf:"bytes,14,opt,name=order_sign,json=orderSign,proto3" json:"order_sign,omitempty"`
	// Do NOT use create time to check which order is accepted earlier. Use order id instead.
	CreateTime           int64    `protobuf:"varint,15,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           int64    `protobuf:"varint,16,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	ExpireTime           int64    `protobuf:"varint,17,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	Status               int32    `protobuf:"varint,18,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_d984b92ea9a558b4, []int{0}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *Order) GetTrader() string {
	if m != nil {
		return m.Trader
	}
	return ""
}

func (m *Order) GetUserFrom() string {
	if m != nil {
		return m.UserFrom
	}
	return ""
}

func (m *Order) GetAction() int32 {
	if m != nil {
		return m.Action
	}
	return 0
}

func (m *Order) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Order) GetStockTokenId() int64 {
	if m != nil {
		return m.StockTokenId
	}
	return 0
}

func (m *Order) GetStockTokenSymbol() string {
	if m != nil {
		return m.StockTokenSymbol
	}
	return ""
}

func (m *Order) GetCashTokenId() int64 {
	if m != nil {
		return m.CashTokenId
	}
	return 0
}

func (m *Order) GetCashTokenSymbol() string {
	if m != nil {
		return m.CashTokenSymbol
	}
	return ""
}

func (m *Order) GetPriceE8() uint64 {
	if m != nil {
		return m.PriceE8
	}
	return 0
}

func (m *Order) GetAmountTotalE8() uint64 {
	if m != nil {
		return m.AmountTotalE8
	}
	return 0
}

func (m *Order) GetAmountLeftE8() uint64 {
	if m != nil {
		return m.AmountLeftE8
	}
	return 0
}

func (m *Order) GetAccumulatedFundE8() uint64 {
	if m != nil {
		return m.AccumulatedFundE8
	}
	return 0
}

func (m *Order) GetOrderSign() string {
	if m != nil {
		return m.OrderSign
	}
	return ""
}

func (m *Order) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *Order) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *Order) GetExpireTime() int64 {
	if m != nil {
		return m.ExpireTime
	}
	return 0
}

func (m *Order) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

// 币种对： coin_pair表。
type CoinPair struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CashTokenId          int64    `protobuf:"varint,2,opt,name=cash_token_id,json=cashTokenId,proto3" json:"cash_token_id,omitempty"`
	CashTokenSymbol      string   `protobuf:"bytes,3,opt,name=cash_token_symbol,json=cashTokenSymbol,proto3" json:"cash_token_symbol,omitempty"`
	StockTokenId         int64    `protobuf:"varint,4,opt,name=stock_token_id,json=stockTokenId,proto3" json:"stock_token_id,omitempty"`
	StockTokenSymbol     string   `protobuf:"bytes,5,opt,name=stock_token_symbol,json=stockTokenSymbol,proto3" json:"stock_token_symbol,omitempty"`
	ChainToken           string   `protobuf:"bytes,6,opt,name=chain_token,json=chainToken,proto3" json:"chain_token,omitempty"`
	ValidDecimals        int32    `protobuf:"varint,7,opt,name=valid_decimals,json=validDecimals,proto3" json:"valid_decimals,omitempty"`
	OrderMinAmountE8     int64    `protobuf:"varint,8,opt,name=order_min_amount_e8,json=orderMinAmountE8,proto3" json:"order_min_amount_e8,omitempty"`
	OrderMinValueE8      int64    `protobuf:"varint,9,opt,name=order_min_value_e8,json=orderMinValueE8,proto3" json:"order_min_value_e8,omitempty"`
	Status               int32    `protobuf:"varint,10,opt,name=status,proto3" json:"status,omitempty"`
	CreateTime           int64    `protobuf:"varint,11,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           int64    `protobuf:"varint,12,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CoinPair) Reset()         { *m = CoinPair{} }
func (m *CoinPair) String() string { return proto.CompactTextString(m) }
func (*CoinPair) ProtoMessage()    {}
func (*CoinPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_d984b92ea9a558b4, []int{1}
}

func (m *CoinPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CoinPair.Unmarshal(m, b)
}
func (m *CoinPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CoinPair.Marshal(b, m, deterministic)
}
func (m *CoinPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoinPair.Merge(m, src)
}
func (m *CoinPair) XXX_Size() int {
	return xxx_messageInfo_CoinPair.Size(m)
}
func (m *CoinPair) XXX_DiscardUnknown() {
	xxx_messageInfo_CoinPair.DiscardUnknown(m)
}

var xxx_messageInfo_CoinPair proto.InternalMessageInfo

func (m *CoinPair) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CoinPair) GetCashTokenId() int64 {
	if m != nil {
		return m.CashTokenId
	}
	return 0
}

func (m *CoinPair) GetCashTokenSymbol() string {
	if m != nil {
		return m.CashTokenSymbol
	}
	return ""
}

func (m *CoinPair) GetStockTokenId() int64 {
	if m != nil {
		return m.StockTokenId
	}
	return 0
}

func (m *CoinPair) GetStockTokenSymbol() string {
	if m != nil {
		return m.StockTokenSymbol
	}
	return ""
}

func (m *CoinPair) GetChainToken() string {
	if m != nil {
		return m.ChainToken
	}
	return ""
}

func (m *CoinPair) GetValidDecimals() int32 {
	if m != nil {
		return m.ValidDecimals
	}
	return 0
}

func (m *CoinPair) GetOrderMinAmountE8() int64 {
	if m != nil {
		return m.OrderMinAmountE8
	}
	return 0
}

func (m *CoinPair) GetOrderMinValueE8() int64 {
	if m != nil {
		return m.OrderMinValueE8
	}
	return 0
}

func (m *CoinPair) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *CoinPair) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *CoinPair) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Order)(nil), "chain.Order")
	proto.RegisterType((*CoinPair)(nil), "chain.CoinPair")
}

func init() { proto.RegisterFile("serverdata.proto", fileDescriptor_d984b92ea9a558b4) }

var fileDescriptor_d984b92ea9a558b4 = []byte{
	// 565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x5d, 0x6b, 0xdb, 0x3c,
	0x14, 0xc7, 0xc9, 0x8b, 0xd3, 0xf8, 0xe4, 0xb5, 0x2a, 0x14, 0x97, 0x87, 0x87, 0x85, 0xb2, 0x8d,
	0xb0, 0xae, 0xc9, 0xc5, 0x6e, 0x7c, 0xbb, 0x17, 0x17, 0x0a, 0x1b, 0x1b, 0x6e, 0xd9, 0xc5, 0x6e,
	0x8c, 0x62, 0x29, 0xa9, 0xa8, 0x25, 0x19, 0x59, 0x2e, 0xed, 0x77, 0xdb, 0x27, 0xd9, 0xa7, 0x19,
	0x3a, 0x72, 0x9b, 0xd0, 0x6e, 0xac, 0x77, 0x3e, 0x3f, 0xfd, 0x72, 0xa2, 0x9c, 0xf3, 0x8f, 0x61,
	0x5a, 0x71, 0x73, 0xc3, 0x0d, 0xa3, 0x96, 0x2e, 0x4a, 0xa3, 0xad, 0x26, 0x41, 0x7e, 0x45, 0x85,
	0x3a, 0xfe, 0xd5, 0x85, 0xe0, 0xab, 0x61, 0xdc, 0x90, 0x23, 0xe8, 0x6b, 0xf7, 0x90, 0x09, 0x16,
	0xb5, 0x66, 0xad, 0x79, 0x27, 0xdd, 0xc3, 0xfa, 0x9c, 0x91, 0x43, 0xe8, 0x59, 0x43, 0x19, 0x37,
	0x51, 0x7b, 0xd6, 0x9a, 0x87, 0x69, 0x53, 0x91, 0xff, 0x20, 0xac, 0x2b, 0x6e, 0xb2, 0xb5, 0xd1,
	0x32, 0xea, 0xe0, 0x51, 0xdf, 0x81, 0x33, 0xa3, 0xa5, 0xfb, 0x10, 0xcd, 0xad, 0xd0, 0x2a, 0xea,
	0xce, 0x5a, 0xf3, 0x20, 0x6d, 0x2a, 0x42, 0xa0, 0x6b, 0xef, 0x4a, 0x1e, 0x05, 0x48, 0xf1, 0x99,
	0xbc, 0x84, 0x71, 0x65, 0x75, 0x7e, 0x9d, 0x59, 0x7d, 0xcd, 0x95, 0xbb, 0x41, 0x0f, 0x6f, 0x30,
	0x44, 0x7a, 0xe9, 0xe0, 0x39, 0x23, 0x6f, 0x81, 0xec, 0x5a, 0xd5, 0x9d, 0x5c, 0xe9, 0x22, 0xda,
	0xc3, 0xef, 0x9d, 0x6e, 0xcd, 0x0b, 0xe4, 0xe4, 0x18, 0x46, 0x39, 0xad, 0xae, 0xb6, 0x2d, 0xfb,
	0xd8, 0x72, 0xe0, 0xe0, 0x7d, 0xc7, 0x37, 0xb0, 0xbf, 0xe3, 0x34, 0x0d, 0x43, 0x6c, 0x38, 0x79,
	0xf0, 0x9a, 0x7e, 0x47, 0xd0, 0x2f, 0x8d, 0xc8, 0x79, 0xc6, 0xe3, 0x08, 0x66, 0xad, 0x79, 0x37,
	0xdd, 0xc3, 0x3a, 0x89, 0xc9, 0x6b, 0x98, 0x50, 0xa9, 0x6b, 0x65, 0x33, 0xab, 0x2d, 0x2d, 0x9c,
	0x31, 0x40, 0x63, 0xe4, 0xf1, 0xa5, 0xa3, 0x49, 0xec, 0x7e, 0x66, 0xe3, 0x15, 0x7c, 0x6d, 0x9d,
	0x36, 0x44, 0x6d, 0xe8, 0xe9, 0x67, 0xbe, 0xb6, 0x49, 0x4c, 0x16, 0x70, 0x40, 0xf3, 0xbc, 0x96,
	0x75, 0x41, 0x2d, 0x67, 0xd9, 0xba, 0x56, 0xcc, 0xa9, 0x23, 0x54, 0xf7, 0x77, 0x8e, 0xce, 0x6a,
	0xc5, 0x92, 0x98, 0xfc, 0x0f, 0xe0, 0x17, 0x57, 0x89, 0x8d, 0x8a, 0xc6, 0x78, 0xfb, 0x10, 0xc9,
	0x85, 0xd8, 0x28, 0xf2, 0x02, 0x06, 0xb9, 0xe1, 0xd4, 0xf2, 0xcc, 0x0a, 0xc9, 0xa3, 0x09, 0x4e,
	0x01, 0x3c, 0xba, 0x14, 0x92, 0x3b, 0xa1, 0x2e, 0xd9, 0x83, 0x30, 0xf5, 0x82, 0x47, 0xf7, 0x02,
	0xbf, 0x2d, 0x85, 0x69, 0x84, 0x7d, 0x2f, 0x78, 0x84, 0xc2, 0x21, 0xf4, 0x2a, 0x4b, 0x6d, 0x5d,
	0x45, 0xc4, 0xaf, 0xda, 0x57, 0xc7, 0x3f, 0x3b, 0xd0, 0xff, 0xa8, 0x85, 0xfa, 0x46, 0x85, 0x21,
	0x63, 0x68, 0x3f, 0x24, 0xab, 0x2d, 0xd8, 0xd3, 0xfd, 0xb4, 0x9f, 0xb9, 0x9f, 0xce, 0x9f, 0xf7,
	0xf3, 0x34, 0x43, 0xdd, 0x67, 0x67, 0x28, 0xf8, 0x4b, 0x86, 0xdc, 0xec, 0xdc, 0xdf, 0xc4, 0xdb,
	0x18, 0xca, 0x30, 0x05, 0x44, 0xa8, 0x91, 0x57, 0x30, 0xbe, 0xa1, 0x85, 0x60, 0x19, 0xe3, 0xb9,
	0x90, 0xb4, 0xa8, 0x30, 0x8e, 0x41, 0x3a, 0x42, 0xfa, 0xa9, 0x81, 0xe4, 0x14, 0x0e, 0xfc, 0x8a,
	0xa4, 0x50, 0x59, 0x13, 0x01, 0x1e, 0x37, 0x89, 0x9c, 0xe2, 0xd1, 0x17, 0xa1, 0xde, 0xe3, 0x41,
	0x12, 0x93, 0x13, 0x20, 0x5b, 0xfd, 0x86, 0x16, 0x35, 0x86, 0x2e, 0x44, 0x7b, 0x72, 0x6f, 0x7f,
	0x77, 0x3c, 0x89, 0x77, 0x86, 0x0f, 0xbb, 0xc3, 0x7f, 0xbc, 0xf7, 0xc1, 0xbf, 0xf6, 0x3e, 0x7c,
	0xbc, 0xf7, 0x0f, 0xa7, 0x3f, 0x4e, 0x36, 0xc2, 0x2e, 0x56, 0x62, 0xa5, 0x6f, 0x17, 0xb9, 0x96,
	0x4b, 0xc6, 0x6f, 0xad, 0x2e, 0x97, 0xb9, 0x96, 0x52, 0xab, 0xc5, 0x46, 0xd8, 0x25, 0xbe, 0x4b,
	0x96, 0x38, 0x90, 0x55, 0x0f, 0x8b, 0x77, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x32, 0xc1, 0xa6,
	0x56, 0x6c, 0x04, 0x00, 0x00,
}
