// Copyright © 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: gogo.proto

package test

import (
	_ "github.com/TheThingsIndustries/protoc-gen-go-json/annotations"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MessageWithGoGoOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EuiWithCustomName                   []byte                 `protobuf:"bytes,1,opt,name=eui_with_custom_name,json=euiWithCustomName,proto3" json:"eui_with_custom_name,omitempty"`
	EuiWithCustomNameAndType            []byte                 `protobuf:"bytes,2,opt,name=eui_with_custom_name_and_type,json=euiWithCustomNameAndType,proto3" json:"eui_with_custom_name_and_type,omitempty"`
	NonNullableEuiWithCustomNameAndType []byte                 `protobuf:"bytes,3,opt,name=non_nullable_eui_with_custom_name_and_type,json=nonNullableEuiWithCustomNameAndType,proto3" json:"non_nullable_eui_with_custom_name_and_type,omitempty"`
	EuisWithCustomNameAndType           [][]byte               `protobuf:"bytes,4,rep,name=euis_with_custom_name_and_type,json=euisWithCustomNameAndType,proto3" json:"euis_with_custom_name_and_type,omitempty"`
	Duration                            *durationpb.Duration   `protobuf:"bytes,5,opt,name=duration,proto3" json:"duration,omitempty"`
	NonNullableDuration                 *durationpb.Duration   `protobuf:"bytes,6,opt,name=non_nullable_duration,json=nonNullableDuration,proto3" json:"non_nullable_duration,omitempty"`
	Timestamp                           *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	NonNullableTimestamp                *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=non_nullable_timestamp,json=nonNullableTimestamp,proto3" json:"non_nullable_timestamp,omitempty"`
}

func (x *MessageWithGoGoOptions) Reset() {
	*x = MessageWithGoGoOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gogo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithGoGoOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithGoGoOptions) ProtoMessage() {}

func (x *MessageWithGoGoOptions) ProtoReflect() protoreflect.Message {
	mi := &file_gogo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithGoGoOptions.ProtoReflect.Descriptor instead.
func (*MessageWithGoGoOptions) Descriptor() ([]byte, []int) {
	return file_gogo_proto_rawDescGZIP(), []int{0}
}

func (x *MessageWithGoGoOptions) GetEuiWithCustomName() []byte {
	if x != nil {
		return x.EuiWithCustomName
	}
	return nil
}

func (x *MessageWithGoGoOptions) GetEuiWithCustomNameAndType() []byte {
	if x != nil {
		return x.EuiWithCustomNameAndType
	}
	return nil
}

func (x *MessageWithGoGoOptions) GetNonNullableEuiWithCustomNameAndType() []byte {
	if x != nil {
		return x.NonNullableEuiWithCustomNameAndType
	}
	return nil
}

func (x *MessageWithGoGoOptions) GetEuisWithCustomNameAndType() [][]byte {
	if x != nil {
		return x.EuisWithCustomNameAndType
	}
	return nil
}

func (x *MessageWithGoGoOptions) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *MessageWithGoGoOptions) GetNonNullableDuration() *durationpb.Duration {
	if x != nil {
		return x.NonNullableDuration
	}
	return nil
}

func (x *MessageWithGoGoOptions) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *MessageWithGoGoOptions) GetNonNullableTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.NonNullableTimestamp
	}
	return nil
}

type SubMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *SubMessage) Reset() {
	*x = SubMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gogo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubMessage) ProtoMessage() {}

func (x *SubMessage) ProtoReflect() protoreflect.Message {
	mi := &file_gogo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubMessage.ProtoReflect.Descriptor instead.
func (*SubMessage) Descriptor() ([]byte, []int) {
	return file_gogo_proto_rawDescGZIP(), []int{1}
}

func (x *SubMessage) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

type SubMessageWithoutMarshalers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OtherField string `protobuf:"bytes,1,opt,name=other_field,json=otherField,proto3" json:"other_field,omitempty"`
}

func (x *SubMessageWithoutMarshalers) Reset() {
	*x = SubMessageWithoutMarshalers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gogo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubMessageWithoutMarshalers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubMessageWithoutMarshalers) ProtoMessage() {}

func (x *SubMessageWithoutMarshalers) ProtoReflect() protoreflect.Message {
	mi := &file_gogo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubMessageWithoutMarshalers.ProtoReflect.Descriptor instead.
func (*SubMessageWithoutMarshalers) Descriptor() ([]byte, []int) {
	return file_gogo_proto_rawDescGZIP(), []int{2}
}

func (x *SubMessageWithoutMarshalers) GetOtherField() string {
	if x != nil {
		return x.OtherField
	}
	return ""
}

type MessageWithNullable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sub       *SubMessage                    `protobuf:"bytes,1,opt,name=sub,proto3" json:"sub,omitempty"`
	Subs      []*SubMessage                  `protobuf:"bytes,2,rep,name=subs,proto3" json:"subs,omitempty"`
	OtherSub  *SubMessageWithoutMarshalers   `protobuf:"bytes,3,opt,name=other_sub,json=otherSub,proto3" json:"other_sub,omitempty"`
	OtherSubs []*SubMessageWithoutMarshalers `protobuf:"bytes,4,rep,name=other_subs,json=otherSubs,proto3" json:"other_subs,omitempty"`
}

func (x *MessageWithNullable) Reset() {
	*x = MessageWithNullable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gogo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithNullable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithNullable) ProtoMessage() {}

func (x *MessageWithNullable) ProtoReflect() protoreflect.Message {
	mi := &file_gogo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithNullable.ProtoReflect.Descriptor instead.
func (*MessageWithNullable) Descriptor() ([]byte, []int) {
	return file_gogo_proto_rawDescGZIP(), []int{3}
}

func (x *MessageWithNullable) GetSub() *SubMessage {
	if x != nil {
		return x.Sub
	}
	return nil
}

func (x *MessageWithNullable) GetSubs() []*SubMessage {
	if x != nil {
		return x.Subs
	}
	return nil
}

func (x *MessageWithNullable) GetOtherSub() *SubMessageWithoutMarshalers {
	if x != nil {
		return x.OtherSub
	}
	return nil
}

func (x *MessageWithNullable) GetOtherSubs() []*SubMessageWithoutMarshalers {
	if x != nil {
		return x.OtherSubs
	}
	return nil
}

type MessageWithEmbedded struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sub      *SubMessage                  `protobuf:"bytes,1,opt,name=sub,proto3" json:"sub,omitempty"`
	OtherSub *SubMessageWithoutMarshalers `protobuf:"bytes,2,opt,name=other_sub,json=otherSub,proto3" json:"other_sub,omitempty"`
}

func (x *MessageWithEmbedded) Reset() {
	*x = MessageWithEmbedded{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gogo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithEmbedded) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithEmbedded) ProtoMessage() {}

func (x *MessageWithEmbedded) ProtoReflect() protoreflect.Message {
	mi := &file_gogo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithEmbedded.ProtoReflect.Descriptor instead.
func (*MessageWithEmbedded) Descriptor() ([]byte, []int) {
	return file_gogo_proto_rawDescGZIP(), []int{4}
}

func (x *MessageWithEmbedded) GetSub() *SubMessage {
	if x != nil {
		return x.Sub
	}
	return nil
}

func (x *MessageWithEmbedded) GetOtherSub() *SubMessageWithoutMarshalers {
	if x != nil {
		return x.OtherSub
	}
	return nil
}

type MessageWithNullableEmbedded struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sub      *SubMessage                  `protobuf:"bytes,1,opt,name=sub,proto3" json:"sub,omitempty"`
	OtherSub *SubMessageWithoutMarshalers `protobuf:"bytes,2,opt,name=other_sub,json=otherSub,proto3" json:"other_sub,omitempty"`
}

func (x *MessageWithNullableEmbedded) Reset() {
	*x = MessageWithNullableEmbedded{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gogo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithNullableEmbedded) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithNullableEmbedded) ProtoMessage() {}

func (x *MessageWithNullableEmbedded) ProtoReflect() protoreflect.Message {
	mi := &file_gogo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithNullableEmbedded.ProtoReflect.Descriptor instead.
func (*MessageWithNullableEmbedded) Descriptor() ([]byte, []int) {
	return file_gogo_proto_rawDescGZIP(), []int{5}
}

func (x *MessageWithNullableEmbedded) GetSub() *SubMessage {
	if x != nil {
		return x.Sub
	}
	return nil
}

func (x *MessageWithNullableEmbedded) GetOtherSub() *SubMessageWithoutMarshalers {
	if x != nil {
		return x.OtherSub
	}
	return nil
}

var File_gogo_proto protoreflect.FileDescriptor

var file_gogo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x74, 0x68,
	0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67,
	0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x0b, 0x0a, 0x16, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x57, 0x69, 0x74, 0x68, 0x47, 0x6f, 0x47, 0x6f, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x46, 0x0a, 0x14, 0x65, 0x75, 0x69, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x15, 0xe2,
	0xde, 0x1f, 0x11, 0x45, 0x55, 0x49, 0x57, 0x69, 0x74, 0x68, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x4e, 0x61, 0x6d, 0x65, 0x52, 0x11, 0x65, 0x75, 0x69, 0x57, 0x69, 0x74, 0x68, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0xbd, 0x02, 0x0a, 0x1d, 0x65, 0x75, 0x69, 0x5f,
	0x77, 0x69, 0x74, 0x68, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42,
	0xfb, 0x01, 0xe2, 0xde, 0x1f, 0x18, 0x45, 0x55, 0x49, 0x57, 0x69, 0x74, 0x68, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x41, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0xda, 0xde,
	0x1f, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x68, 0x65,
	0x54, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x6a,
	0x73, 0x6f, 0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45,
	0x55, 0x49, 0x36, 0x34, 0xea, 0xaa, 0x19, 0x94, 0x01, 0x0a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x68, 0x65, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x49,
	0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x6a, 0x73, 0x6f, 0x6e, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x48,
	0x45, 0x58, 0x12, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54,
	0x68, 0x65, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f,
	0x2d, 0x6a, 0x73, 0x6f, 0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x55, 0x6e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x48, 0x45, 0x58, 0x52, 0x18, 0x65,
	0x75, 0x69, 0x57, 0x69, 0x74, 0x68, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65,
	0x41, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0xe4, 0x02, 0x0a, 0x2a, 0x6e, 0x6f, 0x6e, 0x5f,
	0x6e, 0x75, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x65, 0x75, 0x69, 0x5f, 0x77, 0x69, 0x74,
	0x68, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x61, 0x6e,
	0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x8a, 0x02, 0xe2,
	0xde, 0x1f, 0x23, 0x4e, 0x6f, 0x6e, 0x4e, 0x75, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x55,
	0x49, 0x57, 0x69, 0x74, 0x68, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x41,
	0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0xda, 0xde, 0x1f, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x68, 0x65, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x49, 0x6e,
	0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x6a, 0x73, 0x6f, 0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x55, 0x49, 0x36, 0x34, 0xc8, 0xde, 0x1f, 0x00,
	0xea, 0xaa, 0x19, 0x94, 0x01, 0x0a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x54, 0x68, 0x65, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x49, 0x6e, 0x64, 0x75, 0x73,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2d, 0x6a, 0x73, 0x6f, 0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x48, 0x45, 0x58, 0x12, 0x49,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x68, 0x65, 0x54, 0x68,
	0x69, 0x6e, 0x67, 0x73, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x6a, 0x73, 0x6f,
	0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x55, 0x6e, 0x6d,
	0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x48, 0x45, 0x58, 0x52, 0x23, 0x6e, 0x6f, 0x6e, 0x4e, 0x75,
	0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x75, 0x69, 0x57, 0x69, 0x74, 0x68, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x41, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0xca,
	0x02, 0x0a, 0x1e, 0x65, 0x75, 0x69, 0x73, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0c, 0x42, 0x86, 0x02, 0xe2, 0xde, 0x1f, 0x19, 0x45, 0x55,
	0x49, 0x73, 0x57, 0x69, 0x74, 0x68, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65,
	0x41, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0xda, 0xde, 0x1f, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x68, 0x65, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x49,
	0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x6a, 0x73, 0x6f, 0x6e, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x55, 0x49, 0x36, 0x34, 0xea, 0xaa, 0x19,
	0x9e, 0x01, 0x0a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54,
	0x68, 0x65, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f,
	0x2d, 0x6a, 0x73, 0x6f, 0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x4d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x48, 0x45, 0x58, 0x41, 0x72, 0x72, 0x61, 0x79,
	0x12, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x68, 0x65,
	0x54, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x6a,
	0x73, 0x6f, 0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x55,
	0x6e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x48, 0x45, 0x58, 0x41, 0x72, 0x72, 0x61, 0x79,
	0x52, 0x19, 0x65, 0x75, 0x69, 0x73, 0x57, 0x69, 0x74, 0x68, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x4e, 0x61, 0x6d, 0x65, 0x41, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x04, 0x98, 0xdf, 0x1f, 0x01, 0x52, 0x08,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x57, 0x0a, 0x15, 0x6e, 0x6f, 0x6e, 0x5f,
	0x6e, 0x75, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x08, 0xc8, 0xde, 0x1f, 0x00, 0x98, 0xdf, 0x1f, 0x01, 0x52, 0x13, 0x6e, 0x6f,
	0x6e, 0x4e, 0x75, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x3e, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x04, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x5a, 0x0a, 0x16, 0x6e, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6c, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xc8,
	0xde, 0x1f, 0x00, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x14, 0x6e, 0x6f, 0x6e, 0x4e, 0x75, 0x6c, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x22, 0x0a,
	0x0a, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x22, 0x48, 0x0a, 0x1b, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57,
	0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x4d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x65, 0x72, 0x73,
	0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x3a, 0x08, 0xea, 0xaa, 0x19, 0x04, 0x08, 0x00, 0x10, 0x00, 0x22, 0xb5, 0x02, 0x0a, 0x13,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x4e, 0x75, 0x6c, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x37, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6a, 0x73, 0x6f,
	0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x03, 0x73, 0x75, 0x62, 0x12, 0x39, 0x0a, 0x04,
	0x73, 0x75, 0x62, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x68, 0x65,
	0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x04, 0xc8, 0xde, 0x1f,
	0x00, 0x52, 0x04, 0x73, 0x75, 0x62, 0x73, 0x12, 0x53, 0x0a, 0x09, 0x6f, 0x74, 0x68, 0x65, 0x72,
	0x5f, 0x73, 0x75, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x74, 0x68, 0x65,
	0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x6f,
	0x75, 0x74, 0x4d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x65, 0x72, 0x73, 0x42, 0x04, 0xc8, 0xde,
	0x1f, 0x00, 0x52, 0x08, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x53, 0x75, 0x62, 0x12, 0x55, 0x0a, 0x0a,
	0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x30, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6a, 0x73, 0x6f,
	0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x57, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x4d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x65,
	0x72, 0x73, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x09, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x53,
	0x75, 0x62, 0x73, 0x22, 0xa3, 0x01, 0x0a, 0x13, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57,
	0x69, 0x74, 0x68, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x12, 0x37, 0x0a, 0x03, 0x73,
	0x75, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x68,
	0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53,
	0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x04, 0xd0, 0xde, 0x1f, 0x01, 0x52,
	0x03, 0x73, 0x75, 0x62, 0x12, 0x53, 0x0a, 0x09, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x75,
	0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x68, 0x69,
	0x6e, 0x67, 0x73, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x75,
	0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x4d,
	0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x65, 0x72, 0x73, 0x42, 0x04, 0xd0, 0xde, 0x1f, 0x01, 0x52,
	0x08, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x53, 0x75, 0x62, 0x22, 0xb3, 0x01, 0x0a, 0x1b, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x4e, 0x75, 0x6c, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x12, 0x3b, 0x0a, 0x03, 0x73, 0x75, 0x62,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x68, 0x69, 0x6e,
	0x67, 0x73, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x75, 0x62,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x08, 0xc8, 0xde, 0x1f, 0x00, 0xd0, 0xde, 0x1f,
	0x01, 0x52, 0x03, 0x73, 0x75, 0x62, 0x12, 0x57, 0x0a, 0x09, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f,
	0x73, 0x75, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x74, 0x68, 0x65, 0x74,
	0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x6f, 0x75,
	0x74, 0x4d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x65, 0x72, 0x73, 0x42, 0x08, 0xc8, 0xde, 0x1f,
	0x00, 0xd0, 0xde, 0x1f, 0x01, 0x52, 0x08, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x53, 0x75, 0x62, 0x42,
	0x40, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x68,
	0x65, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d,
	0x6a, 0x73, 0x6f, 0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0xea, 0xaa, 0x19, 0x04, 0x08, 0x01, 0x10,
	0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gogo_proto_rawDescOnce sync.Once
	file_gogo_proto_rawDescData = file_gogo_proto_rawDesc
)

func file_gogo_proto_rawDescGZIP() []byte {
	file_gogo_proto_rawDescOnce.Do(func() {
		file_gogo_proto_rawDescData = protoimpl.X.CompressGZIP(file_gogo_proto_rawDescData)
	})
	return file_gogo_proto_rawDescData
}

var file_gogo_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_gogo_proto_goTypes = []interface{}{
	(*MessageWithGoGoOptions)(nil),      // 0: thethings.json.test.MessageWithGoGoOptions
	(*SubMessage)(nil),                  // 1: thethings.json.test.SubMessage
	(*SubMessageWithoutMarshalers)(nil), // 2: thethings.json.test.SubMessageWithoutMarshalers
	(*MessageWithNullable)(nil),         // 3: thethings.json.test.MessageWithNullable
	(*MessageWithEmbedded)(nil),         // 4: thethings.json.test.MessageWithEmbedded
	(*MessageWithNullableEmbedded)(nil), // 5: thethings.json.test.MessageWithNullableEmbedded
	(*durationpb.Duration)(nil),         // 6: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil),       // 7: google.protobuf.Timestamp
}
var file_gogo_proto_depIdxs = []int32{
	6,  // 0: thethings.json.test.MessageWithGoGoOptions.duration:type_name -> google.protobuf.Duration
	6,  // 1: thethings.json.test.MessageWithGoGoOptions.non_nullable_duration:type_name -> google.protobuf.Duration
	7,  // 2: thethings.json.test.MessageWithGoGoOptions.timestamp:type_name -> google.protobuf.Timestamp
	7,  // 3: thethings.json.test.MessageWithGoGoOptions.non_nullable_timestamp:type_name -> google.protobuf.Timestamp
	1,  // 4: thethings.json.test.MessageWithNullable.sub:type_name -> thethings.json.test.SubMessage
	1,  // 5: thethings.json.test.MessageWithNullable.subs:type_name -> thethings.json.test.SubMessage
	2,  // 6: thethings.json.test.MessageWithNullable.other_sub:type_name -> thethings.json.test.SubMessageWithoutMarshalers
	2,  // 7: thethings.json.test.MessageWithNullable.other_subs:type_name -> thethings.json.test.SubMessageWithoutMarshalers
	1,  // 8: thethings.json.test.MessageWithEmbedded.sub:type_name -> thethings.json.test.SubMessage
	2,  // 9: thethings.json.test.MessageWithEmbedded.other_sub:type_name -> thethings.json.test.SubMessageWithoutMarshalers
	1,  // 10: thethings.json.test.MessageWithNullableEmbedded.sub:type_name -> thethings.json.test.SubMessage
	2,  // 11: thethings.json.test.MessageWithNullableEmbedded.other_sub:type_name -> thethings.json.test.SubMessageWithoutMarshalers
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_gogo_proto_init() }
func file_gogo_proto_init() {
	if File_gogo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gogo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageWithGoGoOptions); i {
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
		file_gogo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubMessage); i {
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
		file_gogo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubMessageWithoutMarshalers); i {
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
		file_gogo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageWithNullable); i {
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
		file_gogo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageWithEmbedded); i {
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
		file_gogo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageWithNullableEmbedded); i {
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
			RawDescriptor: file_gogo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gogo_proto_goTypes,
		DependencyIndexes: file_gogo_proto_depIdxs,
		MessageInfos:      file_gogo_proto_msgTypes,
	}.Build()
	File_gogo_proto = out.File
	file_gogo_proto_rawDesc = nil
	file_gogo_proto_goTypes = nil
	file_gogo_proto_depIdxs = nil
}
