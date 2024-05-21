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
// source: wb/content/v1/content_v2_cards_update_post_request_inner.proto

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

type ContentV2CardsUpdatePostReqInner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Артикул WB
	NmID int32 `protobuf:"varint,1,opt,name=nmID,proto3" json:"nmID,omitempty"`
	// Артикул продавца
	VendorCode string `protobuf:"bytes,2,opt,name=vendorCode,proto3" json:"vendorCode,omitempty"`
	// Бренд
	Brand string `protobuf:"bytes,3,opt,name=brand,proto3" json:"brand,omitempty"`
	// Наименование товара
	Title string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	// Описание товара. Максимальное количество символов зависит от категории товара. Стандарт — 2000, минимум — 1000, максимум — 5000.<br> Подробно о правилах описания в **Правилах заполнения карточки товара** в разделе [Инструкции](https://seller.wildberries.ru/training) на портале продавцов.
	Description string                                      `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Dimensions  *ContentV2CardsUpdatePostReqInnerDimensions `protobuf:"bytes,6,opt,name=dimensions,proto3" json:"dimensions,omitempty"`
	// Характеристики товара
	Characteristics []*ContentV2CardsUpdatePostReqInnerCharacteristicsInner `protobuf:"bytes,7,rep,name=characteristics,proto3" json:"characteristics,omitempty"`
	// Массив размеров артикула. <br> Для безразмерного товара все равно нужно передавать данный массив без параметров (wbSize и techSize), но с баркодом.
	Sizes []*ContentV2CardsUpdatePostReqInnerSizesInner `protobuf:"bytes,8,rep,name=sizes,proto3" json:"sizes,omitempty"`
}

func (x *ContentV2CardsUpdatePostReqInner) Reset() {
	*x = ContentV2CardsUpdatePostReqInner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_content_v1_content_v2_cards_update_post_request_inner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentV2CardsUpdatePostReqInner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentV2CardsUpdatePostReqInner) ProtoMessage() {}

func (x *ContentV2CardsUpdatePostReqInner) ProtoReflect() protoreflect.Message {
	mi := &file_wb_content_v1_content_v2_cards_update_post_request_inner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentV2CardsUpdatePostReqInner.ProtoReflect.Descriptor instead.
func (*ContentV2CardsUpdatePostReqInner) Descriptor() ([]byte, []int) {
	return file_wb_content_v1_content_v2_cards_update_post_request_inner_proto_rawDescGZIP(), []int{0}
}

func (x *ContentV2CardsUpdatePostReqInner) GetNmID() int32 {
	if x != nil {
		return x.NmID
	}
	return 0
}

func (x *ContentV2CardsUpdatePostReqInner) GetVendorCode() string {
	if x != nil {
		return x.VendorCode
	}
	return ""
}

func (x *ContentV2CardsUpdatePostReqInner) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *ContentV2CardsUpdatePostReqInner) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ContentV2CardsUpdatePostReqInner) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ContentV2CardsUpdatePostReqInner) GetDimensions() *ContentV2CardsUpdatePostReqInnerDimensions {
	if x != nil {
		return x.Dimensions
	}
	return nil
}

func (x *ContentV2CardsUpdatePostReqInner) GetCharacteristics() []*ContentV2CardsUpdatePostReqInnerCharacteristicsInner {
	if x != nil {
		return x.Characteristics
	}
	return nil
}

func (x *ContentV2CardsUpdatePostReqInner) GetSizes() []*ContentV2CardsUpdatePostReqInnerSizesInner {
	if x != nil {
		return x.Sizes
	}
	return nil
}

var File_wb_content_v1_content_v2_cards_update_post_request_inner_proto protoreflect.FileDescriptor

var file_wb_content_v1_content_v2_cards_update_post_request_inner_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x77, 0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x32, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73,
	0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x77, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a,
	0x54, 0x77, 0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x32, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x5f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x72, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x49, 0x77, 0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x32, 0x5f,
	0x63, 0x61, 0x72, 0x64, 0x73, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x73,
	0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f,
	0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x4a, 0x77, 0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x32, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73,
	0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x73,
	0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x03, 0x0a,
	0x20, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x56, 0x32, 0x43, 0x61, 0x72, 0x64, 0x73, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x49, 0x6e, 0x6e, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x6e, 0x6d, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x59, 0x0a, 0x0a, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x77, 0x62, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x56,
	0x32, 0x43, 0x61, 0x72, 0x64, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x0a, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x6d,
	0x0a, 0x0f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x77, 0x62, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x56,
	0x32, 0x43, 0x61, 0x72, 0x64, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x0f, 0x63, 0x68,
	0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x4f, 0x0a,
	0x05, 0x73, 0x69, 0x7a, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x77,
	0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x56, 0x32, 0x43, 0x61, 0x72, 0x64, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x53, 0x69, 0x7a,
	0x65, 0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x05, 0x73, 0x69, 0x7a, 0x65, 0x73, 0x42, 0x45,
	0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x77, 0x62,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x62, 0x43, 0x4f,
	0x4e, 0x54, 0x45, 0x4e, 0x54, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wb_content_v1_content_v2_cards_update_post_request_inner_proto_rawDescOnce sync.Once
	file_wb_content_v1_content_v2_cards_update_post_request_inner_proto_rawDescData = file_wb_content_v1_content_v2_cards_update_post_request_inner_proto_rawDesc
)

func file_wb_content_v1_content_v2_cards_update_post_request_inner_proto_rawDescGZIP() []byte {
	file_wb_content_v1_content_v2_cards_update_post_request_inner_proto_rawDescOnce.Do(func() {
		file_wb_content_v1_content_v2_cards_update_post_request_inner_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_content_v1_content_v2_cards_update_post_request_inner_proto_rawDescData)
	})
	return file_wb_content_v1_content_v2_cards_update_post_request_inner_proto_rawDescData
}

var file_wb_content_v1_content_v2_cards_update_post_request_inner_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_content_v1_content_v2_cards_update_post_request_inner_proto_goTypes = []interface{}{
	(*ContentV2CardsUpdatePostReqInner)(nil),                     // 0: wb.content.v1.ContentV2CardsUpdatePostReqInner
	(*ContentV2CardsUpdatePostReqInnerDimensions)(nil),           // 1: wb.content.v1.ContentV2CardsUpdatePostReqInnerDimensions
	(*ContentV2CardsUpdatePostReqInnerCharacteristicsInner)(nil), // 2: wb.content.v1.ContentV2CardsUpdatePostReqInnerCharacteristicsInner
	(*ContentV2CardsUpdatePostReqInnerSizesInner)(nil),           // 3: wb.content.v1.ContentV2CardsUpdatePostReqInnerSizesInner
}
var file_wb_content_v1_content_v2_cards_update_post_request_inner_proto_depIdxs = []int32{
	1, // 0: wb.content.v1.ContentV2CardsUpdatePostReqInner.dimensions:type_name -> wb.content.v1.ContentV2CardsUpdatePostReqInnerDimensions
	2, // 1: wb.content.v1.ContentV2CardsUpdatePostReqInner.characteristics:type_name -> wb.content.v1.ContentV2CardsUpdatePostReqInnerCharacteristicsInner
	3, // 2: wb.content.v1.ContentV2CardsUpdatePostReqInner.sizes:type_name -> wb.content.v1.ContentV2CardsUpdatePostReqInnerSizesInner
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_wb_content_v1_content_v2_cards_update_post_request_inner_proto_init() }
func file_wb_content_v1_content_v2_cards_update_post_request_inner_proto_init() {
	if File_wb_content_v1_content_v2_cards_update_post_request_inner_proto != nil {
		return
	}
	file_wb_content_v1_content_v2_cards_update_post_request_inner_characteristics_inner_proto_init()
	file_wb_content_v1_content_v2_cards_update_post_request_inner_dimensions_proto_init()
	file_wb_content_v1_content_v2_cards_update_post_request_inner_sizes_inner_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wb_content_v1_content_v2_cards_update_post_request_inner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentV2CardsUpdatePostReqInner); i {
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
			RawDescriptor: file_wb_content_v1_content_v2_cards_update_post_request_inner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_content_v1_content_v2_cards_update_post_request_inner_proto_goTypes,
		DependencyIndexes: file_wb_content_v1_content_v2_cards_update_post_request_inner_proto_depIdxs,
		MessageInfos:      file_wb_content_v1_content_v2_cards_update_post_request_inner_proto_msgTypes,
	}.Build()
	File_wb_content_v1_content_v2_cards_update_post_request_inner_proto = out.File
	file_wb_content_v1_content_v2_cards_update_post_request_inner_proto_rawDesc = nil
	file_wb_content_v1_content_v2_cards_update_post_request_inner_proto_goTypes = nil
	file_wb_content_v1_content_v2_cards_update_post_request_inner_proto_depIdxs = nil
}
