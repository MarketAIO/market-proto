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
// source: wb/content/v1/response_content_error1.proto

package wbCONTENT

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResponseContentError1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *anypb.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// Флаг ошибки
	Error bool `protobuf:"varint,2,opt,name=error,proto3" json:"error,omitempty"`
	// Описание ошибки
	ErrorText        string                                 `protobuf:"bytes,3,opt,name=errorText,proto3" json:"errorText,omitempty"`
	AdditionalErrors *ResponseContentError1AdditionalErrors `protobuf:"bytes,4,opt,name=additionalErrors,proto3" json:"additionalErrors,omitempty"`
}

func (x *ResponseContentError1) Reset() {
	*x = ResponseContentError1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_content_v1_response_content_error1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseContentError1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseContentError1) ProtoMessage() {}

func (x *ResponseContentError1) ProtoReflect() protoreflect.Message {
	mi := &file_wb_content_v1_response_content_error1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseContentError1.ProtoReflect.Descriptor instead.
func (*ResponseContentError1) Descriptor() ([]byte, []int) {
	return file_wb_content_v1_response_content_error1_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseContentError1) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ResponseContentError1) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *ResponseContentError1) GetErrorText() string {
	if x != nil {
		return x.ErrorText
	}
	return ""
}

func (x *ResponseContentError1) GetAdditionalErrors() *ResponseContentError1AdditionalErrors {
	if x != nil {
		return x.AdditionalErrors
	}
	return nil
}

var File_wb_content_v1_response_content_error1_proto protoreflect.FileDescriptor

var file_wb_content_v1_response_content_error1_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x77, 0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x77,
	0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x3d, 0x77, 0x62,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x31, 0x5f, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x31,
	0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x65, 0x78, 0x74, 0x12, 0x60,
	0x0a, 0x10, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x77, 0x62, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x31, 0x41, 0x64,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x52, 0x10,
	0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f,
	0x77, 0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x62,
	0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wb_content_v1_response_content_error1_proto_rawDescOnce sync.Once
	file_wb_content_v1_response_content_error1_proto_rawDescData = file_wb_content_v1_response_content_error1_proto_rawDesc
)

func file_wb_content_v1_response_content_error1_proto_rawDescGZIP() []byte {
	file_wb_content_v1_response_content_error1_proto_rawDescOnce.Do(func() {
		file_wb_content_v1_response_content_error1_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_content_v1_response_content_error1_proto_rawDescData)
	})
	return file_wb_content_v1_response_content_error1_proto_rawDescData
}

var file_wb_content_v1_response_content_error1_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_content_v1_response_content_error1_proto_goTypes = []interface{}{
	(*ResponseContentError1)(nil),                 // 0: wb.content.v1.ResponseContentError1
	(*anypb.Any)(nil),                             // 1: google.protobuf.Any
	(*ResponseContentError1AdditionalErrors)(nil), // 2: wb.content.v1.ResponseContentError1AdditionalErrors
}
var file_wb_content_v1_response_content_error1_proto_depIdxs = []int32{
	1, // 0: wb.content.v1.ResponseContentError1.data:type_name -> google.protobuf.Any
	2, // 1: wb.content.v1.ResponseContentError1.additionalErrors:type_name -> wb.content.v1.ResponseContentError1AdditionalErrors
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_wb_content_v1_response_content_error1_proto_init() }
func file_wb_content_v1_response_content_error1_proto_init() {
	if File_wb_content_v1_response_content_error1_proto != nil {
		return
	}
	file_wb_content_v1_response_content_error1_additional_errors_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wb_content_v1_response_content_error1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseContentError1); i {
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
			RawDescriptor: file_wb_content_v1_response_content_error1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_content_v1_response_content_error1_proto_goTypes,
		DependencyIndexes: file_wb_content_v1_response_content_error1_proto_depIdxs,
		MessageInfos:      file_wb_content_v1_response_content_error1_proto_msgTypes,
	}.Build()
	File_wb_content_v1_response_content_error1_proto = out.File
	file_wb_content_v1_response_content_error1_proto_rawDesc = nil
	file_wb_content_v1_response_content_error1_proto_goTypes = nil
	file_wb_content_v1_response_content_error1_proto_depIdxs = nil
}
