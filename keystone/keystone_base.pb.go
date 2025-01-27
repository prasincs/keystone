// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: keystone_base.proto

package keystone

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// keygenAlgorithm is the algorithm that should be used
// for generating a key - typically the EC curve in use
// that the key is related to.
type KeygenAlgorithm int32

const (
	KeygenAlgorithm_KEYGEN_SECP256K1 KeygenAlgorithm = 0
	KeygenAlgorithm_KEYGEN_SECP256R1 KeygenAlgorithm = 1
	KeygenAlgorithm_KEYGEN_ED25519   KeygenAlgorithm = 2
)

// Enum value maps for KeygenAlgorithm.
var (
	KeygenAlgorithm_name = map[int32]string{
		0: "KEYGEN_SECP256K1",
		1: "KEYGEN_SECP256R1",
		2: "KEYGEN_ED25519",
	}
	KeygenAlgorithm_value = map[string]int32{
		"KEYGEN_SECP256K1": 0,
		"KEYGEN_SECP256R1": 1,
		"KEYGEN_ED25519":   2,
	}
)

func (x KeygenAlgorithm) Enum() *KeygenAlgorithm {
	p := new(KeygenAlgorithm)
	*p = x
	return p
}

func (x KeygenAlgorithm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KeygenAlgorithm) Descriptor() protoreflect.EnumDescriptor {
	return file_keystone_base_proto_enumTypes[0].Descriptor()
}

func (KeygenAlgorithm) Type() protoreflect.EnumType {
	return &file_keystone_base_proto_enumTypes[0]
}

func (x KeygenAlgorithm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KeygenAlgorithm.Descriptor instead.
func (KeygenAlgorithm) EnumDescriptor() ([]byte, []int) {
	return file_keystone_base_proto_rawDescGZIP(), []int{0}
}

type SigningProfile int32

const (
	// ECDSA signing, SHA256 prior to signature, low-s normalization,
	// and raw r, s values instead of ASN
	SigningProfile_PROFILE_BC_ECDSA_SHA256 SigningProfile = 0
	// ECDSA signing, SHA512 prior to signature, low-s normalization,
	// and raw r, s values instead of ASN
	SigningProfile_PROFILE_BC_ECDSA_SHA512 SigningProfile = 1
	// ECDSA signing, SHA256 prior to signature, no normalization, and
	// standard ASN1 encoding
	SigningProfile_PROFILE_ECDSA_SHA256 SigningProfile = 2
	// ECDSA signing, caller is expected to hash (or not), standard ASN1
	// encoding
	SigningProfile_PROFILE_ECDSA_NOHASH SigningProfile = 3
)

// Enum value maps for SigningProfile.
var (
	SigningProfile_name = map[int32]string{
		0: "PROFILE_BC_ECDSA_SHA256",
		1: "PROFILE_BC_ECDSA_SHA512",
		2: "PROFILE_ECDSA_SHA256",
		3: "PROFILE_ECDSA_NOHASH",
	}
	SigningProfile_value = map[string]int32{
		"PROFILE_BC_ECDSA_SHA256": 0,
		"PROFILE_BC_ECDSA_SHA512": 1,
		"PROFILE_ECDSA_SHA256":    2,
		"PROFILE_ECDSA_NOHASH":    3,
	}
)

func (x SigningProfile) Enum() *SigningProfile {
	p := new(SigningProfile)
	*p = x
	return p
}

func (x SigningProfile) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SigningProfile) Descriptor() protoreflect.EnumDescriptor {
	return file_keystone_base_proto_enumTypes[1].Descriptor()
}

func (SigningProfile) Type() protoreflect.EnumType {
	return &file_keystone_base_proto_enumTypes[1]
}

func (x SigningProfile) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SigningProfile.Descriptor instead.
func (SigningProfile) EnumDescriptor() ([]byte, []int) {
	return file_keystone_base_proto_rawDescGZIP(), []int{1}
}

type KeySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Label string          `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Size  int32           `protobuf:"fixed32,3,opt,name=size,proto3" json:"size,omitempty"`
	Algo  KeygenAlgorithm `protobuf:"varint,4,opt,name=algo,proto3,enum=keystone.KeygenAlgorithm" json:"algo,omitempty"`
}

func (x *KeySpec) Reset() {
	*x = KeySpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keystone_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeySpec) ProtoMessage() {}

func (x *KeySpec) ProtoReflect() protoreflect.Message {
	mi := &file_keystone_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeySpec.ProtoReflect.Descriptor instead.
func (*KeySpec) Descriptor() ([]byte, []int) {
	return file_keystone_base_proto_rawDescGZIP(), []int{0}
}

func (x *KeySpec) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *KeySpec) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *KeySpec) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *KeySpec) GetAlgo() KeygenAlgorithm {
	if x != nil {
		return x.Algo
	}
	return KeygenAlgorithm_KEYGEN_SECP256K1
}

type KeyMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	KeyId int32           `protobuf:"varint,2,opt,name=keyId,proto3" json:"keyId,omitempty"`
	Size  int32           `protobuf:"fixed32,3,opt,name=size,proto3" json:"size,omitempty"`
	Algo  KeygenAlgorithm `protobuf:"varint,4,opt,name=algo,proto3,enum=keystone.KeygenAlgorithm" json:"algo,omitempty"`
}

func (x *KeyMetadata) Reset() {
	*x = KeyMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keystone_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyMetadata) ProtoMessage() {}

func (x *KeyMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_keystone_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyMetadata.ProtoReflect.Descriptor instead.
func (*KeyMetadata) Descriptor() ([]byte, []int) {
	return file_keystone_base_proto_rawDescGZIP(), []int{1}
}

func (x *KeyMetadata) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *KeyMetadata) GetKeyId() int32 {
	if x != nil {
		return x.KeyId
	}
	return 0
}

func (x *KeyMetadata) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *KeyMetadata) GetAlgo() KeygenAlgorithm {
	if x != nil {
		return x.Algo
	}
	return KeygenAlgorithm_KEYGEN_SECP256K1
}

type KeyRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InResponseTo int32   `protobuf:"varint,1,opt,name=inResponseTo,proto3" json:"inResponseTo,omitempty"`
	Label        *string `protobuf:"bytes,2,opt,name=label,proto3,oneof" json:"label,omitempty"`
}

func (x *KeyRef) Reset() {
	*x = KeyRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keystone_base_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyRef) ProtoMessage() {}

func (x *KeyRef) ProtoReflect() protoreflect.Message {
	mi := &file_keystone_base_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyRef.ProtoReflect.Descriptor instead.
func (*KeyRef) Descriptor() ([]byte, []int) {
	return file_keystone_base_proto_rawDescGZIP(), []int{2}
}

func (x *KeyRef) GetInResponseTo() int32 {
	if x != nil {
		return x.InResponseTo
	}
	return 0
}

func (x *KeyRef) GetLabel() string {
	if x != nil && x.Label != nil {
		return *x.Label
	}
	return ""
}

type PublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InResponseTo int32  `protobuf:"varint,1,opt,name=inResponseTo,proto3" json:"inResponseTo,omitempty"`
	Label        string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	KeyBytes     []byte `protobuf:"bytes,3,opt,name=keyBytes,proto3" json:"keyBytes,omitempty"`
}

func (x *PublicKey) Reset() {
	*x = PublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keystone_base_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicKey) ProtoMessage() {}

func (x *PublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_keystone_base_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicKey.ProtoReflect.Descriptor instead.
func (*PublicKey) Descriptor() ([]byte, []int) {
	return file_keystone_base_proto_rawDescGZIP(), []int{3}
}

func (x *PublicKey) GetInResponseTo() int32 {
	if x != nil {
		return x.InResponseTo
	}
	return 0
}

func (x *PublicKey) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *PublicKey) GetKeyBytes() []byte {
	if x != nil {
		return x.KeyBytes
	}
	return nil
}

type Signable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*Signable_SignableBytes
	//	*Signable_Txref
	Data isSignable_Data `protobuf_oneof:"data"`
}

func (x *Signable) Reset() {
	*x = Signable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keystone_base_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signable) ProtoMessage() {}

func (x *Signable) ProtoReflect() protoreflect.Message {
	mi := &file_keystone_base_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signable.ProtoReflect.Descriptor instead.
func (*Signable) Descriptor() ([]byte, []int) {
	return file_keystone_base_proto_rawDescGZIP(), []int{4}
}

func (m *Signable) GetData() isSignable_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *Signable) GetSignableBytes() []byte {
	if x, ok := x.GetData().(*Signable_SignableBytes); ok {
		return x.SignableBytes
	}
	return nil
}

func (x *Signable) GetTxref() string {
	if x, ok := x.GetData().(*Signable_Txref); ok {
		return x.Txref
	}
	return ""
}

type isSignable_Data interface {
	isSignable_Data()
}

type Signable_SignableBytes struct {
	SignableBytes []byte `protobuf:"bytes,1,opt,name=signableBytes,proto3,oneof"`
}

type Signable_Txref struct {
	Txref string `protobuf:"bytes,2,opt,name=txref,proto3,oneof"`
}

func (*Signable_SignableBytes) isSignable_Data() {}

func (*Signable_Txref) isSignable_Data() {}

type Signed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InResponseTo int32 `protobuf:"varint,1,opt,name=inResponseTo,proto3" json:"inResponseTo,omitempty"`
	// Types that are assignable to Data:
	//	*Signed_SignedBytes
	//	*Signed_SignedTxRef
	//	*Signed_Error
	Data isSigned_Data `protobuf_oneof:"data"`
}

func (x *Signed) Reset() {
	*x = Signed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keystone_base_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signed) ProtoMessage() {}

func (x *Signed) ProtoReflect() protoreflect.Message {
	mi := &file_keystone_base_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signed.ProtoReflect.Descriptor instead.
func (*Signed) Descriptor() ([]byte, []int) {
	return file_keystone_base_proto_rawDescGZIP(), []int{5}
}

func (x *Signed) GetInResponseTo() int32 {
	if x != nil {
		return x.InResponseTo
	}
	return 0
}

func (m *Signed) GetData() isSigned_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *Signed) GetSignedBytes() []byte {
	if x, ok := x.GetData().(*Signed_SignedBytes); ok {
		return x.SignedBytes
	}
	return nil
}

func (x *Signed) GetSignedTxRef() string {
	if x, ok := x.GetData().(*Signed_SignedTxRef); ok {
		return x.SignedTxRef
	}
	return ""
}

func (x *Signed) GetError() int32 {
	if x, ok := x.GetData().(*Signed_Error); ok {
		return x.Error
	}
	return 0
}

type isSigned_Data interface {
	isSigned_Data()
}

type Signed_SignedBytes struct {
	SignedBytes []byte `protobuf:"bytes,2,opt,name=signedBytes,proto3,oneof"`
}

type Signed_SignedTxRef struct {
	SignedTxRef string `protobuf:"bytes,3,opt,name=signedTxRef,proto3,oneof"`
}

type Signed_Error struct {
	Error int32 `protobuf:"varint,4,opt,name=error,proto3,oneof"`
}

func (*Signed_SignedBytes) isSigned_Data() {}

func (*Signed_SignedTxRef) isSigned_Data() {}

func (*Signed_Error) isSigned_Data() {}

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int32          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	KeySpec        *KeySpec       `protobuf:"bytes,2,opt,name=keySpec,proto3" json:"keySpec,omitempty"`
	SigningProfile SigningProfile `protobuf:"varint,3,opt,name=signingProfile,proto3,enum=keystone.SigningProfile" json:"signingProfile,omitempty"`
	Content        *Signable      `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keystone_base_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_keystone_base_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_keystone_base_proto_rawDescGZIP(), []int{6}
}

func (x *Msg) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Msg) GetKeySpec() *KeySpec {
	if x != nil {
		return x.KeySpec
	}
	return nil
}

func (x *Msg) GetSigningProfile() SigningProfile {
	if x != nil {
		return x.SigningProfile
	}
	return SigningProfile_PROFILE_BC_ECDSA_SHA256
}

func (x *Msg) GetContent() *Signable {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_keystone_base_proto protoreflect.FileDescriptor

var file_keystone_base_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x22,
	0x72, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x61, 0x6c, 0x67, 0x6f, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x2e, 0x6b, 0x65,
	0x79, 0x67, 0x65, 0x6e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x04, 0x61,
	0x6c, 0x67, 0x6f, 0x22, 0x76, 0x0a, 0x0b, 0x6b, 0x65, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x2d, 0x0a, 0x04,
	0x61, 0x6c, 0x67, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6b, 0x65, 0x79,
	0x73, 0x74, 0x6f, 0x6e, 0x65, 0x2e, 0x6b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x41, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x04, 0x61, 0x6c, 0x67, 0x6f, 0x22, 0x51, 0x0a, 0x06, 0x6b,
	0x65, 0x79, 0x52, 0x65, 0x66, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x54, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x6f, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x61,
	0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x6f, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x22, 0x52, 0x0a, 0x08, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x26, 0x0a,
	0x0d, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x05, 0x74, 0x78, 0x72, 0x65, 0x66, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x74, 0x78, 0x72, 0x65, 0x66, 0x42, 0x06, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x94, 0x01, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x54, 0x6f, 0x12, 0x22, 0x0a, 0x0b, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0b, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x54, 0x78, 0x52, 0x65, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0b, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x78, 0x52, 0x65, 0x66, 0x12, 0x16, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xb2, 0x01, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x53, 0x70, 0x65, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x6e, 0x65,
	0x2e, 0x6b, 0x65, 0x79, 0x53, 0x70, 0x65, 0x63, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x53, 0x70, 0x65,
	0x63, 0x12, 0x40, 0x0a, 0x0e, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6b, 0x65, 0x79, 0x73,
	0x74, 0x6f, 0x6e, 0x65, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x52, 0x0e, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x2e,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2a, 0x51, 0x0a, 0x0f, 0x6b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x41, 0x6c, 0x67, 0x6f, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x12, 0x14, 0x0a, 0x10, 0x4b, 0x45, 0x59, 0x47, 0x45, 0x4e, 0x5f, 0x53,
	0x45, 0x43, 0x50, 0x32, 0x35, 0x36, 0x4b, 0x31, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x4b, 0x45,
	0x59, 0x47, 0x45, 0x4e, 0x5f, 0x53, 0x45, 0x43, 0x50, 0x32, 0x35, 0x36, 0x52, 0x31, 0x10, 0x01,
	0x12, 0x12, 0x0a, 0x0e, 0x4b, 0x45, 0x59, 0x47, 0x45, 0x4e, 0x5f, 0x45, 0x44, 0x32, 0x35, 0x35,
	0x31, 0x39, 0x10, 0x02, 0x2a, 0x7e, 0x0a, 0x0e, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c,
	0x45, 0x5f, 0x42, 0x43, 0x5f, 0x45, 0x43, 0x44, 0x53, 0x41, 0x5f, 0x53, 0x48, 0x41, 0x32, 0x35,
	0x36, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x42,
	0x43, 0x5f, 0x45, 0x43, 0x44, 0x53, 0x41, 0x5f, 0x53, 0x48, 0x41, 0x35, 0x31, 0x32, 0x10, 0x01,
	0x12, 0x18, 0x0a, 0x14, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x45, 0x43, 0x44, 0x53,
	0x41, 0x5f, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x52,
	0x4f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x45, 0x43, 0x44, 0x53, 0x41, 0x5f, 0x4e, 0x4f, 0x48, 0x41,
	0x53, 0x48, 0x10, 0x03, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f,
	0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_keystone_base_proto_rawDescOnce sync.Once
	file_keystone_base_proto_rawDescData = file_keystone_base_proto_rawDesc
)

func file_keystone_base_proto_rawDescGZIP() []byte {
	file_keystone_base_proto_rawDescOnce.Do(func() {
		file_keystone_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_keystone_base_proto_rawDescData)
	})
	return file_keystone_base_proto_rawDescData
}

var file_keystone_base_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_keystone_base_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_keystone_base_proto_goTypes = []interface{}{
	(KeygenAlgorithm)(0), // 0: keystone.keygenAlgorithm
	(SigningProfile)(0),  // 1: keystone.signingProfile
	(*KeySpec)(nil),      // 2: keystone.keySpec
	(*KeyMetadata)(nil),  // 3: keystone.keyMetadata
	(*KeyRef)(nil),       // 4: keystone.keyRef
	(*PublicKey)(nil),    // 5: keystone.publicKey
	(*Signable)(nil),     // 6: keystone.signable
	(*Signed)(nil),       // 7: keystone.signed
	(*Msg)(nil),          // 8: keystone.msg
}
var file_keystone_base_proto_depIdxs = []int32{
	0, // 0: keystone.keySpec.algo:type_name -> keystone.keygenAlgorithm
	0, // 1: keystone.keyMetadata.algo:type_name -> keystone.keygenAlgorithm
	2, // 2: keystone.msg.keySpec:type_name -> keystone.keySpec
	1, // 3: keystone.msg.signingProfile:type_name -> keystone.signingProfile
	6, // 4: keystone.msg.content:type_name -> keystone.signable
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_keystone_base_proto_init() }
func file_keystone_base_proto_init() {
	if File_keystone_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_keystone_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeySpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_keystone_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_keystone_base_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyRef); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_keystone_base_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicKey); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_keystone_base_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signable); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_keystone_base_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signed); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_keystone_base_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_keystone_base_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_keystone_base_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Signable_SignableBytes)(nil),
		(*Signable_Txref)(nil),
	}
	file_keystone_base_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*Signed_SignedBytes)(nil),
		(*Signed_SignedTxRef)(nil),
		(*Signed_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_keystone_base_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_keystone_base_proto_goTypes,
		DependencyIndexes: file_keystone_base_proto_depIdxs,
		EnumInfos:         file_keystone_base_proto_enumTypes,
		MessageInfos:      file_keystone_base_proto_msgTypes,
	}.Build()
	File_keystone_base_proto = out.File
	file_keystone_base_proto_rawDesc = nil
	file_keystone_base_proto_goTypes = nil
	file_keystone_base_proto_depIdxs = nil
}
