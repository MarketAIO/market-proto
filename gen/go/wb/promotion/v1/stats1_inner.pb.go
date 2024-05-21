//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document:
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: wb/promotion/v1/stats1_inner.proto

package wbPROMOTION

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

type Stats1Inner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Количество просмотров
	Views int32 `protobuf:"varint,1,opt,name=views,proto3" json:"views,omitempty"`
	// Количество кликов
	Clicks int32 `protobuf:"varint,2,opt,name=clicks,proto3" json:"clicks,omitempty"`
	// Количество добавлений товаров в корзину
	Atbs int32 `protobuf:"varint,3,opt,name=atbs,proto3" json:"atbs,omitempty"`
	// CTR (click-through rate) — показатель кликабельности, отношение числа кликов к количеству показов в рамках медиакампании
	Ctr float32 `protobuf:"fixed32,4,opt,name=ctr,proto3" json:"ctr,omitempty"`
}

func (x *Stats1Inner) Reset() {
	*x = Stats1Inner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_promotion_v1_stats1_inner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stats1Inner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stats1Inner) ProtoMessage() {}

func (x *Stats1Inner) ProtoReflect() protoreflect.Message {
	mi := &file_wb_promotion_v1_stats1_inner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stats1Inner.ProtoReflect.Descriptor instead.
func (*Stats1Inner) Descriptor() ([]byte, []int) {
	return file_wb_promotion_v1_stats1_inner_proto_rawDescGZIP(), []int{0}
}

func (x *Stats1Inner) GetViews() int32 {
	if x != nil {
		return x.Views
	}
	return 0
}

func (x *Stats1Inner) GetClicks() int32 {
	if x != nil {
		return x.Clicks
	}
	return 0
}

func (x *Stats1Inner) GetAtbs() int32 {
	if x != nil {
		return x.Atbs
	}
	return 0
}

func (x *Stats1Inner) GetCtr() float32 {
	if x != nil {
		return x.Ctr
	}
	return 0
}

var File_wb_promotion_v1_stats1_inner_proto protoreflect.FileDescriptor

var file_wb_promotion_v1_stats1_inner_proto_rawDesc = []byte{
	0x0a, 0x22, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x31, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x77, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0x61, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x73, 0x31, 0x49,
	0x6e, 0x6e, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c,
	0x69, 0x63, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x63,
	0x6b, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x74, 0x62, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x61, 0x74, 0x62, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x74, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x03, 0x63, 0x74, 0x72, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f,
	0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x62, 0x50, 0x52, 0x4f, 0x4d, 0x4f, 0x54,
	0x49, 0x4f, 0x4e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wb_promotion_v1_stats1_inner_proto_rawDescOnce sync.Once
	file_wb_promotion_v1_stats1_inner_proto_rawDescData = file_wb_promotion_v1_stats1_inner_proto_rawDesc
)

func file_wb_promotion_v1_stats1_inner_proto_rawDescGZIP() []byte {
	file_wb_promotion_v1_stats1_inner_proto_rawDescOnce.Do(func() {
		file_wb_promotion_v1_stats1_inner_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_promotion_v1_stats1_inner_proto_rawDescData)
	})
	return file_wb_promotion_v1_stats1_inner_proto_rawDescData
}

var file_wb_promotion_v1_stats1_inner_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_promotion_v1_stats1_inner_proto_goTypes = []interface{}{
	(*Stats1Inner)(nil), // 0: wb.promotion.v1.Stats1Inner
}
var file_wb_promotion_v1_stats1_inner_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wb_promotion_v1_stats1_inner_proto_init() }
func file_wb_promotion_v1_stats1_inner_proto_init() {
	if File_wb_promotion_v1_stats1_inner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wb_promotion_v1_stats1_inner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stats1Inner); i {
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
			RawDescriptor: file_wb_promotion_v1_stats1_inner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_promotion_v1_stats1_inner_proto_goTypes,
		DependencyIndexes: file_wb_promotion_v1_stats1_inner_proto_depIdxs,
		MessageInfos:      file_wb_promotion_v1_stats1_inner_proto_msgTypes,
	}.Build()
	File_wb_promotion_v1_stats1_inner_proto = out.File
	file_wb_promotion_v1_stats1_inner_proto_rawDesc = nil
	file_wb_promotion_v1_stats1_inner_proto_goTypes = nil
	file_wb_promotion_v1_stats1_inner_proto_depIdxs = nil
}