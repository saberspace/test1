// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account/model.proto

package models

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
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

type Account struct {
	AccountId            string               `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Username             string               `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email                string               `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PasswordHash         string               `protobuf:"bytes,4,opt,name=password_hash,json=passwordHash,proto3" json:"password_hash,omitempty"`
	PasswordSalt         string               `protobuf:"bytes,5,opt,name=password_salt,json=passwordSalt,proto3" json:"password_salt,omitempty"`
	Description          string               `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	VerificationCode     string               `protobuf:"bytes,7,opt,name=verification_code,json=verificationCode,proto3" json:"verification_code,omitempty"`
	AccountType          string               `protobuf:"bytes,8,opt,name=account_type,json=accountType,proto3" json:"account_type,omitempty"`
	Verified             bool                 `protobuf:"varint,9,opt,name=verified,proto3" json:"verified,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8c387d18afeb6b2, []int{0}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Account) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Account) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Account) GetPasswordHash() string {
	if m != nil {
		return m.PasswordHash
	}
	return ""
}

func (m *Account) GetPasswordSalt() string {
	if m != nil {
		return m.PasswordSalt
	}
	return ""
}

func (m *Account) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Account) GetVerificationCode() string {
	if m != nil {
		return m.VerificationCode
	}
	return ""
}

func (m *Account) GetAccountType() string {
	if m != nil {
		return m.AccountType
	}
	return ""
}

func (m *Account) GetVerified() bool {
	if m != nil {
		return m.Verified
	}
	return false
}

func (m *Account) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Account) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Account)(nil), "account.Account")
}

func init() { proto.RegisterFile("account/model.proto", fileDescriptor_d8c387d18afeb6b2) }

var fileDescriptor_d8c387d18afeb6b2 = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x86, 0xe5, 0x7c, 0x69, 0xe2, 0x1c, 0xa7, 0x3f, 0xdf, 0x94, 0x85, 0x9d, 0x05, 0x44, 0xb0,
	0x20, 0x01, 0x25, 0x6e, 0x83, 0x5a, 0xd1, 0xb2, 0x80, 0x84, 0x0d, 0x6c, 0x4d, 0x57, 0x6c, 0xa2,
	0x89, 0x7d, 0xe2, 0x0c, 0xb2, 0x7d, 0xac, 0x99, 0x49, 0x4b, 0x2e, 0xcd, 0x5e, 0x71, 0x01, 0xdc,
	0x05, 0x37, 0x82, 0xe2, 0x9f, 0xd6, 0x48, 0xa8, 0xac, 0xec, 0x63, 0x3f, 0xcf, 0xeb, 0xd1, 0xd1,
	0x6b, 0x38, 0xe5, 0xbe, 0x4f, 0xdb, 0x44, 0xbb, 0x31, 0x05, 0x18, 0x4d, 0x53, 0x49, 0x9a, 0x58,
	0xb7, 0x7a, 0x38, 0xb8, 0x0e, 0x85, 0xde, 0x6c, 0x57, 0x53, 0x9f, 0x62, 0x57, 0x24, 0x6b, 0x5a,
	0x45, 0xf4, 0x9d, 0x52, 0x4c, 0xdc, 0x82, 0xf3, 0x27, 0x21, 0x26, 0x93, 0x90, 0x64, 0xec, 0x52,
	0xaa, 0x05, 0x25, 0xca, 0xdd, 0x0f, 0x65, 0xc8, 0xe0, 0x59, 0x48, 0x14, 0x46, 0x58, 0xa2, 0xab,
	0xed, 0xda, 0xd5, 0x22, 0x46, 0xa5, 0x79, 0x9c, 0x56, 0xc0, 0xa4, 0x11, 0x1e, 0x52, 0x48, 0x0f,
	0xe4, 0x7e, 0x2a, 0x86, 0xe2, 0xae, 0xc4, 0x9f, 0xff, 0x6c, 0x43, 0x77, 0x5e, 0x9e, 0x8b, 0xbd,
	0x04, 0xa8, 0x8e, 0xb8, 0x14, 0x81, 0x6d, 0x0c, 0x8d, 0x51, 0x6f, 0x61, 0xe6, 0x99, 0xd3, 0x86,
	0xd6, 0xc8, 0xf0, 0x7a, 0xd5, 0xbb, 0xcf, 0x01, 0x3b, 0x03, 0x73, 0xab, 0x50, 0x26, 0x3c, 0x46,
	0xbb, 0x55, 0x60, 0x4f, 0xf2, 0xcc, 0x39, 0x81, 0x23, 0xd6, 0xbf, 0xe5, 0xd2, 0xdf, 0x70, 0x39,
	0x3a, 0x9f, 0x9d, 0x8d, 0xbd, 0x7b, 0x8a, 0xbd, 0x82, 0x03, 0x8c, 0xb9, 0x88, 0xec, 0xff, 0x1e,
	0xc1, 0x4b, 0x84, 0x5d, 0xc1, 0x61, 0xca, 0x95, 0xba, 0x23, 0x19, 0x2c, 0x37, 0x5c, 0x6d, 0xec,
	0xf6, 0x23, 0x4e, 0xbf, 0x46, 0x3f, 0x71, 0xb5, 0x61, 0x2f, 0x1a, 0xaa, 0xe2, 0x91, 0xb6, 0x0f,
	0xf6, 0xea, 0x03, 0xf4, 0x85, 0x47, 0x9a, 0x5d, 0x82, 0x15, 0xa0, 0xf2, 0xa5, 0x28, 0xb6, 0x6b,
	0x77, 0xfe, 0x9e, 0x7e, 0x71, 0x3e, 0x1b, 0x7b, 0x4d, 0x90, 0xbd, 0x87, 0xff, 0x6f, 0x51, 0x8a,
	0xb5, 0xf0, 0xf9, 0x7e, 0x5e, 0xfa, 0x14, 0xa0, 0xdd, 0x2d, 0x6c, 0x96, 0x67, 0xce, 0x11, 0xf4,
	0x19, 0xd4, 0xf6, 0xdb, 0xb1, 0x77, 0xd2, 0x84, 0x3f, 0x52, 0x80, 0xec, 0x12, 0xfa, 0xf5, 0x7e,
	0xf5, 0x2e, 0x45, 0xdb, 0x2c, 0xdc, 0xd3, 0x3c, 0x73, 0x8e, 0xe1, 0x90, 0x59, 0xb5, 0x3b, 0xbb,
	0x18, 0x7b, 0x56, 0x05, 0xde, 0xec, 0x52, 0x64, 0xaf, 0xc1, 0x2c, 0xb3, 0x30, 0xb0, 0x7b, 0x43,
	0x63, 0x64, 0x2e, 0x8e, 0xf3, 0xcc, 0xb1, 0xa0, 0xc7, 0xba, 0x2b, 0xa2, 0x08, 0x79, 0xe2, 0xdd,
	0x03, 0xec, 0x0a, 0xc0, 0x97, 0xc8, 0x35, 0x06, 0x4b, 0xae, 0x6d, 0x18, 0x1a, 0x23, 0x6b, 0x36,
	0x98, 0x96, 0xad, 0x99, 0xd6, 0x5d, 0x98, 0xde, 0xd4, 0xad, 0xf1, 0x7a, 0x15, 0x3d, 0xd7, 0x7b,
	0x75, 0x9b, 0x06, 0xb5, 0x6a, 0xfd, 0x5b, 0xad, 0xe8, 0xb9, 0xbe, 0xee, 0xe4, 0x99, 0xd3, 0x32,
	0x8d, 0xc5, 0x87, 0x1f, 0xbf, 0x9e, 0x1a, 0x5f, 0x9b, 0x05, 0x4f, 0x84, 0xe4, 0xdf, 0x42, 0x24,
	0x19, 0x0a, 0x9e, 0xb8, 0x21, 0xd7, 0x78, 0xc7, 0x77, 0xae, 0x92, 0xbe, 0xfb, 0xc7, 0x3f, 0xa2,
	0xde, 0x95, 0x97, 0x55, 0xa7, 0xf8, 0xd0, 0x9b, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcb, 0xad,
	0xda, 0xba, 0x43, 0x03, 0x00, 0x00,
}
