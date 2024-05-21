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
// source: wb/promotion/v1/response_info_advert_type9_united_params_inner.proto

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

type ResponseInfoAdvertType9UnitedParamsInner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject *ResponseInfoAdvertType9UnitedParamsInnerSubject      `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Menus   []*ResponseInfoAdvertType9UnitedParamsInnerMenusInner `protobuf:"bytes,2,rep,name=menus,proto3" json:"menus,omitempty"`
	// Артикулы WB
	Nms []int32 `protobuf:"varint,3,rep,packed,name=nms,proto3" json:"nms,omitempty"`
	// Ставка в Поиске.
	SearchCPM int32 `protobuf:"varint,4,opt,name=searchCPM,proto3" json:"searchCPM,omitempty"`
	// Ставка в Каталоге, при наличии.
	CatalogCPM int32 `protobuf:"varint,5,opt,name=catalogCPM,proto3" json:"catalogCPM,omitempty"`
}

func (x *ResponseInfoAdvertType9UnitedParamsInner) Reset() {
	*x = ResponseInfoAdvertType9UnitedParamsInner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_promotion_v1_response_info_advert_type9_united_params_inner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseInfoAdvertType9UnitedParamsInner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseInfoAdvertType9UnitedParamsInner) ProtoMessage() {}

func (x *ResponseInfoAdvertType9UnitedParamsInner) ProtoReflect() protoreflect.Message {
	mi := &file_wb_promotion_v1_response_info_advert_type9_united_params_inner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseInfoAdvertType9UnitedParamsInner.ProtoReflect.Descriptor instead.
func (*ResponseInfoAdvertType9UnitedParamsInner) Descriptor() ([]byte, []int) {
	return file_wb_promotion_v1_response_info_advert_type9_united_params_inner_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseInfoAdvertType9UnitedParamsInner) GetSubject() *ResponseInfoAdvertType9UnitedParamsInnerSubject {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *ResponseInfoAdvertType9UnitedParamsInner) GetMenus() []*ResponseInfoAdvertType9UnitedParamsInnerMenusInner {
	if x != nil {
		return x.Menus
	}
	return nil
}

func (x *ResponseInfoAdvertType9UnitedParamsInner) GetNms() []int32 {
	if x != nil {
		return x.Nms
	}
	return nil
}

func (x *ResponseInfoAdvertType9UnitedParamsInner) GetSearchCPM() int32 {
	if x != nil {
		return x.SearchCPM
	}
	return 0
}

func (x *ResponseInfoAdvertType9UnitedParamsInner) GetCatalogCPM() int32 {
	if x != nil {
		return x.CatalogCPM
	}
	return 0
}

var File_wb_promotion_v1_response_info_advert_type9_united_params_inner_proto protoreflect.FileDescriptor

var file_wb_promotion_v1_response_info_advert_type9_united_params_inner_proto_rawDesc = []byte{
	0x0a, 0x44, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f,
	0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x39, 0x5f, 0x75, 0x6e, 0x69,
	0x74, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x77, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x50, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x39, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x6e, 0x75, 0x73, 0x5f, 0x69, 0x6e,
	0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4c, 0x77, 0x62, 0x2f, 0x70, 0x72,
	0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x39, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x02, 0x0a, 0x28, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x39, 0x55, 0x6e, 0x69, 0x74, 0x65, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x49,
	0x6e, 0x6e, 0x65, 0x72, 0x12, 0x5a, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x77, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x39, 0x55,
	0x6e, 0x69, 0x74, 0x65, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72,
	0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x59, 0x0a, 0x05, 0x6d, 0x65, 0x6e, 0x75, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x43, 0x2e, 0x77, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x41, 0x64,
	0x76, 0x65, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x39, 0x55, 0x6e, 0x69, 0x74, 0x65, 0x64, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x49,
	0x6e, 0x6e, 0x65, 0x72, 0x52, 0x05, 0x6d, 0x65, 0x6e, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6e,
	0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x6d, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x50, 0x4d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x50, 0x4d, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x43, 0x50, 0x4d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x43, 0x50, 0x4d, 0x42, 0x49, 0x5a, 0x47, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x41, 0x49, 0x4f, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x77, 0x62, 0x2f, 0x70, 0x72,
	0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x62, 0x50, 0x52, 0x4f,
	0x4d, 0x4f, 0x54, 0x49, 0x4f, 0x4e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wb_promotion_v1_response_info_advert_type9_united_params_inner_proto_rawDescOnce sync.Once
	file_wb_promotion_v1_response_info_advert_type9_united_params_inner_proto_rawDescData = file_wb_promotion_v1_response_info_advert_type9_united_params_inner_proto_rawDesc
)

func file_wb_promotion_v1_response_info_advert_type9_united_params_inner_proto_rawDescGZIP() []byte {
	file_wb_promotion_v1_response_info_advert_type9_united_params_inner_proto_rawDescOnce.Do(func() {
		file_wb_promotion_v1_response_info_advert_type9_united_params_inner_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_promotion_v1_response_info_advert_type9_united_params_inner_proto_rawDescData)
	})
	return file_wb_promotion_v1_response_info_advert_type9_united_params_inner_proto_rawDescData
}

var file_wb_promotion_v1_response_info_advert_type9_united_params_inner_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_promotion_v1_response_info_advert_type9_united_params_inner_proto_goTypes = []interface{}{
	(*ResponseInfoAdvertType9UnitedParamsInner)(nil),           // 0: wb.promotion.v1.ResponseInfoAdvertType9UnitedParamsInner
	(*ResponseInfoAdvertType9UnitedParamsInnerSubject)(nil),    // 1: wb.promotion.v1.ResponseInfoAdvertType9UnitedParamsInnerSubject
	(*ResponseInfoAdvertType9UnitedParamsInnerMenusInner)(nil), // 2: wb.promotion.v1.ResponseInfoAdvertType9UnitedParamsInnerMenusInner
}
var file_wb_promotion_v1_response_info_advert_type9_united_params_inner_proto_depIdxs = []int32{
	1, // 0: wb.promotion.v1.ResponseInfoAdvertType9UnitedParamsInner.subject:type_name -> wb.promotion.v1.ResponseInfoAdvertType9UnitedParamsInnerSubject
	2, // 1: wb.promotion.v1.ResponseInfoAdvertType9UnitedParamsInner.menus:type_name -> wb.promotion.v1.ResponseInfoAdvertType9UnitedParamsInnerMenusInner
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_wb_promotion_v1_response_info_advert_type9_united_params_inner_proto_init() }
func file_wb_promotion_v1_response_info_advert_type9_united_params_inner_proto_init() {
	if File_wb_promotion_v1_response_info_advert_type9_united_params_inner_proto != nil {
		return
	}
	file_wb_promotion_v1_response_info_advert_type9_united_params_inner_menus_inner_proto_init()
	file_wb_promotion_v1_response_info_advert_type9_united_params_inner_subject_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wb_promotion_v1_response_info_advert_type9_united_params_inner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseInfoAdvertType9UnitedParamsInner); i {
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
			RawDescriptor: file_wb_promotion_v1_response_info_advert_type9_united_params_inner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_promotion_v1_response_info_advert_type9_united_params_inner_proto_goTypes,
		DependencyIndexes: file_wb_promotion_v1_response_info_advert_type9_united_params_inner_proto_depIdxs,
		MessageInfos:      file_wb_promotion_v1_response_info_advert_type9_united_params_inner_proto_msgTypes,
	}.Build()
	File_wb_promotion_v1_response_info_advert_type9_united_params_inner_proto = out.File
	file_wb_promotion_v1_response_info_advert_type9_united_params_inner_proto_rawDesc = nil
	file_wb_promotion_v1_response_info_advert_type9_united_params_inner_proto_goTypes = nil
	file_wb_promotion_v1_response_info_advert_type9_united_params_inner_proto_depIdxs = nil
}
