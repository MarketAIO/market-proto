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
// source: wb/promotion/v1/adv_v1_stat_words_get200_response_stat_inner.proto

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

type AdvV1StatWordsGet200ResponseStatInner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Идентификатор кампании в системе Wildberries
	AdvertId int32 `protobuf:"varint,1,opt,name=advertId,proto3" json:"advertId,omitempty"`
	// Ключевая фраза
	Keyword string `protobuf:"bytes,2,opt,name=keyword,proto3" json:"keyword,omitempty"`
	// Поле перманентно отключено
	AdvertName string `protobuf:"bytes,3,opt,name=advertName,proto3" json:"advertName,omitempty"`
	// Название кампании
	CampaignName string `protobuf:"bytes,4,opt,name=campaignName,proto3" json:"campaignName,omitempty"`
	// Дата запуска кампании
	Begin string `protobuf:"bytes,5,opt,name=begin,proto3" json:"begin,omitempty"`
	// Дата завершения кампании
	End string `protobuf:"bytes,6,opt,name=end,proto3" json:"end,omitempty"`
	// Количество просмотров
	Views int32 `protobuf:"varint,7,opt,name=views,proto3" json:"views,omitempty"`
	// Количество кликов
	Clicks int32 `protobuf:"varint,8,opt,name=clicks,proto3" json:"clicks,omitempty"`
	// Частота (отношение количества просмотров к количеству уникальных пользователей)
	Frq int32 `protobuf:"varint,9,opt,name=frq,proto3" json:"frq,omitempty"`
	// Показатель кликабельности (отношение числа кликов к количеству показов. Выражается в процентах).
	Ctr float32 `protobuf:"fixed32,10,opt,name=ctr,proto3" json:"ctr,omitempty"`
	// Стоимость клика, ₽
	Cpc float32 `protobuf:"fixed32,11,opt,name=cpc,proto3" json:"cpc,omitempty"`
	// Длительность кампании, в секундах
	Duration int32 `protobuf:"varint,12,opt,name=duration,proto3" json:"duration,omitempty"`
	// Затраты, ₽.
	Sum float32 `protobuf:"fixed32,13,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *AdvV1StatWordsGet200ResponseStatInner) Reset() {
	*x = AdvV1StatWordsGet200ResponseStatInner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_promotion_v1_adv_v1_stat_words_get200_response_stat_inner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdvV1StatWordsGet200ResponseStatInner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdvV1StatWordsGet200ResponseStatInner) ProtoMessage() {}

func (x *AdvV1StatWordsGet200ResponseStatInner) ProtoReflect() protoreflect.Message {
	mi := &file_wb_promotion_v1_adv_v1_stat_words_get200_response_stat_inner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdvV1StatWordsGet200ResponseStatInner.ProtoReflect.Descriptor instead.
func (*AdvV1StatWordsGet200ResponseStatInner) Descriptor() ([]byte, []int) {
	return file_wb_promotion_v1_adv_v1_stat_words_get200_response_stat_inner_proto_rawDescGZIP(), []int{0}
}

func (x *AdvV1StatWordsGet200ResponseStatInner) GetAdvertId() int32 {
	if x != nil {
		return x.AdvertId
	}
	return 0
}

func (x *AdvV1StatWordsGet200ResponseStatInner) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *AdvV1StatWordsGet200ResponseStatInner) GetAdvertName() string {
	if x != nil {
		return x.AdvertName
	}
	return ""
}

func (x *AdvV1StatWordsGet200ResponseStatInner) GetCampaignName() string {
	if x != nil {
		return x.CampaignName
	}
	return ""
}

func (x *AdvV1StatWordsGet200ResponseStatInner) GetBegin() string {
	if x != nil {
		return x.Begin
	}
	return ""
}

func (x *AdvV1StatWordsGet200ResponseStatInner) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

func (x *AdvV1StatWordsGet200ResponseStatInner) GetViews() int32 {
	if x != nil {
		return x.Views
	}
	return 0
}

func (x *AdvV1StatWordsGet200ResponseStatInner) GetClicks() int32 {
	if x != nil {
		return x.Clicks
	}
	return 0
}

func (x *AdvV1StatWordsGet200ResponseStatInner) GetFrq() int32 {
	if x != nil {
		return x.Frq
	}
	return 0
}

func (x *AdvV1StatWordsGet200ResponseStatInner) GetCtr() float32 {
	if x != nil {
		return x.Ctr
	}
	return 0
}

func (x *AdvV1StatWordsGet200ResponseStatInner) GetCpc() float32 {
	if x != nil {
		return x.Cpc
	}
	return 0
}

func (x *AdvV1StatWordsGet200ResponseStatInner) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *AdvV1StatWordsGet200ResponseStatInner) GetSum() float32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

var File_wb_promotion_v1_adv_v1_stat_words_get200_response_stat_inner_proto protoreflect.FileDescriptor

var file_wb_promotion_v1_adv_v1_stat_words_get200_response_stat_inner_proto_rawDesc = []byte{
	0x0a, 0x42, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x64, 0x76, 0x5f, 0x76, 0x31, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x5f, 0x67, 0x65, 0x74, 0x32, 0x30, 0x30, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x77, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0xdb, 0x02, 0x0a, 0x25, 0x41, 0x64, 0x76, 0x56, 0x31, 0x53,
	0x74, 0x61, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x47, 0x65, 0x74, 0x32, 0x30, 0x30, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x64, 0x76, 0x65, 0x72,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x6d,
	0x70, 0x61, 0x69, 0x67, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x65, 0x67,
	0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x63, 0x6b,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x66, 0x72, 0x71, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x66, 0x72,
	0x71, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x74, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03,
	0x63, 0x74, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x63, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x03, 0x63, 0x70, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03,
	0x73, 0x75, 0x6d, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f, 0x2f, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x77, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b,
	0x77, 0x62, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_wb_promotion_v1_adv_v1_stat_words_get200_response_stat_inner_proto_rawDescOnce sync.Once
	file_wb_promotion_v1_adv_v1_stat_words_get200_response_stat_inner_proto_rawDescData = file_wb_promotion_v1_adv_v1_stat_words_get200_response_stat_inner_proto_rawDesc
)

func file_wb_promotion_v1_adv_v1_stat_words_get200_response_stat_inner_proto_rawDescGZIP() []byte {
	file_wb_promotion_v1_adv_v1_stat_words_get200_response_stat_inner_proto_rawDescOnce.Do(func() {
		file_wb_promotion_v1_adv_v1_stat_words_get200_response_stat_inner_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_promotion_v1_adv_v1_stat_words_get200_response_stat_inner_proto_rawDescData)
	})
	return file_wb_promotion_v1_adv_v1_stat_words_get200_response_stat_inner_proto_rawDescData
}

var file_wb_promotion_v1_adv_v1_stat_words_get200_response_stat_inner_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_promotion_v1_adv_v1_stat_words_get200_response_stat_inner_proto_goTypes = []interface{}{
	(*AdvV1StatWordsGet200ResponseStatInner)(nil), // 0: wb.promotion.v1.AdvV1StatWordsGet200ResponseStatInner
}
var file_wb_promotion_v1_adv_v1_stat_words_get200_response_stat_inner_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wb_promotion_v1_adv_v1_stat_words_get200_response_stat_inner_proto_init() }
func file_wb_promotion_v1_adv_v1_stat_words_get200_response_stat_inner_proto_init() {
	if File_wb_promotion_v1_adv_v1_stat_words_get200_response_stat_inner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wb_promotion_v1_adv_v1_stat_words_get200_response_stat_inner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdvV1StatWordsGet200ResponseStatInner); i {
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
			RawDescriptor: file_wb_promotion_v1_adv_v1_stat_words_get200_response_stat_inner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_promotion_v1_adv_v1_stat_words_get200_response_stat_inner_proto_goTypes,
		DependencyIndexes: file_wb_promotion_v1_adv_v1_stat_words_get200_response_stat_inner_proto_depIdxs,
		MessageInfos:      file_wb_promotion_v1_adv_v1_stat_words_get200_response_stat_inner_proto_msgTypes,
	}.Build()
	File_wb_promotion_v1_adv_v1_stat_words_get200_response_stat_inner_proto = out.File
	file_wb_promotion_v1_adv_v1_stat_words_get200_response_stat_inner_proto_rawDesc = nil
	file_wb_promotion_v1_adv_v1_stat_words_get200_response_stat_inner_proto_goTypes = nil
	file_wb_promotion_v1_adv_v1_stat_words_get200_response_stat_inner_proto_depIdxs = nil
}
