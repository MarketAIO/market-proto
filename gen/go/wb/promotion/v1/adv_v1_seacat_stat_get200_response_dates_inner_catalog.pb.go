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
// source: wb/promotion/v1/adv_v1_seacat_stat_get200_response_dates_inner_catalog.proto

package wbPromotion

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

type AdvV1SeacatStatGet200ResponseDatesInnerCatalog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Количество просмотров
	Views int32 `protobuf:"varint,1,opt,name=views,proto3" json:"views,omitempty"`
	// Количество кликов
	Clicks int32 `protobuf:"varint,2,opt,name=clicks,proto3" json:"clicks,omitempty"`
	// Количество заказов
	Orders int32 `protobuf:"varint,3,opt,name=orders,proto3" json:"orders,omitempty"`
	// Затраты, ₽.
	Sum int32 `protobuf:"varint,4,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *AdvV1SeacatStatGet200ResponseDatesInnerCatalog) Reset() {
	*x = AdvV1SeacatStatGet200ResponseDatesInnerCatalog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_promotion_v1_adv_v1_seacat_stat_get200_response_dates_inner_catalog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdvV1SeacatStatGet200ResponseDatesInnerCatalog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdvV1SeacatStatGet200ResponseDatesInnerCatalog) ProtoMessage() {}

func (x *AdvV1SeacatStatGet200ResponseDatesInnerCatalog) ProtoReflect() protoreflect.Message {
	mi := &file_wb_promotion_v1_adv_v1_seacat_stat_get200_response_dates_inner_catalog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdvV1SeacatStatGet200ResponseDatesInnerCatalog.ProtoReflect.Descriptor instead.
func (*AdvV1SeacatStatGet200ResponseDatesInnerCatalog) Descriptor() ([]byte, []int) {
	return file_wb_promotion_v1_adv_v1_seacat_stat_get200_response_dates_inner_catalog_proto_rawDescGZIP(), []int{0}
}

func (x *AdvV1SeacatStatGet200ResponseDatesInnerCatalog) GetViews() int32 {
	if x != nil {
		return x.Views
	}
	return 0
}

func (x *AdvV1SeacatStatGet200ResponseDatesInnerCatalog) GetClicks() int32 {
	if x != nil {
		return x.Clicks
	}
	return 0
}

func (x *AdvV1SeacatStatGet200ResponseDatesInnerCatalog) GetOrders() int32 {
	if x != nil {
		return x.Orders
	}
	return 0
}

func (x *AdvV1SeacatStatGet200ResponseDatesInnerCatalog) GetSum() int32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

var File_wb_promotion_v1_adv_v1_seacat_stat_get200_response_dates_inner_catalog_proto protoreflect.FileDescriptor

var file_wb_promotion_v1_adv_v1_seacat_stat_get200_response_dates_inner_catalog_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x64, 0x76, 0x5f, 0x76, 0x31, 0x5f, 0x73, 0x65, 0x61, 0x63, 0x61, 0x74, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x5f, 0x67, 0x65, 0x74, 0x32, 0x30, 0x30, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72,
	0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x77, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22,
	0x88, 0x01, 0x0a, 0x2e, 0x41, 0x64, 0x76, 0x56, 0x31, 0x53, 0x65, 0x61, 0x63, 0x61, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x47, 0x65, 0x74, 0x32, 0x30, 0x30, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x44, 0x61, 0x74, 0x65, 0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x43, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x63,
	0x6b, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x41,
	0x49, 0x4f, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x62, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wb_promotion_v1_adv_v1_seacat_stat_get200_response_dates_inner_catalog_proto_rawDescOnce sync.Once
	file_wb_promotion_v1_adv_v1_seacat_stat_get200_response_dates_inner_catalog_proto_rawDescData = file_wb_promotion_v1_adv_v1_seacat_stat_get200_response_dates_inner_catalog_proto_rawDesc
)

func file_wb_promotion_v1_adv_v1_seacat_stat_get200_response_dates_inner_catalog_proto_rawDescGZIP() []byte {
	file_wb_promotion_v1_adv_v1_seacat_stat_get200_response_dates_inner_catalog_proto_rawDescOnce.Do(func() {
		file_wb_promotion_v1_adv_v1_seacat_stat_get200_response_dates_inner_catalog_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_promotion_v1_adv_v1_seacat_stat_get200_response_dates_inner_catalog_proto_rawDescData)
	})
	return file_wb_promotion_v1_adv_v1_seacat_stat_get200_response_dates_inner_catalog_proto_rawDescData
}

var file_wb_promotion_v1_adv_v1_seacat_stat_get200_response_dates_inner_catalog_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_promotion_v1_adv_v1_seacat_stat_get200_response_dates_inner_catalog_proto_goTypes = []interface{}{
	(*AdvV1SeacatStatGet200ResponseDatesInnerCatalog)(nil), // 0: wb.promotion.v1.AdvV1SeacatStatGet200ResponseDatesInnerCatalog
}
var file_wb_promotion_v1_adv_v1_seacat_stat_get200_response_dates_inner_catalog_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wb_promotion_v1_adv_v1_seacat_stat_get200_response_dates_inner_catalog_proto_init() }
func file_wb_promotion_v1_adv_v1_seacat_stat_get200_response_dates_inner_catalog_proto_init() {
	if File_wb_promotion_v1_adv_v1_seacat_stat_get200_response_dates_inner_catalog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wb_promotion_v1_adv_v1_seacat_stat_get200_response_dates_inner_catalog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdvV1SeacatStatGet200ResponseDatesInnerCatalog); i {
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
			RawDescriptor: file_wb_promotion_v1_adv_v1_seacat_stat_get200_response_dates_inner_catalog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_promotion_v1_adv_v1_seacat_stat_get200_response_dates_inner_catalog_proto_goTypes,
		DependencyIndexes: file_wb_promotion_v1_adv_v1_seacat_stat_get200_response_dates_inner_catalog_proto_depIdxs,
		MessageInfos:      file_wb_promotion_v1_adv_v1_seacat_stat_get200_response_dates_inner_catalog_proto_msgTypes,
	}.Build()
	File_wb_promotion_v1_adv_v1_seacat_stat_get200_response_dates_inner_catalog_proto = out.File
	file_wb_promotion_v1_adv_v1_seacat_stat_get200_response_dates_inner_catalog_proto_rawDesc = nil
	file_wb_promotion_v1_adv_v1_seacat_stat_get200_response_dates_inner_catalog_proto_goTypes = nil
	file_wb_promotion_v1_adv_v1_seacat_stat_get200_response_dates_inner_catalog_proto_depIdxs = nil
}
