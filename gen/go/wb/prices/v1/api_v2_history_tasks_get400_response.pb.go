//
//API цен и скидок
//
//С помощью этих методов можно устанавливать цены и скидки. Максимум — 10 запросов за 6 секунд суммарно для всех методов.
//
//The version of the OpenAPI document:
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: wb/prices/v1/api_v2_history_tasks_get400_response.proto

package wbPRICES

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

type ApiV2HistoryTasksGet400Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data      *anypb.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Error     bool       `protobuf:"varint,2,opt,name=error,proto3" json:"error,omitempty"`
	ErrorText string     `protobuf:"bytes,3,opt,name=errorText,proto3" json:"errorText,omitempty"`
}

func (x *ApiV2HistoryTasksGet400Response) Reset() {
	*x = ApiV2HistoryTasksGet400Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_prices_v1_api_v2_history_tasks_get400_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiV2HistoryTasksGet400Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiV2HistoryTasksGet400Response) ProtoMessage() {}

func (x *ApiV2HistoryTasksGet400Response) ProtoReflect() protoreflect.Message {
	mi := &file_wb_prices_v1_api_v2_history_tasks_get400_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiV2HistoryTasksGet400Response.ProtoReflect.Descriptor instead.
func (*ApiV2HistoryTasksGet400Response) Descriptor() ([]byte, []int) {
	return file_wb_prices_v1_api_v2_history_tasks_get400_response_proto_rawDescGZIP(), []int{0}
}

func (x *ApiV2HistoryTasksGet400Response) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ApiV2HistoryTasksGet400Response) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *ApiV2HistoryTasksGet400Response) GetErrorText() string {
	if x != nil {
		return x.ErrorText
	}
	return ""
}

var File_wb_prices_v1_api_v2_history_tasks_get400_response_proto protoreflect.FileDescriptor

var file_wb_prices_v1_api_v2_history_tasks_get400_response_proto_rawDesc = []byte{
	0x0a, 0x37, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x70, 0x69, 0x5f, 0x76, 0x32, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x5f, 0x67, 0x65, 0x74, 0x34, 0x30, 0x30, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x77, 0x62, 0x2e, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x1f, 0x41, 0x70, 0x69, 0x56, 0x32, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x47, 0x65, 0x74, 0x34, 0x30, 0x30, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x65,
	0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x54,
	0x65, 0x78, 0x74, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f, 0x2f, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x70, 0x62, 0x2f, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b,
	0x77, 0x62, 0x50, 0x52, 0x49, 0x43, 0x45, 0x53, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wb_prices_v1_api_v2_history_tasks_get400_response_proto_rawDescOnce sync.Once
	file_wb_prices_v1_api_v2_history_tasks_get400_response_proto_rawDescData = file_wb_prices_v1_api_v2_history_tasks_get400_response_proto_rawDesc
)

func file_wb_prices_v1_api_v2_history_tasks_get400_response_proto_rawDescGZIP() []byte {
	file_wb_prices_v1_api_v2_history_tasks_get400_response_proto_rawDescOnce.Do(func() {
		file_wb_prices_v1_api_v2_history_tasks_get400_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_prices_v1_api_v2_history_tasks_get400_response_proto_rawDescData)
	})
	return file_wb_prices_v1_api_v2_history_tasks_get400_response_proto_rawDescData
}

var file_wb_prices_v1_api_v2_history_tasks_get400_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_prices_v1_api_v2_history_tasks_get400_response_proto_goTypes = []interface{}{
	(*ApiV2HistoryTasksGet400Response)(nil), // 0: wb.prices.v1.ApiV2HistoryTasksGet400Response
	(*anypb.Any)(nil),                       // 1: google.protobuf.Any
}
var file_wb_prices_v1_api_v2_history_tasks_get400_response_proto_depIdxs = []int32{
	1, // 0: wb.prices.v1.ApiV2HistoryTasksGet400Response.data:type_name -> google.protobuf.Any
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_wb_prices_v1_api_v2_history_tasks_get400_response_proto_init() }
func file_wb_prices_v1_api_v2_history_tasks_get400_response_proto_init() {
	if File_wb_prices_v1_api_v2_history_tasks_get400_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wb_prices_v1_api_v2_history_tasks_get400_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiV2HistoryTasksGet400Response); i {
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
			RawDescriptor: file_wb_prices_v1_api_v2_history_tasks_get400_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_prices_v1_api_v2_history_tasks_get400_response_proto_goTypes,
		DependencyIndexes: file_wb_prices_v1_api_v2_history_tasks_get400_response_proto_depIdxs,
		MessageInfos:      file_wb_prices_v1_api_v2_history_tasks_get400_response_proto_msgTypes,
	}.Build()
	File_wb_prices_v1_api_v2_history_tasks_get400_response_proto = out.File
	file_wb_prices_v1_api_v2_history_tasks_get400_response_proto_rawDesc = nil
	file_wb_prices_v1_api_v2_history_tasks_get400_response_proto_goTypes = nil
	file_wb_prices_v1_api_v2_history_tasks_get400_response_proto_depIdxs = nil
}
