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
// source: wb/promotion/v1/adv_v1_stats_post200_response_inner.proto

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

type AdvV1StatsPost200ResponseInner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interval *StatIntervalInterval `protobuf:"bytes,1,opt,name=interval,proto3" json:"interval,omitempty"`
	// Блок статистики
	Stats []*StatsBlok1 `protobuf:"bytes,2,rep,name=stats,proto3" json:"stats,omitempty"`
	// Даты, за которые необходимо выдать информацию.
	Dates []string `protobuf:"bytes,3,rep,name=dates,proto3" json:"dates,omitempty"`
}

func (x *AdvV1StatsPost200ResponseInner) Reset() {
	*x = AdvV1StatsPost200ResponseInner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_promotion_v1_adv_v1_stats_post200_response_inner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdvV1StatsPost200ResponseInner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdvV1StatsPost200ResponseInner) ProtoMessage() {}

func (x *AdvV1StatsPost200ResponseInner) ProtoReflect() protoreflect.Message {
	mi := &file_wb_promotion_v1_adv_v1_stats_post200_response_inner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdvV1StatsPost200ResponseInner.ProtoReflect.Descriptor instead.
func (*AdvV1StatsPost200ResponseInner) Descriptor() ([]byte, []int) {
	return file_wb_promotion_v1_adv_v1_stats_post200_response_inner_proto_rawDescGZIP(), []int{0}
}

func (x *AdvV1StatsPost200ResponseInner) GetInterval() *StatIntervalInterval {
	if x != nil {
		return x.Interval
	}
	return nil
}

func (x *AdvV1StatsPost200ResponseInner) GetStats() []*StatsBlok1 {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *AdvV1StatsPost200ResponseInner) GetDates() []string {
	if x != nil {
		return x.Dates
	}
	return nil
}

var File_wb_promotion_v1_adv_v1_stats_post200_response_inner_proto protoreflect.FileDescriptor

var file_wb_promotion_v1_adv_v1_stats_post200_response_inner_proto_rawDesc = []byte{
	0x0a, 0x39, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x64, 0x76, 0x5f, 0x76, 0x31, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x70,
	0x6f, 0x73, 0x74, 0x32, 0x30, 0x30, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x77, 0x62, 0x2e,
	0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1a, 0x77, 0x62,
	0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x77, 0x62, 0x2f, 0x70, 0x72,
	0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c,
	0x77, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x77, 0x62,
	0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x6b, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xac, 0x01, 0x0a, 0x1e, 0x41, 0x64, 0x76, 0x56, 0x31, 0x53, 0x74, 0x61, 0x74, 0x73, 0x50, 0x6f,
	0x73, 0x74, 0x32, 0x30, 0x30, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x6e,
	0x65, 0x72, 0x12, 0x41, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x77, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x08, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x31, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x77, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x42, 0x6c, 0x6f, 0x6b,
	0x31, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x61, 0x74, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x64, 0x61, 0x74, 0x65, 0x73, 0x42, 0x49,
	0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x77, 0x62,
	0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x62,
	0x50, 0x52, 0x4f, 0x4d, 0x4f, 0x54, 0x49, 0x4f, 0x4e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_wb_promotion_v1_adv_v1_stats_post200_response_inner_proto_rawDescOnce sync.Once
	file_wb_promotion_v1_adv_v1_stats_post200_response_inner_proto_rawDescData = file_wb_promotion_v1_adv_v1_stats_post200_response_inner_proto_rawDesc
)

func file_wb_promotion_v1_adv_v1_stats_post200_response_inner_proto_rawDescGZIP() []byte {
	file_wb_promotion_v1_adv_v1_stats_post200_response_inner_proto_rawDescOnce.Do(func() {
		file_wb_promotion_v1_adv_v1_stats_post200_response_inner_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_promotion_v1_adv_v1_stats_post200_response_inner_proto_rawDescData)
	})
	return file_wb_promotion_v1_adv_v1_stats_post200_response_inner_proto_rawDescData
}

var file_wb_promotion_v1_adv_v1_stats_post200_response_inner_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_promotion_v1_adv_v1_stats_post200_response_inner_proto_goTypes = []interface{}{
	(*AdvV1StatsPost200ResponseInner)(nil), // 0: wb.promotion.v1.AdvV1StatsPost200ResponseInner
	(*StatIntervalInterval)(nil),           // 1: wb.promotion.v1.StatIntervalInterval
	(*StatsBlok1)(nil),                     // 2: wb.promotion.v1.StatsBlok1
}
var file_wb_promotion_v1_adv_v1_stats_post200_response_inner_proto_depIdxs = []int32{
	1, // 0: wb.promotion.v1.AdvV1StatsPost200ResponseInner.interval:type_name -> wb.promotion.v1.StatIntervalInterval
	2, // 1: wb.promotion.v1.AdvV1StatsPost200ResponseInner.stats:type_name -> wb.promotion.v1.StatsBlok1
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_wb_promotion_v1_adv_v1_stats_post200_response_inner_proto_init() }
func file_wb_promotion_v1_adv_v1_stats_post200_response_inner_proto_init() {
	if File_wb_promotion_v1_adv_v1_stats_post200_response_inner_proto != nil {
		return
	}
	file_wb_promotion_v1_stat_proto_init()
	file_wb_promotion_v1_stat_date_proto_init()
	file_wb_promotion_v1_stat_interval_proto_init()
	file_wb_promotion_v1_stat_interval_interval_proto_init()
	file_wb_promotion_v1_stats_blok1_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wb_promotion_v1_adv_v1_stats_post200_response_inner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdvV1StatsPost200ResponseInner); i {
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
			RawDescriptor: file_wb_promotion_v1_adv_v1_stats_post200_response_inner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_promotion_v1_adv_v1_stats_post200_response_inner_proto_goTypes,
		DependencyIndexes: file_wb_promotion_v1_adv_v1_stats_post200_response_inner_proto_depIdxs,
		MessageInfos:      file_wb_promotion_v1_adv_v1_stats_post200_response_inner_proto_msgTypes,
	}.Build()
	File_wb_promotion_v1_adv_v1_stats_post200_response_inner_proto = out.File
	file_wb_promotion_v1_adv_v1_stats_post200_response_inner_proto_rawDesc = nil
	file_wb_promotion_v1_adv_v1_stats_post200_response_inner_proto_goTypes = nil
	file_wb_promotion_v1_adv_v1_stats_post200_response_inner_proto_depIdxs = nil
}
