// Copyright 2016 Google Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or at
// https://developers.google.com/open-source/licenses/bsd
//
// This file defines Widevine Pssh Data proto format.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: widevine_pssh_data.proto

package proto

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

type WidevinePsshData_Algorithm int32

const (
	WidevinePsshData_UNENCRYPTED WidevinePsshData_Algorithm = 0
	WidevinePsshData_AESCTR      WidevinePsshData_Algorithm = 1
)

// Enum value maps for WidevinePsshData_Algorithm.
var (
	WidevinePsshData_Algorithm_name = map[int32]string{
		0: "UNENCRYPTED",
		1: "AESCTR",
	}
	WidevinePsshData_Algorithm_value = map[string]int32{
		"UNENCRYPTED": 0,
		"AESCTR":      1,
	}
)

func (x WidevinePsshData_Algorithm) Enum() *WidevinePsshData_Algorithm {
	p := new(WidevinePsshData_Algorithm)
	*p = x
	return p
}

func (x WidevinePsshData_Algorithm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WidevinePsshData_Algorithm) Descriptor() protoreflect.EnumDescriptor {
	return file_widevine_pssh_data_proto_enumTypes[0].Descriptor()
}

func (WidevinePsshData_Algorithm) Type() protoreflect.EnumType {
	return &file_widevine_pssh_data_proto_enumTypes[0]
}

func (x WidevinePsshData_Algorithm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *WidevinePsshData_Algorithm) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = WidevinePsshData_Algorithm(num)
	return nil
}

// Deprecated: Use WidevinePsshData_Algorithm.Descriptor instead.
func (WidevinePsshData_Algorithm) EnumDescriptor() ([]byte, []int) {
	return file_widevine_pssh_data_proto_rawDescGZIP(), []int{0, 0}
}

type WidevinePsshData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Algorithm *WidevinePsshData_Algorithm `protobuf:"varint,1,opt,name=algorithm,enum=shaka.media.WidevinePsshData_Algorithm" json:"algorithm,omitempty"`
	KeyId     [][]byte                    `protobuf:"bytes,2,rep,name=key_id,json=keyId" json:"key_id,omitempty"`
	// Content provider name.
	Provider *string `protobuf:"bytes,3,opt,name=provider" json:"provider,omitempty"`
	// A content identifier, specified by content provider.
	ContentId []byte `protobuf:"bytes,4,opt,name=content_id,json=contentId" json:"content_id,omitempty"`
	// The name of a registered policy to be used for this asset.
	Policy *string `protobuf:"bytes,6,opt,name=policy" json:"policy,omitempty"`
	// Crypto period index, for media using key rotation.
	CryptoPeriodIndex *uint32 `protobuf:"varint,7,opt,name=crypto_period_index,json=cryptoPeriodIndex" json:"crypto_period_index,omitempty"`
	// Optional protected context for group content. The grouped_license is a
	// serialized SignedMessage.
	GroupedLicense []byte `protobuf:"bytes,8,opt,name=grouped_license,json=groupedLicense" json:"grouped_license,omitempty"`
	// Protection scheme identifying the encryption algorithm. Represented as one
	// of the following 4CC values: 'cenc' (AES-CTR), 'cbc1' (AES-CBC),
	// 'cens' (AES-CTR subsample), 'cbcs' (AES-CBC subsample).
	ProtectionScheme *uint32 `protobuf:"varint,9,opt,name=protection_scheme,json=protectionScheme" json:"protection_scheme,omitempty"`
}

func (x *WidevinePsshData) Reset() {
	*x = WidevinePsshData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_widevine_pssh_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WidevinePsshData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WidevinePsshData) ProtoMessage() {}

func (x *WidevinePsshData) ProtoReflect() protoreflect.Message {
	mi := &file_widevine_pssh_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WidevinePsshData.ProtoReflect.Descriptor instead.
func (*WidevinePsshData) Descriptor() ([]byte, []int) {
	return file_widevine_pssh_data_proto_rawDescGZIP(), []int{0}
}

func (x *WidevinePsshData) GetAlgorithm() WidevinePsshData_Algorithm {
	if x != nil && x.Algorithm != nil {
		return *x.Algorithm
	}
	return WidevinePsshData_UNENCRYPTED
}

func (x *WidevinePsshData) GetKeyId() [][]byte {
	if x != nil {
		return x.KeyId
	}
	return nil
}

func (x *WidevinePsshData) GetProvider() string {
	if x != nil && x.Provider != nil {
		return *x.Provider
	}
	return ""
}

func (x *WidevinePsshData) GetContentId() []byte {
	if x != nil {
		return x.ContentId
	}
	return nil
}

func (x *WidevinePsshData) GetPolicy() string {
	if x != nil && x.Policy != nil {
		return *x.Policy
	}
	return ""
}

func (x *WidevinePsshData) GetCryptoPeriodIndex() uint32 {
	if x != nil && x.CryptoPeriodIndex != nil {
		return *x.CryptoPeriodIndex
	}
	return 0
}

func (x *WidevinePsshData) GetGroupedLicense() []byte {
	if x != nil {
		return x.GroupedLicense
	}
	return nil
}

func (x *WidevinePsshData) GetProtectionScheme() uint32 {
	if x != nil && x.ProtectionScheme != nil {
		return *x.ProtectionScheme
	}
	return 0
}

// Derived from WidevinePsshData. The JSON format of this proto is used in
// Widevine HLS DRM signaling v1.
// We cannot build JSON from WidevinePsshData as |key_id| is required to be in
// hex format, while |bytes| type is translated to base64 by JSON formatter, so
// we have to use |string| type and do hex conversion in the code.
type WidevineHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyIds []string `protobuf:"bytes,2,rep,name=key_ids,json=keyIds" json:"key_ids,omitempty"`
	// Content provider name.
	Provider *string `protobuf:"bytes,3,opt,name=provider" json:"provider,omitempty"`
	// A content identifier, specified by content provider.
	ContentId []byte `protobuf:"bytes,4,opt,name=content_id,json=contentId" json:"content_id,omitempty"`
}

func (x *WidevineHeader) Reset() {
	*x = WidevineHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_widevine_pssh_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WidevineHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WidevineHeader) ProtoMessage() {}

func (x *WidevineHeader) ProtoReflect() protoreflect.Message {
	mi := &file_widevine_pssh_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WidevineHeader.ProtoReflect.Descriptor instead.
func (*WidevineHeader) Descriptor() ([]byte, []int) {
	return file_widevine_pssh_data_proto_rawDescGZIP(), []int{1}
}

func (x *WidevineHeader) GetKeyIds() []string {
	if x != nil {
		return x.KeyIds
	}
	return nil
}

func (x *WidevineHeader) GetProvider() string {
	if x != nil && x.Provider != nil {
		return *x.Provider
	}
	return ""
}

func (x *WidevineHeader) GetContentId() []byte {
	if x != nil {
		return x.ContentId
	}
	return nil
}

var File_widevine_pssh_data_proto protoreflect.FileDescriptor

var file_widevine_pssh_data_proto_rawDesc = []byte{
	0x0a, 0x18, 0x77, 0x69, 0x64, 0x65, 0x76, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x73, 0x73, 0x68, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x68, 0x61, 0x6b,
	0x61, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x22, 0xf3, 0x02, 0x0a, 0x10, 0x57, 0x69, 0x64, 0x65,
	0x76, 0x69, 0x6e, 0x65, 0x50, 0x73, 0x73, 0x68, 0x44, 0x61, 0x74, 0x61, 0x12, 0x45, 0x0a, 0x09,
	0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x27, 0x2e, 0x73, 0x68, 0x61, 0x6b, 0x61, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x57, 0x69,
	0x64, 0x65, 0x76, 0x69, 0x6e, 0x65, 0x50, 0x73, 0x73, 0x68, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x41,
	0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x12, 0x15, 0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0c, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2e, 0x0a,
	0x13, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x27, 0x0a,
	0x0f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x65, 0x64, 0x5f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x65, 0x64, 0x4c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x65, 0x22, 0x28, 0x0a, 0x09, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d,
	0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x45, 0x4e, 0x43, 0x52, 0x59, 0x50, 0x54, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x45, 0x53, 0x43, 0x54, 0x52, 0x10, 0x01, 0x22, 0x64, 0x0a,
	0x0e, 0x57, 0x69, 0x64, 0x65, 0x76, 0x69, 0x6e, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x17, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x49, 0x64,
}

var (
	file_widevine_pssh_data_proto_rawDescOnce sync.Once
	file_widevine_pssh_data_proto_rawDescData = file_widevine_pssh_data_proto_rawDesc
)

func file_widevine_pssh_data_proto_rawDescGZIP() []byte {
	file_widevine_pssh_data_proto_rawDescOnce.Do(func() {
		file_widevine_pssh_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_widevine_pssh_data_proto_rawDescData)
	})
	return file_widevine_pssh_data_proto_rawDescData
}

var file_widevine_pssh_data_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_widevine_pssh_data_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_widevine_pssh_data_proto_goTypes = []interface{}{
	(WidevinePsshData_Algorithm)(0), // 0: shaka.media.WidevinePsshData.Algorithm
	(*WidevinePsshData)(nil),        // 1: shaka.media.WidevinePsshData
	(*WidevineHeader)(nil),          // 2: shaka.media.WidevineHeader
}
var file_widevine_pssh_data_proto_depIdxs = []int32{
	0, // 0: shaka.media.WidevinePsshData.algorithm:type_name -> shaka.media.WidevinePsshData.Algorithm
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_widevine_pssh_data_proto_init() }
func file_widevine_pssh_data_proto_init() {
	if File_widevine_pssh_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_widevine_pssh_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WidevinePsshData); i {
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
		file_widevine_pssh_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WidevineHeader); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_widevine_pssh_data_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_widevine_pssh_data_proto_goTypes,
		DependencyIndexes: file_widevine_pssh_data_proto_depIdxs,
		EnumInfos:         file_widevine_pssh_data_proto_enumTypes,
		MessageInfos:      file_widevine_pssh_data_proto_msgTypes,
	}.Build()
	File_widevine_pssh_data_proto = out.File
	file_widevine_pssh_data_proto_rawDesc = nil
	file_widevine_pssh_data_proto_goTypes = nil
	file_widevine_pssh_data_proto_depIdxs = nil
}
