// Code generated by protoc-gen-go. DO NOT EDIT.
// source: userdata.proto

package dex2

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

// 用户信息。
type UserInfo struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Trader               string   `protobuf:"bytes,2,opt,name=trader,proto3" json:"trader,omitempty"`
	Chain                string   `protobuf:"bytes,3,opt,name=chain,proto3" json:"chain,omitempty"`
	UserFrom             string   `protobuf:"bytes,4,opt,name=user_from,json=userFrom,proto3" json:"user_from,omitempty"`
	CreateTime           int64    `protobuf:"varint,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           int64    `protobuf:"varint,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b77b85936f4a4076, []int{0}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfo) GetTrader() string {
	if m != nil {
		return m.Trader
	}
	return ""
}

func (m *UserInfo) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *UserInfo) GetUserFrom() string {
	if m != nil {
		return m.UserFrom
	}
	return ""
}

func (m *UserInfo) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *UserInfo) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

// 用户操作记录。
type UserLog struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Trader               string   `protobuf:"bytes,3,opt,name=trader,proto3" json:"trader,omitempty"`
	Chain                string   `protobuf:"bytes,4,opt,name=chain,proto3" json:"chain,omitempty"`
	UserFrom             string   `protobuf:"bytes,5,opt,name=user_from,json=userFrom,proto3" json:"user_from,omitempty"`
	Action               int32    `protobuf:"varint,6,opt,name=action,proto3" json:"action,omitempty"`
	Sig                  string   `protobuf:"bytes,7,opt,name=sig,proto3" json:"sig,omitempty"`
	CreateTime           int64    `protobuf:"varint,8,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserLog) Reset()         { *m = UserLog{} }
func (m *UserLog) String() string { return proto.CompactTextString(m) }
func (*UserLog) ProtoMessage()    {}
func (*UserLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_b77b85936f4a4076, []int{1}
}

func (m *UserLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserLog.Unmarshal(m, b)
}
func (m *UserLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserLog.Marshal(b, m, deterministic)
}
func (m *UserLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLog.Merge(m, src)
}
func (m *UserLog) XXX_Size() int {
	return xxx_messageInfo_UserLog.Size(m)
}
func (m *UserLog) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLog.DiscardUnknown(m)
}

var xxx_messageInfo_UserLog proto.InternalMessageInfo

func (m *UserLog) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserLog) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *UserLog) GetTrader() string {
	if m != nil {
		return m.Trader
	}
	return ""
}

func (m *UserLog) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *UserLog) GetUserFrom() string {
	if m != nil {
		return m.UserFrom
	}
	return ""
}

func (m *UserLog) GetAction() int32 {
	if m != nil {
		return m.Action
	}
	return 0
}

func (m *UserLog) GetSig() string {
	if m != nil {
		return m.Sig
	}
	return ""
}

func (m *UserLog) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

// 用户Session 记录
type UserSession struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Enabled              int32    `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	CreateTime           int64    `protobuf:"varint,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Tid                  string   `protobuf:"bytes,4,opt,name=tid,proto3" json:"tid,omitempty"`
	Ua                   string   `protobuf:"bytes,5,opt,name=ua,proto3" json:"ua,omitempty"`
	Ip                   string   `protobuf:"bytes,6,opt,name=ip,proto3" json:"ip,omitempty"`
	Sub                  string   `protobuf:"bytes,7,opt,name=sub,proto3" json:"sub,omitempty"`
	ExpireTime           int64    `protobuf:"varint,8,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserSession) Reset()         { *m = UserSession{} }
func (m *UserSession) String() string { return proto.CompactTextString(m) }
func (*UserSession) ProtoMessage()    {}
func (*UserSession) Descriptor() ([]byte, []int) {
	return fileDescriptor_b77b85936f4a4076, []int{2}
}

func (m *UserSession) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSession.Unmarshal(m, b)
}
func (m *UserSession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSession.Marshal(b, m, deterministic)
}
func (m *UserSession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSession.Merge(m, src)
}
func (m *UserSession) XXX_Size() int {
	return xxx_messageInfo_UserSession.Size(m)
}
func (m *UserSession) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSession.DiscardUnknown(m)
}

var xxx_messageInfo_UserSession proto.InternalMessageInfo

func (m *UserSession) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserSession) GetEnabled() int32 {
	if m != nil {
		return m.Enabled
	}
	return 0
}

func (m *UserSession) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *UserSession) GetTid() string {
	if m != nil {
		return m.Tid
	}
	return ""
}

func (m *UserSession) GetUa() string {
	if m != nil {
		return m.Ua
	}
	return ""
}

func (m *UserSession) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *UserSession) GetSub() string {
	if m != nil {
		return m.Sub
	}
	return ""
}

func (m *UserSession) GetExpireTime() int64 {
	if m != nil {
		return m.ExpireTime
	}
	return 0
}

func init() {
	proto.RegisterType((*UserInfo)(nil), "dex.UserInfo")
	proto.RegisterType((*UserLog)(nil), "dex.UserLog")
	proto.RegisterType((*UserSession)(nil), "dex.UserSession")
}

func init() { proto.RegisterFile("userdata.proto", fileDescriptor_b77b85936f4a4076) }

var fileDescriptor_b77b85936f4a4076 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x49, 0xd2, 0xa4, 0xed, 0x29, 0x94, 0x12, 0x2e, 0x25, 0x70, 0x17, 0xb7, 0x74, 0x55,
	0x2e, 0x92, 0x82, 0xbe, 0x81, 0x0b, 0x41, 0x70, 0x15, 0x75, 0xe3, 0xa6, 0x4c, 0x32, 0xd3, 0x78,
	0xc0, 0xc9, 0x84, 0xc9, 0x0c, 0xe4, 0x8d, 0x7c, 0x0a, 0x17, 0xbe, 0x99, 0x9c, 0x99, 0x58, 0x34,
	0xd5, 0xdd, 0x39, 0x7f, 0x7e, 0xce, 0xff, 0xfd, 0x61, 0x60, 0x69, 0x3b, 0xa1, 0x39, 0x33, 0x2c,
	0x6f, 0xb5, 0x32, 0x2a, 0x8d, 0xb8, 0xe8, 0xb7, 0xaf, 0x01, 0xcc, 0x1e, 0x3b, 0xa1, 0x6f, 0x9b,
	0xa3, 0x4a, 0x97, 0x10, 0x22, 0xcf, 0x82, 0x4d, 0xb0, 0x8b, 0x8a, 0x10, 0x79, 0xba, 0x86, 0xc4,
	0x68, 0xc6, 0x85, 0xce, 0xc2, 0x4d, 0xb0, 0x9b, 0x17, 0xc3, 0x96, 0xfe, 0x81, 0xb8, 0x7a, 0x66,
	0xd8, 0x64, 0x91, 0x93, 0xfd, 0x92, 0xfe, 0x85, 0x39, 0x25, 0x1c, 0x8e, 0x5a, 0xc9, 0x6c, 0xe2,
	0xbe, 0xcc, 0x48, 0xb8, 0xd1, 0x4a, 0xa6, 0xff, 0x60, 0x51, 0x69, 0xc1, 0x8c, 0x38, 0x18, 0x94,
	0x22, 0x8b, 0x5d, 0x06, 0x78, 0xe9, 0x01, 0xa5, 0x20, 0x83, 0x6d, 0xf9, 0xc9, 0x90, 0x78, 0x83,
	0x97, 0xc8, 0xb0, 0x7d, 0x0f, 0x60, 0x4a, 0xa4, 0x77, 0xaa, 0x3e, 0x03, 0x5d, 0x41, 0x64, 0x91,
	0x0f, 0x94, 0x34, 0x7e, 0x41, 0x8f, 0x7e, 0x46, 0x9f, 0xfc, 0x8a, 0x1e, 0x8f, 0xd0, 0xd7, 0x90,
	0xb0, 0xca, 0xa0, 0x6a, 0x1c, 0x54, 0x5c, 0x0c, 0x1b, 0x85, 0x76, 0x58, 0x67, 0x53, 0x1f, 0xda,
	0x61, 0x3d, 0x2e, 0x39, 0x1b, 0x97, 0xdc, 0xbe, 0x05, 0xb0, 0xa0, 0x0e, 0xf7, 0xa2, 0xeb, 0xe8,
	0xc4, 0xb8, 0x47, 0x06, 0x53, 0xd1, 0xb0, 0xf2, 0x45, 0xf8, 0x2e, 0x71, 0xf1, 0xb9, 0x8e, 0x4f,
	0x47, 0x67, 0xff, 0x6f, 0x05, 0x91, 0x41, 0x3e, 0xd4, 0xa2, 0x91, 0x8e, 0x5b, 0x36, 0xb4, 0x09,
	0x2d, 0x73, 0x61, 0xad, 0xeb, 0x30, 0x2f, 0x42, 0x6c, 0x1d, 0xbf, 0x2d, 0x4f, 0xfc, 0xb6, 0xa4,
	0x10, 0xd1, 0xb7, 0xa8, 0xbf, 0xf3, 0x7b, 0x89, 0x42, 0xae, 0x2f, 0x9e, 0xfe, 0xd7, 0x68, 0xf2,
	0x12, 0x4b, 0xd5, 0xe7, 0x95, 0x92, 0x7b, 0x2e, 0x7a, 0xa3, 0xda, 0x7d, 0xa5, 0xa4, 0x54, 0x4d,
	0x5e, 0xa3, 0xd9, 0xbb, 0xb7, 0x45, 0xfa, 0x65, 0x99, 0xb8, 0xf9, 0xea, 0x23, 0x00, 0x00, 0xff,
	0xff, 0x51, 0xa2, 0x0a, 0x47, 0x79, 0x02, 0x00, 0x00,
}
