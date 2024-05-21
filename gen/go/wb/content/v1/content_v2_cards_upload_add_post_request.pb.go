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
// source: wb/content/v1/content_v2_cards_upload_add_post_request.proto

package wbCONTENT

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

type ContentV2CardsUploadAddPostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// imtID КТ, к которой добавляется НМ
	ImtID int32 `protobuf:"varint,1,opt,name=imtID,proto3" json:"imtID,omitempty"`
	// Структура добавляемой НМ
	CardsToAdd []*ContentV2CardsUploadAddPostReqCardsToAddInner `protobuf:"bytes,2,rep,name=cardsToAdd,proto3" json:"cardsToAdd,omitempty"`
}

func (x *ContentV2CardsUploadAddPostReq) Reset() {
	*x = ContentV2CardsUploadAddPostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_content_v1_content_v2_cards_upload_add_post_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentV2CardsUploadAddPostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentV2CardsUploadAddPostReq) ProtoMessage() {}

func (x *ContentV2CardsUploadAddPostReq) ProtoReflect() protoreflect.Message {
	mi := &file_wb_content_v1_content_v2_cards_upload_add_post_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentV2CardsUploadAddPostReq.ProtoReflect.Descriptor instead.
func (*ContentV2CardsUploadAddPostReq) Descriptor() ([]byte, []int) {
	return file_wb_content_v1_content_v2_cards_upload_add_post_request_proto_rawDescGZIP(), []int{0}
}

func (x *ContentV2CardsUploadAddPostReq) GetImtID() int32 {
	if x != nil {
		return x.ImtID
	}
	return 0
}

func (x *ContentV2CardsUploadAddPostReq) GetCardsToAdd() []*ContentV2CardsUploadAddPostReqCardsToAddInner {
	if x != nil {
		return x.CardsToAdd
	}
	return nil
}

var File_wb_content_v1_content_v2_cards_upload_add_post_request_proto protoreflect.FileDescriptor

var file_wb_content_v1_content_v2_cards_upload_add_post_request_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x77, 0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x32, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73,
	0x5f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x5f, 0x70, 0x6f, 0x73, 0x74,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x77, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x4f, 0x77,
	0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x32, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x5f, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x61,
	0x64, 0x64, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94,
	0x01, 0x0a, 0x1e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x56, 0x32, 0x43, 0x61, 0x72, 0x64,
	0x73, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x69, 0x6d, 0x74, 0x49, 0x44, 0x12, 0x5c, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x73,
	0x54, 0x6f, 0x41, 0x64, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x77, 0x62,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x56, 0x32, 0x43, 0x61, 0x72, 0x64, 0x73, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x41, 0x64, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x43, 0x61, 0x72, 0x64, 0x73, 0x54,
	0x6f, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x73,
	0x54, 0x6f, 0x41, 0x64, 0x64, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f, 0x2f, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x77, 0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x3b, 0x77, 0x62, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wb_content_v1_content_v2_cards_upload_add_post_request_proto_rawDescOnce sync.Once
	file_wb_content_v1_content_v2_cards_upload_add_post_request_proto_rawDescData = file_wb_content_v1_content_v2_cards_upload_add_post_request_proto_rawDesc
)

func file_wb_content_v1_content_v2_cards_upload_add_post_request_proto_rawDescGZIP() []byte {
	file_wb_content_v1_content_v2_cards_upload_add_post_request_proto_rawDescOnce.Do(func() {
		file_wb_content_v1_content_v2_cards_upload_add_post_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_content_v1_content_v2_cards_upload_add_post_request_proto_rawDescData)
	})
	return file_wb_content_v1_content_v2_cards_upload_add_post_request_proto_rawDescData
}

var file_wb_content_v1_content_v2_cards_upload_add_post_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_content_v1_content_v2_cards_upload_add_post_request_proto_goTypes = []interface{}{
	(*ContentV2CardsUploadAddPostReq)(nil),                // 0: wb.content.v1.ContentV2CardsUploadAddPostReq
	(*ContentV2CardsUploadAddPostReqCardsToAddInner)(nil), // 1: wb.content.v1.ContentV2CardsUploadAddPostReqCardsToAddInner
}
var file_wb_content_v1_content_v2_cards_upload_add_post_request_proto_depIdxs = []int32{
	1, // 0: wb.content.v1.ContentV2CardsUploadAddPostReq.cardsToAdd:type_name -> wb.content.v1.ContentV2CardsUploadAddPostReqCardsToAddInner
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_wb_content_v1_content_v2_cards_upload_add_post_request_proto_init() }
func file_wb_content_v1_content_v2_cards_upload_add_post_request_proto_init() {
	if File_wb_content_v1_content_v2_cards_upload_add_post_request_proto != nil {
		return
	}
	file_wb_content_v1_content_v2_cards_upload_add_post_request_cards_to_add_inner_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wb_content_v1_content_v2_cards_upload_add_post_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentV2CardsUploadAddPostReq); i {
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
			RawDescriptor: file_wb_content_v1_content_v2_cards_upload_add_post_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_content_v1_content_v2_cards_upload_add_post_request_proto_goTypes,
		DependencyIndexes: file_wb_content_v1_content_v2_cards_upload_add_post_request_proto_depIdxs,
		MessageInfos:      file_wb_content_v1_content_v2_cards_upload_add_post_request_proto_msgTypes,
	}.Build()
	File_wb_content_v1_content_v2_cards_upload_add_post_request_proto = out.File
	file_wb_content_v1_content_v2_cards_upload_add_post_request_proto_rawDesc = nil
	file_wb_content_v1_content_v2_cards_upload_add_post_request_proto_goTypes = nil
	file_wb_content_v1_content_v2_cards_upload_add_post_request_proto_depIdxs = nil
}
