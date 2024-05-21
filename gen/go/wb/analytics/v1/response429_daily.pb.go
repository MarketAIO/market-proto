//
//Описание API Аналитика
//
//Сервис предоставляет публичный API для получения аналитических данных. С помощью этих методов вы можете получать аналитические отчёты. [Часть методов](./#tag/Voronka-prodazh-(Dzhem)/) доступна только с [подпиской на расширенную аналитику Джем](https://seller.wildberries.ru/dynamic-product-categories/jam).
//
//The version of the OpenAPI document:
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: wb/analytics/v1/response429_daily.proto

package wbANALYTICS

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

type Response429Daily struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*anypb.Any `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	// Флаг ошибки
	Error bool `protobuf:"varint,2,opt,name=error,proto3" json:"error,omitempty"`
	// Описание ошибки
	ErrorText string `protobuf:"bytes,3,opt,name=errorText,proto3" json:"errorText,omitempty"`
	// Дополнительные ошибки
	AdditionalErrors []*ResponseErrorAdditionalErrorsInner `protobuf:"bytes,4,rep,name=additionalErrors,proto3" json:"additionalErrors,omitempty"`
}

func (x *Response429Daily) Reset() {
	*x = Response429Daily{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_analytics_v1_response429_daily_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response429Daily) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response429Daily) ProtoMessage() {}

func (x *Response429Daily) ProtoReflect() protoreflect.Message {
	mi := &file_wb_analytics_v1_response429_daily_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response429Daily.ProtoReflect.Descriptor instead.
func (*Response429Daily) Descriptor() ([]byte, []int) {
	return file_wb_analytics_v1_response429_daily_proto_rawDescGZIP(), []int{0}
}

func (x *Response429Daily) GetData() []*anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Response429Daily) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *Response429Daily) GetErrorText() string {
	if x != nil {
		return x.ErrorText
	}
	return ""
}

func (x *Response429Daily) GetAdditionalErrors() []*ResponseErrorAdditionalErrorsInner {
	if x != nil {
		return x.AdditionalErrors
	}
	return nil
}

var File_wb_analytics_v1_response429_daily_proto protoreflect.FileDescriptor

var file_wb_analytics_v1_response429_daily_proto_rawDesc = []byte{
	0x0a, 0x27, 0x77, 0x62, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x34, 0x32, 0x39, 0x5f, 0x64, 0x61,
	0x69, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x77, 0x62, 0x2e, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x3c, 0x77, 0x62, 0x2f, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x5f, 0x69, 0x6e, 0x6e,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x34, 0x32, 0x39, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x54, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x54, 0x65, 0x78, 0x74, 0x12, 0x5f, 0x0a, 0x10, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x33, 0x2e, 0x77, 0x62, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x10, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f, 0x2f,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x77, 0x62, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x62, 0x41, 0x4e, 0x41, 0x4c, 0x59, 0x54, 0x49,
	0x43, 0x53, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wb_analytics_v1_response429_daily_proto_rawDescOnce sync.Once
	file_wb_analytics_v1_response429_daily_proto_rawDescData = file_wb_analytics_v1_response429_daily_proto_rawDesc
)

func file_wb_analytics_v1_response429_daily_proto_rawDescGZIP() []byte {
	file_wb_analytics_v1_response429_daily_proto_rawDescOnce.Do(func() {
		file_wb_analytics_v1_response429_daily_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_analytics_v1_response429_daily_proto_rawDescData)
	})
	return file_wb_analytics_v1_response429_daily_proto_rawDescData
}

var file_wb_analytics_v1_response429_daily_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_analytics_v1_response429_daily_proto_goTypes = []interface{}{
	(*Response429Daily)(nil),                   // 0: wb.analytics.v1.Response429Daily
	(*anypb.Any)(nil),                          // 1: google.protobuf.Any
	(*ResponseErrorAdditionalErrorsInner)(nil), // 2: wb.analytics.v1.ResponseErrorAdditionalErrorsInner
}
var file_wb_analytics_v1_response429_daily_proto_depIdxs = []int32{
	1, // 0: wb.analytics.v1.Response429Daily.data:type_name -> google.protobuf.Any
	2, // 1: wb.analytics.v1.Response429Daily.additionalErrors:type_name -> wb.analytics.v1.ResponseErrorAdditionalErrorsInner
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_wb_analytics_v1_response429_daily_proto_init() }
func file_wb_analytics_v1_response429_daily_proto_init() {
	if File_wb_analytics_v1_response429_daily_proto != nil {
		return
	}
	file_wb_analytics_v1_response_error_additional_errors_inner_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wb_analytics_v1_response429_daily_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response429Daily); i {
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
			RawDescriptor: file_wb_analytics_v1_response429_daily_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_analytics_v1_response429_daily_proto_goTypes,
		DependencyIndexes: file_wb_analytics_v1_response429_daily_proto_depIdxs,
		MessageInfos:      file_wb_analytics_v1_response429_daily_proto_msgTypes,
	}.Build()
	File_wb_analytics_v1_response429_daily_proto = out.File
	file_wb_analytics_v1_response429_daily_proto_rawDesc = nil
	file_wb_analytics_v1_response429_daily_proto_goTypes = nil
	file_wb_analytics_v1_response429_daily_proto_depIdxs = nil
}
