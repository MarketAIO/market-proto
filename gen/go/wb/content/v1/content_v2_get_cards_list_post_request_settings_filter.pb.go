//
//Описание API Контента
//
//<dl> <dt>Словарь сокращений:</dt> <dd>КТ — карточка товара</dd> <dd>НМ — номенклатура</dd> </dl> Ограничения по количеству запросов: <dd>Допускается максимум 100 запросов в минуту на методы контента в целом.</dd>  <br> Публичное API Контента создано для синхронизации данных между серверами Wildberries и серверами продавцов. <br> Вы загружаете данные на свои носители, работаете с ними на своих мощностях и синхронизируетесь с нашими серверами по мере необходимости. <br> <code>Не допускается использование API Контента в качестве внешней базы данных. При превышении лимитов на запросы доступ к API будет ограничен.</code>
//
//The version of the OpenAPI document:
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: wb/content/v1/content_v2_get_cards_list_post_request_settings_filter.proto

package wbContent

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

type ContentV2GetCardsListPostReqSettingsFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Фильтр по фото:      * `0` — только карточки без фото   * `1` — только карточки с фото   * `-1` — все карточки товара
	WithPhoto int32 `protobuf:"varint,1,opt,name=withPhoto,proto3" json:"withPhoto,omitempty"`
	// Поиск по артикулу продавца, артикулу WB, баркоду
	TextSearch string `protobuf:"bytes,2,opt,name=textSearch,proto3" json:"textSearch,omitempty"`
	// Поиск по ID тегов
	TagIDs []int32 `protobuf:"varint,3,rep,packed,name=tagIDs,proto3" json:"tagIDs,omitempty"`
	// Фильтр по категории. `true` - только разрешённые, `false` - все. Не используется в песочнице.
	AllowedCategoriesOnly bool `protobuf:"varint,4,opt,name=allowedCategoriesOnly,proto3" json:"allowedCategoriesOnly,omitempty"`
	// Поиск по id предметов
	ObjectIDs []int32 `protobuf:"varint,5,rep,packed,name=objectIDs,proto3" json:"objectIDs,omitempty"`
	// Поиск по брендам
	Brands []string `protobuf:"bytes,6,rep,name=brands,proto3" json:"brands,omitempty"`
	// Поиск по идентификатору КТ
	ImtID int32 `protobuf:"varint,7,opt,name=imtID,proto3" json:"imtID,omitempty"`
}

func (x *ContentV2GetCardsListPostReqSettingsFilter) Reset() {
	*x = ContentV2GetCardsListPostReqSettingsFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_content_v1_content_v2_get_cards_list_post_request_settings_filter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentV2GetCardsListPostReqSettingsFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentV2GetCardsListPostReqSettingsFilter) ProtoMessage() {}

func (x *ContentV2GetCardsListPostReqSettingsFilter) ProtoReflect() protoreflect.Message {
	mi := &file_wb_content_v1_content_v2_get_cards_list_post_request_settings_filter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentV2GetCardsListPostReqSettingsFilter.ProtoReflect.Descriptor instead.
func (*ContentV2GetCardsListPostReqSettingsFilter) Descriptor() ([]byte, []int) {
	return file_wb_content_v1_content_v2_get_cards_list_post_request_settings_filter_proto_rawDescGZIP(), []int{0}
}

func (x *ContentV2GetCardsListPostReqSettingsFilter) GetWithPhoto() int32 {
	if x != nil {
		return x.WithPhoto
	}
	return 0
}

func (x *ContentV2GetCardsListPostReqSettingsFilter) GetTextSearch() string {
	if x != nil {
		return x.TextSearch
	}
	return ""
}

func (x *ContentV2GetCardsListPostReqSettingsFilter) GetTagIDs() []int32 {
	if x != nil {
		return x.TagIDs
	}
	return nil
}

func (x *ContentV2GetCardsListPostReqSettingsFilter) GetAllowedCategoriesOnly() bool {
	if x != nil {
		return x.AllowedCategoriesOnly
	}
	return false
}

func (x *ContentV2GetCardsListPostReqSettingsFilter) GetObjectIDs() []int32 {
	if x != nil {
		return x.ObjectIDs
	}
	return nil
}

func (x *ContentV2GetCardsListPostReqSettingsFilter) GetBrands() []string {
	if x != nil {
		return x.Brands
	}
	return nil
}

func (x *ContentV2GetCardsListPostReqSettingsFilter) GetImtID() int32 {
	if x != nil {
		return x.ImtID
	}
	return 0
}

var File_wb_content_v1_content_v2_get_cards_list_post_request_settings_filter_proto protoreflect.FileDescriptor

var file_wb_content_v1_content_v2_get_cards_list_post_request_settings_filter_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x77, 0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x32, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x63,
	0x61, 0x72, 0x64, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x5f,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x77, 0x62,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x84, 0x02, 0x0a, 0x2a,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x56, 0x32, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x64,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x69,
	0x74, 0x68, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x77,
	0x69, 0x74, 0x68, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x78, 0x74,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65,
	0x78, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x67, 0x49,
	0x44, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x74, 0x61, 0x67, 0x49, 0x44, 0x73,
	0x12, 0x34, 0x0a, 0x15, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x4f, 0x6e, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x15, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x44, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x44, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6d, 0x74, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6d, 0x74,
	0x49, 0x44, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x77,
	0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x62, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wb_content_v1_content_v2_get_cards_list_post_request_settings_filter_proto_rawDescOnce sync.Once
	file_wb_content_v1_content_v2_get_cards_list_post_request_settings_filter_proto_rawDescData = file_wb_content_v1_content_v2_get_cards_list_post_request_settings_filter_proto_rawDesc
)

func file_wb_content_v1_content_v2_get_cards_list_post_request_settings_filter_proto_rawDescGZIP() []byte {
	file_wb_content_v1_content_v2_get_cards_list_post_request_settings_filter_proto_rawDescOnce.Do(func() {
		file_wb_content_v1_content_v2_get_cards_list_post_request_settings_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_content_v1_content_v2_get_cards_list_post_request_settings_filter_proto_rawDescData)
	})
	return file_wb_content_v1_content_v2_get_cards_list_post_request_settings_filter_proto_rawDescData
}

var file_wb_content_v1_content_v2_get_cards_list_post_request_settings_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_content_v1_content_v2_get_cards_list_post_request_settings_filter_proto_goTypes = []interface{}{
	(*ContentV2GetCardsListPostReqSettingsFilter)(nil), // 0: wb.content.v1.ContentV2GetCardsListPostReqSettingsFilter
}
var file_wb_content_v1_content_v2_get_cards_list_post_request_settings_filter_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wb_content_v1_content_v2_get_cards_list_post_request_settings_filter_proto_init() }
func file_wb_content_v1_content_v2_get_cards_list_post_request_settings_filter_proto_init() {
	if File_wb_content_v1_content_v2_get_cards_list_post_request_settings_filter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wb_content_v1_content_v2_get_cards_list_post_request_settings_filter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentV2GetCardsListPostReqSettingsFilter); i {
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
			RawDescriptor: file_wb_content_v1_content_v2_get_cards_list_post_request_settings_filter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_content_v1_content_v2_get_cards_list_post_request_settings_filter_proto_goTypes,
		DependencyIndexes: file_wb_content_v1_content_v2_get_cards_list_post_request_settings_filter_proto_depIdxs,
		MessageInfos:      file_wb_content_v1_content_v2_get_cards_list_post_request_settings_filter_proto_msgTypes,
	}.Build()
	File_wb_content_v1_content_v2_get_cards_list_post_request_settings_filter_proto = out.File
	file_wb_content_v1_content_v2_get_cards_list_post_request_settings_filter_proto_rawDesc = nil
	file_wb_content_v1_content_v2_get_cards_list_post_request_settings_filter_proto_goTypes = nil
	file_wb_content_v1_content_v2_get_cards_list_post_request_settings_filter_proto_depIdxs = nil
}
