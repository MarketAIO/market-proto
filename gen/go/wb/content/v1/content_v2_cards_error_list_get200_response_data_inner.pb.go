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
// source: wb/content/v1/content_v2_cards_error_list_get200_response_data_inner.proto

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

type ContentV2CardsErrorListGet200ResponseDataInner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Категория для который создавалось КТ с данной НМ
	Object string `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	// Артикул продавца
	VendorCode string `protobuf:"bytes,2,opt,name=vendorCode,proto3" json:"vendorCode,omitempty"`
	// Дата и время запроса на создание КТ с данным НМ
	UpdateAt string `protobuf:"bytes,3,opt,name=updateAt,proto3" json:"updateAt,omitempty"`
	// Список ошибок из-за которых не обработался запрос на создание КТ с данным НМ
	Errors []string `protobuf:"bytes,4,rep,name=errors,proto3" json:"errors,omitempty"`
	// Идентификатор предмета
	ObjectID int32 `protobuf:"varint,5,opt,name=objectID,proto3" json:"objectID,omitempty"`
}

func (x *ContentV2CardsErrorListGet200ResponseDataInner) Reset() {
	*x = ContentV2CardsErrorListGet200ResponseDataInner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_content_v1_content_v2_cards_error_list_get200_response_data_inner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentV2CardsErrorListGet200ResponseDataInner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentV2CardsErrorListGet200ResponseDataInner) ProtoMessage() {}

func (x *ContentV2CardsErrorListGet200ResponseDataInner) ProtoReflect() protoreflect.Message {
	mi := &file_wb_content_v1_content_v2_cards_error_list_get200_response_data_inner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentV2CardsErrorListGet200ResponseDataInner.ProtoReflect.Descriptor instead.
func (*ContentV2CardsErrorListGet200ResponseDataInner) Descriptor() ([]byte, []int) {
	return file_wb_content_v1_content_v2_cards_error_list_get200_response_data_inner_proto_rawDescGZIP(), []int{0}
}

func (x *ContentV2CardsErrorListGet200ResponseDataInner) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *ContentV2CardsErrorListGet200ResponseDataInner) GetVendorCode() string {
	if x != nil {
		return x.VendorCode
	}
	return ""
}

func (x *ContentV2CardsErrorListGet200ResponseDataInner) GetUpdateAt() string {
	if x != nil {
		return x.UpdateAt
	}
	return ""
}

func (x *ContentV2CardsErrorListGet200ResponseDataInner) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *ContentV2CardsErrorListGet200ResponseDataInner) GetObjectID() int32 {
	if x != nil {
		return x.ObjectID
	}
	return 0
}

var File_wb_content_v1_content_v2_cards_error_list_get200_response_data_inner_proto protoreflect.FileDescriptor

var file_wb_content_v1_content_v2_cards_error_list_get200_response_data_inner_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x77, 0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x32, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73,
	0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x67, 0x65, 0x74, 0x32,
	0x30, 0x30, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x77, 0x62,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x22, 0xb8, 0x01, 0x0a, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x56, 0x32, 0x43, 0x61, 0x72, 0x64, 0x73, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x65, 0x74, 0x32, 0x30, 0x30, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f, 0x2f, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x77, 0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x3b, 0x77, 0x62, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_wb_content_v1_content_v2_cards_error_list_get200_response_data_inner_proto_rawDescOnce sync.Once
	file_wb_content_v1_content_v2_cards_error_list_get200_response_data_inner_proto_rawDescData = file_wb_content_v1_content_v2_cards_error_list_get200_response_data_inner_proto_rawDesc
)

func file_wb_content_v1_content_v2_cards_error_list_get200_response_data_inner_proto_rawDescGZIP() []byte {
	file_wb_content_v1_content_v2_cards_error_list_get200_response_data_inner_proto_rawDescOnce.Do(func() {
		file_wb_content_v1_content_v2_cards_error_list_get200_response_data_inner_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_content_v1_content_v2_cards_error_list_get200_response_data_inner_proto_rawDescData)
	})
	return file_wb_content_v1_content_v2_cards_error_list_get200_response_data_inner_proto_rawDescData
}

var file_wb_content_v1_content_v2_cards_error_list_get200_response_data_inner_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_content_v1_content_v2_cards_error_list_get200_response_data_inner_proto_goTypes = []interface{}{
	(*ContentV2CardsErrorListGet200ResponseDataInner)(nil), // 0: wb.content.v1.ContentV2CardsErrorListGet200ResponseDataInner
}
var file_wb_content_v1_content_v2_cards_error_list_get200_response_data_inner_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wb_content_v1_content_v2_cards_error_list_get200_response_data_inner_proto_init() }
func file_wb_content_v1_content_v2_cards_error_list_get200_response_data_inner_proto_init() {
	if File_wb_content_v1_content_v2_cards_error_list_get200_response_data_inner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wb_content_v1_content_v2_cards_error_list_get200_response_data_inner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentV2CardsErrorListGet200ResponseDataInner); i {
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
			RawDescriptor: file_wb_content_v1_content_v2_cards_error_list_get200_response_data_inner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_content_v1_content_v2_cards_error_list_get200_response_data_inner_proto_goTypes,
		DependencyIndexes: file_wb_content_v1_content_v2_cards_error_list_get200_response_data_inner_proto_depIdxs,
		MessageInfos:      file_wb_content_v1_content_v2_cards_error_list_get200_response_data_inner_proto_msgTypes,
	}.Build()
	File_wb_content_v1_content_v2_cards_error_list_get200_response_data_inner_proto = out.File
	file_wb_content_v1_content_v2_cards_error_list_get200_response_data_inner_proto_rawDesc = nil
	file_wb_content_v1_content_v2_cards_error_list_get200_response_data_inner_proto_goTypes = nil
	file_wb_content_v1_content_v2_cards_error_list_get200_response_data_inner_proto_depIdxs = nil
}
