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
// source: wb/prices/v1/task_already_exists_error_data.proto

package wbPrices

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

type TaskAlreadyExistsErrorData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID загрузки
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Флаг дублирования загрузки: `true` — такая загрузка уже есть
	AlreadyExists bool `protobuf:"varint,2,opt,name=alreadyExists,proto3" json:"alreadyExists,omitempty"`
}

func (x *TaskAlreadyExistsErrorData) Reset() {
	*x = TaskAlreadyExistsErrorData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_prices_v1_task_already_exists_error_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskAlreadyExistsErrorData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskAlreadyExistsErrorData) ProtoMessage() {}

func (x *TaskAlreadyExistsErrorData) ProtoReflect() protoreflect.Message {
	mi := &file_wb_prices_v1_task_already_exists_error_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskAlreadyExistsErrorData.ProtoReflect.Descriptor instead.
func (*TaskAlreadyExistsErrorData) Descriptor() ([]byte, []int) {
	return file_wb_prices_v1_task_already_exists_error_data_proto_rawDescGZIP(), []int{0}
}

func (x *TaskAlreadyExistsErrorData) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TaskAlreadyExistsErrorData) GetAlreadyExists() bool {
	if x != nil {
		return x.AlreadyExists
	}
	return false
}

var File_wb_prices_v1_task_already_exists_error_data_proto protoreflect.FileDescriptor

var file_wb_prices_v1_task_already_exists_error_data_proto_rawDesc = []byte{
	0x0a, 0x31, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x61, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x5f, 0x65, 0x78, 0x69, 0x73,
	0x74, 0x73, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x77, 0x62, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x22, 0x52, 0x0a, 0x1a, 0x54, 0x61, 0x73, 0x6b, 0x41, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x24, 0x0a, 0x0d, 0x61, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x61, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x45,
	0x78, 0x69, 0x73, 0x74, 0x73, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f, 0x2f, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x77,
	0x62, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wb_prices_v1_task_already_exists_error_data_proto_rawDescOnce sync.Once
	file_wb_prices_v1_task_already_exists_error_data_proto_rawDescData = file_wb_prices_v1_task_already_exists_error_data_proto_rawDesc
)

func file_wb_prices_v1_task_already_exists_error_data_proto_rawDescGZIP() []byte {
	file_wb_prices_v1_task_already_exists_error_data_proto_rawDescOnce.Do(func() {
		file_wb_prices_v1_task_already_exists_error_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_prices_v1_task_already_exists_error_data_proto_rawDescData)
	})
	return file_wb_prices_v1_task_already_exists_error_data_proto_rawDescData
}

var file_wb_prices_v1_task_already_exists_error_data_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_prices_v1_task_already_exists_error_data_proto_goTypes = []interface{}{
	(*TaskAlreadyExistsErrorData)(nil), // 0: wb.prices.v1.TaskAlreadyExistsErrorData
}
var file_wb_prices_v1_task_already_exists_error_data_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wb_prices_v1_task_already_exists_error_data_proto_init() }
func file_wb_prices_v1_task_already_exists_error_data_proto_init() {
	if File_wb_prices_v1_task_already_exists_error_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wb_prices_v1_task_already_exists_error_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskAlreadyExistsErrorData); i {
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
			RawDescriptor: file_wb_prices_v1_task_already_exists_error_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_prices_v1_task_already_exists_error_data_proto_goTypes,
		DependencyIndexes: file_wb_prices_v1_task_already_exists_error_data_proto_depIdxs,
		MessageInfos:      file_wb_prices_v1_task_already_exists_error_data_proto_msgTypes,
	}.Build()
	File_wb_prices_v1_task_already_exists_error_data_proto = out.File
	file_wb_prices_v1_task_already_exists_error_data_proto_rawDesc = nil
	file_wb_prices_v1_task_already_exists_error_data_proto_goTypes = nil
	file_wb_prices_v1_task_already_exists_error_data_proto_depIdxs = nil
}
