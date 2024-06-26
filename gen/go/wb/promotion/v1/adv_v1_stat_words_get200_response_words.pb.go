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
// source: wb/promotion/v1/adv_v1_stat_words_get200_response_words.proto

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

type AdvV1StatWordsGet200ResponseWords struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Фразовое соответствие (минус фразы)
	Phrase []string `protobuf:"bytes,1,rep,name=phrase,proto3" json:"phrase,omitempty"`
	// Точное соответствие (минус фразы)
	Strong []string `protobuf:"bytes,2,rep,name=strong,proto3" json:"strong,omitempty"`
	// Минус фразы из поиска
	Excluded []string `protobuf:"bytes,3,rep,name=excluded,proto3" json:"excluded,omitempty"`
	// Фиксированные фразы
	Pluse []string `protobuf:"bytes,4,rep,name=pluse,proto3" json:"pluse,omitempty"`
	// Блок со статистикой по ключевым фразам
	Keywords []*AdvV1StatWordsGet200ResponseWordsKeywordsInner `protobuf:"bytes,5,rep,name=keywords,proto3" json:"keywords,omitempty"`
	// Фиксированные ключевые фразы (`true` - включены, `false` - выключены)
	Fixed bool `protobuf:"varint,6,opt,name=fixed,proto3" json:"fixed,omitempty"`
}

func (x *AdvV1StatWordsGet200ResponseWords) Reset() {
	*x = AdvV1StatWordsGet200ResponseWords{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_promotion_v1_adv_v1_stat_words_get200_response_words_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdvV1StatWordsGet200ResponseWords) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdvV1StatWordsGet200ResponseWords) ProtoMessage() {}

func (x *AdvV1StatWordsGet200ResponseWords) ProtoReflect() protoreflect.Message {
	mi := &file_wb_promotion_v1_adv_v1_stat_words_get200_response_words_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdvV1StatWordsGet200ResponseWords.ProtoReflect.Descriptor instead.
func (*AdvV1StatWordsGet200ResponseWords) Descriptor() ([]byte, []int) {
	return file_wb_promotion_v1_adv_v1_stat_words_get200_response_words_proto_rawDescGZIP(), []int{0}
}

func (x *AdvV1StatWordsGet200ResponseWords) GetPhrase() []string {
	if x != nil {
		return x.Phrase
	}
	return nil
}

func (x *AdvV1StatWordsGet200ResponseWords) GetStrong() []string {
	if x != nil {
		return x.Strong
	}
	return nil
}

func (x *AdvV1StatWordsGet200ResponseWords) GetExcluded() []string {
	if x != nil {
		return x.Excluded
	}
	return nil
}

func (x *AdvV1StatWordsGet200ResponseWords) GetPluse() []string {
	if x != nil {
		return x.Pluse
	}
	return nil
}

func (x *AdvV1StatWordsGet200ResponseWords) GetKeywords() []*AdvV1StatWordsGet200ResponseWordsKeywordsInner {
	if x != nil {
		return x.Keywords
	}
	return nil
}

func (x *AdvV1StatWordsGet200ResponseWords) GetFixed() bool {
	if x != nil {
		return x.Fixed
	}
	return false
}

var File_wb_promotion_v1_adv_v1_stat_words_get200_response_words_proto protoreflect.FileDescriptor

var file_wb_promotion_v1_adv_v1_stat_words_get200_response_words_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x64, 0x76, 0x5f, 0x76, 0x31, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x5f, 0x67, 0x65, 0x74, 0x32, 0x30, 0x30, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x77, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x1a, 0x4c, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x64, 0x76, 0x5f, 0x76, 0x31, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x5f, 0x67, 0x65, 0x74, 0x32, 0x30, 0x30, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8,
	0x01, 0x0a, 0x21, 0x41, 0x64, 0x76, 0x56, 0x31, 0x53, 0x74, 0x61, 0x74, 0x57, 0x6f, 0x72, 0x64,
	0x73, 0x47, 0x65, 0x74, 0x32, 0x30, 0x30, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x57,
	0x6f, 0x72, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x72, 0x6f, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x75, 0x73, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x6c, 0x75, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x77, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x76, 0x56, 0x31,
	0x53, 0x74, 0x61, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x47, 0x65, 0x74, 0x32, 0x30, 0x30, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x4b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x78, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x66, 0x69, 0x78, 0x65, 0x64, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x49,
	0x4f, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x62, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wb_promotion_v1_adv_v1_stat_words_get200_response_words_proto_rawDescOnce sync.Once
	file_wb_promotion_v1_adv_v1_stat_words_get200_response_words_proto_rawDescData = file_wb_promotion_v1_adv_v1_stat_words_get200_response_words_proto_rawDesc
)

func file_wb_promotion_v1_adv_v1_stat_words_get200_response_words_proto_rawDescGZIP() []byte {
	file_wb_promotion_v1_adv_v1_stat_words_get200_response_words_proto_rawDescOnce.Do(func() {
		file_wb_promotion_v1_adv_v1_stat_words_get200_response_words_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_promotion_v1_adv_v1_stat_words_get200_response_words_proto_rawDescData)
	})
	return file_wb_promotion_v1_adv_v1_stat_words_get200_response_words_proto_rawDescData
}

var file_wb_promotion_v1_adv_v1_stat_words_get200_response_words_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_promotion_v1_adv_v1_stat_words_get200_response_words_proto_goTypes = []interface{}{
	(*AdvV1StatWordsGet200ResponseWords)(nil),              // 0: wb.promotion.v1.AdvV1StatWordsGet200ResponseWords
	(*AdvV1StatWordsGet200ResponseWordsKeywordsInner)(nil), // 1: wb.promotion.v1.AdvV1StatWordsGet200ResponseWordsKeywordsInner
}
var file_wb_promotion_v1_adv_v1_stat_words_get200_response_words_proto_depIdxs = []int32{
	1, // 0: wb.promotion.v1.AdvV1StatWordsGet200ResponseWords.keywords:type_name -> wb.promotion.v1.AdvV1StatWordsGet200ResponseWordsKeywordsInner
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_wb_promotion_v1_adv_v1_stat_words_get200_response_words_proto_init() }
func file_wb_promotion_v1_adv_v1_stat_words_get200_response_words_proto_init() {
	if File_wb_promotion_v1_adv_v1_stat_words_get200_response_words_proto != nil {
		return
	}
	file_wb_promotion_v1_adv_v1_stat_words_get200_response_words_keywords_inner_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wb_promotion_v1_adv_v1_stat_words_get200_response_words_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdvV1StatWordsGet200ResponseWords); i {
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
			RawDescriptor: file_wb_promotion_v1_adv_v1_stat_words_get200_response_words_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_promotion_v1_adv_v1_stat_words_get200_response_words_proto_goTypes,
		DependencyIndexes: file_wb_promotion_v1_adv_v1_stat_words_get200_response_words_proto_depIdxs,
		MessageInfos:      file_wb_promotion_v1_adv_v1_stat_words_get200_response_words_proto_msgTypes,
	}.Build()
	File_wb_promotion_v1_adv_v1_stat_words_get200_response_words_proto = out.File
	file_wb_promotion_v1_adv_v1_stat_words_get200_response_words_proto_rawDesc = nil
	file_wb_promotion_v1_adv_v1_stat_words_get200_response_words_proto_goTypes = nil
	file_wb_promotion_v1_adv_v1_stat_words_get200_response_words_proto_depIdxs = nil
}
