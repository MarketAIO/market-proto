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
// source: wb/analytics/v1/by_selected_nm_id_req.proto

package wbAnalytics

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

type BySelectedNmIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Идентификатор отчёта в UUID-формате. Генерируется продавцом самостоятельно
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Тип отчёта — `DETAIL_HISTORY_REPORT`
	ReportType string `protobuf:"bytes,2,opt,name=reportType,proto3" json:"reportType,omitempty"`
	// Название отчёта (если не указано, сформируется автоматически)
	UserReportName string                   `protobuf:"bytes,3,opt,name=userReportName,proto3" json:"userReportName,omitempty"`
	Params         *BySelectedNmIDReqParams `protobuf:"bytes,4,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *BySelectedNmIDReq) Reset() {
	*x = BySelectedNmIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_analytics_v1_by_selected_nm_id_req_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BySelectedNmIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BySelectedNmIDReq) ProtoMessage() {}

func (x *BySelectedNmIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_wb_analytics_v1_by_selected_nm_id_req_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BySelectedNmIDReq.ProtoReflect.Descriptor instead.
func (*BySelectedNmIDReq) Descriptor() ([]byte, []int) {
	return file_wb_analytics_v1_by_selected_nm_id_req_proto_rawDescGZIP(), []int{0}
}

func (x *BySelectedNmIDReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BySelectedNmIDReq) GetReportType() string {
	if x != nil {
		return x.ReportType
	}
	return ""
}

func (x *BySelectedNmIDReq) GetUserReportName() string {
	if x != nil {
		return x.UserReportName
	}
	return ""
}

func (x *BySelectedNmIDReq) GetParams() *BySelectedNmIDReqParams {
	if x != nil {
		return x.Params
	}
	return nil
}

var File_wb_analytics_v1_by_selected_nm_id_req_proto protoreflect.FileDescriptor

var file_wb_analytics_v1_by_selected_nm_id_req_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x77, 0x62, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x62, 0x79, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x6e, 0x6d,
	0x5f, 0x69, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x77,
	0x62, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x32,
	0x77, 0x62, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x79, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x6e, 0x6d, 0x5f, 0x69,
	0x64, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a, 0x11, 0x42, 0x79, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x4e, 0x6d, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x40, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x77, 0x62, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x79, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4e, 0x6d, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x77,
	0x62, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x77,
	0x62, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_wb_analytics_v1_by_selected_nm_id_req_proto_rawDescOnce sync.Once
	file_wb_analytics_v1_by_selected_nm_id_req_proto_rawDescData = file_wb_analytics_v1_by_selected_nm_id_req_proto_rawDesc
)

func file_wb_analytics_v1_by_selected_nm_id_req_proto_rawDescGZIP() []byte {
	file_wb_analytics_v1_by_selected_nm_id_req_proto_rawDescOnce.Do(func() {
		file_wb_analytics_v1_by_selected_nm_id_req_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_analytics_v1_by_selected_nm_id_req_proto_rawDescData)
	})
	return file_wb_analytics_v1_by_selected_nm_id_req_proto_rawDescData
}

var file_wb_analytics_v1_by_selected_nm_id_req_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_analytics_v1_by_selected_nm_id_req_proto_goTypes = []interface{}{
	(*BySelectedNmIDReq)(nil),       // 0: wb.analytics.v1.BySelectedNmIDReq
	(*BySelectedNmIDReqParams)(nil), // 1: wb.analytics.v1.BySelectedNmIDReqParams
}
var file_wb_analytics_v1_by_selected_nm_id_req_proto_depIdxs = []int32{
	1, // 0: wb.analytics.v1.BySelectedNmIDReq.params:type_name -> wb.analytics.v1.BySelectedNmIDReqParams
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_wb_analytics_v1_by_selected_nm_id_req_proto_init() }
func file_wb_analytics_v1_by_selected_nm_id_req_proto_init() {
	if File_wb_analytics_v1_by_selected_nm_id_req_proto != nil {
		return
	}
	file_wb_analytics_v1_by_selected_nm_id_req_params_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wb_analytics_v1_by_selected_nm_id_req_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BySelectedNmIDReq); i {
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
			RawDescriptor: file_wb_analytics_v1_by_selected_nm_id_req_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_analytics_v1_by_selected_nm_id_req_proto_goTypes,
		DependencyIndexes: file_wb_analytics_v1_by_selected_nm_id_req_proto_depIdxs,
		MessageInfos:      file_wb_analytics_v1_by_selected_nm_id_req_proto_msgTypes,
	}.Build()
	File_wb_analytics_v1_by_selected_nm_id_req_proto = out.File
	file_wb_analytics_v1_by_selected_nm_id_req_proto_rawDesc = nil
	file_wb_analytics_v1_by_selected_nm_id_req_proto_goTypes = nil
	file_wb_analytics_v1_by_selected_nm_id_req_proto_depIdxs = nil
}
