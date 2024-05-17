// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package go_micro_service_user

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

type ReqSignup struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqSignup) Reset()         { *m = ReqSignup{} }
func (m *ReqSignup) String() string { return proto.CompactTextString(m) }
func (*ReqSignup) ProtoMessage()    {}
func (*ReqSignup) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *ReqSignup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSignup.Unmarshal(m, b)
}
func (m *ReqSignup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSignup.Marshal(b, m, deterministic)
}
func (m *ReqSignup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSignup.Merge(m, src)
}
func (m *ReqSignup) XXX_Size() int {
	return xxx_messageInfo_ReqSignup.Size(m)
}
func (m *ReqSignup) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSignup.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSignup proto.InternalMessageInfo

func (m *ReqSignup) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ReqSignup) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RespSignup struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespSignup) Reset()         { *m = RespSignup{} }
func (m *RespSignup) String() string { return proto.CompactTextString(m) }
func (*RespSignup) ProtoMessage()    {}
func (*RespSignup) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *RespSignup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespSignup.Unmarshal(m, b)
}
func (m *RespSignup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespSignup.Marshal(b, m, deterministic)
}
func (m *RespSignup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespSignup.Merge(m, src)
}
func (m *RespSignup) XXX_Size() int {
	return xxx_messageInfo_RespSignup.Size(m)
}
func (m *RespSignup) XXX_DiscardUnknown() {
	xxx_messageInfo_RespSignup.DiscardUnknown(m)
}

var xxx_messageInfo_RespSignup proto.InternalMessageInfo

func (m *RespSignup) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespSignup) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ReqSignin struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqSignin) Reset()         { *m = ReqSignin{} }
func (m *ReqSignin) String() string { return proto.CompactTextString(m) }
func (*ReqSignin) ProtoMessage()    {}
func (*ReqSignin) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *ReqSignin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSignin.Unmarshal(m, b)
}
func (m *ReqSignin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSignin.Marshal(b, m, deterministic)
}
func (m *ReqSignin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSignin.Merge(m, src)
}
func (m *ReqSignin) XXX_Size() int {
	return xxx_messageInfo_ReqSignin.Size(m)
}
func (m *ReqSignin) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSignin.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSignin proto.InternalMessageInfo

func (m *ReqSignin) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ReqSignin) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RespSignin struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespSignin) Reset()         { *m = RespSignin{} }
func (m *RespSignin) String() string { return proto.CompactTextString(m) }
func (*RespSignin) ProtoMessage()    {}
func (*RespSignin) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *RespSignin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespSignin.Unmarshal(m, b)
}
func (m *RespSignin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespSignin.Marshal(b, m, deterministic)
}
func (m *RespSignin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespSignin.Merge(m, src)
}
func (m *RespSignin) XXX_Size() int {
	return xxx_messageInfo_RespSignin.Size(m)
}
func (m *RespSignin) XXX_DiscardUnknown() {
	xxx_messageInfo_RespSignin.DiscardUnknown(m)
}

var xxx_messageInfo_RespSignin proto.InternalMessageInfo

func (m *RespSignin) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespSignin) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RespSignin) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ReqUserInfo struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqUserInfo) Reset()         { *m = ReqUserInfo{} }
func (m *ReqUserInfo) String() string { return proto.CompactTextString(m) }
func (*ReqUserInfo) ProtoMessage()    {}
func (*ReqUserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *ReqUserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqUserInfo.Unmarshal(m, b)
}
func (m *ReqUserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqUserInfo.Marshal(b, m, deterministic)
}
func (m *ReqUserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqUserInfo.Merge(m, src)
}
func (m *ReqUserInfo) XXX_Size() int {
	return xxx_messageInfo_ReqUserInfo.Size(m)
}
func (m *ReqUserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqUserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReqUserInfo proto.InternalMessageInfo

func (m *ReqUserInfo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type RespUserInfo struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	SignupAt             string   `protobuf:"bytes,6,opt,name=signupAt,proto3" json:"signupAt,omitempty"`
	LastActiveAt         string   `protobuf:"bytes,7,opt,name=lastActiveAt,proto3" json:"lastActiveAt,omitempty"`
	Status               int32    `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespUserInfo) Reset()         { *m = RespUserInfo{} }
func (m *RespUserInfo) String() string { return proto.CompactTextString(m) }
func (*RespUserInfo) ProtoMessage()    {}
func (*RespUserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *RespUserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespUserInfo.Unmarshal(m, b)
}
func (m *RespUserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespUserInfo.Marshal(b, m, deterministic)
}
func (m *RespUserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespUserInfo.Merge(m, src)
}
func (m *RespUserInfo) XXX_Size() int {
	return xxx_messageInfo_RespUserInfo.Size(m)
}
func (m *RespUserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RespUserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RespUserInfo proto.InternalMessageInfo

func (m *RespUserInfo) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespUserInfo) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RespUserInfo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RespUserInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RespUserInfo) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *RespUserInfo) GetSignupAt() string {
	if m != nil {
		return m.SignupAt
	}
	return ""
}

func (m *RespUserInfo) GetLastActiveAt() string {
	if m != nil {
		return m.LastActiveAt
	}
	return ""
}

func (m *RespUserInfo) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type ReqUserFile struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqUserFile) Reset()         { *m = ReqUserFile{} }
func (m *ReqUserFile) String() string { return proto.CompactTextString(m) }
func (*ReqUserFile) ProtoMessage()    {}
func (*ReqUserFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *ReqUserFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqUserFile.Unmarshal(m, b)
}
func (m *ReqUserFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqUserFile.Marshal(b, m, deterministic)
}
func (m *ReqUserFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqUserFile.Merge(m, src)
}
func (m *ReqUserFile) XXX_Size() int {
	return xxx_messageInfo_ReqUserFile.Size(m)
}
func (m *ReqUserFile) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqUserFile.DiscardUnknown(m)
}

var xxx_messageInfo_ReqUserFile proto.InternalMessageInfo

func (m *ReqUserFile) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ReqUserFile) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type RespUserFile struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	FileData             []byte   `protobuf:"bytes,3,opt,name=fileData,proto3" json:"fileData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespUserFile) Reset()         { *m = RespUserFile{} }
func (m *RespUserFile) String() string { return proto.CompactTextString(m) }
func (*RespUserFile) ProtoMessage()    {}
func (*RespUserFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *RespUserFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespUserFile.Unmarshal(m, b)
}
func (m *RespUserFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespUserFile.Marshal(b, m, deterministic)
}
func (m *RespUserFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespUserFile.Merge(m, src)
}
func (m *RespUserFile) XXX_Size() int {
	return xxx_messageInfo_RespUserFile.Size(m)
}
func (m *RespUserFile) XXX_DiscardUnknown() {
	xxx_messageInfo_RespUserFile.DiscardUnknown(m)
}

var xxx_messageInfo_RespUserFile proto.InternalMessageInfo

func (m *RespUserFile) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespUserFile) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RespUserFile) GetFileData() []byte {
	if m != nil {
		return m.FileData
	}
	return nil
}

type ReqUserFileRename struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Filehash             string   `protobuf:"bytes,2,opt,name=filehash,proto3" json:"filehash,omitempty"`
	NewFileName          string   `protobuf:"bytes,3,opt,name=newFileName,proto3" json:"newFileName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqUserFileRename) Reset()         { *m = ReqUserFileRename{} }
func (m *ReqUserFileRename) String() string { return proto.CompactTextString(m) }
func (*ReqUserFileRename) ProtoMessage()    {}
func (*ReqUserFileRename) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{8}
}

func (m *ReqUserFileRename) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqUserFileRename.Unmarshal(m, b)
}
func (m *ReqUserFileRename) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqUserFileRename.Marshal(b, m, deterministic)
}
func (m *ReqUserFileRename) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqUserFileRename.Merge(m, src)
}
func (m *ReqUserFileRename) XXX_Size() int {
	return xxx_messageInfo_ReqUserFileRename.Size(m)
}
func (m *ReqUserFileRename) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqUserFileRename.DiscardUnknown(m)
}

var xxx_messageInfo_ReqUserFileRename proto.InternalMessageInfo

func (m *ReqUserFileRename) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ReqUserFileRename) GetFilehash() string {
	if m != nil {
		return m.Filehash
	}
	return ""
}

func (m *ReqUserFileRename) GetNewFileName() string {
	if m != nil {
		return m.NewFileName
	}
	return ""
}

type RespUserFileRename struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	FileData             []byte   `protobuf:"bytes,3,opt,name=fileData,proto3" json:"fileData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespUserFileRename) Reset()         { *m = RespUserFileRename{} }
func (m *RespUserFileRename) String() string { return proto.CompactTextString(m) }
func (*RespUserFileRename) ProtoMessage()    {}
func (*RespUserFileRename) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{9}
}

func (m *RespUserFileRename) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespUserFileRename.Unmarshal(m, b)
}
func (m *RespUserFileRename) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespUserFileRename.Marshal(b, m, deterministic)
}
func (m *RespUserFileRename) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespUserFileRename.Merge(m, src)
}
func (m *RespUserFileRename) XXX_Size() int {
	return xxx_messageInfo_RespUserFileRename.Size(m)
}
func (m *RespUserFileRename) XXX_DiscardUnknown() {
	xxx_messageInfo_RespUserFileRename.DiscardUnknown(m)
}

var xxx_messageInfo_RespUserFileRename proto.InternalMessageInfo

func (m *RespUserFileRename) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespUserFileRename) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RespUserFileRename) GetFileData() []byte {
	if m != nil {
		return m.FileData
	}
	return nil
}

func init() {
	proto.RegisterType((*ReqSignup)(nil), "go.micro.service.user.ReqSignup")
	proto.RegisterType((*RespSignup)(nil), "go.micro.service.user.RespSignup")
	proto.RegisterType((*ReqSignin)(nil), "go.micro.service.user.ReqSignin")
	proto.RegisterType((*RespSignin)(nil), "go.micro.service.user.RespSignin")
	proto.RegisterType((*ReqUserInfo)(nil), "go.micro.service.user.ReqUserInfo")
	proto.RegisterType((*RespUserInfo)(nil), "go.micro.service.user.RespUserInfo")
	proto.RegisterType((*ReqUserFile)(nil), "go.micro.service.user.ReqUserFile")
	proto.RegisterType((*RespUserFile)(nil), "go.micro.service.user.RespUserFile")
	proto.RegisterType((*ReqUserFileRename)(nil), "go.micro.service.user.ReqUserFileRename")
	proto.RegisterType((*RespUserFileRename)(nil), "go.micro.service.user.RespUserFileRename")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x53, 0xba, 0x74, 0xed, 0x6b, 0x85, 0x84, 0x55, 0x90, 0xd5, 0x53, 0x31, 0x97, 0xed,
	0x92, 0x03, 0xdc, 0xb8, 0xa0, 0x0a, 0x84, 0xc4, 0x05, 0x90, 0xa7, 0x21, 0x4e, 0x48, 0xa6, 0x7b,
	0x6b, 0x2d, 0x12, 0x3b, 0x8b, 0xdd, 0xed, 0x23, 0xf2, 0x25, 0xf8, 0x30, 0xc8, 0x76, 0x9c, 0xa6,
	0xd3, 0x96, 0x30, 0xb4, 0x5b, 0xfe, 0xcf, 0xcf, 0x3f, 0xff, 0xfd, 0xfc, 0x5e, 0x00, 0x76, 0x06,
	0xab, 0xac, 0xac, 0xb4, 0xd5, 0xe4, 0xf9, 0x46, 0x67, 0x85, 0x5c, 0x57, 0x3a, 0x33, 0x58, 0x5d,
	0xcb, 0x35, 0x66, 0x6e, 0x91, 0xbd, 0x87, 0x09, 0xc7, 0xab, 0x33, 0xb9, 0x51, 0xbb, 0x92, 0x2c,
	0x60, 0xec, 0x82, 0x4a, 0x14, 0x48, 0x07, 0xcb, 0xc1, 0xc9, 0x84, 0x37, 0xda, 0xad, 0x95, 0xc2,
	0x98, 0x1b, 0x5d, 0x5d, 0xd0, 0x27, 0x61, 0x2d, 0x6a, 0xf6, 0x16, 0x80, 0xa3, 0x29, 0x6b, 0x0a,
	0x81, 0xa3, 0xb5, 0xbe, 0x08, 0x84, 0x94, 0xfb, 0x6f, 0x42, 0xe1, 0xb8, 0x40, 0x63, 0xc4, 0x06,
	0xeb, 0xcd, 0x51, 0xb6, 0x0c, 0x48, 0xf5, 0xdf, 0x06, 0xbe, 0xee, 0x0d, 0x48, 0x75, 0xa7, 0x81,
	0x39, 0xa4, 0x56, 0xff, 0x42, 0x55, 0x6f, 0x0d, 0xa2, 0x6d, 0x6b, 0x78, 0x68, 0xeb, 0x14, 0xa6,
	0x1c, 0xaf, 0xce, 0x0d, 0x56, 0x9f, 0xd4, 0xa5, 0xee, 0x32, 0xc6, 0xfe, 0x0c, 0x60, 0xe6, 0x4e,
	0x6f, 0x92, 0x1f, 0x54, 0x80, 0x03, 0xf4, 0xf0, 0xd6, 0x9d, 0xe7, 0x90, 0x62, 0x21, 0x64, 0x4e,
	0x8f, 0x82, 0x6b, 0x2f, 0x5c, 0xb4, 0xdc, 0x6a, 0x85, 0x34, 0x0d, 0x51, 0x2f, 0x1c, 0xc7, 0xf8,
	0x07, 0x58, 0x59, 0x3a, 0x0a, 0x9c, 0xa8, 0x09, 0x83, 0x59, 0x2e, 0x8c, 0x5d, 0xad, 0xad, 0xbc,
	0xc6, 0x95, 0xa5, 0xc7, 0x7e, 0xfd, 0x20, 0x46, 0x5e, 0xc0, 0xc8, 0x58, 0x61, 0x77, 0x86, 0x8e,
	0xbd, 0xef, 0x5a, 0xb1, 0x77, 0x4d, 0x25, 0x3e, 0xca, 0x1c, 0x3b, 0x9f, 0x68, 0x0e, 0x69, 0x2e,
	0x0b, 0x69, 0xfd, 0x15, 0x53, 0x1e, 0x04, 0xfb, 0xbe, 0x2f, 0x8f, 0x27, 0x3c, 0xb8, 0x3c, 0x97,
	0x32, 0xc7, 0x0f, 0xc2, 0x0a, 0x5f, 0x9e, 0x19, 0x6f, 0x34, 0x2b, 0xe0, 0x59, 0xcb, 0x1a, 0xc7,
	0xd8, 0x27, 0x5d, 0x3d, 0xe4, 0x36, 0x6f, 0x85, 0xd9, 0xc6, 0x1e, 0x8a, 0x9a, 0x2c, 0x61, 0xaa,
	0xf0, 0xc6, 0x81, 0x3e, 0xef, 0x9f, 0xa2, 0x1d, 0x62, 0x3f, 0x80, 0xb4, 0x2f, 0x52, 0x9f, 0xf7,
	0x68, 0xd7, 0x79, 0xfd, 0x7b, 0x08, 0x53, 0x07, 0x3f, 0x0b, 0x03, 0x4a, 0xbe, 0xc0, 0xa8, 0x1e,
	0xa9, 0x65, 0x76, 0xe7, 0xf4, 0x66, 0xcd, 0xe8, 0x2e, 0x5e, 0xde, 0x9b, 0x11, 0xe7, 0x92, 0x25,
	0x11, 0x28, 0x55, 0x1f, 0x50, 0xaa, 0x5e, 0xa0, 0x54, 0x2c, 0x21, 0xe7, 0x30, 0x6e, 0xba, 0x9e,
	0xdd, 0x8f, 0x8c, 0x39, 0x8b, 0x57, 0x1d, 0xd0, 0x98, 0xc4, 0x12, 0xf2, 0x0d, 0x26, 0xb1, 0xc8,
	0xa6, 0x8f, 0xeb, 0x92, 0x7a, 0xb9, 0x2e, 0x89, 0x25, 0x64, 0x03, 0x4f, 0x6f, 0x3d, 0xde, 0x49,
	0x3f, 0x3c, 0x64, 0x2e, 0x4e, 0xff, 0xe1, 0x88, 0x90, 0xca, 0x92, 0x9f, 0x23, 0xff, 0xcf, 0x7d,
	0xf3, 0x37, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x5b, 0xb8, 0x64, 0x81, 0x05, 0x00, 0x00,
}