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
// source: wb/promotion/v1/adv_v1_promotion_count_get200_response_adverts_inner.proto

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

type AdvV1PromotionCountGet200ResponseAdvertsInner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Тип кампании
	Type int32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	// Статус кампании
	Status int32 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	// Количество кампаний
	Count int32 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	// Список кампаний
	AdvertList []*AdvV1PromotionCountGet200ResponseAdvertsInnerAdvertListInner `protobuf:"bytes,4,rep,name=advert_list,json=advertList,proto3" json:"advert_list,omitempty"`
}

func (x *AdvV1PromotionCountGet200ResponseAdvertsInner) Reset() {
	*x = AdvV1PromotionCountGet200ResponseAdvertsInner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_promotion_v1_adv_v1_promotion_count_get200_response_adverts_inner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdvV1PromotionCountGet200ResponseAdvertsInner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdvV1PromotionCountGet200ResponseAdvertsInner) ProtoMessage() {}

func (x *AdvV1PromotionCountGet200ResponseAdvertsInner) ProtoReflect() protoreflect.Message {
	mi := &file_wb_promotion_v1_adv_v1_promotion_count_get200_response_adverts_inner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdvV1PromotionCountGet200ResponseAdvertsInner.ProtoReflect.Descriptor instead.
func (*AdvV1PromotionCountGet200ResponseAdvertsInner) Descriptor() ([]byte, []int) {
	return file_wb_promotion_v1_adv_v1_promotion_count_get200_response_adverts_inner_proto_rawDescGZIP(), []int{0}
}

func (x *AdvV1PromotionCountGet200ResponseAdvertsInner) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *AdvV1PromotionCountGet200ResponseAdvertsInner) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AdvV1PromotionCountGet200ResponseAdvertsInner) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *AdvV1PromotionCountGet200ResponseAdvertsInner) GetAdvertList() []*AdvV1PromotionCountGet200ResponseAdvertsInnerAdvertListInner {
	if x != nil {
		return x.AdvertList
	}
	return nil
}

var File_wb_promotion_v1_adv_v1_promotion_count_get200_response_adverts_inner_proto protoreflect.FileDescriptor

var file_wb_promotion_v1_adv_v1_promotion_count_get200_response_adverts_inner_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x64, 0x76, 0x5f, 0x76, 0x31, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x65, 0x74, 0x32, 0x30, 0x30, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x73,
	0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x77, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x5c, 0x77,
	0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x64, 0x76, 0x5f, 0x76, 0x31, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x65, 0x74, 0x32, 0x30, 0x30, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x73, 0x5f, 0x69, 0x6e,
	0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f,
	0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x01, 0x0a, 0x2d,
	0x41, 0x64, 0x76, 0x56, 0x31, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x32, 0x30, 0x30, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x6e, 0x0a, 0x0b, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x4d, 0x2e, 0x77, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x76, 0x56, 0x31, 0x50, 0x72, 0x6f, 0x6d,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x32, 0x30, 0x30,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x73, 0x49,
	0x6e, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e,
	0x6e, 0x65, 0x72, 0x52, 0x0a, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x77,
	0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x77,
	0x62, 0x50, 0x52, 0x4f, 0x4d, 0x4f, 0x54, 0x49, 0x4f, 0x4e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_wb_promotion_v1_adv_v1_promotion_count_get200_response_adverts_inner_proto_rawDescOnce sync.Once
	file_wb_promotion_v1_adv_v1_promotion_count_get200_response_adverts_inner_proto_rawDescData = file_wb_promotion_v1_adv_v1_promotion_count_get200_response_adverts_inner_proto_rawDesc
)

func file_wb_promotion_v1_adv_v1_promotion_count_get200_response_adverts_inner_proto_rawDescGZIP() []byte {
	file_wb_promotion_v1_adv_v1_promotion_count_get200_response_adverts_inner_proto_rawDescOnce.Do(func() {
		file_wb_promotion_v1_adv_v1_promotion_count_get200_response_adverts_inner_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_promotion_v1_adv_v1_promotion_count_get200_response_adverts_inner_proto_rawDescData)
	})
	return file_wb_promotion_v1_adv_v1_promotion_count_get200_response_adverts_inner_proto_rawDescData
}

var file_wb_promotion_v1_adv_v1_promotion_count_get200_response_adverts_inner_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_promotion_v1_adv_v1_promotion_count_get200_response_adverts_inner_proto_goTypes = []interface{}{
	(*AdvV1PromotionCountGet200ResponseAdvertsInner)(nil),                // 0: wb.promotion.v1.AdvV1PromotionCountGet200ResponseAdvertsInner
	(*AdvV1PromotionCountGet200ResponseAdvertsInnerAdvertListInner)(nil), // 1: wb.promotion.v1.AdvV1PromotionCountGet200ResponseAdvertsInnerAdvertListInner
}
var file_wb_promotion_v1_adv_v1_promotion_count_get200_response_adverts_inner_proto_depIdxs = []int32{
	1, // 0: wb.promotion.v1.AdvV1PromotionCountGet200ResponseAdvertsInner.advert_list:type_name -> wb.promotion.v1.AdvV1PromotionCountGet200ResponseAdvertsInnerAdvertListInner
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_wb_promotion_v1_adv_v1_promotion_count_get200_response_adverts_inner_proto_init() }
func file_wb_promotion_v1_adv_v1_promotion_count_get200_response_adverts_inner_proto_init() {
	if File_wb_promotion_v1_adv_v1_promotion_count_get200_response_adverts_inner_proto != nil {
		return
	}
	file_wb_promotion_v1_adv_v1_promotion_count_get200_response_adverts_inner_advert_list_inner_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wb_promotion_v1_adv_v1_promotion_count_get200_response_adverts_inner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdvV1PromotionCountGet200ResponseAdvertsInner); i {
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
			RawDescriptor: file_wb_promotion_v1_adv_v1_promotion_count_get200_response_adverts_inner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_promotion_v1_adv_v1_promotion_count_get200_response_adverts_inner_proto_goTypes,
		DependencyIndexes: file_wb_promotion_v1_adv_v1_promotion_count_get200_response_adverts_inner_proto_depIdxs,
		MessageInfos:      file_wb_promotion_v1_adv_v1_promotion_count_get200_response_adverts_inner_proto_msgTypes,
	}.Build()
	File_wb_promotion_v1_adv_v1_promotion_count_get200_response_adverts_inner_proto = out.File
	file_wb_promotion_v1_adv_v1_promotion_count_get200_response_adverts_inner_proto_rawDesc = nil
	file_wb_promotion_v1_adv_v1_promotion_count_get200_response_adverts_inner_proto_goTypes = nil
	file_wb_promotion_v1_adv_v1_promotion_count_get200_response_adverts_inner_proto_depIdxs = nil
}